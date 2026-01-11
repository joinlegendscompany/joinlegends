# Authentication API Documentation

This document outlines the authentication endpoints provided by the API. All endpoints are versioned and follow RESTful principles.

## Base URL
`/v1`

## Overview

The authentication system supports two main flows:
1.  **Stateless JWT Authentication**: Standard JWT token issuance.
2.  **Stateful Session Authentication**: JWT token issuance backed by a server-side session, allowing for session management (e.g., listing active sessions, revoking access).

Additionally, it provides a comprehensive password recovery mechanism.

---

## 1. Standard Authentication

These endpoints provide standard JWT-based authentication without server-side session tracking.

### 1.1. Sign Up
Registers a new user in the system.

- **URL:** `/v1/auth/sign-up`
- **Method:** `POST`
- **Content-Type:** `application/json`

**Request Body:**

| Field | Type | Description | Required |
|---|---|---|---|
| `name` | string | Full name of the user | Yes |
| `email` | string | Valid email address | Yes |
| `password` | string | User's password | Yes |

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Response (200 OK):**

```json
{
  "message": "User registered successfully",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2023-10-27T10:00:00Z"
  },
  "access_token": {
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires": "2023-10-28T10:00:00Z"
  }
}
```

---

### 1.2. Sign In
Authenticates a user using email and password, returning a JWT.

- **URL:** `/v1/auth/sign-in`
- **Method:** `POST`
- **Content-Type:** `application/json`

**Request Body:**

| Field | Type | Description | Required |
|---|---|---|---|
| `email` | string | Registered email address | Yes |
| `password` | string | User's password | Yes |

```json
{
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Response (200 OK):**

```json
{
  "message": "User logged in successfully",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com"
  },
  "access_token": {
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": "24h"
  }
}
```

---

## 2. Session-Based Authentication

These endpoints create a database-backed session record upon login/registration. This allows users to view their active logins and revoke them remotely.

### 2.1. Sign Up with Session
Registers a new user and automatically creates a tracked session.

- **URL:** `/v1/auth/session/sign-up`
- **Method:** `POST`
- **Content-Type:** `application/json`

**Request Body:**
*Same as Standard Sign Up.*

```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "anotherpassword"
}
```

**Response (200 OK):**

Includes `session` details in addition to user and token data.

```json
{
  "message": "User registered successfully",
  "user": { ... },
  "access_token": { ... },
  "session": {
    "id": "session-uuid",
    "user_id": "user-uuid",
    "browser": "Mozilla/5.0 ...",
    "ip": "127.0.0.1",
    "created_at": "...",
    "expires_at": "..."
  }
}
```

---

### 2.2. Sign In with Session
Authenticates a user and creates a tracked session.

- **URL:** `/v1/auth/session/sign-in`
- **Method:** `POST`
- **Content-Type:** `application/json`

**Request Body:**
*Same as Standard Sign In.*

```json
{
  "email": "jane@example.com",
  "password": "anotherpassword"
}
```

**Response (200 OK):**

Returns the session object alongside the standard login response.

---

### 2.3. Validate Session
Checks if the provided session token is valid and active.

- **URL:** `/v1/auth/session/`
- **Method:** `GET`
- **Headers:** `Authorization: Bearer <token>`

**Response (200 OK):**

```json
{
  "message": "Valid session",
  "user_id": "user-uuid"
}
```

**Response (401 Unauthorized):**
If the session is expired or invalid.

---

### 2.4. List All Sessions
Retrieves a list of all active sessions for the currently authenticated user. Useful for "Security" settings pages.

- **URL:** `/v1/auth/session/all`
- **Method:** `GET`
- **Headers:** `Authorization: Bearer <token>`

**Response (200 OK):**

```json
{
  "message": "Sessions retrieved successfully",
  "sessions": [
    {
      "id": "session-uuid-1",
      "browser": "Chrome on Windows",
      "ip": "192.168.1.1",
      "created_at": "...",
      "expires_at": "..."
    },
    {
      "id": "session-uuid-2",
      "browser": "Safari on iPhone",
      "ip": "10.0.0.1",
      "created_at": "...",
      "expires_at": "..."
    }
  ]
}
```

---

### 2.5. Revoke Session
Revokes (deletes) a specific session. This forces the client using that session to log out.

- **URL:** `/v1/auth/session/:sessionId`
- **Method:** `DELETE`
- **Headers:** `Authorization: Bearer <token>`
- **URL Params:** `sessionId` (string) - The ID of the session to revoke.

**Response (200 OK):**

```json
{
  "message": "Session deleted successfully"
}
```

---

## 3. Password Recovery

Flow for users who have forgotten their passwords.

### 3.1. Request Recovery Code
Initiates the recovery process. The server generates a code and sends it to the user's email.

- **URL:** `/v1/auth/recovery/request/:email`
- **Method:** `POST`
- **URL Params:** `email` (string) - The email address of the user.

**Response (200 OK):**

```json
{
  "message": "Recovery code sent to email user@example.com successfully"
}
```

**Response (404 Not Found):**
If the email is not registered.

---

### 3.2. Change Password
Completes the recovery process by setting a new password using the valid recovery code.

- **URL:** `/v1/auth/recovery/change-password`
- **Method:** `POST`
- **Content-Type:** `application/json`

**Request Body:**

| Field | Type | Description | Required |
|---|---|---|---|
| `email` | string | User's email address | Yes |
| `code` | string | The recovery code received via email | Yes |
| `new_password` | string | The new password | Yes |

```json
{
  "email": "john@example.com",
  "code": "123456",
  "new_password": "newSecurePassword2024"
}
```

**Response (200 OK):**

```json
{
  "message": "Password changed successfully"
}
```
