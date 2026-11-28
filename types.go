package main

import (
	"time"

	"github.com/google/uuid"
)

type returnErr struct {
	Error string `json:"error"`
}

type returnCleaned struct {
	Cleaned_body string `json:"cleaned_body"`
}

type reqEmailPass struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

type reqCreateChirp struct {
	Body   string    `json:"body"`
	UserId uuid.UUID `json:"user_id"`
}
