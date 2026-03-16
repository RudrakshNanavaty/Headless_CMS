# Headless CMS (Go, Gin, GORM)

A lightweight, opinionated **headless CMS API** built with **Go**, **Gin**, and **GORM**.

It provides:

- JWT-based authentication
- Role-based access control (Super Admin / Admin / User)
- Generic CRUD APIs to define content types, attributes, content entries, child entities, and data (including file uploads to AWS S3)

---

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture Overview](#architecture-overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Environment Variables](#environment-variables)
  - [Run Locally](#run-locally)
  - [Run with Docker](#run-with-docker)
- [API Overview](#api-overview)
  - [Base Paths](#base-paths)
  - [Authentication](#authentication)
  - [Super Admin Endpoints](#super-admin-endpoints)
  - [CMS Endpoints](#cms-endpoints)
- [Project Structure](#project-structure)
- [Database & Migrations](#database--migrations)
- [Security Notes](#security-notes)
- [Extending the CMS](#extending-the-cms)
- [License](#license)

---

## Features

- **JWT Authentication**
  - Sign up, login, logout.
  - JWT stored in an `Authorization` cookie (24h expiry).
  - `RequireAuth` middleware protects authenticated routes.

- **Role-Based Access Control**
  - Roles defined in `config/roles`:
    - `SuperAdmin`
    - `Admin`
    - `User`
  - **Super Admin**:
    - Seeded at startup from environment variables.
    - Manages admins and users (create, list, promote, demote, delete).
  - **Admin**:
    - Full control over CMS entities (types, attributes, content, children, data).
  - **User**:
    - Authenticated access with restricted capabilities.

- **CMS CRUD**
  - **Types**: High-level content types (e.g. `BlogPost`, `Product`).
  - **Attributes**: Fields / schema for each type.
  - **Content**: Instances of a type (e.g. a blog post record).
  - **Children**: Child entities related to content items.
  - **Data**: Arbitrary structured data, including file uploads.
  - File uploads are handled via AWS S3.

- **Dockerized**
  - Multi-stage `Dockerfile`.
  - `docker compose` support for local development.

---

## Tech Stack

- **Language**: Go 1.22
- **Web Framework**: [`gin-gonic/gin`](https://github.com/gin-gonic/gin)
- **ORM**: [`gorm.io/gorm`](https://gorm.io) with Postgres driver
- **Auth**: [`golang-jwt/jwt/v5`](https://github.com/golang-jwt/jwt)
- **Config**: [`joho/godotenv`](https://github.com/joho/godotenv)
- **Database**: PostgreSQL
- **File Storage**: AWS S3 (`aws-sdk-go`)
- **Containerization**: Docker, Docker Compose

---

## Architecture Overview

The application entrypoint is `main.go`:

- Loads environment variables and connects to the database:
  - `initializers.LoadEnvVariables()`
  - `initializers.ConnectToDB()`
  - `initializers.CreateSuperAdminIfNotExists()`
- Sets up a Gin HTTP server:
  - Base group: `/api/v1`
  - Health route: `GET /api/v1/` → `{"message": "Hello World!"}`
- Registers route groups:
  - `/api/v1/super-admin` (Super admin management; protected by `RequireSuperAdmin`)
  - `/api/v1/auth` (Authentication)
  - `/api/v1/cms` (CMS CRUD; protected by `RequireAuth`, with additional `RequireAdmin` on admin-only actions)

Entities and controllers live under:

- `types/` (GORM models)
- `controllers/auth/`
- `controllers/super_admin/`
- `controllers/cms_crud/`

---

## Getting Started

### Prerequisites

- **Go** `>= 1.22`
- **PostgreSQL** (local or remote)
- **Docker & Docker Compose** (optional but recommended for local dev)
- **AWS S3** bucket & credentials (for file uploads)

### Environment Variables

1. Copy the example environment file:

   ```bash
   cp example.env .env
   ```

2. Fill in the `.env` file:

| Variable               | Description                           | Example            |
| ---------------------- | ------------------------------------- | ------------------ |
| `DB_USER`              | Postgres username                     | `postgres`         |
| `DB_PASSWORD`          | Postgres password                     | `password`         |
| `DB_PORT`              | Postgres port                         | `5432`             |
| `DB_NAME`              | Postgres database name                | `headless_cms`     |
| `DB_HOST`              | Postgres host                         | `localhost`        |
| `DB_SSLMODE`           | SSL mode (`disable`, `require`, etc.) | `disable`          |
| `PORT`                 | App port (inside container / app)     | `4000`             |
| `JWT_SECRET`           | Secret key for signing JWTs           | `super-secret-key` |
| `SUPER_ADMIN_USERNAME` | Initial super admin username          | `root`             |
| `SUPER_ADMIN_PASSWORD` | Initial super admin password          | `change-me`        |
| `AWS_BUCKET_NAME`      | S3 bucket name                        | `my-cms-bucket`    |
| `AWS_BUCKET_REGION`    | S3 region                             | `us-east-1`        |
| `AWS_ACCESS_KEY`       | AWS access key ID                     | `AKIA...`          |
| `AWS_SECRET`           | AWS secret access key                 | `abc123...`        |

> The super admin credentials are used at startup to ensure a Super Admin user exists.

### Run Locally

1. Install dependencies:

   ```bash
   go mod tidy
   ```

2. Make sure PostgreSQL is running and accessible with your `.env` configuration.

3. Run the application:

   ```bash
   go run main.go
   ```

4. The server should be available at:
   - `http://localhost:4000/api/v1/`
   - Health check: `GET /api/v1/` → `{"message":"Hello World!"}`

### Run with Docker

1. Ensure `.env` exists in the project root.

2. Start the app using Docker Compose:

   ```bash
   docker compose up --build
   ```

3. Access the API at:
   - `http://localhost:4000/api/v1/`

For more Docker-specific notes, see `README.Docker.md`.

---

## API Overview

### Base Paths

- Base URL: `http://localhost:4000`
- API prefix: `/api/v1`
- Health: `GET /api/v1/` → `{"message": "Hello World!"}`

### Authentication

Routes: `/api/v1/auth`

- `POST /auth/signup`
  - Body:

    ```json
    {
      "username": "john",
      "password": "secret"
    }
    ```

  - Creates a new user with role `User`.

- `POST /auth/login`
  - Body:

    ```json
    {
      "username": "john",
      "password": "secret"
    }
    ```

  - Verifies credentials and sets an `Authorization` cookie with a JWT (`SameSite=Lax`, `HttpOnly`).
  - Token expiration: 24 hours.

- `POST /auth/logout`
  - Clears the `Authorization` cookie.

- `GET /auth/auth`
  - Protected by `RequireAuth`.
  - Returns `{"message": "Authenticated"}` if the JWT is valid.

### Super Admin Endpoints

Routes: `/api/v1/super-admin`
Protected by `middlewares.RequireSuperAdmin`.

- `POST /super-admin/register-admin`
  - Register a new Admin user.

- `GET /super-admin/list/admins`
  - List all admins.

- `GET /super-admin/list/users`
  - List all normal users.

- `GET /super-admin/account/:id`
  - Get account info (id, username, role_type) for a user.

- `PUT /super-admin/update/:id`
  - Update admin account details.

- `PATCH /super-admin/promote/:id`
  - Promote a user to admin.

- `PATCH /super-admin/demote/:id`
  - Demote an admin to user.

- `DELETE /super-admin/delete/:id`
  - Delete an admin user
    (cannot delete a `SuperAdmin`).

### CMS Endpoints

Routes: `/api/v1/cms`
Group protected by `middlewares.RequireAuth`; specific routes also use `middlewares.RequireAdmin` for admin-only operations.

#### Attributes

- `GET /cms/attributes`
- `POST /cms/attribute` (Admin only)
- `GET /cms/attribute/:id`
- `PUT /cms/attribute/:id` (Admin only)
- `DELETE /cms/attribute/:id` (Admin only)

#### Content

- `GET /cms/contents`
- `POST /cms/content` (Admin only)
- `GET /cms/content/:id`
- `PUT /cms/content/:id` (Admin only)
- `DELETE /cms/content/:id` (Admin only)

#### Data

- `GET /cms/alldata`
- `POST /cms/data` (Admin only)
- `GET /cms/data/:id`
- `PUT /cms/data/:id` (Admin only)
- `DELETE /cms/data/:id` (Admin only)
- `PATCH /cms/data/:id` (Admin only)
  - File upload endpoint (handled by `cms_crud.UploadFile` and `utils/fileupload.go`).
  - Files are uploaded to AWS S3 using the configuration from `initializers/awssession.go`.

#### Types

- `GET /cms/types`
- `POST /cms/type` (Admin only)
- `GET /cms/type/:id`
- `PUT /cms/type/:id` (Admin only)

#### Children

- `GET /cms/children`
- `POST /cms/child` (Admin only)
- `GET /cms/child/:id`
- `PUT /cms/child/:id` (Admin only)
- `DELETE /cms/child/:id` (Admin only)

> Refer to the `types/` package to see the GORM models and JSON structure for `User`, `Type`, `Attribute`, `Content`, `Child`, and `Data`.

---

## Project Structure

```text
.
├── main.go                      # Gin server setup, route groups
├── go.mod / go.sum              # Go module and dependencies
├── Dockerfile                   # Multi-stage Docker build
├── compose.yaml                 # Docker Compose config
├── example.env                  # Environment variable template
├── README.md                    # Project documentation
├── README.Docker.md             # Docker usage notes
├── controllers/
│   ├── auth/                    # SignUp, Login, Logout
│   ├── cms_crud/                # CRUD for Types, Content, Data, Children, Attributes
│   └── super_admin/             # Admin & user management
├── routes/                      # Route registrations (auth, cms, super-admin)
├── middlewares/                 # Auth and role-based middlewares
├── initializers/                # Env loading, DB, AWS session, super admin seeding
├── types/                       # GORM models for users and CMS entities
├── config/
│   └── roles/                   # Role definitions (User, Admin, SuperAdmin)
└── utils/
    └── fileupload.go            # Helpers for file upload to S3
```

---

## Database & Migrations

- Database connection is configured in `initializers/database.go` using PostgreSQL.
- Environment variables are loaded in `initializers/load_environment_variables.go`.
- `initializers.ConnectToDB()` creates the DB connection and (typically) runs GORM auto-migrations.
- `initializers.CreateSuperAdminIfNotExists()` uses `SUPER_ADMIN_USERNAME` and `SUPER_ADMIN_PASSWORD` to seed a Super Admin at startup.

A `migrate/migrate.go` file is available for manual migrations or schema updates; you can run it with:

```bash
go run migrate/migrate.go
```

(Adapt its behavior to your needs.)

---

## Security Notes

- **JWT Secret**
  - Use a long, random value for `JWT_SECRET`.
  - Never commit real secrets to version control.

- **Cookies**
  - `Authorization` cookie is set with:
    - `HttpOnly: true`
    - `SameSite: Lax`
  - For production, also enable `Secure` (HTTPS only) on cookies.

- **Roles & Authorization**
  - Ensure admin and super admin routes are only called from trusted clients.
  - Carefully manage who has `SuperAdmin` access; they can promote/demote roles and delete admins.

- **CORS**
  - Configure CORS in `main.go` / Gin middleware if the API is used from browsers on a different origin.

---

## Extending the CMS

Possible extensions:

- **Versioning / Drafts**
  - Add fields for version, status (draft/published), and publish scheduling.

- **Filtering & Pagination**
  - Add query parameters for pagination, sorting, and filtering content.

- **Public API Tokens**
  - Introduce API keys or tokens for server-to-server integrations.

- **OpenAPI / Swagger Docs**
  - Generate and expose an OpenAPI spec for better client integration.

- **Observability**
  - Add request logging, metrics, and tracing.

---

## License

MIT License
