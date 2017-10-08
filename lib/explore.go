package lib

import (
	"net/http"
	"io"
)

//登录
func Explore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "NanjingYouzi")
	defer func() {
		io.WriteString(w, "boom1")

	}()
	}