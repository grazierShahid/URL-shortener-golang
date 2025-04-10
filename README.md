# URL Shortener with Analytics

A high-performance URL shortener service built with Go, featuring real-time analytics, geographical tracking, and a scalable architecture.

## 🚀 Features

- **URL Shortening**: Generate short URLs for long links with ease.
- **Real-Time Analytics**: Track clicks with detailed analytics, including:
  - Total clicks
  - Geographical location (country, city)
  - Click timestamps
- **Geolocation Support**: Integrated with IP2Location for accurate geolocation data.
- **Scalable Architecture**: Uses Redis for caching and PostgreSQL for persistent storage.

## 🛠️ Tech Stack

- **Backend**: Go (Gin Web Framework)
- **Database**: PostgreSQL
- **Caching**: Redis
- **Geolocation**: IP2Location
- **Deployment**: Docker-ready

## 📖 API Endpoints

### 1. URL Shortening

- **POST** `/api/shorten`
  - Request: `{ "url": "https://example.com" }`
  - Response: `{ "short_url": "http://localhost:8080/abc123" }`

### 2. Redirect to Original URL

- **GET** `/:key`
  - Redirects to the original URL.

### 3. URL Analytics

- **GET** `/api/analytics/:key`
  - Response:
    ```json
    {
      "short_key": "abc123",
      "total_clicks": 42,
      "country_stats": [
        { "country": "United States", "count": 20 },
        { "country": "India", "count": 15 }
      ],
      "time_stats": [
        { "date": "2025-04-09", "count": 10 },
        { "date": "2025-04-08", "count": 12 }
      ]
    }
    ```

### 4. Overall Statistics

- **GET** `/api/stats`
  - Response:
    ```json
    {
      "stats": [
        { "short_url": "abc123", "total_clicks": 42 },
        { "short_url": "xyz789", "total_clicks": 30 }
      ]
    }
    ```

## ⚙️ Setup Instructions

### Prerequisites

- Go 1.21+
- PostgreSQL
- Redis
- IP2Location database (`IP2LOCATION-LITE-DB5.BIN`)

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/grazierShahid/URL-shortener-golang.git
   cd URL-shortener-golang
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up environment variables in a `.env` file:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=url_shortener
   DB_SSLMODE=disable
   REDIS_HOST=localhost
   REDIS_PORT=6379
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. Test the APIs using tools like Postman or `curl`.

## 📊 Database Schema

### Table: `url_clicks`

| Column       | Type         | Description            |
| ------------ | ------------ | ---------------------- |
| `id`         | SERIAL       | Primary key            |
| `short_key`  | VARCHAR(255) | Shortened URL key      |
| `click_time` | TIMESTAMP    | Timestamp of the click |
| `ip_address` | VARCHAR(255) | IP address of the user |
| `country`    | VARCHAR(255) | Country of the user    |
| `city`       | VARCHAR(255) | City of the user       |
| `region`     | VARCHAR(255) | Region of the user     |
