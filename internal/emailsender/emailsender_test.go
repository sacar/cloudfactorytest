package emailsender

import "testing"

func TestCreateUserList(t *testing.T) {
	data := [][]string{
		{"1", "sakar", "ghimire", "random2@test.com"},
		{"2", "saransh", "ghimire", "random@test.com"},
		{"3", "sanjay", "bhatta", "test@test.com"}}
	userList := CreateUserList(data)
	if len(userList) != 3 {
		t.Errorf(`expected 3 users, but got %v users `, len(userList))
	}

	data = [][]string{
		{"1", "sakar", "ghimire", "random2@test.com"},
		{"2", "saransh", "ghimire", "random@test.com"},
		{"3", "sanjay", "bhatta"}}
	userList = CreateUserList(data)
	if len(userList) != 2 {
		t.Errorf(`expected 2 users, but got %v users `, len(userList))
	}

}

func TestValidEmail(t *testing.T) {
	email := "78#ar.@test.com"
	if validEmail(email) {
		t.Errorf("Treating invalid email as valid email")
	}

	email = "random@test.com"
	if !validEmail(email) {
		t.Errorf("Treating valid email as invalid email")
	}
}
