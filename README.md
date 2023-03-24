# Go API

Go API é um projeto de exemplo de uma API RESTful em Go. Ele utiliza o framework [Gin](https://github.com/gin-gonic/gin) para tratar requisições HTTP e a biblioteca [GORM](https://gorm.io/) para interagir com o banco de dados.

## Como executar

Para executar o projeto, siga os passos abaixo:

1. Clone este repositório para sua máquina local
2. Certifique-se de que o Go está instalado em sua máquina
3. Configure as variáveis de ambiente no arquivo `.env` (exemplo fornecido em `.env.example`)
4. Execute o comando `go run cmd/main.go`

O servidor será iniciado em http://localhost:8080.

## Endpoints

A API oferece os seguintes endpoints:

### POST /api/register

Endpoint responsável por registrar um novo usuário na aplicação.

Requisição:

```bash
{
  "email": "email@example.com",
  "password": "123456"
}
```

Resposta:

```bash
{
  "status": "success",
  "message": "User created successfully"
}
```


### POST /api/login

Endpoint responsável por autenticar um usuário na aplicação.

Requisição:

```bash
{
  "email": "email@example.com",
  "password": "123456"
}
```

Resposta:

```bash
{
  "status": "success",
  "message": "Logged in successfully",
  "data": {
    "user_id": 1,
    "access_token": "<jwt-token>"
  }
}
```

### GET /

Endpoint de teste.

Resposta:

```bash
{
  "status": "success",
  "message": "API is working"
}
```

## Estrutura do projeto

A estrutura do projeto é baseada no padrão MVC. O diretório `api` contém o código relacionado à API. O diretório `config` contém configurações da aplicação e o arquivo `config.database` contém o código relacionado ao banco de dados.


