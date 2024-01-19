# First-goland-no-databse-API
This is a simple RESTful API for managing books. It allows you to retrieve a list of books, get details of a specific book by ID, create a new book, check out a book, and return a book.

Prerequisites
Go (Golang) installed on your machine
Gin package installed
Getting Started
Clone the repository:

bash
Copy code
git clone https://github.com/emmanuelalo52/First-goland-no-databse-API.git
cd your-repo
Install dependencies:

bash
Copy code
go get -u github.com/gin-gonic/gin
Run the application:

bash
Copy code
go run main.go
The API will be accessible at http://localhost:8080.

API Endpoints
1. Get All Books
Endpoint: GET /books
Description: Retrieve a list of all books.
2. Get Book by ID
Endpoint: GET /books/:id
Description: Retrieve details of a specific book by its ID.
3. Create a New Book
Endpoint: POST /books
Description: Create a new book.
json
Copy code
{
  "id": 5,
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "quantity": 3
}
4. Check Out a Book
Endpoint: PATCH /checkout?id=<book_id>
Description: Check out a book by reducing its quantity.
5. Return a Book
Endpoint: PATCH /return?id=<book_id>
Description: Return a book by increasing its quantity.
Contributing
Contributions are welcome! Please feel free to open issues or pull requests.

License
This project is licensed under the MIT License - see the LICENSE file for details.


