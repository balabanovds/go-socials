package web

import (
	"balabanovds/go-social/internal/data"
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(w http.ResponseWriter, r *http.Request, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	_, err := session(w, r)
	templates.Funcs(template.FuncMap{
		// TODO implement
		"isLoggedIn": func() bool {
			return err == nil
		},
	})
	_ = templates.ExecuteTemplate(w, "layout", data)
}

func errMessage(w http.ResponseWriter, r *http.Request, msg string) {
	http.Redirect(w, r, fmt.Sprintf("/err?msg=%s", msg), 302)
}

func session(w http.ResponseWriter, r *http.Request) (session data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		session = data.Session{
			Uuid: cookie.Value,
		}
		if ok, _ := session.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}
