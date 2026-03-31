# SpendSense

An AI-powered expense tracker REST API built in Go. Users describe purchases in plain English, and the app automatically parses, categorizes, and stores them using OpenAI. Supports monthly budgets per category with overspending alerts, and provides spending insights like summaries, trends, and top merchants.

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL
- **ORM/Driver:** sqlx
- **Migrations:** golang-migrate
- **Auth:** JWT + bcrypt
- **AI:** OpenAI API
- **Containerization:** Docker + docker-compose

## Architecture

```
Handler → Service → Repository
```

Layered architecture with dependency injection and interfaces for testability.

## Features

- **User Authentication** — Register, login, and JWT-protected routes
- **Natural Language Expense Input** — Describe a purchase in plain English and let AI parse it into structured data (amount, merchant, category)
- **Manual Expense Tracking** — Create, read, filter, and delete expenses with ownership enforcement
- **Category Management** — Default seeded categories + user-created custom categories
- **Monthly Budgets** — Set spending limits per category
- **Overspending Alerts** — Get notified when a category exceeds its budget
- **Spending Insights** — Monthly summaries, top merchants, and spending trends over time

## API Endpoints

### Auth
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/auth/register` | Create a new account |
| POST | `/auth/login` | Login and receive JWT |

### Categories
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/categories` | List all categories |
| POST | `/categories` | Create a custom category |

### Expenses
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/expenses` | Create an expense manually |
| GET | `/expenses` | List expenses (with filters) |
| GET | `/expenses/:id` | Get a single expense |
| DELETE | `/expenses/:id` | Delete an expense (owner only) |
| POST | `/expenses/parse` | AI-powered: parse a plain English expense |

### Budgets
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/budgets` | Set a monthly budget for a category |
| GET | `/budgets` | List all budgets |
| GET | `/budgets/alerts` | Get overspending alerts |

### Insights
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/insights/summary` | Monthly spending summary |
| GET | `/insights/top-merchants` | Top merchants by spend |
| GET | `/insights/trend` | Spending trend over time |

## Getting Started

### Prerequisites
- Docker and docker-compose installed

### Run
```bash
git clone https://github.com/yourusername/spendsense.git
cd spendsense
docker-compose up --build
```

The API will be available at `http://localhost:8080`.

### Environment Variables

| Variable | Description |
|----------|-------------|
| `DB_HOST` | PostgreSQL host |
| `DB_PORT` | PostgreSQL port |
| `DB_USER` | Database username |
| `DB_PASSWORD` | Database password |
| `DB_NAME` | Database name |
| `JWT_SECRET` | Secret key for JWT signing |
| `OPENAI_API_KEY` | OpenAI API key for expense parsing |

## Project Structure

```
spendsense/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   ├── middleware/
│   └── parser/
├── migrations/
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── README.md
```

## What I Learned

This is the second backend project I built in Go (after [RepLog](https://github.com/yourusername/replog)). SpendSense added:

- Integrating the OpenAI API from a Go backend
- Designing structured extraction prompts for LLMs
- Validating and sanitizing AI-generated output before persisting
- Seeding default data via migrations
- Budget threshold logic and alert generation
- Aggregation queries for insights (summaries, trends, top merchants)
- Filtering with multiple optional query parameters
