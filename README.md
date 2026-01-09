# Moka POS Backend with Go Fiber
Moka POS Backend

## Getting Started

### Prerequisites
- Go 1.19+
- PostgreSQL 12+
- Make (optional)

### Installation

1. Clone the repository
2. Install dependencies:
```bash
$ go mod download
```

3. Set up environment variables (create `.env` file)
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=moka_pos
DB_PORT=5432
```

### Running the Application

```bash
$ go run ./cmd/main.go
```

The application will automatically run all pending database migrations on startup.

## Database Migrations

This project uses `golang-migrate/migrate` for database version control.

### How Migrations Work

Migrations are automatically executed when the application starts. The migration files are located in the `migrations/` directory.

### Creating New Migrations

1. **Create migration files** in the `migrations/` directory following the naming convention:
   ```
   000002_create_posts_table.up.sql
   000002_create_posts_table.down.sql
   ```
   - `up.sql` - Creates/modifies database schema
   - `down.sql` - Rolls back the changes

2. **Example migration files:**

   `migrations/000002_create_posts_table.up.sql`:
   ```sql
   CREATE TABLE IF NOT EXISTS posts (
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
       title VARCHAR(255) NOT NULL,
       content TEXT NOT NULL,
       user_id UUID NOT NULL REFERENCES users(id),
       status VARCHAR(20) NOT NULL DEFAULT 'draft',
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
   );
   
   CREATE INDEX idx_posts_user_id ON posts(user_id);
   ```

   `migrations/000002_create_posts_table.down.sql`:
   ```sql
   DROP TABLE IF EXISTS posts;
   ```

3. **Run migrations** - Restart the application to apply new migrations:
   ```bash
   $ go run ./cmd/main.go
   ```

### Migration Naming Convention

- Start with a 6-digit number (000001, 000002, etc.)
- Use underscores to separate words
- Must have `.up.sql` and `.down.sql` files
- Numbers must be sequential and unique

### Manual Migration (CLI)

You can also run migrations manually using the migrate CLI:

```bash
# Apply all pending migrations
migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

# Rollback one migration
migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down

# Check migration status
migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" version
```



