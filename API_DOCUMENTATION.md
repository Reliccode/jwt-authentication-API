
---

# JWT Authentication API v1.0

The JWT Authentication API provides endpoints for user signup, login, and validation utilizing JSON Web Tokens (JWT) for authentication.

## Base URL

```url
http://your-api-domain.com
```

## Endpoints

### 1. User Signup

- **URL:** `/signup`
- **Method:** `POST`
- **Description:** Creates a new user account.
- **Request Body:**
  
  ```json
  {
      "Email": "string",
      "Password": "string"
  }
  ```
- **Responses:**
  - **Success Response:** 
    - **Status Code:** `200 OK`
  - **Error Responses:**
    - **Status Code:** `400 Bad Request`
      - **Content:**
        ```json
        {
            "error": "Failed to read body"
        }
        ```
        OR
        ```json
        {
            "error": "Failed to hash password"
        }
        ```
        OR
        ```json
        {
            "error": "Failed to create user"
        }
        ```

### 2. User Login

- **URL:** `/login`
- **Method:** `POST`
- **Description:** Authenticates user and generates a JWT token.
- **Request Body:**
  
  ```json
  {
      "Email": "string",
      "Password": "string"
  }
  ```
- **Responses:**
  - **Success Response:** 
    - **Status Code:** `200 OK`
      - **Content:**
        ```json
        {
            "token": "string"
        }
        ```
  - **Error Responses:**
    - **Status Code:** `400 Bad Request`
      - **Content:**
        ```json
        {
            "error": "Invalid email or password"
        }
        ```
        OR
        ```json
        {
            "error": "Failed to create token"
        }
        ```

### 3. Validate User

- **URL:** `/validate`
- **Method:** `GET`
- **Description:** Validates user token.
- **Request Headers:**
  - `Authorization: Bearer <token>`
- **Responses:**
  - **Success Response:** 
    - **Status Code:** `200 OK`
      - **Content:**
        ```json
        {
            "message": "user details"
        }
        ```
  - **Error Response:**
    - **Status Code:** `401 Unauthorized`

## Authentication

This API uses JWT tokens for authentication. The JWT token is generated upon successful login and must be included in the request headers for authorized endpoints.

## Error Handling

- **400 Bad Request:** Returned when there is an issue with the request body or processing the request.
- **401 Unauthorized:** Returned when the request lacks valid authentication credentials or token is invalid/expired.

## Release Notes

### Version 1.0 (March 20, 2024)

- Initial release of the JWT Authentication API.
- Features:
  - User signup endpoint (`/signup`).
  - User login endpoint (`/login`) with JWT token generation.
  - User validation endpoint (`/validate`) for token validation.
- Authentication using JWT tokens.
- Error handling for bad requests and unauthorized access.

---
