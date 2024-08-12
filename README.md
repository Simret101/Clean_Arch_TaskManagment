# Task Management Clean Architecture

Task Manager is a web application built with Go, designed to manage tasks efficiently. The application uses the Gin framework for HTTP routing, JWT for authentication, and secure password hashing. The project is organized with a clean, modular architecture to ensure maintainability and scalability.

## Features

- User Registration and Authentication (JWT-based)
- Task Creation, Retrieval, Update, and Deletion
- Role-based Access Control (Admin and User roles)
- Secure Password Storage (Hashing)
- Middleware for Authentication and Role Verification

## Folder Structure

```plaintext
task-manager/
├── api/
│   ├── middleware/
│   │   ├── admin.go
│   │   └── auth.go
│   ├── router/
│   │   └── router.go
│   └── controller/
│       ├── task_controller.go
│       └── user_controller.go
├── Domain/
│   ├── task.go
│   └── user.go 
├── cmd/
│   └── main.go
├── config/
│   ├── config.go
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repository/
│   ├── task_repository.go
│   └── user_repository.go
├── usecase/
│   ├── task_usecases.go
│   └── user_usecases.go
```

### Components

- **api/**: Contains HTTP controllers, middleware, and router definitions.
- **Domain/**: Contains domain models (entities) like `Task` and `User`.
- **cmd/**: Entry point of the application.
- **config/**: Configuration settings, JWT, and password services.
- **Repository/**: Data access layer for interacting with the database or storage.
- **usecase/**: Business logic for task and user management.

## API Endpoints

### User Endpoints

- **POST `/register`**: Register a new user.
- **POST `/login`**: Authenticate a user and receive a JWT token.

### Task Endpoints

- **GET `/tasks`**: Retrieve all tasks (admin) or tasks for the logged-in user.
- **GET `/tasks/:id`**: Retrieve a specific task by ID.
- **POST `/tasks`**: Create a new task.
- **PUT `/tasks/:id`**: Update a specific task by ID.
- **DELETE `/tasks/:id`**: Delete a specific task by ID.

## Middleware

### Authentication Middleware

- Verifies the JWT token sent in the `Authorization` header.
- Ensures that only authenticated users can access protected routes.

### Admin Middleware

- Restricts access to certain routes to users with the admin role.
- Ensures that only admins can manage all tasks.

## Models

### User Model

```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"`
}
```

### Task Model

```go
type Task struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" bson:"title"`
    Description string             `json:"description" bson:"description"`
    DueDate     time.Time          `json:"duedate" bson:"duedate"`
    Status      TaskStatus         `json:"status" bson:"status"`
    UserID      int                `json:"userID" bson:"userID"`
}
```

## Installation

1. **Clone the repository:**

   ```sh
   git clone [https://github.com/Simret101/Clean_Arch_TaskManagment]
   ```

2. **Navigate to the project directory:**

   ```sh
   cd task-manager
   ```

3. **Install dependencies:**

   ```sh
   go mod tidy
   ```

4. **Run the application:**

   ```sh
   go run cmd/main.go
   ```

## Configuration

Configuration settings are managed in the `config/` directory. Modify `config.go` to update application settings such as database connections, JWT secret, and more.

## Testing

Unit and integration tests should be added to ensure the application functions as expected. You can run the tests using the following command:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature/bugfix.
3. Make your changes and commit them with clear and descriptive messages.
4. Push your changes to your fork.
5. Open a pull request describing your changes.



## Postman Documentation

For detailed API documentation, refer to the Postman collection: [https://documenter.getpostman.com/view/37289771/2sA3rzKsPp]

