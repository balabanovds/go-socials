package web

import (
	"balabanovds/go-social/internal/data"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	profiles, err := data.Profiles()
	if err != nil {
		errMessage(w, r, "Can not get profiles")
		return
	}

	_, err = session(w, r)

	if err != nil {
		generateHTML(w, r, profiles, "layout", "pub.navbar", "index")
	} else {
		generateHTML(w, r, profiles, "layout", "priv.navbar", "index")
	}

}

func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, r, vals.Get("msg"), "layout", "pub.navbar", "error")
	} else {
		generateHTML(w, r, vals.Get("msg"), "layout", "priv.navbar", "error")
	}
}
