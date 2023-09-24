package model

import "github.com/google/uuid"

type Post struct {
	PostID uint64
	CustomerID uuid.UUID
}

