# Go gRPC Microservice - Employee API

This project is a simple gRPC microservice written in Go that provides employee data from a MySQL database. It implements a clean architecture using layered structure: handler → service → repository.

## 🧱 Tech Stack

- Language: Go (Golang)
- Framework: gRPC
- Database: MySQL
- ORM: GORM
- Protobuf: Protocol Buffers (v3)
- Tools: grpcurl, protoc, protoc-gen-go, protoc-gen-go-grpc

## 📁 Project Structure

go-grpc/
├── cmd/
│   └── server/            # Main entry point
│       └── main.go
├── config/                # Database config
│   └── db.go
├── internal/
│   └── employee/          # Business logic
│       ├── handler.go
│       ├── service.go
│       └── repository.go
├── proto/                 # Proto files & generated code
|   └── employee/
│       ├── employee.pb.go
│       └── employee_grpc.pb.go
│   └── employee.proto
├── go.mod
└── README.md

## ⚙️ Prerequisites

- Go 1.18+
- MySQL
- protoc (Protocol Buffers compiler)
- Plugins (install via terminal):
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

- Optional: gRPC CLI tool
  - Windows: choco install grpcurl
  - macOS: brew install grpcurl

## 🛠️ Setup

1. Clone the Repository

  git clone https://your-bitbucket-or-github-url/go-grpc.git
  cd go-grpc

2. Configure MySQL

Create a database and table:

  CREATE DATABASE go_grpc_db;

  USE go_grpc_db;

  CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    position VARCHAR(100)
  );

  INSERT INTO employees (name, position) VALUES
  ('Bayu Indrawan', 'Golang Developer'),
  ('Alice', 'QA Engineer');

Update DB credentials in config/db.go.

3. Generate Proto Files

  protoc --go_out=. --go-grpc_out=. proto/employee.proto

4. Run the Server

  go run cmd/server/main.go

Expected output:

  gRPC server running on port 50051...

## 🚀 API Usage

1. Using grpcurl

  grpcurl -plaintext localhost:50051 employee.EmployeeService/GetEmployees

Sample Response:

  {
    "employees": [
      {
        "id": 1,
        "name": "Bayu Indrawan",
        "position": "Golang Developer"
      },
      {
        "id": 2,
        "name": "Alice",
        "position": "QA Engineer"
      }
    ]
  }

## 🧪 Debug with VS Code

Add file: .vscode/launch.json

  {
    "version": "0.2.0",
    "configurations": [
      {
        "name": "Launch gRPC Server",
        "type": "go",
        "request": "launch",
        "mode": "auto",
        "program": "${workspaceFolder}/cmd/server"
      }
    ]
  }

Then press F5 to start debugging.

## 📦 Future Improvements

- Add CreateEmployee, UpdateEmployee, DeleteEmployee
- Add REST API using gRPC Gateway
- Add Dockerfile & docker-compose support
- Add authentication / JWT

## 👨‍💻 Author

Bayu Indrawan  
📧 bayuindrawan95@gmail.com

