package birthday_greeting_kata

type Birthday struct {
	Year  int
	Month int
	Day   int
}

func (birthday Birthday) isSame(other Birthday) bool {
	if birthday.Day == other.Day &&
		birthday.Month == other.Month &&
		birthday.Year == other.Year {
		return true
	}
	return false
}

func NewBirthday(year int, monty int, day int) Birthday {
	return Birthday{Year: year, Month: monty, Day: day}
}
