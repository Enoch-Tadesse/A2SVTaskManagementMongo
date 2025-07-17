# Task Manager API Documentation

Welcome to the **Task Manager API**! This API allows you to manage tasks efficiently with endpoints for creating, retrieving, updating, and deleting tasks. Built with Go and Gin, it connects to a MongoDB database for persistent storage.

---

## Base URL
All endpoints are relative to the base URL:  
**`http://localhost:8080`**

---

## Endpoints

### 1. Get All Tasks
Retrieve a list of all tasks.

- **Method**: GET
- **Path**: `/tasks`
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:
    ```json
    {
      "data": [
        {
          "id": "507f1f77bcf86cd799439011",
          "title": "Complete project",
          "description": "Finish the API documentation",
          "due_date": "2025-08-01",
          "status": "pending"
        },
        ...
      ]
    }
    ```
  - **Error**:
    - **500 Internal Server Error**: `{ "error": "Failed to fetch tasks: <error message>" }`

---

### 2. Get Task by ID
Retrieve a specific task by its ID.

- **Method**: GET
- **Path**: `/tasks/:id`
- **Parameters**:
  - `id` (path): MongoDB ObjectID of the task (e.g., `507f1f77bcf86cd799439011`)
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:
    ```json
    {
      "task": {
        "id": "507f1f77bcf86cd799439011",
        "title": "Complete project",
        "description": "Finish the API documentation",
        "due_date": "2025-08-01",
        "status": "pending"
      }
    }
    ```
  - **Errors**:
    - **400 Bad Request**: `{ "error": "id can not be empty" }` or `{ "error": "invalid id format" }`
    - **404 Not Found**: `{ "error": "failed to retrieve task: <error message>" }`

---

### 3. Create a Task
Add a new task to the collection.

- **Method**: POST
- **Path**: `/tasks`
- **Request Body**:
  ```json
  {
    "title": "New Task",
    "description": "Task description",
    "due_date": "2025-08-01",
    "status": "pending"
  }
  ```
- **Response**:
  - **Status Code**: 201 Created
  - **Body**:
    ```json
    {
      "message": "task created successfully",
      "data": {
        "id": "507f1f77bcf86cd799439011",
        "title": "New Task",
        "description": "Task description",
        "due_date": "2025-08-01",
        "status": "pending"
      }
    }
    ```
  - **Errors**:
    - **400 Bad Request**:
      - `{ "error": "invalid request body" }`
      - `{ "error": "title can not be empty" }`
      - `{ "error": "due date can not be in the past" }`
      - `{ "error": "invalid status" }`
    - **404 Not Found**: `{ "error": "<error message>" }`

---

### 4. Update a Task
Fully replace an existing task with new data.

- **Method**: PUT
- **Path**: `/tasks/:id`
- **Parameters**:
  - `id` (path): MongoDB ObjectID of the task
- **Request Body**:
  ```json
  {
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2025-09-01",
    "status": "completed"
  }
  ```
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:
    ```json
    {
      "message": "task updated successfully"
    }
    ```
  - **Errors**:
    - **400 Bad Request**:
      - `{ "error": "invalid id format" }`
      - `{ "error": "invalid request body" }`
      - `{ "error": "title can not be empty" }`
      - `{ "error": "due date can not be in the past" }`
      - `{ "error": "invalid status" }`
    - **404 Not Found**: `{ "error": "task not found" }`

---

### 5. Delete a Task
Remove a task by its ID.

- **Method**: DELETE
- **Path**: `/tasks/:id`
- **Parameters**:
  - `id` (path): MongoDB ObjectID of the task
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:
    ```json
    {
      "message": "task deleted successfully"
    }
    ```
  - **Errors**:
    - **400 Bad Request**:
      - `{ "error": "id can not be empty" }`
      - `{ "error": "invalid id format" }`
    - **404 Not Found**: `{ "error": "task with id <id> is not found" }`

---

## Data Model
The API uses the following task structure:

```json
{
  "id": "MongoDB ObjectID",
  "title": "string",
  "description": "string",
  "due_date": "YYYY-MM-DD",
  "status": "pending | completed | missed"
}
```

- **id**: Unique identifier (MongoDB ObjectID).
- **title**: Task title (required, non-empty).
- **description**: Task description (optional).
- **due_date**: Due date in `YYYY-MM-DD` format (must be in the future).
- **status**: Task status (must be `pending`, `completed`, or `missed`).

---

## Error Handling
The API returns errors in the following format:

```json
{
  "error": "<error message>"
}
```

Common error codes:
- **400 Bad Request**: Invalid input or request format.
- **404 Not Found**: Task or resource not found.
- **500 Internal Server Error**: Server-side issue.

---

## Notes
- All endpoints require a valid MongoDB ObjectID for operations involving task IDs.
- The `status` field is case-insensitive and automatically converted to lowercase.
- Due dates must be in the future and in `YYYY-MM-DD` format.
- The API uses a MongoDB database, so ensure the database connection is properly configured.
