Project Description
This project is a simple user authentication service built in Go (Golang) using the Gorilla Mux router for HTTP routing, GORM for database interactions, and bcrypt for password hashing. It allows users to sign up and log in, securely storing their credentials in a MySQL database.

Features
User Signup: Allows users to create an account with a unique username, email, and password. Passwords are hashed for security.
User Login: Authenticates users by checking the entered username and password against the stored credentials.
Database Interaction: Utilizes GORM for seamless interaction with a MySQL database, including automatic table migration.
Technologies Used
Go (Golang): The programming language used to build the service.
Gorilla Mux: A powerful HTTP router and URL matcher for building Go web applications.
GORM: An ORM library for Go that simplifies database operations.
MySQL: The database used to store user information.
bcrypt: A library for hashing passwords securely.
