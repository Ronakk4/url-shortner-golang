# URL Shortener in Golang with Fiber, Redis, and API Rate Limiting

This project is a URL shortener service built using [Fiber](https://gofiber.io/) (a fast HTTP framework for Golang) and [Redis](https://redis.io/) (an in-memory key-value store). The service includes user-level API rate limiting to prevent abuse by restricting the number of API calls within a given time frame.

---

## Features

- **Shorten URLs**: Generate short and unique aliases for long URLs.
- **Redirect URLs**: Redirect users to the original URL using the short link.
- **User Rate Limiting**: Enforce API usage limits per user using IP-based or API-key-based restrictions.
- **Fast and Scalable**: Built with Fiber for speed and Redis for performance.

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (v1.18+ recommended)
- [Redis](https://redis.io/download) (installed and running)
- [Docker](https://www.docker.com/) (optional, for containerized deployment
