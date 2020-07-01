package handlers

import (
	"net/http"
	"webCli/dirBox"
)

func Index(w http.ResponseWriter, r *http.Request) {
	box := dirBox.NewBox()

	index, err := box.Find("html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't open file: " + err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html")
	w.Write(index)
}
