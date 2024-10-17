package handlers

import (
	"net/http"
	"user-service/rest/utils"
)

func (handlers *Handlers) Hello(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, http.StatusOK, map[string]any{
		"success": true,
	})
}
