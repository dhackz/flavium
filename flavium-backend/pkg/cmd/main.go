package main

import (
	"flavium-backend/pkg/session"
	"flavium-backend/pkg/server"
	pb "flavium-backend/pkg/torrents"
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	grpcPort = "10000"
	COOKIE_NAME = "sessioncookie"
)

var (
	getEndpoint  = flag.String("get", "localhost:"+grpcPort, "endpoint of TorrentService")
	postEndpoint = flag.String("post", "localhost:"+grpcPort, "endpoint of TorrentService")

	dryRun = flag.Bool("dry_run", false, "Print commands instead of running them")

	authServer = session.NewServer(*dryRun)
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
	credentials := []string{"Access-Control-Allow-Credentials", "true"}
	headers := []string{"Content-Type", "Accept",credentials[0]}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	w.Header().Set(credentials[0], credentials[1])
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	fmt.Printf("preflight request for %s \n", r.URL.Path)
	return
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		credentials := []string{"Access-Control-Allow-Credentials", "true"}
		expose := []string{"Access-Control-Expose-Headers", "Location"}
		w.Header().Set(credentials[0], credentials[1])
		w.Header().Set(expose[0], expose[1])
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		if r.URL.Path != "/login" && r.URL.Path != "/callback" {
			secret, err := r.Cookie(COOKIE_NAME)
			if err != nil {
				fmt.Println(err.Error())
				http.Redirect(w,r,"/401", http.StatusUnauthorized)
				return
			}
			if !authServer.ValidateSecret(secret.Value) {
				fmt.Println("User not authenticated")
				cookie := http.Cookie{Name: COOKIE_NAME, Value: "", MaxAge: 0, Expires: time.Now(), Path: "/"}
				http.SetCookie(w, &cookie)
				http.Redirect(w,r,"/401", http.StatusUnauthorized)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func Run(address string, opts ...runtime.ServeMuxOption) error {
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
	mux.HandleFunc("/auth", handleAuth)
	mux.Handle("/", gw)

	return http.ListenAndServe(address, allowCORS(mux))
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	secret, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w,r,"/401", http.StatusUnauthorized)
		return
	}
	if !authServer.ValidateSecret(secret.Value) {
		fmt.Println("Cookie not valid, sign in again")
		cookie := http.Cookie{Name: COOKIE_NAME, Value: "", MaxAge: 0, Expires: time.Now(), Path: "/"}
		http.SetCookie(w, &cookie)
		http.Redirect(w,r,"/401", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "ok")
	return
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := authServer.GenerateSession()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	if secret, err := authServer.AuthenticateUser(r.FormValue("state"), r.FormValue("code")); err == nil {
		storeCookie(w, secret)
		http.Redirect(w,r,"http://localhost/", http.StatusTemporaryRedirect)
		return
	} else {
		fmt.Printf("Authentication failed: %s\n", err.Error())
		http.Redirect(w,r,"/", http.StatusUnauthorized)
		return
	}
}

func storeCookie(w http.ResponseWriter, state string) {
	var expiration = time.Now().Add(12 * time.Hour)

	cookie := http.Cookie{Name: COOKIE_NAME, Value: state, Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
}

func main(){
	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterTorrentServer(s, &server.TorrentServer{IsDryRun: *dryRun})

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
