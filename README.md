
---

## 📝 Go Task Manager API

A clean, secure RESTful API built with **Go** and the **Echo framework** for user authentication and task management. Implements **JWT-based authentication** and follows clean architecture principles.

---

### 🚀 Features

* ✅ User Registration & Login
* 🔐 JWT Authentication Middleware
* 📋 Create, Read, Update, and Delete Tasks (CRUD)
* ✅ Mark Tasks as Completed
* 📂 Structured Route Grouping with Middleware
* 🧼 Clean and Modular Codebase

---

### 📁 API Endpoints

#### 🔐 Authentication Routes

| Method | Endpoint    | Description          |
| ------ | ----------- | -------------------- |
| POST   | `/register` | Register a new user  |
| POST   | `/login`    | Login and obtain JWT |

---

#### ✅ Protected Task Routes (`/api` prefix)

> **Note:** All these routes require a valid **Bearer JWT token** in the `Authorization` header.

| Method | Endpoint                   | Description            |
| ------ | -------------------------- | ---------------------- |
| POST   | `/api/tasks`               | Create a new task      |
| GET    | `/api/tasks`               | List all tasks         |
| GET    | `/api/tasks/:id`           | Get task by ID         |
| PUT    | `/api/tasks/:id`           | Update a task          |
| PUT    | `/api/tasks/:id/completed` | Mark task as completed |
| DELETE | `/api/tasks/:id`           | Delete a task by ID    |

---

### 🔐 JWT Authentication

Include the JWT token in the `Authorization` header for all protected routes:

```
Authorization: Bearer <your-token>
```

---

### 📦 Project Structure (Simplified)

```
go-task-manager/
├── cmd/
│   └── main.go
├── internal/
│   ├── infrastructure/
│   │   └── server/
│   │       ├── handler/
│   │       ├── middlewares/
│   │       └── routes/
│   └── utils/
│       └── jwtutil/
```

---

### 🛠️ Getting Started

```bash
# Clone the repository
git clone https://github.com/bhushan-aruto/go-task-manager.git
cd go-task-manager

# Set environment variables (example)
export JWT_SECRET="your_secret_key"
export DATABASE_URL="your_db_url"
export SERVER_ADDRESS="localhost:8080"
export DATABASE_NAME="task_manager"

# Run the application
go run cmd/main.go
```

---

### 📚 Tech Stack

* **Go (Golang)**
* **Echo Web Framework**
* **JWT (JSON Web Tokens)**
* **Clean Architecture Pattern**

---

### 🙌 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

---


