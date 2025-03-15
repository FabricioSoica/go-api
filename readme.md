
# CRUD em Golang com MySQL e Swagger

Este é um exemplo de um CRUD simples implementado em Golang com integração ao MySQL. A API oferece funcionalidades básicas de criação, leitura, atualização e exclusão de usuários. A documentação da API é gerada automaticamente via Swagger.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação.
- **MySQL**: Banco de dados para armazenar as informações dos usuários.
- **Swagger**: Para a geração automática da documentação da API.
- **GORM**: ORM para interação com o MySQL.
- **Dotenv**: Para carregar variáveis de ambiente a partir de um arquivo `.env`.

## Funcionalidades

- **Cadastro de usuário**: Criação de novos usuários.
- **Leitura de usuários**: Consulta de todos os usuários ou um usuário específico.
- **Atualização de usuário**: Modificação das informações de um usuário existente.
- **Exclusão de usuário**: Remoção de um usuário do banco de dados.
- **Login e autenticação**: Implementação de login simples com autenticação via JWT.

## Instalação

### Pré-requisitos

1. **Golang**: Verifique se o Golang está instalado em seu sistema. Caso não, siga a [documentação oficial](https://golang.org/doc/install).
2. **MySQL**: Instale o MySQL localmente e configure a porta 3306 (ou qualquer outra porta que preferir).
3. **Swagger**: A documentação da API será gerada automaticamente com Swagger.

### Passos para Execução

1. Clone o repositório para o seu computador:
   ```bash
   git clone https://github.com/seuusuario/crud-golang-mysql-swagger.git
   cd crud-golang-mysql-swagger
   ```

2. Crie o arquivo `.env` na raiz do projeto com as configurações do banco de dados:
   ```
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=goapi
   DB_NAME=go_api
   ```

3. Instale as dependências do projeto:
   ```bash
   go mod tidy
   ```

4. Execute as migrations para criar as tabelas no banco de dados:
   ```bash
   go run migrations.go
   ```

5. Inicie o servidor:
   ```bash
   go run main.go
   ```

6. Acesse a documentação Swagger da API em: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Endpoints

- `POST /users` - Criar um novo usuário
- `GET /users` - Listar todos os usuários
- `GET /users/{id}` - Obter um usuário pelo ID
- `PUT /users/{id}` - Atualizar um usuário
- `DELETE /users/{id}` - Excluir um usuário

## Exemplo de Requisição

### Criar Usuário
```bash
POST /users
Content-Type: application/json

{
  "name": "João da Silva",
  "email": "joao@example.com",
  "password": "123456"
}
```

### Resposta:
```json
{
  "id": 1,
  "name": "João da Silva",
  "email": "joao@example.com"
}
```

## Swagger

A documentação da API é gerada automaticamente utilizando o Swagger. Você pode acessá-la localmente em [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

## Mascote Go

<img src="https://go.dev/images/gophers/motorcycle.svg" style="width: 40%;" alt="Go Mascote">

---

Projeto desenvolvido apenas para fins de aprendizado ^-^
