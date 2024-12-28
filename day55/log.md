# Project: To-Do List API

Create a simple RESTful API for a To-Do List application.

## Features to Implement:
CRUD Operations:

1. Create a new task (POST /tasks)
Retrieve all tasks (GET /tasks)
Retrieve a specific task by ID (GET /tasks/:id)
Update a task (PUT /tasks/:id)
Delete a task (DELETE /tasks/:id)

2. Task Model: Each task should have:

ID (integer, auto-incremented)
Title (string)
Description (string)
Completed (boolean)

2B. In-Memory Storage: Use a simple in-memory data structure (e.g., a map) to store tasks.

3. Validation:

Ensure Title is not empty during creation or update.

4. Error Handling:

Return appropriate HTTP status codes for errors (e.g., 404 for task not found).

5. Optional Extras:

Add filtering (e.g., GET /tasks?completed=true).
Add basic authentication for protected routes.
