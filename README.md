# All-Boomer REST API

Example Album Database REST API written entirely in Go.

## Built With

- [Go](https://golang.org/) - Go programming language
- [Cloud SQL Proxy](github.com/GoogleCloudPlatform/cloudsql-proxy) - Cloud SQL proxy client for connecting to MySQL database
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
- [ ] Add Redis caching
- [ ] Implement JWT authentication