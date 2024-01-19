package jwttoken

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	cookieName = "token"
)

type (
	HttpUserIDInHandler  func(string, http.ResponseWriter, *http.Request)
	HttpUserIDOutHandler func(http.ResponseWriter, *http.Request) (string, error)
)

type JWT struct {
	secretKey []byte
}

func New(key string) *JWT {
	return &JWT{
		secretKey: []byte(key),
	}
}

func (j JWT) Add(next HttpUserIDOutHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := next(w, r)
		if err != nil {
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
		tokenString, err := token.SignedString(j.secretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:     cookieName,
			Value:    tokenString,
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}

		http.SetCookie(w, &cookie)
	}
}

func (j JWT) Check(next HttpUserIDInHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token method")
			}

			return j.secretKey, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		val, ok := claims["exp"]
		if !ok && float64(time.Now().Unix()) > val.(float64) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		val, ok = claims["sub"]
		if !ok {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next(val.(string), w, r)
	}
}
