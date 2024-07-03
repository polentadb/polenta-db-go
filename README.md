# Polenta DB

Polenta DB is an in-memory database. 

This version is implemented in GoLang, and uses TCP as communication protocol between clients and server.

Either Java driver or JDBC driver can be used by clients, avoiding them to have to handle TCP details.

# How to run and test

1. Clone polenta-db-go
2. Run ```go run main.go statement.go```
3. Clone polenta-db-go-test
4. Run ```go run main.go```