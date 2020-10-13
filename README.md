# Our school project

The goal of this project is to create a simple wiki and an introduction to MVC, Strategies design patern and Unit tests in Go.
We used Golang for our API, and Mysql as a database.

## Features:

- CRUD articles
- Add comment
- Export article's data
- Some unit tests

### Start with Docker

After cloning the repo, `cd` into the project, create the .env according to the .env.example and run following commands

```bash
docker-compose up --build
```

The app will be accessible at localhost:8080 !

## Start manualy

The project requires Golang v 1.14.4

Install the dependencies and start the server.

```sh
$ git clone https://github.com/HETIC-MT-P2021/go-wiki-group-4.git
$ go get
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/cosmtrek/air
$ go mod vendor
$ air
```

### Technical Choices

Feel free to discuss with any contributor about the technical choices that were made.
Go version : 1.14.4
MySQL : 8.X

### Authors

- [Myouuu](https://github.com/myouuu)
- [Tsabot](https://github.com/Tsabot)
- [Jean-Jacques](https://github.com/gensjaak)

# Licence

The code is available under the MIT license.
