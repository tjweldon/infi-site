package web

import (
	"net/http"
	"tjweldon/infi-site/page"
)

func echoPath(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	content := page.ListFromCurrentPath(path, 16).String()
	myPage := page.NewPage(content)
	_, err := w.Write(myPage.Bytes())
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/html")
}

var EchoPath = (Controller)(echoPath)
