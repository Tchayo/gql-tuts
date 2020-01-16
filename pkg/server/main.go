package server

import (
	"encoding/json"
	"fmt"
	//"github.com/Tchayo/gql-tuts.git/internal/handlers"
	//"github.com/Tchayo/gql-tuts.git/pkg/utils"
	"log"

	//"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

//var host, port string
//
//func init() {
//	host = utils.MustGet("GQL_SERVER_HOST")
//	port = utils.MustGet("GQL_SERVER_PORT")
//}

// Run : run server
func Run() {
	//r := gin.Default()

	// Handlers
	// Simple keep-alive/ping handler
	//r.GET("/ping", handlers.Ping())
	//log.Println(host + "Running @ http://" + ":" + port)
	//log.Fatalln(r.Run(host + ":" + port))

	// GraphQL
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type:              graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, erro: %v", err)
	}

	// Query
	query := `{
		hello
	}`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJson, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJson)

}
