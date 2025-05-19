📘 Read this documentation in [Português 🇧🇷](./README.pt-BR.md)

# 📊 Finance API - Personal Finance Management in Go

This project is a REST API for managing personal finances developed in Go. It allows recording income and expenses, automatic balance calculation, local data persistence using a JSON file, and includes user management features.

## Features

- Transaction recording (income and expenses)
- Automatic balance calculation
- Local data persistence using JSON file
- User management (registration, authentication, and deletion)
- Transactions linked to specific users
- API documentation with Swagger

## Technologies Used

- Go (Golang)
- Gin Gonic (web framework)
- Encoding/JSON for data persistence
- Modular architecture
- Swagger for API documentation

## Project Structure

The project follows a modular structure, organized as follows:


```
finance-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   │   ├── finance.go
│   │   └── user.go
│   ├── model/
│   │   ├── dateOnly.go
│   │   ├── transaction.go
│   │   └── user.go
│   ├── service/
│   │   ├── finance.go
│   │   ├── finance_test.go
│   │   ├── user.go
│   │   └── user_test.go
│   └── storage/
│       ├── commonStorage.go
│       ├── financesStorage.go
│       └── usersStorage.go
├── go.mod
├── go.sum
├── .gitignore
├── README.md
└── README.pt-BR.md
```
## How to Run

1. Clone the repository:

```bash
  git clone https://github.com/mth-ribeiro-dev/finance-api.git
```

2. Navigate to the project directory:

```bash
  cd finance-api
```

3. Install dependencies:

```bash
  go get -v ./...
```

4. Run the application:

```bash
   go run cmd/server/main.go
```

5. Access the Swagger documentation:

   Open your browser and go to `http://localhost:8081/swagger/index.html`

## API Endpoints

For a detailed and interactive documentation of all API endpoints, please refer to the Swagger UI available at `http://localhost:8081/swagger/index.html` when the application is running.

### Users
- `POST /user/register`: Registers a new user
- `POST /user/login`: Authenticates a user
- `DELETE /user/:id`: Deactivates a user

### Financial Transactions
- `POST /finance/transaction`: Adds a new transaction
- `GET /finance/transactions/:userId`: Returns all transactions for a user
- `GET /finance/balance/:userId`: Returns a user's current balance
- `PUT /finance/:id`: Update an existing transaction
- `DELETE /finance/:id`: Delete a transaction

## Testing

The project includes comprehensive unit tests for the service layer. To run the tests:
```bash
   go test ./...
```

## Documentation

This project uses Swagger for API documentation. The documentation includes detailed information about all endpoints, request/response models, and allows for interactive API testing.

To view the Swagger documentation:
1. Start the application
2. Open your web browser
3. Navigate to `http://localhost:8081/swagger/index.html`

The code also includes comprehensive comments, enhancing readability and maintainability.

## Best Practices Applied

- Modular code organization
- Use of `go.mod` for dependency management
- Implementation of unit tests
- Use of interfaces for decoupling (e.g., storage)
- Robust input validation

## License

This project is licensed under the [Creative Commons BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/) license for educational and non-commercial purposes.

## Author

Developed by Matheus Ribeiro
- Email: matheus.junio159@gmail.com
- GitHub: [https://github.com/mth-ribeiro-dev](https://github.com/mth-ribeiro-dev)