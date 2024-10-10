# Go Authentication Service

This project is a simple user authentication service built in Go (Golang) using the Gorilla Mux router for HTTP routing, GORM for database interactions, and bcrypt for password hashing. It allows users to sign up and log in, securely storing their credentials in a MySQL database.

## Features

- **User Signup**: Users can create an account with a unique username, email, and password. Passwords are hashed for security.
- **User Login**: Users can log in by verifying their username and password against the stored credentials.
- **Database Interaction**: Utilizes GORM for seamless interaction with a MySQL database, including automatic table migration.

## Technologies Used

- **Go (Golang)**: The programming language used to build the service.
- **Gorilla Mux**: A powerful HTTP router and URL matcher for building Go web applications.
- **GORM**: An ORM library for Go that simplifies database operations.
- **MySQL**: The database used to store user information.
- **bcrypt**: A library for hashing passwords securely.

## Setup Instructions

1. **Install Go**: Make sure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

2. **Install MySQL**: Ensure that MySQL is installed and running on your local machine. Create a database named `gorm` if it doesn't exist.

3. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/repo-name.git
   cd repo-name

