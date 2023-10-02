# JSON CRUD API in Go

This project is an example of implementing CRUD (Create, Read, Update, Delete) operations in Go language to work with data in PostgreSQL database. 

## Project structure
- `main.go`: Server main file
- `models/task.go`: Defining a model for the database
- `migrate/postModel.go`/: Database migration
- `initializers/database.go`/: Connecting to the database
- `initializers/loadEnvVariables.go`/: Environment variables
- `controllers/postsController.go`/: Realization of CRUD queries


## Functionality

+ Creating new records
+ Updating existing records
+ Reading data from the database
+ Deleting records from the database

## Getting Started

1. Make sure you have Go installed. If not, you can download it from [official website](https://golang.org/dl/)

2. Clone the repository to your computer:
```
git clone https://github.com/rovezuka/go-crud
```

3. Install the required dependencies:

Go to the root directory of the project and run ``go mod tidy`` to install all required dependencies


4. Run the application:

Execute the command go run `main.go` to run the project on the local server. By default, the project will be available at `http://localhost:8080``

## Usage

1. Navigate to the main page of the application

2. Add a new task by entering its description and clicking the "Add" button

3. Mark tasks as completed by pressing the corresponding button

4. Delete tasks by clicking on the "Delete" button

## API Usage

The project provides the following routes for working with data:

- `POST /posts`: Create a new record. Send a JSON request with the record data
- `PUT /posts/:id`: Update an existing record by id. Send a JSON request with the updated record data
- `GET /posts`: Get a list of all records
- `GET /posts/:id`: Get record information by id
- `DELETE /posts/:id`: Delete a record by identifier

For each request, the API expects a JSON object with data in the request body and returns a JSON response with the result of the operation.

Example of a request to create a new record using Postman:
```
http://localhost:8080/posts

{
    "title": "Record title",
    "content": "Content of the record"
}
```
Example of a response in Postman:
```
{
    "id": 1,
    "title": "Record title",
    "content": "Content of the record"
}
```

## Configuring the database

The project uses PostgreSQL database. To customize the database connection, specify the appropriate parameters in the .env file. Example .env file:
```
PORT=5432
DB_URL="YOUR_URL"
```

## License

This project is distributed under the MIT license. See the ``LICENSE`` file for details.