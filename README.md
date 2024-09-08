# Todo List API

This is a practice project for [Todo List API](https://roadmap.sh/projects/todo-list-api).

## Usage

### Get Started

```bash
git clone https://github.com/moneychien19/todo-list-api-go.git
cd todo-list-api-go
go build -o main
./main # and the server will run on localhost:8080
```

### User Registration

```
POST http://localhost:8080/register
{
  "name": "John Doe",
  "email": "john@doe.com"
  "password": "password"
}
```

### User Login

```
POST http://localhost:8080/login
{
  "email": "john@doe.com",
  "password": "password"
}
```

### Create a Todo Item

```
POST http://localhost:8080/todos
Authorization: token
{
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread"
}
```

### Update a Todo Item

```
PUT http://localhost:8080/todos/1
Authorization: <token>
{
  "title": "Buy groceries",
  "description": "Buy milk, eggs, bread, and cheese"
}
```

### Delete a Todo Item

```
DELETE http://localhost:8080/todos/1
Authorization: <token>
```

### Get Todo Items

```
GET http://localhost:8080/todos?page=1&limit=10
Authorization: <token>
```
