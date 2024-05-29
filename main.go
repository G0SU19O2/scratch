package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/G0SU19O2/rssagg/internal/database"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	_ "github.com/lib/pq"
)
type apiConfig struct {
	db *database.Queries
}
func main() {
	godotenv.Load()

	dbURL := os.Getenv(("DB_URL"))
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	apiConfig := apiConfig{
		db: database.New(conn),
	}

	port := os.Getenv("PORT")
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", readinessHandler)
	v1Router.Get("/err", handleError)
	v1Router.Get("/users", apiConfig.handlerGetUser)
	v1Router.Post("/users", apiConfig.userHandler)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	srv.ListenAndServe()
}
