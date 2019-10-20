package session

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var GOOGLE_AUTHENTICATION_URL string = "https://www.googleapis.com/auth/userinfo.email"
var GOOGLE_AUTHORIZATION_URL string = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

const (
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	SECRET_LENGTH = 24
	ID_LENGTH = 10
)

type SessionServer struct {
	approvedEmails []string
	dryRun bool
	googleOauthConfig *oauth2.Config
	activeSessions map[string]*Session
}

type User struct {
	Email string `json:"email"`
}

type Session struct {
	secret string
	validated bool
	createdAt time.Time
}

func NewServer(dryRun bool) SessionServer {
	s := SessionServer{
		dryRun: dryRun,
	}
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  os.Getenv("BACKEND_URL") + "/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{GOOGLE_AUTHENTICATION_URL},
		Endpoint:     google.Endpoint,
	}
	s.googleOauthConfig = googleOauthConfig

	envList := os.Getenv("APPROVED_EMAILS")
	s.approvedEmails = strings.Split(envList,",")

	s.activeSessions = make(map[string]*Session)

	return s
}

func (s SessionServer) extractUser(id string, code string) (User, error) {
	user := User{}
	token, err := s.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return user, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get(GOOGLE_AUTHORIZATION_URL + token.AccessToken)
	if err != nil {
		return user, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return user, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	err = json.Unmarshal(contents, &user)
	if err != nil {
		return user, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return user, nil
}

func randomString(n int) string {
	s := make([]byte, n)
	for  i := range s {
		s[i] = characters[rand.Int63() % int64(len(characters))]
	}
	return string(s)
}

func (s SessionServer) GenerateSession() string {
	session := Session{
		secret:    randomString(SECRET_LENGTH),
		validated: false,
		createdAt: time.Now(),
	}
	id := randomString(ID_LENGTH)
	s.activeSessions[id] = &session
	return s.googleOauthConfig.AuthCodeURL(id)
}

func (s SessionServer) AuthenticateUser(id string, code string) (string, error) {
	session := s.activeSessions[id]
	if session == nil {
		return "", errors.New(fmt.Sprintf("Session with id '%s' not found \n", id))
	}

	user, err := s.extractUser(id, code)
	if err != nil {
		return "", err
	}
	if !s.approvedUser(user.Email) {
		return "", errors.New(fmt.Sprintf("User '%s' not approved", user.Email))
	}

	sessionTimeout := session.createdAt.Add(12 * time.Hour)
	if !inTimeSpan(session.createdAt, sessionTimeout, time.Now()) {
		delete(s.activeSessions, id)
		return "", errors.New(fmt.Sprintln("Stale session, please sign in again"))
	}

	session.validated = true

	return session.secret, nil
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (s SessionServer) approvedUser(email string) bool {
	for i := range s.approvedEmails {
		if s.approvedEmails[i] == email {
			return true
		}
	}
	return false
}

func (s SessionServer) ValidateSecret(secret string) bool {
	for i := range s.activeSessions {
		if s.activeSessions[i].secret == secret {
			if s.activeSessions[i].validated {
				sessionTimeout := s.activeSessions[i].createdAt.Add(12* time.Hour)
				if inTimeSpan(s.activeSessions[i].createdAt, sessionTimeout, time.Now()) {
					return true
				}
				delete(s.activeSessions, i)
			}
		}
	}
	return false
}
