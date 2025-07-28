# 🔗 URL Shortener

A simple and modern URL shortener built as a **monorepo** using:

- ⚛️ **Frontend**: React + TypeScript (Vite)
- 🦫 **Backend**: Golang + Gin
- 🐳 Docker + Docker Compose for development

---

## 🚀 Features

- Create short URLs with ease
- Redirect to original links using short codes
- Copy-to-clipboard button
- Fully containerized setup
- Built with clean, modular code

---

## 🧰 Tech Stack

| Layer     | Stack                             |
|-----------|------------------------------------|
| Frontend  | React, TypeScript, Vite            |
| Backend   | Golang, Gin                        |
| Database  | MongoDB, Redis                     |
| DevOps    | Docker, Docker Compose             |

---

## 🧪 API Endpoints

### POST `/api/shorten`

Create a short link.

**Request body:**

```json
{
  "url": "https://example.com"
}
```

### GET `/:shortCode`

Redirect to the original URL based on the short code.

---

## 🧑‍💻 Local Development

### Prerequisites

- Docker
- Node.js (for running frontend independently)
- Go 1.21+

### Clone the repository

```bash
git clone https://github.com/JonasOli/url-shortener.git
cd url-shortener
```

### Run with Docker and Make (recommended)

``` shell
make start
```
