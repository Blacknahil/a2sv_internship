# Task Manager API

## Overview
The Task Manager API is a RESTful API built using Go and Gin, with MongoDB as the persistent data storage. This API allows users to create, retrieve, update, and delete tasks efficiently.

## Features
- Create a new task
- Retrieve all tasks
- Retrieve a specific task by ID
- Update a task (partially or fully)
- Delete a task

## Technologies Used
- **Go**: Programming language
- **Gin**: Web framework for handling HTTP requests
- **MongoDB**: NoSQL database for persistent storage
- **MongoDB Go Driver**: Official MongoDB driver for Go

## Folder Structure
```
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── services/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── README.md
└── go.mod
```
- **main.go**: Entry point of the application.
- **controllers/**: Contains HTTP handlers for managing tasks.
- **models/**: Defines data structures used in the API.
- **services/**: Contains business logic for managing tasks.
- **router/**: Defines API routes using Gin.
- **docs/**: Contains project documentation.

## Installation & Setup
### Prerequisites
Ensure you have the following installed:
- [Go](https://golang.org/doc/install)
- [MongoDB](https://www.mongodb.com/try/download/community)

### Steps
1. Clone the repository:
   ```sh
   cd task-manager-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Start MongoDB locally or use a cloud-based MongoDB instance.
4. Run the application:
   ```sh
   go run main.go
   ```

## API Endpoints
### 1. Create a Task
**Endpoint:** `POST /tasks`
- **Request Body:**
  ```json
  {
    "title": "New Task",
    "description": "Task details",
    "status": "pending",
    "due_date": "2025-04-01T10:00:00Z"
  }
  ```
- **Response:**
  ```json
  {
    "id": "66035ae3e123456789abcdef",
    "title": "New Task",
    "description": "Task details",
    "status": "pending",
    "due_date": "2025-04-01T10:00:00Z"
  }
  ```

### 2. Get All Tasks
**Endpoint:** `GET /tasks`
- **Response:**
  ```json
  [
    {
      "id": "66035ae3e123456789abcdef",
      "title": "Complete Go API",
      "description": "Implement MongoDB persistence",
      "status": "pending",
      "due_date": "2025-04-01T10:00:00Z"
    }
  ]
  ```

### 3. Get Task by ID
**Endpoint:** `GET /tasks/{id}`
- **Response:**
  ```json
  {
    "id": "66035ae3e123456789abcdef",
    "title": "Complete Go API",
    "description": "Implement MongoDB persistence",
    "status": "pending",
    "due_date": "2025-04-01T10:00:00Z"
  }
  ```

### 4. Update a Task
**Endpoint:** `PUT /tasks/{id}`
- **Request Body:**
  ```json
  {
    "title": "Updated Task Title"
  }
  ```
- **Response:**
  ```json
  {
    "message": "Task updated successfully"
  }
  ```

### 5. Delete a Task
**Endpoint:** `DELETE /tasks/{id}`
- **Response:**
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```

## Contribution and Feedback
Feel free to contribute to this project and giving code reviews(even if harsh I don't mind).



