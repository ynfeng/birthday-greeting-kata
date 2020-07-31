package birthday_greeting_kata

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type FileAddressBook struct {
	filePath string
}

func (addressBook FileAddressBook) friends() []Friend {
	fi, _ := os.Open(addressBook.filePath)
	defer fi.Close()

	result := make([]Friend, 0)

	br := bufio.NewReader(fi)
	lineNo := 0
	for {
		readBytes, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if lineNo == 0 {
			lineNo++
			continue
		}
		lineNo++
		line := string(readBytes)
		lineItems := strings.Split(line, ",")
		birthdayItems := strings.Split(lineItems[2], "/")
		year, _ := strconv.Atoi(strings.Trim(birthdayItems[0], " "))
		month, _ := strconv.Atoi(strings.Trim(birthdayItems[1], " "))
		day, _ := strconv.Atoi(strings.Trim(birthdayItems[2], " "))
		birthday := NewBirthday(year, month, day)
		lastName := strings.Trim(lineItems[0], " ")
		firstName := strings.Trim(lineItems[1], " ")
		email := strings.Trim(lineItems[3], " ")
		friend := NewFriend(lastName, firstName, birthday, email)
		result = append(result, friend)
	}
	return result
}

func NewFileAddressBook(filePath string) FileAddressBook {
	return FileAddressBook{filePath: filePath}
}
