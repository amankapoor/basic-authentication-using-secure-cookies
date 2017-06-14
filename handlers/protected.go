package handlers

import (
	"fmt"
	"net/http"

	"github.com/amankapoor/basic-authentication-using-secure-cookies/common"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))

func Protected(w http.ResponseWriter, r *http.Request) {

	if cookie, err := r.Cookie("cookie-name"); err == nil {
		value := make(map[string]string)

		if err = sc.Decode("cookie-name", cookie.Value, &value); err == nil {
			w.Write([]byte(fmt.Sprintf("The value of name is %v \n", value["name"])))
		}
	}
	common.RenderTemplate(w, "templates/protected.html", nil)
}
