package main

import (
	"encoding/json"
	"errors"
	"github.com/ghthor/gowol"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func wakeScott(w http.ResponseWriter, r *http.Request) {
	err := wol.MagicWake("30:5a:3a:05:da:20", "255.255.255.255")
	if err != nil {
		log.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request")

	t, err := template.ParseGlob("/home/pi/go/src/github.com/Laremere/home/templates/*")
	if err != nil {
		serveErrorPage(w, "Failed to parse templates", err)
		return
	}

	data := struct {
		Auth authInfo
	}{}

	data.Auth, err = checkUserLogin(r)
	if err != nil {
		serveErrorPage(w, "Failed user auth", err)
		return
	}

	t.ExecuteTemplate(w, "index.html", data)
}

func serveErrorPage(w http.ResponseWriter, part string, err error) {
	log.Println(part, err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

type authInfo struct {
	Authorized bool
	LoggedIn   bool
	Email      string
}

func checkUserLogin(r *http.Request) (authInfo, error) {
	authTokenCookie, err := r.Cookie("token")
	if err == http.ErrNoCookie {
		return authInfo{}, nil
	}
	if err != nil {
		return authInfo{}, err
	}
	if authTokenCookie.Value == "" {
		return authInfo{}, nil
	}

	authUrl := "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=" + authTokenCookie.Value
	resp, err := http.Get(authUrl)
	if err != nil {
		return authInfo{}, err
	}

	jsonVals := struct {
		Aud   string `json:"aud"`
		Email string `json:"email"`
	}{}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return authInfo{}, err
	}

	err = json.Unmarshal(bytes, &jsonVals)
	if err != nil {
		return authInfo{}, err
	}

	if jsonVals.Aud != "141488749003-audttelm23ke99cmd1qgc4utd9hpqopu.apps.googleusercontent.com" {
		return authInfo{}, errors.New("Invalid aud used to authenticate: " + string(bytes) + "\nauthUrl: " + authUrl)
	}

	return authInfo{
		LoggedIn:   true,
		Authorized: "laremere@gmail.com" == jsonVals.Email, // lol
		Email:      jsonVals.Email,
	}, nil
}

func main() {
	log.Println("Starting server")
	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/home/pi/go/src/github.com/Laremere/home/static"))))
	http.HandleFunc("/wake/scott", wakeScott)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
