# RESTful CRUD API written in Go 

## Stack

- [Go](https://go.dev/) 
- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin) 

## MVP Objectives

* [X] as a API, I have at least one endpoint `/users`
* [X] as a user, I can ceate a `User` instance by sending a `POST` request to `/users`
* [X] as a user, I can udpate a `User` instance by sending a `PATCH` request to `/users/:id`
* [X] as a user, I can read one or many `User` instance by sending a `GET` request `/users/:id`
* [X] as a user, I can delete a `User` instance by sending a `POST` request `/users/:id`

## Data persistance

* [X] In-memory
* [ ] relational-database

## Nice-To-Have

* [ ] Authentication
* [ ] Have a one-to-many relationship with a second model
* [ ] Readable and horizontally scallable file structure
* [ ] Explore popular ORM packages and implement one
* [ ] Unit test using native `testing` package