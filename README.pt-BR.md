> 📘 Read this documentation in [English 🇺🇸](./README.md)

# 📊 Finance API - Controle Financeiro em Go

Este projeto é uma API REST para gerenciamento de finanças pessoais desenvolvida em Go. Ela permite o registro de receitas e despesas, cálculo automático de saldo, persistência local de dados em arquivo JSON, e inclui gerenciamento de usuários.

## Funcionalidades

- Registro de transações (receitas e despesas)
- Cálculo automático de saldo
- Persistência de dados em arquivo JSON local
- Gerenciamento de usuários (registro, autenticação e exclusão)
- Associação de transações a usuários específicos

## Tecnologias Utilizadas

- Go (Golang)
- Gin Gonic (framework web)
- Encoding/JSON para persistência de dados
- Arquitetura modular

## Estrutura do Projeto

O projeto segue uma estrutura modular, organizada da seguinte forma:

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

## Como Executar

1. Clone o repositório:

```bash
  git clone https://github.com/mth-ribeiro-dev/finance-api.git
```
2. Navegue até o diretório do projeto:

```bash
  cd finance-api
```
3. Instale as dependências:

```bash
  go get -v ./...
```
4. Execute a apalicação
```bash
   go run cmd/server/main.go
```

## Endpoints da API

### Usuários
- `POST /user/register`: Registra um novo usuário
- `POST /user/login`: Autentica um usuário
- `DELETE /user/:id`: Desativa um usuário

### Transações Financeiras
- `POST /finance/transaction`: Adiciona uma nova transação
- `GET /finance/transactions/:userId`: Retorna todas as transações de um usuário
- `GET /finance/balance/:userId`: Retorna o saldo atual de um usuário
- `PUT /finance/:id`: Atualiza uma transação existente
- `DELETE /finance/:id`: Remove uma transação

## Testes

O projeto inclui testes unitários abrangentes para a camada de serviço. Para executar os testes:
```bash
   go test ./...
```
## Boas Práticas Aplicadas

- Organização modular do código
- Uso de `go.mod` para gerenciamento de dependências
- Implementação de testes unitários
- Uso de interfaces para desacoplamento (ex: storage)
- Validação robusta de entradas

## Licença

Este projeto está licenciado sob a [Creative Commons BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/) para fins educacionais e não comerciais.

## Autor

Desenvolvido por Matheus Ribeiro
- Email: matheus.junio159@gmail.com
- GitHub: [https://github.com/mth-ribeiro-dev](https://github.com/mth-ribeiro-dev)