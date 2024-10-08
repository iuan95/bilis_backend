package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iuan95/bilis_backend/internal/entity"
	"github.com/iuan95/bilis_backend/internal/usecase"
)

type UserHandler struct{
	service *usecase.UserService
}

func NewUserHandler(service *usecase.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user entity.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.service.Create(context.Background(),&user); err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)
    id,err:=strconv.Atoi(params["id"])
    if err !=nil{
        http.Error(w, "Invalid user id", http.StatusBadRequest)
        return
    }
    user,err := h.service.GetById(context.Background(),id);
    if err != nil {
        http.Error(w, "Failed to get user by id", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

