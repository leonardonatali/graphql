package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/leonardonatali/graphql/graph"
	"github.com/leonardonatali/graphql/graph/generated"
	"github.com/leonardonatali/graphql/persistence"
	"github.com/leonardonatali/graphql/pkg/generators"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := logrus.New()

	mdb := persistence.NewMemoryPersistence(generators.NewUUIDGenenerator())

	resolver := graph.NewResolver(mdb, mdb, mdb)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}
