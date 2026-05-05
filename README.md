# Task Manager with Go

A simple, lightweight Task Manager API built using Go's standard library. This project demonstrates a clean architecture with separate layers for handlers, services, and models.

## Features

- **Get All Tasks**: Retrieve a list of all tasks.
- **Create Task**: Add a new task to the list.
- **In-Memory Storage**: Tasks are stored in memory for simplicity (resets on server restart).

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: Standard Library (`net/http`)
- **Data Format**: JSON

## Project Structure

```text
task-manager-with-go/
├── handler/            # HTTP request handlers
│   └── task_handler.go
├── model/              # Data models/structs
│   └── task.go
├── service/            # Business logic and data management
│   └── task_service.go
├── go.mod              # Go module definition
└── main.go             # Application entry point and routing
```

## API Endpoints

### 1. Get All Tasks
- **URL**: `/tasks`
- **Method**: `GET`
- **Success Response**: `200 OK`
- **Sample Output**:
  ```json
  [
    {
      "id": "1",
      "title": "Learn Go",
      "done": false
    }
  ]
  ```

### 2. Create a Task
- **URL**: `/tasks`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "id": "2",
    "title": "Build a REST API",
    "done": false
  }
  ```
- **Success Response**: `200 OK`

## How to Run

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd task-manager-with-go
   ```

2. **Run the application**:
   ```bash
   go run main.go
   ```

3. **Access the API**:
   The server will start on `http://localhost:8080`. You can use tools like `curl`, Postman, or Insomnia to interact with the endpoints.

   **Example CURL to get tasks**:
   ```bash
   curl http://localhost:8080/tasks
   ```

   **Example CURL to create a task**:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"id":"1", "title":"Task One", "done":false}' http://localhost:8080/tasks
   ```
