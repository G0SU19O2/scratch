package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/G0SU19O2/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	feed, err := apiConfig.db.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:       uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name:     params.Name,
		Url:      params.Url,
		UserID:   user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiConfig *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.db.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}