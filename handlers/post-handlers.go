package handlers

import (
	"fmt"
	"net/http"

	"github.com/mmvergara/go-simple-api/repository/post"
)

type Post struct {
	Repo *post.RedisRepo
}


func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating an order")
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