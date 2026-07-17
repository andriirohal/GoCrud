# 🎓 GoCrud

A REST API built with Go, Gin, and pgx, featuring a clean layered architecture and PostgreSQL as the persistence layer.

---

## ✨ Features

- Get all student
- Get student by ID
- Create new student
- Update student by ID
- Delete student by ID

---

## 🛠️ Tech Stack

- Go — fast, compiled language
- Gin — lightweight Go web framework
- PostgreSQL — relational database
- pgx — Go driver for PostgreSQL

---

## 📁 Structure

```
cmd/
├── main.go          # entry point & routes
├── controllers/      # HTTP handlers
├── services/         # business logic & queries
├── helpers/           # scanning & validation
├── models/            # structs
└── db/                 # DB connection pool
```