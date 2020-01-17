package server

import (
	"fmt"
	"github.com/Tchayo/gql-tuts.git/internal/gql/mutations"
	"github.com/Tchayo/gql-tuts.git/internal/gql/queries"
	"github.com/Tchayo/gql-tuts.git/internal/handlers"
	"github.com/Tchayo/gql-tuts.git/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"

	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

var Host, Port, DbHost, DbPort, DbUser, DbName, DbPassword string

func init() {
	Host = utils.MustGet("SERVER_HOST")
	Port = utils.MustGet("SERVER_PORT")
	DbHost = utils.MustGet("DB_HOST")
	DbPort = utils.MustGet("DB_PORT")
	DbUser = utils.MustGet("DB_USER")
	DbName = utils.MustGet("DB_NAME")
	DbPassword = utils.MustGet("DB_NAME")
}

func initializeApi() (*gorm.DB, error) {
	var dbErr error
	Dbdriver := "postgres"

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("Connected to the %s database\n", Dbdriver)
		//db.Debug().AutoMigrate(&models.Author{}, &models.Message{}) //database migration
	}

	// Create our root query for graphql
	rootQuery := queries.NewRoot(db)
	rootMutation := mutations.NewRootMutation(db)
	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery.Query,
			Mutation: rootMutation,
		},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := handlers.Server{
		GqlSchema: &sc,
	}

	r := gin.Default()
	r.GET("/ping", handlers.Ping())
	r.GET("/graph-get", s.GraphqlHandler())
	r.POST("/graph", s.GraphqlHandler())
	r.POST("/new-message", s.GraphqlHandler())

	log.Println(DbHost + " Running @ http://" + DbHost + ":" + DbPort)
	log.Fatalln(r.Run(Host + ":" + Port))

	return db, dbErr

}

// Run : run server
func Run() {
	//r := gin.Default()

	// Handlers
	// Simple keep-alive/ping handler
	//r.GET("/ping", handlers.Ping())
	//log.Println(host + "Running @ http://" + ":" + port)
	//log.Fatalln(r.Run(host + ":" + port))

	db, err := initializeApi()
	if err != nil {
		log.Fatalf("Database Error: %v", err)
	}

	defer db.Close()

	// Listen on port 4000 and if there's an error log it and exit
	//log.Fatal(http.ListenAndServe(":4000", router))

	// GraphQL
	// Schema

}
