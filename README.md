[![codecov](https://codecov.io/gh/Josimar/GoLangDevBook/branch/main/graph/badge.svg?token=O10HKLLA5N)](https://codecov.io/gh/Josimar/GoLangDevBook)

[![CircleCI](https://circleci.com/gh/Josimar/GoLangDevBook/tree/main.svg?style=svg)](https://circleci.com/gh/Josimar/GoLangDevBook/tree/main)

# GoLangDevBook
Shared history between users

# Create directory API
Enter directory

## Criar estrutura 
go mod init api

###  Create file
main.go

#### Rodar a aplicação
go run main.go

## intalar dependência
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get github.com/go-sql-driver/mysql
go get github.com/badoux/checkmail
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go

# Create directory WebApp
Enter directory

## Start application
go mod init webapp

## start API with web
dir - API
- go build
- ./api

