# what is this project?
this is a todo API that can store todos in database, edit them and delete them when needed.
---
## how to run?
makes sure you're in the list-manager foulder, then to install the requierments, run this command:
`go mod download`
after downloading the files, first go to the connect file and change the database data to your own username and password, then, to run the server, run this command:
`go run main.go`
---
## notes:
the main files contains the endpoints, the handler definitions are in the handler file, the connect file is for connecting to database and the todo file is the main struct for the columns.
