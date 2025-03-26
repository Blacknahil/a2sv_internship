# Task Manager API

## Overview
The Task Manager API is a RESTful service for managing tasks. It allows users to create, read, update, and delete tasks. The API is built using Go with the Gin framework.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
  - [Get All Tasks](#2-get-all-tasks)
  - [Get Task by ID](#3-get-task-by-id)
  - [Create a Task](#4-create-a-task)
  - [Update a Task](#5-update-a-task)
  - [Delete a Task](#6-delete-a-task)
- [Error Handling](#error-handling)
- [Running the API](#running-the-api)
- [Conclusion](#conclusion)

## Installation
1. Clone the repository:

2. Navigate to the project directory:
   ```sh
   cd task-manager-api
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage
Run the application using:
```sh
go run main.go
```
The server will start at `http://localhost:8080`.

## Endpoints


---

### 1. Get All Tasks
**Endpoint:** `GET /tasks`

**Description:**
- Retrieves a list of all tasks.

**Response:**
```json
{
  "tasks": [
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task",
      "due_date": "2025-03-24T00:00:00Z",
      "status": "Pending"
    }
  ]
}
```

---

### 2. Get Task by ID
**Endpoint:** `GET /tasks/:id`

**Description:**
- Retrieves a specific task by ID.

**Response (Success):**
```json
{
  "id": "1",
  "title": "Task 1",
  "description": "First task",
  "due_date": "2025-03-24T00:00:00Z",
  "status": "Pending"
}
```

**Response (Failure):**
```json
{
  "error": "Task not found"
}
```

---

### 3. Create a Task
**Endpoint:** `POST /tasks`

**Description:**
- Creates a new task.

**Request Body:**
```json
{
  "id": "2",
  "title": "Task 2",
  "description": "Second task",
  "due_date": "2025-03-25T00:00:00Z",
  "status": "In Progress"
}
```

**Response:**
```json
{
  "message": "Task created"
}
```

---

### 4. Update a Task
**Endpoint:** `PUT /tasks/:id`

**Description:**
- Updates an existing task.

**Request Body:**
```json
{
  "title": "Updated Task Title",
  "description": "Updated description"
}
```

**Response (Success):**
```json
{
  "message": "Task updated"
}
```

**Response (Failure):**
```json
{
  "error": "Task not found"
}
```

---

### 5. Delete a Task
**Endpoint:** `DELETE /tasks/:id`

**Description:**
- Deletes a task by ID.

**Response (Success):**
```json
{
  "message": "Task deleted"
}
```

**Response (Failure):**
```json
{
  "error": "Task not found"
}
```

---

## Error Handling
All error responses follow this format:
```json
{
  "error": "Error message here"
}
```

## Conclusion
This API provides a simple way to manage tasks with CRUD operations. Future enhancements can include authentication, persistence (database storage), and task prioritization.

