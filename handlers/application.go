package handlers

import (
	"net/http"

	"gitlab.com/jorgechavezrnd/go_rest/utils"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "application/index", nil)
}
