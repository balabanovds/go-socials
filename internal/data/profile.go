package data

import (
	"time"
)

type Gender int

const (
	Male Gender = iota + 1
	Female
)

type Profile struct {
	Id        int64
	Uuid      string
	UserId    int64
	FirstName string
	LastName  string
	Age       int
	Gender    Gender
	Interests string
	City      string
	CreatedAT time.Time
}

func (user *User) CreateProfile(p Profile) (profile Profile, err error) {
	stmt, err := db.Prepare("INSERT INTO profiles (uuid, user_id, first_name, last_name, age, gender, interests, city) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(createUUID(), user.Id, p.FirstName, p.LastName, p.Age, p.Gender, p.Interests, p.City)
	if err != nil {
		return

	}
	profile, err = user.GetProfile()
	return
}

func (user *User) GetProfile() (profile Profile, err error) {
	profile, err = ProfileByUserId(user.Id)
	return
}

func (user *User) UpdateProfile(p Profile) (err error) {
	stmt, err := db.Prepare("UPDATE profiles SET first_name = $2, last_name = $3, age = $4, gender = $5, interests = $6, city = $7 WHERE user_id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, p.FirstName, p.LastName, p.Age, p.Gender, p.Interests, p.City)
	return
}

func (user *User) DeleteProfile() (err error) {
	stmt, err := db.Prepare("DELETE FROM profiles WHERE user_id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

func Profiles() (profiles []Profile, err error) {
	rows, err := db.Query("SELECT id, uuid, first_name, last_name, age, gender, interests, city, created_at FROM profiles")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := Profile{}
		if err = rows.Scan(&p.Id, &p.Uuid, &p.FirstName, &p.LastName, &p.Age, &p.Gender, &p.Interests, &p.City, &p.CreatedAT); err != nil {
			return
		}
		profiles = append(profiles, p)
	}

	return
}

func ProfileByUserId(userId int64) (p Profile, err error) {
	p = Profile{}
	err = db.QueryRow("SELECT id, uuid, user_id, first_name, last_name, age, gender, interests, city FROM profiles WHERE user_id = ?", userId).
		Scan(&p.Id, &p.Uuid, &p.UserId, &p.FirstName, &p.LastName, &p.Age, &p.Gender, &p.Interests, &p.City)

	return
}

func ProfileDeleteAll() (err error) {
	_, err = db.Exec("DELETE FROM profiles")
	return
}
