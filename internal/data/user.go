package data

import (
	"time"
)

type User struct {
	Id        int64
	Uuid      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (user *User) Create() (err error) {
	stmt, err := db.Prepare("INSERT INTO users (uuid, email, password, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(createUUID(), user.Email, Encrypt(user.Password), time.Now())

	return
}

func (user *User) Delete() (err error) {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return
	}

	err = user.DeleteProfile()
	return
}

func (user *User) Update() (err error) {
	stmt, err := db.Prepare("UPDATE users SET email = $2, password = $3 WHERE id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Email, Encrypt(user.Password))
	return
}

func Users() (users []User, err error) {
	rows, err := db.Query("SELECT id, uuid, email, created_at FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Email, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}

	return
}

func UserById(id int64) (user User, err error) {
	user = User{}
	err = db.QueryRow("SELECT id, uuid, email, password, created_at FROM users WHERE id = ?", id).
		Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = db.QueryRow("SELECT id, uuid, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = db.QueryRow("SELECT id, uuid, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserDeleteAll() (err error) {
	_, err = db.Exec("DELETE FROM users")
	return
}
