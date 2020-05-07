package main

import (
	"database/sql"
	"fmt"
	"net"
	"context"
	"github.com/pkg/errors"

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

func (s *server) AddUser(ctx context.Context, request *proto.AddUserRequest) (*proto.User, error) {
	age, email, uid := request.GetAge(), request.GetEmail(), request.GetUid();


	stmt, err := db.Prepare(`INSERT INTO users (age, email, uid) VALUES ($1, $2, $3) returning id`)
	if (err != nil){
		return nil, errors.Wrap(err, "Backend: Unable to prepare sql statement for inserting user")
	}
	var id int
	defer stmt.Close()

	err = stmt.QueryRow(age, email, uid).Scan(&id)
	if (err != nil){
		return nil, errors.Wrap(err, "Backend: Unable to create new user")
	} else {
		return &proto.User{Id: int64(id), Age : int64(age), Email : email, Uid: uid}, nil
	}
}


func (s *server) GetAllUsers(ctx context.Context, request *proto.GetAllUsersRequest) (*proto.GetAllUsersResponse, error) {

	result, err := db.Query(`select * from users`)

	if err != nil {
		return nil, errors.Wrap(err, "Backend: Unable to retrieve users")
	} else {
		defer result.Close()

		var users []*proto.User  // an empty list

		for result.Next() {
			var id int
			var uid string
			var age int
			var email string
			err = result.Scan(&id, &age, &email, &uid)

			users = append(users,&proto.User{Id: int64(id), Age: int64(age), Email : email, Uid: uid})
		}
		
		return &proto.GetAllUsersResponse{Users: users}, nil
	}
}

func (s *server) GetSingleUser(ctx context.Context, request *proto.GetSingleUserRequest) (*proto.User, error) {
	
	
	statement := `select * from users where uid = $1`
	row := db.QueryRow(statement, request.GetUid())
	
	var id int
	var uid string
	var age int
	var email string
	// row.Scan(&id, &age, &email, &uid)

	switch err := row.Scan(&id, &age, &email, &uid); err {
	case sql.ErrNoRows:
	  return nil, errors.Wrap(err, "Backend: User with uid " + request.GetUid() + " does not exist")
	case nil:
		return &proto.User{Id: int64(id), Age : int64(age), Email : email, Uid: uid}, nil
	default:
		return nil, errors.Wrap(err, "Backend: Error retrieving user with uid " + request.GetUid())
	}
}	



