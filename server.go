package main

import (
	"context"
	"encoding/json"
	fields "github.com/gbaptista/requested-fields"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	graphql "github.com/graph-gophers/graphql-go"
	"net/http"
)

const schema = `
	schema {
		query: Query
	}

	type Query {
		user: User
	}

	type User {
		name: String
		address: Address
	}

	type Address {
		city: String
		street: String
	}
`

var parsedSchema = graphql.MustParseSchema(schema, &Query{})

type GraphqlParams struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func graphqlRequestHandler(writer http.ResponseWriter, request *http.Request) {
	var params GraphqlParams

	json.NewDecoder(request.Body).Decode(&params)

	ctx := context.WithValue(request.Context(),
		fields.ContextKey, fields.BuildTree(params.Query))

	response := parsedSchema.Exec(
		ctx, params.Query, params.OperationName, params.Variables)

	jsonResponse, _ := json.Marshal(response)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonResponse)
}

func main() {
	router := chi.NewRouter()

	cors := cors.New(cors.Options{AllowedOrigins: []string{"*"}})

	router.Use(
		cors.Handler, middleware.Logger,
		render.SetContentType(render.ContentTypeJSON),
	)

	router.Post("/graphql", graphqlRequestHandler)

	http.ListenAndServe(":3000", router)
}
