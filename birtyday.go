package birthday_greeting_kata

type Birthday struct {
	Year  int
	Month int
	Day   int
}

func NewBirthday(year int, monty int, day int) Birthday {
	return Birthday{Year: year, Month: monty, Day: day}
}
