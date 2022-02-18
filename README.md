# this simple Api

Simple Api for learning Course using Golang and Mysql for the database. Implementation Json Web Token (JWT) and CRUD using GoFiber.

Step by step:<br>
First make your directory and make a git clone: `https://github.com/AnggaDanarP/Online-Learning-Plathform.git`

run terminal `go mod init`
install dependencies that we need in terminal using `go get`:
1. `github.com/gofiber/fiber/v2`
2. `github.com/dgrijalva/jwt-go`
3. `github.com/go-sql-driver/mysql`

Open your Mysql Workbench and maka a schema on there, the program automatically will make table. The code is from models folder and would be deploy by `connection.go` from database folder

Open database folder and open `connection.go`. Change the value above `the comment` with your username andd password by yor database mysql workbench

Run the program using terminal `go run main.go`

The program will run by `localhost:8000` and use Postman to interact with it.

###### HAPPY HACKING

