package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vokhanh12/RongchoiApp/backend/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	opt := option.WithCredentialsFile("rongchoi-e9690-firebase-adminsdk-jw6np-5bef5c0766.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error auth app : %v", err)
	}

	customToken, err := auth.CustomToken(context.Background(), "-LDLrwZnJ0kI4kvyIb2Q")
	if err != nil {
		log.Fatalf("error CustomToken app : %v", err)
	}
	log.Fatalf("token : %s", customToken)

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/ready", handlerReadiness)
	v1Router.HandleFunc("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)

}
