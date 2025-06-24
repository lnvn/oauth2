package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

const VERSION = "0.1.0"

var (
	showVersion  = flag.Bool("version", false, "print version string")
	clientID     = flag.String("client-id", "", "client id")
	clientSecret = flag.String("client-secret", "", "client secret")
	cookieSecret = flag.String("cookie-secret", "", "cookie secret")
	redirectUrl  = flag.String("redirect-url", "", "the OAuth Redirect URL")
	upstreams    = StringArray{}
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

	var upstreamUrls []*url.URL
	for _, u := range upstreams {
		upstreamUrl, err := url.Parse(u)
		if err != nil {
			log.Fatalf("error parsing --upstream %s", err.Error())
		}
		upstreamUrls = append(upstreamUrls, upstreamUrl)
	}
	// redirectUrl, err := url.Parse(*redirectUrl)
	// if err != nil {
	// 	log.Fatalf("error parsing --redirect-url %s", err.Error())
	// }

	fmt.Printf("client-id: %s\n", *clientID)
	fmt.Printf("client-secret: %s\n", *clientSecret)
	fmt.Printf("cookie-secret: %s\n", *cookieSecret)
}
