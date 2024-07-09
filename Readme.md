# Movie CRUD without actual DB have used  struct and slices for implementation used mux for routes

## Dependencies  
go get github.com/gorilla/mux

## Module initialized
go mod init example.com/gomoviescrud

## To run application follow below once you clone the project
go build
go run main.go 

## test endpoints either via postman or restclient http in vscode 
### GET request for fetching movies
GET http://localhost:8000/movies 

### Get request to fetch one movie
GET http://localhost:8000/movies/1
Content-Type: application/json


### POST request to create another movie
POST http://localhost:8000/movies
Content-Type: application/json

{
 "isbn": "33333",
  "title": "Movie three",
  "director": {
    "firstname": "postreqfirstname",
    "lastname": "postreqlastname"
  }
}

### UPDATE request to updated the movie record say for Id 2
PUT http://localhost:8000/movies/2
Content-Type: application/json

{
  
  "isbn": "update",
  "title": "up three",
  "director": {
    "firstname": "u",
    "lastname": "u"
  }
}

### DELETE request for records say for id 2 as above we have updated
DELETE http://localhost:8000/movies/2
Content-Type: application/json
