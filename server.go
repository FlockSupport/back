package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net"
	

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"context"
	"flock-support/back/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	

)

type server struct{}



var router *chi.Mux
var db *sql.DB

func routers() *chi.Mux {
	router.Get("/posts", AllPosts)
	return router
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "graphql_demo"
)

func init() {
	var err error
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected to db")
	}
}


// func main() {
// 	defer db.Close()
// 	routers()
// 	http.ListenAndServe(":8005", router)	
// }
func main() {
	//listener, err := net.Listen("tcp", ":4040")
	listener, err := net.Listen("tcp", ":8005")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}


func (s *server) AddUser(ctx context.Context, request *proto.AddUserRequest) (*proto.AddUserResponse, error) {
	id, quantity, name := request.GetId(), request.GetQuantity(), request.GetName();

	insertStatement := `INSERT INTO graphql (id, quantity, name) VALUES ($1, $2, $3)`
	result, err := db.Exec(insertStatement, id, quantity, name)

	if err != nil {
		fmt.Println(err);
		return &proto.AddUserResponse{Result: "Error! "}, nil
	} else {
		fmt.Println(result);
		return &proto.AddUserResponse{Result: "Success!"}, nil
	}

	
}


func AllPosts(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query(`SELECT * from graphql`)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var quantity int
		var id int
		var name string
		err = result.Scan(&id, &name, &quantity)
		fmt.Printf("%v%s", quantity, name)
		panic(err)

	}
	err = result.Err()
}


func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}

