package handlers

import (
	"net/http"

	"github.com/amankapoor/basic-authentication-using-secure-cookies/common"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("Aman")
	if err == http.ErrNoCookie {
		w.Write([]byte("You are not registered."))
	}
	cookie = &http.Cookie{
		Name:   "Aman",
		Value:  "Logged Out",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	common.RenderTemplate(w, "templates/logout.html", nil)
}
