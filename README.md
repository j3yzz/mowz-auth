# Mowz 🍌 - Authentication Service

---

Mowz Authentication service is a user authentication and authorization system for Mowz. It provides a secure and efficient way to manage user accounts, passwords, and access to Mowz services.

## APIs
Register a new user with JSON request as follows:

```bash
curl --location '127.0.0.1:8082/api/v1/register' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data-raw '{"name": "amirreza","email": "amir@gmail.com","password": "123123124"}'
```
```json
{
    "data": {
        "email": "amirreza@gmail.com",
        "name": "amirreza"
    },
    "success": true
}
```
---
Login into application ang get the JWT Token:
```bash
curl --location '127.0.0.1:8082/api/v1/login' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data-raw '{"email": "amir@gmail.com","password": "123123124"}'
```
```json
{
    "jwt": "<your-jwt-token>",
    "success": true
}
```
---
Get authenticated user with JSON request as follows:

```bash
curl -vvv 127.0.0.1:8082/api/v1/profile -X GET -H 'Authorization: Bearer <your-jwt-token>'
```

```json
{
    "data": {
        "email": "amir@gmail.com",
        "password": "$2a$08$2Hj/yWZiU4W9o1g7x8ZRg.sqlY99oOW67vmbRyOrQlmTX8wVIY5py",
        "name": "amirreza",
        "id": 1
    },
    "success": true
}
```