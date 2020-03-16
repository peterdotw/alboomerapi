# All-Boomer REST API

Example Album Database REST API written entirely in Go.

## Built With

- [Go](https://golang.org/) - Go programming language
- [Go-MySQL-Driver](github.com/go-sql-driver/mysql) - A MySQL-Driver for Go's database/sql package
- [dotsql](https://github.com/gchaincl/dotsql) - Used for reading queries from SQL files
- [gorilla/mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers
- [CORS](https://github.com/rs/cors) - Go CORS handler
- [GoDotEnv](https://github.com/joho/godotenv) - Dotenv for Go

## To do

- [x] Handle PUT/DELETE requests
- [x] Connect API with a MySQL database on Google Cloud SQL
- [x] Make database queries for all routes
- [x] Handle all security stuff
- [x] Expand Album structure with more info and DB tables
- [x] Add Redis caching for GET requests
- [ ] Implement JWT authentication