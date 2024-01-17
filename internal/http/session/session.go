package session

import "github.com/gorilla/sessions"

type Storage struct {
	store *sessions.CookieStore
}

func New(secretKey string) *Storage {
	return &Storage{
		store: sessions.NewCookieStore([]byte(secretKey)),
	}
}
