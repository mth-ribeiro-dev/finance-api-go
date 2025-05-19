📘 Read this documentation in [Português 🇧🇷](./README.pt-BR.md)

# 📊 Finance API - Personal Finance Management in Go

This project is a REST API for managing personal finances developed in Go. It allows recording income and expenses, automatic balance calculation, local data persistence using a JSON file, and includes user management features.

## Features

- Transaction recording (income and expenses)
- Automatic balance calculation
- Local data persistence using JSON file
- User management (registration, authentication, and deletion)
- Transactions linked to specific users

## Technologies Used

- Go (Golang)
- Gin Gonic (web framework)
- Encoding/JSON for data persistence
- Modular architecture

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

## API Endpoints

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