package user

import "net/http"

func (r Router) Get(w http.ResponseWriter, rq *http.Request) {
	w.Header().Add("Content-Type", "application/json")

}
