# Chirpy

A Twitter-like REST API built with Go and PostgreSQL. Users can register, post short messages called "chirps" (up to 140 characters), and upgrade to Chirpy Red via webhook integration.

## What it does

- User registration and authentication with JWT access tokens and refresh tokens
- Create, read, and delete chirps — short posts capped at 140 characters with automatic profanity filtering
- Role-based access: Chirpy Red membership unlocked via Polka payment webhook
- Admin endpoints for metrics and data reset

## Tech Stack

- **Go** — standard library `net/http`, no web framework
- **PostgreSQL** — schema managed with [Goose](https://github.com/pressly/goose), queries generated with [sqlc](https://sqlc.dev/)
- **JWT** — access tokens (1h expiry) + refresh tokens
- **Argon2id** — password hashing

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/api/users` | Register a new user |
| `PUT` | `/api/users` | Update email/password (auth required) |
| `POST` | `/api/login` | Login — returns JWT + refresh token |
| `POST` | `/api/refresh` | Exchange refresh token for new access token |
| `POST` | `/api/revoke` | Revoke a refresh token |
| `POST` | `/api/chirps` | Create a chirp (auth required) |
| `GET` | `/api/chirps` | List all chirps (supports `sort` and `author_id` query params) |
| `GET` | `/api/chirps/{id}` | Get a single chirp |
| `DELETE` | `/api/chirps/{id}` | Delete a chirp (auth required, owner only) |
| `POST` | `/api/polka/webhooks` | Polka webhook to upgrade user to Chirpy Red |
| `GET` | `/api/healthz` | Health check |
| `GET` | `/admin/metrics` | Request count metrics |
| `POST` | `/admin/reset` | Reset database (dev only) |

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL
- [Goose](https://github.com/pressly/goose) for migrations

### Setup

1. **Clone the repo**
   ```bash
   git clone https://github.com/viniciusLambert/bootdevCourseBackendGo.git
   cd bootdevCourseBackendGo
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Create the database**
   ```bash
   createdb chirpy
   ```

4. **Run migrations**
   ```bash
   goose -dir sql/schema postgres "your-db-url" up
   ```

5. **Configure environment** — create a `.env` file:
   ```env
   DB_URL=postgres://user:password@localhost:5432/chirpy?sslmode=disable
   PLATFORM=dev
   JWT_TOKEN=your-secret-key
   POLKA_SECRET=your-polka-api-key
   ```

6. **Run the server**
   ```bash
   go run .
   ```

   The server starts on `http://localhost:8080`.

## Example Usage

```bash
# Register
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret"}'

# Login
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret"}'

# Post a chirp
curl -X POST http://localhost:8080/api/chirps \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{"body":"Hello, Chirpy!"}'
```
