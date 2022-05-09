# Books Data Back End
It was my typing exercise from https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/. It takes around 1 hour 47 minutes.
## Endpoint
|Request method|Path|Controller|
|:---:|:---:|:---:|
|GET|/books|FindBooks|
|POST|/books|CreateBook|
|GET|/books/:id|FindBookById|
|PATCH|/books/:id|UpdateBookById|
|DELETE|/books/:id|DeleteBookById|
## Struct type definition
|File path|Struct name|Description|
|:---:|:---:|:---|
|models/book.go|Book|Book model, the entity in this project have its fields/attributes/properties|
|controllers/books.go|CreateBookInput|User input validation schema for CreateBook|
|controllers.books.go|UpdateBookInput|User input validation schema for UpdateBookById|