package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/iuan95/bilis_backend/internal/entity"
	"github.com/iuan95/bilis_backend/internal/usecase"
)


type PostHandler struct{
	service *usecase.PostService
}

func NewPostHandler(service *usecase.PostService) *PostHandler {
    return &PostHandler{service: service}
}

func (h *PostHandler) GetById(w http.ResponseWriter, r *http.Request){
    params:= mux.Vars(r)
    id,err:=strconv.Atoi(params["id"])
    if err !=nil{
        http.Error(w, "Invalid post id", http.StatusBadRequest)
        return
    }
    post,err := h.service.GetPostById(id);
    if err != nil {
        http.Error(w, "Failed to get post by id", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
    var post entity.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.service.CreatePost(&post); err != nil {
        http.Error(w, "Failed to create post", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(post)
}
