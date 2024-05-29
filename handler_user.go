package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/G0SU19O2/rssagg/internal/auth"
	"github.com/G0SU19O2/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) userHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user, err := apiConfig.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name:     params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
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
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}