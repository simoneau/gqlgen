package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/example/selection"
	"github.com/99designs/gqlgen/handler"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.GraphQL(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
