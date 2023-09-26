package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mmvergara/go-simple-api/model"
	"github.com/mmvergara/go-simple-api/repository/post"
)

type Post struct {
	Repo *post.RedisRepo
}


func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		AuthorID uuid.UUID `json:"author_id"`
		PostTitle string `json:"post_title"`
		PostDescription string `json:"post_description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	post := model.Post{
		PostID:  uuid.Must(uuid.NewRandom()),
		AuthorID: body.AuthorID,
		PostTitle: body.PostTitle,
		PostDescription: body.PostDescription,
		CreatedAt: &now,
	}

	if err := p.Repo.Insert(r.Context(), post); err != nil {
		fmt.Println("Error inserting post", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res , err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error marshalling post", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

func (p *Post) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders")
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (p *Post) UpdateByID(w http.ResponseWriter, r *http.Request) {

}

func (p *Post) DeleteByID(w http.ResponseWriter, r *http.Request) {

}