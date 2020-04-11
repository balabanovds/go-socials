package web

import (
	"balabanovds/go-social/internal/data"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "pub.navbar", "login")
	_ = t.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, nil, "login.layout", "pub.navbar", "signup")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errMessage(w, r, err.Error())
		return
	}
	user := data.User{
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		errMessage(w, r, "Can not create user: "+err.Error())
		return
	}
	http.Redirect(w, r, "/login", 302)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		errMessage(w, r, "Can not find user: "+err.Error())
		return
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			errMessage(w, r, "Can not create session: "+err.Error())
			return
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		session := data.Session{Uuid: cookie.Value}
		_ = session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
