package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sacar/covidreport/internal/emailsender"
)

func main() {
	fileName := readArguments()
	data, err := readFile(fileName)
	if err != nil {
		log.Print(err.Error())
		return
	}
	userList := emailsender.CreateUserList(data)
	emailsender.SendEmail(userList)
}

func readArguments() string {
	csvFile := flag.String("csvfile", "emailList.csv", "A csv file in format Id, First Name, Last Name, Email")
	flag.Parse()
	return *csvFile
}

func readFile(fileName string) ([][]string, error) {

	// open file if exists
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read csv file and check for data
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	if l := len(data); l == 0 {
		return nil, fmt.Errorf("no data in csv file")
	}

	return data, nil

}
