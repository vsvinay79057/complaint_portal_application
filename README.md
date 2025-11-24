                                                                       Complaint Portal Application

A complete backend application built using Golang, Echo Framework, GORM, JWT Authentication, PostgreSQL, and a clean layered architecture (Controller → Usecase → Repository).

This project provides a structured Complaint Management Portal where users can register, log in, and raise complaints, while the system securely handles authentication and database operations.

🚀 Features
🔐 Authentication & Authorization

User registration & login

JWT-based authentication

Auth middleware for securing routes

🧑‍💼 User Module

Create user (register)

User login

Fetch user profile

📝 Complaint Module

Create complaint

Get all complaints

Get complaint by ID

Update complaint status (if applicable)

🏗️ Clean Project Architecture

Organized into:

controller

usecase

repository

models

middleware

utils

config

Follows best practices for scalable Go backend projects.

📂 Project Structure
complaint_portal_application/
│── main.go
│── go.mod
│── go.sum
│── .env
│
├── config/
│   └── db.go
│
├── controller/
│   ├── user_controller.go
│   └── complaint_controller.go
│
├── models/
│   ├── models.go
│   ├── user.go
│   └── complaint.go
│
├── middleware/
│   └── auth.go
│
├── repository/
│   ├── user_repo.go
│   └── complaint_repo.go
│
├── usecase/
│   ├── user_usecase.go
│   └── complaint_usecase.go
│
└── utils/
    └── helpers.go

🛠️ Technologies Used
Purpose	Tech
Backend Framework	Echo (Golang)
ORM	GORM
Database	PostgreSQL
Authentication	JWT
Config	.env variables
Architecture	Clean MVC + Domain layers
⚙️ Environment Variables

Create a .env file in root:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=complaintdb
JWT_SECRET=yourjwtsecret

▶️ How to Run the Project
1️⃣ Clone the project
git clone https://github.com/yourusername/complaint_portal_application.git
cd complaint_portal_application

2️⃣ Install dependencies
go mod tidy

3️⃣ Run the application
go run main.go


Server starts at:

http://localhost:8080

📡 API Endpoints
🔐 Auth / User
Method	Endpoint	Description
POST	/register	Register new user
POST	/login	Login and get JWT token
GET	/user/profile	Get user profile (protected)
📝 Complaint
Method	Endpoint	Description
POST	/complaints	Create a new complaint
GET	/complaints	Get all complaints
GET	/complaints/:id	Fetch complaint by ID
