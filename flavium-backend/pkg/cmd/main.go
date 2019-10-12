package main

import (
	"../server"
	pb "../torrents"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	grpcPort = "10000"
)

var (
	getEndpoint  = flag.String("get", "localhost:"+grpcPort, "endpoint of TorrentService")
	postEndpoint = flag.String("post", "localhost:"+grpcPort, "endpoint of TorrentService")

	oauthStateString = "pseudo-random"

	googleOauthConfig *oauth2.Config
)

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterTorrentHandlerFromEndpoint(ctx, mux, *getEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	err = pb.RegisterTorrentHandlerFromEndpoint(ctx, mux, *postEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		if r.URL.Path != "/login" && r.URL.Path != "/callback" {
			state, err := r.Cookie("oauthstate")
			if err != nil {
				fmt.Println(err.Error())
				http.Redirect(w,r,"/401", http.StatusUnauthorized)
				return
			}
			if state.Value != oauthStateString {
				fmt.Println("User not authenticated")
				http.Redirect(w,r,"/401", http.StatusUnauthorized)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func Run(address string, opts ...runtime.ServeMuxOption) error {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}
	mux.HandleFunc("/callback", handleGoogleCallback)
	mux.HandleFunc("/login", handleGoogleLogin)
	mux.Handle("/", gw)

	return http.ListenAndServe(address, allowCORS(mux))

}


func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	user, err := extractUser(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w,r,"/", http.StatusTemporaryRedirect)
		return
	}
	//TODO(Vilddjur): fetch list of approved emails
	if user.Email == "approvedemail@example.com" {
		storeCookie(w, r.FormValue("state"))
		fmt.Fprintf(w, "Content: %s\n", user)
		return
	} else {
		fmt.Println("User not approved")
		http.Redirect(w,r,"/", http.StatusUnauthorized)
		return
	}
}

func storeCookie(w http.ResponseWriter, state string) {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
}

type User struct {
	Id string `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:verified_email`
	Picture string `json:"picture"`
}

func extractUser(state string, code string) (User, error) {
	user := User{}
	if state != oauthStateString {
		return user, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return user, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
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

func main(){
	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterTorrentServer(s, &server.TorrentServer{})

		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	flag.Parse()
	defer glog.Flush()

	if err := Run(":8080"); err != nil {
		glog.Fatal(err)
	}
}
