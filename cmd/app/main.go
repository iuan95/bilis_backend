package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	h "github.com/iuan95/bilis_backend/internal/delivery/http"
	"github.com/iuan95/bilis_backend/internal/usecase"
	"github.com/iuan95/bilis_backend/pkg/postgres"
)

//  migrate create -ext sql -dir migrations -seq add_mood_to_users
// / migrations <- папка где находится миграция, add_mood_to_users <- название миграции

// migrate -database "postgres://postgres:admin@localhost:5432/test?sslmode=disable" -path migrations up N
// migrate -database "postgres://postgres:admin@localhost:5432/test?sslmode=disable" -path migrations down N
// migrate -database "postgres://postgres:admin@localhost:5432/test?sslmode=disable" -path migrations goto N (поменять версию на N)
// migrate -database "postgres://postgres:admin@localhost:5432/test?sslmode=disable" -path migrations drop (удалить все внутри бд)
// migrate -database "postgres://postgres:admin@localhost:5432/test?sslmode=disable" -path migrations version (текущая версия)


func init(){
    err := godotenv.Load(".env")
    if err != nil {
      log.Fatal("Error loading .env file")
    }
}

func main(){

    db_postgres:=postgres.New()
    defer db_postgres.Close()

    postRepo:=postgres.NewPostRepository(db_postgres) 
    postService:=usecase.NewPostService(postRepo)
    postHandler:= h.NewPostHandler(postService)


    userRepo:=postgres.NewUserRepository(db_postgres)
    userService:=usecase.NewUserService(userRepo)
    userHandler:=h.NewUserHandler(userService)



    router := mux.NewRouter()

    router.HandleFunc("/post", postHandler.CreatePost).Methods("POST")
    router.HandleFunc("/post/{id}", postHandler.GetById).Methods("GET")
    router.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
    router.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")


    log.Println("Server is running on port 8080...")

    log.Fatal(http.ListenAndServe(":8080", router))
   
}