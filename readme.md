# 📘 AuthTemplateGO – Documentation

## 📌 Base URL

```
http{s}://{APPLICATION_URL}/
````

---

## 🛡️ Auth Routes

### POST `/user/sign-up`

* **Description**: Register a new user.
* **Request JSON**:

```json
{
  "name": "user name",
  "email": "user@example.com",
  "password": "strongPass123!",
  "role": 2
}
````

* **Success Response**:

```json
{
  "success": "Verification link sent"
}
```

* **Status**: `201 Created`

---

### POST `/user/sign-in`

* **Description**: Log in a user.
* **Request JSON**:

```json
{
  "email": "user@example.com",
  "password": "strongPass123!"
}
```

* **Success Response**:

```json
{
  "token": "JWT_TOKEN_HERE"
}
```

* **Status**: `200 OK`

---

## 👥 User Route

### GET `/user/get-me`

* **Description**: Get current authenticated user.
* **Authorization**: Bearer token required.
* **Response**:

```json
{
  "name": "user name",
  "email": "admin@example.com",
  "is_active": true,
  "created_at": "user created time",
  "updated_at": "user updated time",
  "role": {
    "id": 1,
    "name": "role name",
    "created_at": "role created time",
    "updated_at": "role updated time"
  }
}
```

* **Status**: `200 OK`

---

## 🔐 Role Routes

### GET `/role`

* **Description**: List all roles.
* **Authorization**: Bearer token required.
* **Response**:

```json
[
  {
    "id": 1,
    "name": "role name",
    "created_at": "role created time",
    "updated_at": "role updated time"
  }
]
```

---

### POST `/role`

* **Description**: Create a new role.
* **Authorization**: Bearer token required.
* **Request JSON**:

```json
{
  "name": "manager"
}
```

* **Response**:

```json
{
  "id": 1,
  "name": "created role name",
  "created_at": "role created time",
  "updated_at": "role updated time"
}
```

---

### PUT `/role/:id`

* **Description**: Update an existing role name.
* **Authorization**: Bearer token required.
* **Request JSON**:

```json
{
  "name": "superadmin"
}
```

* **Response**:

```json
{
  "success": "role updated"
}
```

* **Status**: `200 OK`

---

### DELETE `/role/:id`

* **Description**: Delete a role by ID.
* **Authorization**: Bearer token required.
* **Response**: `204 No Content`

---

## 🛠️ How to Run

### Create Super User

```shell
go run main.go createsuperuser
```

### Generate Encryption Key

```shell
go run main.go randomtoken
```

### Run Project

```bash
docker compose build
docker compose up
```
