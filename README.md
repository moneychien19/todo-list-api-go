# Todo List API

This is a practice project of [Todo List API](https://roadmap.sh/projects/todo-list-api).

## Usage

### Get Started

```bash
git clone https://github.com/moneychien19/todo-list-api-go.git
```

### User Registration

```
POST /register
{
  "name": "John Doe",
  "email": "john@doe.com"
  "password": "password"
}
```

### User Login

```
POST /login
{
  "email": "john@doe.com",
  "password": "password"
}
```

### Create a Todo Item

```
POST /todos
{
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread"
}
```

### Update a Todo Item

```
PUT /todos/1
{
  "title": "Buy groceries",
  "description": "Buy milk, eggs, bread, and cheese"
}
```

### Delete a Todo Item

```
DELETE /todos/1
```

### Get Todo Items

```
GET /todos?page=1&limit=10
```
