# Go movies CRUD
This is a simple movie CRUD on Go just to train the language and my "muscular memory" on how to write code properly with Go

## Table of contents
- [Credits](#credits)
- [Requirements](#requirements)
- [Installation](#installation)
- [Running the project](#running-the-project)
- [Endpoints](#endpoints)
- [Swagger](#swagger)
- [Author](#author)
- [License](#license)

## Credits
This project was based on the following tutorial: [Learn Go Programming by Building 11 Projects – Full Course
](https://youtu.be/jFfo23yIWac?t=1236) by [freeCodeCamp.org](https://www.youtube.com/channel/UC8butISFwT-Wl7EV0hUK0BQ)

## Requirements
- [Go](https://golang.org/dl/)

## Installation
You can install the project using the following commands:
  
  ```bash
  git clone https://github.com/ologbonowiwi/go-movies-crud.git
  cd go-movies-crud
  go mod download
  ```

## Running the project
You can run the project using the following commands:
  
  ```bash
  go run main.go
  ```
  or
  ```bash
  go build
  ./go-movies-crud
  ```

## Endpoints
The endpoints are the following:

  - GET `/movies`
  - GET `/movies/{id}`
  - POST `/movies`
  - PUT `/movies/{id}`
  - DELETE `/movies/{id}`

## Swagger
You can access the swagger documentation on the following link: http://localhost:8000/swagger/index.html

## Author
My name is [Wesley Matos](https://github.com/ologbonowiwi), and you can find me here (at GitHub 😁), on my [LinkedIn](https://www.linkedin.com/in/ologbonowiwi/) or by [e-mail](mailto:ologbonowiwi520@gmail.com).

## License
[MIT](https://choosealicense.com/licenses/mit/)
