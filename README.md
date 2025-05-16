# 📊 Finance API - Controle Financeiro em Go

Este projeto implementa uma API REST simples para controle de finanças pessoais, permitindo o registro de receitas e despesas, cálculo automático de saldo e persistência local dos dados em arquivo JSON.

---

## 🚀 Tecnologias Utilizadas

- **Go (Golang)** — linguagem principal
- **Gin Gonic** — framework web para a API REST
- **JSON** — formato de armazenamento dos dados
- **Arquitetura Modular** — separação em `cmd`, `internal` e `pkg` (boa prática Go)

---

## 📂 Estrutura do Projeto

```
finance-api/
├── cmd/               # Entrada principal da aplicação
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/       # Controladores HTTP (API)
│   ├── model/         # Definição das estruturas de dados
│   ├── service/       # Lógica de negócios e regras
│   └── storage/       # Persistência em arquivo JSON
├── go.mod             # Gerenciamento de dependências Go
├── .gitignore         # Arquivos ignorados pelo Git
└── README.md          # Documentação do projeto
```

---

## 🧠 Como Funciona

A API permite registrar transações financeiras com os seguintes campos:

```json
{
  "type": "income" | "expense",
  "amount": 100.0
}
```

Cada transação é armazenada com um ID único e salva em:

```
C:\Users\<seu_usuario>\financeiro\arquivo\transactions.json
```

Esse caminho é gerado automaticamente com base no usuário atual do sistema.

---

## 📌 Endpoints Disponíveis

| Método | Rota             | Função                         |
|--------|------------------|--------------------------------|
| POST   | `/transaction`   | Adiciona uma nova transação    |
| GET    | `/transactions`  | Retorna o histórico completo   |
| GET    | `/balance`       | Retorna o saldo atual          |

---

## ▶️ Executando o Projeto

### Pré-requisitos

- Go 1.20+ instalado e configurado
- Git (opcional para clonar)

### Passos

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/finance-api.git
cd finance-api
```

2. Baixe as dependências:

```bash
go mod tidy
```

3. Rode o projeto:

```bash
go run ./cmd/server
```

A API estará disponível em:  
`http://localhost:8080`

---

## 🧪 Testando com cURL ou Postman

### ✅ Adicionar receita

```bash
curl -X POST http://localhost:8080/transaction \
  -H "Content-Type: application/json" \
  -d "{\"type\":\"income\",\"amount\":1000}"
```

### ❌ Adicionar despesa

```bash
curl -X POST http://localhost:8080/transaction \
  -H "Content-Type: application/json" \
  -d "{\"type\":\"expense\",\"amount\":200}"
```

### 📋 Ver transações

```bash
curl http://localhost:8080/transactions
```

### 💰 Ver saldo atual

```bash
curl http://localhost:8080/balance
```

---

## 🧱 Boas Práticas Adotadas

- Organização modular: `cmd/`, `internal/`, `handler/`, `model/`, `service/`, `storage/`
- Uso de `go.mod` para controle de dependências
- Arquivos sensíveis e binários ignorados via `.gitignore`
- Caminho de dados baseado no `user.HomeDir` (portável e seguro)
- Dados salvos localmente com `os.MkdirAll` e `encoding/json`

---

## 📬 Licença

Este projeto foi desenvolvido para fins educacionais e uso pessoal. Livre para estudar, modificar e reutilizar.

---

## ✍️ Autor

Desenvolvido por Matheus Ribeiro
Contato: matheus.junio159@gmail.com  
GitHub: [https://github.com/mth-ribeiro-dev](https://github.com/seu-usuario)
