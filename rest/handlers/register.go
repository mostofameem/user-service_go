package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"user-service/rest/utils"
	"user-service/route"
)

type userReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Dob      string `json:"dob"`
	Type     string `json:"type"`
}

func (handlers *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var req userReq
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		slog.Error(err.Error())
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", utils.ParseValidationErrors(err))
		return
	}

	err := handlers.routeSvc.Register(r.Context(), route.User{
		Name:  req.Username,
		Email: req.Email,
		Dob:   req.Dob,
		Type:  req.Type,
	})
	if err != nil {
		slog.Error(err.Error())
		utils.SendError(w, http.StatusInternalServerError, "Internal server error", utils.ParseValidationErrors(err))
		return
	}

	utils.SendData(w, "Success")
}
