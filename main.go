package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// This struct is used to parse the user info from Google's API
type GoogleUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

var (
	// oauth2Config is the configuration for our OAuth2 client.
	oauth2Config *oauth2.Config

	// A random string generated for each login flow to protect against CSRF attacks.
	oauthStateString string
)

func main() {
	// 1. INITIALIZE OAUTH2 CONFIG
	// It's best practice to load these from environment variables.
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("FATAL: GOOGLE_CLIENT_ID and GOOGLE_CLIENT_SECRET environment variables must be set.")
	}

	oauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/callback", // Must match the one in Google Cloud Console
		Scopes: []string{ // The permissions we are requesting
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	// 2. SETUP HTTP ROUTES
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	fmt.Println("✅ Server started. Go to http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleRoot is the handler for the home page.
// It simply shows a welcome message and a link to login.
func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Welcome!</h1><p>This is an example OAuth2 proxy.</p><a href="/login">Log in with Google to continue</a>`)
}

// handleLogin starts the OAuth2 flow by redirecting the user to Google.
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Generate a random state string for CSRF protection.
	b := make([]byte, 16)
	rand.Read(b)
	oauthStateString = base64.URLEncoding.EncodeToString(b)

	// Redirect user to Google's consent page.
	url := oauth2Config.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// handleCallback is where Google redirects the user back to after authentication.
func handleCallback(w http.ResponseWriter, r *http.Request) {
	// 1. VALIDATE STATE
	// Ensure the state returned by Google matches the one we set.
	if r.FormValue("state") != oauthStateString {
		log.Printf("invalid oauth google state, got: %s, want: %s", r.FormValue("state"), oauthStateString)
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	// 2. EXCHANGE CODE FOR TOKEN
	// The code is used to get the access token.
	code := r.FormValue("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// The token contains the access token as well as an expiry time and a refresh token.
	// In a real app, you would want to store this token in a session or database.

	// 3. USE TOKEN TO GET USER INFO
	// The access token allows us to make requests to Google's APIs on behalf of the user.
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("http.Get() failed with '%s'\n", err)
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("io.ReadAll() failed with '%s'\n", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// 4. PARSE USER INFO AND DISPLAY IT (PROXY LOGIC GOES HERE)
	var userInfo GoogleUserInfo
	json.Unmarshal(contents, &userInfo)

	// For this simple example, we just display the user's info.
	// In a real proxy, you would now:
	//   a. Create a session for the user (e.g., set a secure cookie).
	//   b. Forward the original request to your backend service,
	//      adding the user's identity (like email) in a header.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Login Successful!</h1>")
	fmt.Fprintf(w, "<p>Hello, %s!</p>", userInfo.Name)
	fmt.Fprintf(w, "<p>Your email is: %s</p>", userInfo.Email)
	fmt.Fprintf(w, "<p><b>Next Step:</b> Now you would forward this request to your actual application.</p>")

}