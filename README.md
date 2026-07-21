# 🧩 Go CRUD

A REST API built with Go, Gin, and PostgreSQL, featuring a clean layered architecture with PostgreSQL as the persistence layer.

---

## ✨ Features

- Get all students
- Get student by ID
- Create new student
- Update student by ID
- Delete student by ID
- Type-safe with Go
- PostgreSQL persistence

---

## 🛠️ Tech Stack

- Go
- Gin
- PostgreSQL
- pgx (PostgreSQL driver)

---

## 📁 Project Structure

```txt
├── controllers/
├── db/
├── helpers/
├── middlewares/
├── models/
├── services/
└── main.go
```

---

## 🗄️ Database Setup

Create a PostgreSQL database:

```sql
CREATE DATABASE students_db;
```

Create the students table:

```sql
CREATE TABLE students (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL CHECK (char_length(trim(name)) > 0),
  age INTEGER NOT NULL CHECK (age BETWEEN 0 AND 100),
  grade NUMERIC(5, 2) NOT NULL CHECK (grade BETWEEN 0 AND 100)
);
```

Install PostgreSQL driver:

```console
go get github.com/jackc/pgx/v5
```

Install environment variables package:

```console
go get github.com/joho/godotenv
```

Create a `.env` file:

```env
PGUSER=postgres
PGPASSWORD=password
PGHOST=localhost
PGPORT=5432
PGDATABASE=students_db
```

---

## ⚙️ Installation

Install dependencies:

```console
go mod tidy
```

Start development server:

```console
go run main.go
```

Build the project:

```console
go build
```

---

## 📄 Example Request

### Create Student

```json
{
  "name": "Andrii",
  "age": 18,
  "grade": 95
}
```

Example response:

```json
{
  "id": "c7a2d9f4-8c1d-4a5a-8b67-1d2e3f4a5b6c",
  "name": "Andrii",
  "age": 18,
  "grade": 95
}
```
