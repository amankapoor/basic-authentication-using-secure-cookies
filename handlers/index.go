package handlers

import (
	"net/http"

	"github.com/amankapoor/basic-authentication-using-secure-cookies/common"
)

func Index(w http.ResponseWriter, r *http.Request) {
	common.RenderTemplate(w, "templates/index.html", nil)
	common.RenderTemplate(w, "templates/login.html", nil)
}
