package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthCredentials struct {
	ClientId     string `json:"id"`
	ClientSecret string `json:"secret"`
}

type User struct {
	Email      string `json:"email"`
	FamilyName string `json:"family_name"`
	Gender     string `json:"gender"`
	GivenName  string `json:"given_name"`
	Name       string `json:"name"`
	Picture    string `json:"picture"`
	Profile    string `json:"profile"`
	Sub        string `json:"sub"`
}

const credentialsFile = "./secrets/oauth-creds.json"

var conf *oauth2.Config

/* Public */
func AuthHandler(c *gin.Context) {
	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	if retrievedState != c.Query("state") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
		return
	}

	tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	userJSON, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer userJSON.Body.Close()
	data, err := ioutil.ReadAll(userJSON.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := User{}
	if err = json.Unmarshal(data, &user); err != nil {
		log.Println(err)
		return
	}

	session.Set("user-id", user.Email)
	err = session.Save()
	if err != nil {
		log.Println(err)
		return
	}

	// resp, err := http.Get("http://localhost:9000/login/" + user.Email)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	log.Println("User data: ", string(data))
	c.HTML(http.StatusOK, "index.html", nil)
}

/* Package */
func init() {
	creds := loadCredentialsFile(credentialsFile)
	conf = &oauth2.Config{
		ClientID:     creds.ClientId,
		ClientSecret: creds.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		RedirectURL: "http://localhost:8080/auth",
		Endpoint:    google.Endpoint,
	}
}

func loadCredentialsFile(filePath string) OauthCredentials {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var creds OauthCredentials
	json.Unmarshal(file, &creds)
	return creds
}
