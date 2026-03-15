# TASK API

![Go](https://img.shields.io/badge/Language-Go-blue)
![API](https://img.shields.io/badge/Type-API-orange)
![In Memory](https://img.shields.io/badge/Database-In_Memory-yellow)
![Active](https://img.shields.io/badge/Status-Active-green)


Simple REST API for managing tasks.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /tasks | Get all tasks |
| GET | /tasks/{id} | Get task by id |
| POST | /tasks | Create new task |
| PUT | /tasks/{id} | Replace existing task |
| DELETE | /tasks/{id} | Delete task |

## Data Models

### Task
```json
{
  "id":"string",
  "title":"string",
  "description":"string",
  "completed":false
}
```

## Examples

[Examples](./docs/examples.md)

## How To Run

```bash
go run ./cmd/api
```

**Server starts on:**

`http://localhost:8080`

## Project Structure

```
cmd/api
  main.go
  router.go
  
internal/todo
  handler.go
  model.go
  repository.go
```

## Technical Notes

- Built using Go standard library (`net/http`)
- Custom `http.Server` with configured timeouts
- Thread-safe repository using `sync.RWMutex`
