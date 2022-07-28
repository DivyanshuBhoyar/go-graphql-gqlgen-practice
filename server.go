package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DivyanshuBhoyar/gqlgen-prac/graph"
	"github.com/DivyanshuBhoyar/gqlgen-prac/graph/generated"
	"github.com/DivyanshuBhoyar/gqlgen-prac/internal/auth"
	_ "github.com/DivyanshuBhoyar/gqlgen-prac/internal/auth"
	database "github.com/DivyanshuBhoyar/gqlgen-prac/internal/pkg/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}