
> 📘 Leia esta documentação em [português 🇧🇷](./README.pt-BR.md)


# 📊 Finance API - Personal Finance Control in Go

This project implements a simple REST API for personal finance management, allowing the registration of income and expenses, automatic balance calculation, and local data persistence in a JSON file.

---

## 🚀 Technologies Used

- **Go (Golang)** — main programming language
- **Gin Gonic** — web framework for the REST API
- **JSON** — data storage format
- **Modular Architecture** — structured in `cmd`, `internal`, and `pkg` following Go best practices

---

## 📂 Project Structure

```
finance-api/
├── cmd/               # Application entry point
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/       # HTTP controllers (API)
│   ├── model/         # Data structure definitions
│   ├── service/       # Business logic and rules
│   └── storage/       # JSON file persistence
├── go.mod             # Go dependency management
├── .gitignore         # Files ignored by Git
└── README.md          # Project documentation
```

---

## 🧪 Automated Tests

This project includes comprehensive unit tests for the following methods in the service layer (`FinanceService`):

- `AddTransaction` — adds transactions with multiple validations
- `GetAll` — returns all stored transactions
- `GetBalance` — calculates balance based on income and expenses
- `DeleteTransaction` — removes a transaction by ID
- `UpdateTransaction` — updates an existing transaction

### ▶️ Running Tests

To run all tests from the project root:

```bash
go test ./...
```

To check test coverage:

```bash
go test -cover ./...
```

---

## 🧱 Best Practices Applied

- Modular organization: `cmd/`, `internal/`, `handler/`, `model/`, `service/`, `storage/`
- Use of `go.mod` for dependency management
- Sensitive files and binaries ignored via `.gitignore`
- Cross-platform safe file path using `user.HomeDir`
- Data saved locally using `os.MkdirAll` and `encoding/json`
- Robust validation: transaction type, date (`yyyy-mm-dd`), and required fields

---

## 📄 License

This project is licensed under the [Creative Commons BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/) license.  
Use is permitted only for **educational and non-commercial purposes**, with proper credit to the author.

---

## ✍️ Author

Developed by Matheus Ribeiro  
Contact: matheus.junio159@gmail.com  
GitHub: [https://github.com/mth-ribeiro-dev](https://github.com/mth-ribeiro-dev)