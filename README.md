This project contains two projects: a simple rest api, that shows the covid related data and a go program to send email from mailing list. The docs folder contains the architecture diagram for the authentication system.

# Rest API
Run cmd/covidreport/main.go for starting the web server.
The web server contains two api: a get api /country/{name} to get the countries covid data and a post /compare which takes list of countries and shows which countries has the maximum and minimmum cases of covid and deaths

# Email sender
Build the cmd/emailreader/main.go for starting the web server. While running the binary file, you can provide specify which file to use. A file is already present and is used by default.
