package data

var users = []User{
	{
		Email:    "vasya@mail.ru",
		Password: "vasya_pass",
	},
	{
		Email:    "petya@mail.ru",
		Password: "petya_pass",
	},
}

var profile = Profile{
	FirstName: "Vasya",
	LastName:  "Pupkin",
	Age:       30,
	Gender:    Male,
	Interests: "ps4 fishing",
	City:      "SPb",
}

var profile2 = Profile{
	FirstName: "Petya",
}

func setup() {
	UserDeleteAll()
	SessionDeleteAll()
	ProfileDeleteAll()
}
