package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net"
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
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


func main() {
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

func (s *server) AddUser(ctx context.Context, request *proto.AddUserRequest) (*proto.AddUserResponse, error) {
	id, age, name := request.GetId(), request.GetAge(), request.GetName();

	insertStatement := `INSERT INTO users (id, age, name) VALUES ($1, $2, $3)`
	result, err := db.Exec(insertStatement, id, age, name)

	if err != nil {
		fmt.Println(err);
		return &proto.AddUserResponse{Result: "Error! "}, nil
	} else {
		fmt.Println(result);
		return &proto.AddUserResponse{Result: "Success!"}, nil
	}
}


func (s *server) GetAllUsers(ctx context.Context, request *proto.GetAllUsersRequest) (*proto.GetAllUsersResponse, error) {

	result, err := db.Query(`select * from users`)

	if err != nil {
		panic(err)
	} else {
		defer result.Close()

		var users []*proto.User  // an empty list

		for result.Next() {
			var id int
			var age int
			var name string
			err = result.Scan(&id, &age, &name)

			users = append(users,&proto.User{Id: int64(id), Age: int64(age), Name: name})
			fmt.Printf("%v%s", age, name)
		}
		
		return &proto.GetAllUsersResponse{Users: users}, nil
	}
}

func (s *server) GetSingleUser(ctx context.Context, request *proto.GetSingleUserRequest) (*proto.User, error) {
	
	
	statement := `select * from users where id = $1`
	row := db.QueryRow(statement, request.GetId())
	
	var id int
	var age int
	var name string
	row.Scan(&id, &age, &name)
	return &proto.User{Id: int64(id), Age : int64(age), Name : name}, nil



}	



// keeping for reference
func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
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



