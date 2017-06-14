package handlers

import (
	"net/http"

	"github.com/gorilla/securecookie"

	"log"
)

var sc = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "aman.sk95@gmail.com" && password != "" {

			value := map[string]string{
				"foo": "bar",
			}

			if encoded, err := sc.Encode("cookie-name", value); err == nil {
				cookie := &http.Cookie{
					Name:  "cookie-name",
					Value: encoded,
					Path:  "/",
				}

				http.SetCookie(w, cookie)

				log.Printf("&cookie gives: %v", &cookie)
				log.Printf("*cookie gives: %v", *cookie)
				log.Printf("cookie gives: %v", cookie)
			}

			http.Redirect(w, r, "http://localhost:8080/protected", http.StatusSeeOther)
		}
	}
}
