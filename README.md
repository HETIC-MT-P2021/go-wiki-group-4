# Our school project

The goal of this project is to create a simple wiki.
We used Golang for our API, and Mysql as a database.

## Features:

- CRUD articles
- add comment
- Export article's data

## Installation

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

### With Docker

After cloning the repo, `cd` into the project and run following commands

```bash
docker-compose up --build
```

The app will be accessible at localhost:8080 !

### Technical Choices

Feel free to discuss with any contributor about the technical choices that were made.
Go version : 1.14.4
MySQL : 8.X

# Licence

The code is available under the MIT license.
