package handlers

import (
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./common/html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't open file: " + err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html")
	w.Write(file)
}
