package birthday_greeting_kata

type Friend struct {
	LastName  string
	FirstName string
	Birthday  Birthday
	Email     string
}

func NewFriend(lastName string, firstName string, birthday Birthday, email string) Friend {
	return Friend{LastName: lastName, FirstName: firstName, Birthday: birthday, Email: email}
}
