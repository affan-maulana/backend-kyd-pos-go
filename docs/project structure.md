# Project Structure - Go Fiber JWT

Dokumentasi struktur folder project dengan implementasi modular architecture, service-repository pattern, dan pemisahan request/response.

## Folder Structure

```
.
├── cmd/
│   └── main.go                          # Entry point aplikasi
│
├── config/
│   ├── db.go                            # Konfigurasi koneksi database
│   └── loadEnv.go                       # Load environment variables
│
├── routes/
│   ├── routes.go                        # Main router setup + healthchecker + 404 handler
│   ├── auth.routes.go                   # Auth module routes
│   └── user.routes.go                   # User module routes
│
├── internal/
│   ├── middleware/
│   │   └── deserialize-user.go          # JWT middleware untuk authenticate user
│   │
│   └── modules/                         # Folder utama pengelompokan fitur/domain
│       ├── entity/                      # Semua model database global
│       │   └── user.go                  # User entity model
│       │
│       ├── auth/                        # Auth module
│       │   ├── auth.controller.go       # Auth controller - handle HTTP request/response
│       │   ├── auth.service.go          # Auth service - business logic
│       │   ├── auth.repository.go       # Auth repository - database operations
│       │   └── dto/
│       │       ├── auth.request.go      # Request DTO (SignUpInput, SignInInput)
│       │       └── auth.response.go     # Response DTO (UserResponse, FilterUserRecord)
│       │
│       └── user/                        # User module
│           ├── user.controller.go       # User controller
│           ├── user.service.go          # User service
│           ├── user.repository.go       # User repository
│           └── dto/
│               ├── user.request.go      # Request DTO
│               └── user.response.go     # Response DTO
│
├── pkg/                                 # Utilities umum yang dipakai lintas modul
│   ├── response/                        # Response wrapper/formatter
│   └── hashing/                         # Password hashing utilities
│
├── docker-compose.yml                   # Docker compose untuk database
├── app.env                              # Environment variables
├── go.mod                               # Go module dependencies
├── go.sum                               # Go module checksums
├── Makefile                             # Build scripts
└── README.md                            # Project documentation
```

## Architecture Pattern

### 1. Modular Architecture
Setiap fitur/domain (auth, user, master, dll) memiliki folder sendiri di `internal/modules/` dengan semua layer di dalamnya.

### 2. Service-Repository Pattern
```
Controller (HTTP Handler)
    ↓
Service (Business Logic)
    ↓
Repository (Database Operations)
    ↓
Entity (Database Model)
```

### 3. Request & Response Separation
- **DTO (Data Transfer Object)**: Struct khusus untuk request dan response di folder `dto/`
- **Entity**: Model database di folder `entity/`

## File Naming Convention

### Controllers
- File: `{module}.controller.go`
- Package: `controllers` atau sesuai modul
- Function: `SignUpUser()`, `GetMe()`, dll

### Services
- File: `{module}.service.go`
- Package: `service`
- Struct: `{Module}Service`
- Function: `(s *{Module}Service) Method()`

### Repositories
- File: `{module}.repository.go`
- Package: `repository`
- Interface: `{Module}Repository`
- Struct: `{module}Repository`

### DTOs
- File: `{module}.request.go` atau `{module}.response.go`
- Package: `dto`
- Struct: `{Action}Request`, `{Action}Response`

### Models/Entities
- File: `{module}.go`
- Package: `entity`
- Struct: `{Module}` (User, Product, dll)

## Routing Structure

Routes diorganisir per modul untuk maintainability:
- `routes.go`: Main router dan global endpoints (healthchecker, 404 handler)
- `auth.routes.go`: Auth module routes
- `user.routes.go`: User module routes
- Dst untuk module lain

## Dependency Flow

```
main.go
  ↓
routes.go (SetupRoutes)
  ↓
{module}.routes.go (SetupAuthRoutes, SetupUserRoutes)
  ↓
{module}.controller.go
  ↓
{module}.service.go
  ↓
{module}.repository.go
  ↓
config/db.go (GORM DB)
```

## Adding New Module

Untuk menambah modul baru (misal: `product`):

1. Buat folder di `internal/modules/product/`
2. Buat file:
   - `product.controller.go`
   - `product.service.go`
   - `product.repository.go`
   - `dto/product.request.go`
   - `dto/product.response.go`
3. Tambah entity di `internal/modules/entity/product.go`
4. Buat route file di `routes/product.routes.go`
5. Setup route di `routes.go`:
   ```go
   SetupProductRoutes(micro)
   ```
6. Update migration di `config/db.go`:
   ```go
   DB.AutoMigrate(&entity.Product{})
   ```

## Transaction & Commit

Database operations yang critical menggunakan transaction:
```go
tx := config.DB.Begin()
defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
}()

// database operations
result := tx.Create(&newUser)

if result.Error != nil {
    tx.Rollback()
    return c.Status(fiber.StatusBadRequest).JSON(...)
}

if err := tx.Commit().Error; err != nil {
    tx.Rollback()
    return c.Status(fiber.StatusInternalServerError).JSON(...)
}
```
