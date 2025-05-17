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
## 🧪 Testes Automatizados

Este projeto possui testes unitários completos para os seguintes métodos da camada de serviço (`FinanceService`):

- `AddTransaction` — adiciona transações com diferentes validações
- `GetAll` — retorna todas as transações
- `GetBalance` — calcula o saldo com base em receitas e despesas
- `DeleteTransaction` — remove transações por ID
- `UpdateTransaction` — atualiza uma transação existente
- `GetMaxID` — retorna o maior ID registrado

### ▶️ Rodando os testes

Execute os testes com o seguinte comando na raiz do projeto:

```bash
go test ./...
```

Para ver a cobertura de testes:

```bash
go test -cover ./...
```


## 🧱 Boas Práticas Adotadas

- Organização modular: `cmd/`, `internal/`, `handler/`, `model/`, `service/`, `storage/`
- Uso de `go.mod` para controle de dependências
- Arquivos sensíveis e binários ignorados via `.gitignore`
- Caminho de dados baseado no `user.HomeDir` (portável e seguro)
- Dados salvos localmente com `os.MkdirAll` e `encoding/json`
- Validação robusta de dados: tipo, data (`yyyy-mm-dd`), e campos obrigatórios

---

## 📄 Licença

Este projeto está licenciado sob a [Creative Commons BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/).  
Uso permitido apenas para fins **educacionais e não comerciais**, com atribuição ao autor.

---

## ✍️ Autor

Desenvolvido por Matheus Ribeiro  
Contato: matheus.junio159@gmail.com  
GitHub: [https://github.com/mth-ribeiro-dev](https://github.com/mth-ribeiro-dev)