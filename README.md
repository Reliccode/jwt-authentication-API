# jwt-authentication-API

Welcome to our JWT Authentication project! This repository serves as a demonstration of how to implement JSON Web Token (JWT) authentication in a Go web application, making your applications secure and user-friendly. We are Glen Maritim and Naniwet Maritim.

## 📑 Table of Contents

- [Overview](#-overview)
- [Features](#-features)
- [Requirements](#️️-requirements)
- [Installation](#️️-installation)
- [Usage](#️️-usage)
- [API Endpoints](#️️-api-endpoints)
- [License](#️️-license)
- [Acknowledgments](#️️-acknowledgments)
- [Authors](#️️-authors)
- [Contributing](#️️-contributing)

## 🔭 Overview

Are you ready to supercharge your Go web applications with robust authentication? Look no further! Our project showcases the power of JWT authentication, providing features such as user registration, login, and protected routes. We've leveraged the flexibility of GORM as the ORM library and the blazing-fast Gin framework to create a seamless user experience.

## ✨ Features

- **User Registration:** Allow users to sign up effortlessly by providing a username and password.
- **User Login:** Enable users to access their accounts securely with their credentials.
- **Authentication Middleware:** Safeguard your routes with JWT token verification, ensuring only authenticated users gain access.
- **Password Hashing:** Protect user passwords using bcrypt hashing, making them virtually impossible to crack.
- **JWT Generation and Validation:** Generate JWT tokens upon successful login and validate them for authenticated routes.
- **Database Integration:** Seamlessly interact with a PostgreSQL database using GORM, ensuring efficient storage and retrieval of user information.

## 📋Requirements

To get started, ensure you have the following prerequisites installed:

- Go (version 1.15 or higher)
- PostgreSQL (version 9.6 or higher)

## 🛠️ Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Reliccode/jwt-authentication-API.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd jwt-authentication-API
   ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Set up the PostgreSQL database:**

   Execute the SQL script located in `database.sql` to initialize the required tables.

5. **Configure the database:**

   Update the database configuration in `config/config.go` with your PostgreSQL database credentials.

6. **Run the application:**

   ```bash
   go run main.go
   ```

Voila! Your application is now up and running on `http://localhost:8080`.

## 💻 Usage

1. **Register a new user:**

   Send a `POST` request to `/api/register` with the following JSON payload:

   ```json
   {
       "username": "yourusername",
       "password": "yourpassword"
   }
   ```

2. **Log in:**

   Log in with the registered user credentials by making a `POST` request to `/api/login` with the same JSON payload.

3. **Access protected routes:**

   Upon successful login, you will receive a JWT token in the response. Use this token to access protected routes by including it in the `Authorization` header of your requests:

   ```bash
   Authorization: Bearer your-jwt-token
   ```

## 📝 API Endpoints

- `POST /api/register`: Register a new user.
- `POST /api/login`: Log in and receive a JWT token.
- `GET /api/user/profile`: Get the profile of the currently logged-in user (protected route).

## 📜 License

This project is licensed under the MIT License. For more details, check out the [LICENSE](LICENSE) file.

## 🙌 Acknowledgments

A big shoutout to the fantastic tools and libraries that made this project possible:

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM - The fantastic ORM library for Golang](https://gorm.io/)
- [bcrypt - Password hashing library for Go](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [jwt-go - Golang implementation of JSON Web Tokens (JWT)](https://github.com/dgrijalva/jwt-go)

- [African Leadership Group](https://algroup.org/) & [ALX Africa](https://www.alxafrica.com/) for the life-changing opportunity to become worldclass engineers and thinkers.

## 🖋️ Authors

- Naniwet Maritim ([@Reliccode](https://github.com/Reliccode))
- Glen Maritim ([@ExitMirth](https://github.com/ExitMirth))

## 🤝 Contributing

Contributions are more than welcome! If you have any ideas, suggestions, or bug fixes, feel free to open an issue or submit a pull request. Let's make this project even better together!

---
