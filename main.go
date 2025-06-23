package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const VERSION = "0.1.0"

var (
	showVersion  = flag.Bool("version", false, "print version string")
	clientID     = flag.String("client-id", "", "client id")
	clientSecret = flag.String("client-secret", "", "client secret")
	cookieSecret = flag.String("cookie-secret", "", "cookie secret")
)

func main() {
	flag.Parse()

	if *clientID == "" {
		*clientID = os.Getenv("google_auth_client_id")
	}
	if *clientSecret == "" {
		*clientSecret = os.Getenv("google_auth_client_secret")
	}
	if *cookieSecret == "" {
		*cookieSecret = os.Getenv("google_auth_cookie_secret")
	}
	if *showVersion {
		fmt.Printf("google_auth_proxy v%s\n", VERSION)
		return
	}

	if *clientID == "" {
		log.Fatal("--client-id is missing")
	}
	if *clientSecret == "" {
		log.Fatal("--client-secret is missing")
	}
	if *cookieSecret == "" {
		log.Fatal("--cookie-secret is missing")
	}

	fmt.Printf("client-id: %s\n", *clientID)
	fmt.Printf("client-secret: %s\n", *clientSecret)
	fmt.Printf("cookie-secret: %s\n", *cookieSecret)
}
