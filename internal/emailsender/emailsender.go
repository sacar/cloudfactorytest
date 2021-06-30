package emailsender

import (
	"fmt"
)

type users struct {
	id        string
	firstName string
	lastName  string
	email     string
}

type smtpConfig struct {
	smtpHost string
	smtpPort string
	username string
	password string
}

type email struct {
	from    string
	to      []string
	message []byte
}

func CreateUserList(data [][]string) []users {

	// extract data from each line
	var userList []users
	for _, line := range data {

		if l := len(line); l != 4 {
			continue
		}

		u := users{
			id:        line[0],
			firstName: line[1],
			lastName:  line[2],
			email:     line[3],
		}

		userList = append(userList, u)

	}

	return userList

}

func SendEmail(userList []users) {
	config := getSMTPConfing()
	server := ConfigureServer(config)

	for _, usr := range userList {
		eml := email{}
		eml.from = "test@mailtrap.io"
		eml.to = []string{usr.email}
		eml.message = []byte(fmt.Sprintf("From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: Cloudfactory test\r\n"+
			"\r\n"+
			"Hello %s,\nTest email from cloudfactory\r\n", eml.from, usr.email, usr.firstName))

		server.send(&eml)

	}
}

/* all config should be in a sepearte config file,
   but it is kept here for simplicity
*/
func getSMTPConfing() *smtpConfig {

	return &smtpConfig{
		smtpHost: "smtp.mailtrap.io",
		smtpPort: "2525",
		username: "d481b85664f042",
		password: "73af6a3ce17d58",
	}
}
