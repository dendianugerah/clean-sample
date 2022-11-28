# Clean Architecture

This is a sample project to demonstrate how to use Clean Architecture.

File structure:

```
├── main
│   ├── config
│   │   ├── connection.go
│   ├── docs: swaggo docs
│   ├── domain
│   │   ├── model: this is where the domain model is defined
│   │   │   ├── example.go
│   │   ├── repository: interface for CRUD operations
│   │   │   ├── example.go
│   ├── inftastructure
|   |   ├── example.go
│   ├── interface: implementation of repository interface
│   |   ├── handler
│   │   │   ├── example.go
│   |   ├── route: about routing
│   │   │   ├── example.go
│   ├── usecase: this is where the business logic is defined
│   │   ├── example.go

```
