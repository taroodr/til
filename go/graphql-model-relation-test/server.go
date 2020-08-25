package main

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/taroodr/til/go/graphql-model-relation-test/model"

	"github.com/taroodr/til/go/graphql-model-relation-test/domain"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/taroodr/til/go/graphql-model-relation-test/graph"
	"github.com/taroodr/til/go/graphql-model-relation-test/graph/generated"
)

const dataSource = "user:admin@tcp(127.0.0.1:3306)/todo_db?charset=utf8&parseTime=True&loc=Local"
const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("mysql", dataSource)
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"innoDB", "UTF8"}}
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	defer dbmap.Db.Close()

	dbmap.AddTableWithName(model.User{}, "user").SetKeys(false, "id")

	todoRepository := domain.NewTodoRepository(dbmap)
	userRepository := domain.NewUserRepository(dbmap)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{TodoRepository: todoRepository, UserRepository: userRepository}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

