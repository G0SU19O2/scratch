package main

import (
	"fmt"
	"net/http"

	"github.com/G0SU19O2/scratch/internal/auth"
	"github.com/G0SU19O2/scratch/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Invalid API key: %s", err))
			return
		}
		user, err := apiConfig.db.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("User not found: %s", err))
			return
		}
		handler(w, r, user)
	}
}
