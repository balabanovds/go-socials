package data

import (
	"time"
)

type Session struct {
	Id        int64
	Uuid      string
	Email     string
	UserId    int64
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	stmt, err := db.Prepare("INSERT INTO sessions (uuid, email, user_id, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(createUUID(), user.Email, user.Id, time.Now())
	session, err = user.GetSession()
	return
}

func (user *User) GetSession() (session Session, err error) {
	session = Session{}
	err = db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	return
}

func (session *Session) Check() (valid bool, err error) {
	err = db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteByUUID() (err error) {
	stmt, err := db.Prepare("DELETE FROM sessions WHERE uuid = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)

	return
}

func (session *Session) GetUser() (user User, err error) {
	user = User{}
	err = db.QueryRow("SELECT id, uuid, email, created_at FROM users WHERE id = ?", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Email, &user.CreatedAt)
	return
}

func SessionDeleteAll() (err error) {
	_, err = db.Exec("DELETE FROM sessions")
	return
}
