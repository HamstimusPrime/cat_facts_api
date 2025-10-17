# Cat Facts API ⚙️

A simple and elegant Go REST API that fetches random cat facts and returns them along with user metadata.

## 🚀 Features

- **Random Cat Facts**: Fetches interesting cat facts from an external API
- **User Metadata**: Returns configurable user information (name, email, tech stack)
- **Clean JSON Responses**: Well-structured API responses with timestamps
- **Environment Configuration**: Flexible configuration using environment variables
- **Error Handling**: Robust error handling with appropriate HTTP status codes
- **Middleware Support**: Clean middleware implementation for request processing

## 📡 API Endpoints

### `GET /me`
Returns a random cat fact along with developer information.

**Response Format:**
```json
{
  "status": "success",
  "user": {
    "email": "your.email@example.com",
    "name": "Your Name",
    "stack": "Go, Docker, AWS"
  },
  "timestamp": "2024-10-17T10:30:00Z",
  "fact": "A cat's heart beats nearly twice as fast as a human heart."
}
```

## Tech Stack

- **Languages**: Go 1.24.3
- **Dependencies**: Go 1.24.3
    - [github.com/joho/godotenv](https://github.com/joho/godotenv) — Environment variable management
- **Architecture**: RESTFUL API 
- **HTTP Server**: Native Go ```net/http```

## Setup & Installation

1. **Clone the repository**
```bash
git clone https://github.com/HamstimusPrime/cat_facts_api.git
cd cat_facts_api
```


2. **Create environment file**
```bash
touch .env
```
&nbsp;&nbsp;&nbsp;&nbsp;**Configure your ```⚙️.env``` file:**

```
PORT=8080
API_URL=https://catfact.ninja/fact
EMAIL=your.email@example.com
NAME=Your Name
STACK=Golang
STATUS=success
TIMEOUT_DURATION=5
```
3. **Install dependencies**
```bash
go mod tidy 
```

4. **Run the application**
```bash
go run . 
```
&nbsp;&nbsp;&nbsp;&nbsp;The API will be available at http://localhost:8080:

## 🧪 Usage Example
```
# Fetch a random cat fact
curl http://localhost:8080/me
```

## 📁 Project Structure
```cat_facts_api/
├── main.go          # Application entry point and server setup
├── handlers.go      # HTTP request handlers and middleware
├── models.go        # Data structures and types
├── utils/
│   └── utils.go     # Utility functions for error handling
├── go.mod           # Go module dependencies
└── README.md        # Project documentation
```
## 🔍 Key Features Explained
- **Middleware Pattern:** Clean separation of concerns using middleware for request processing
- **Timeout Handling:** HTTP client with dynamic timeout configuration for external API requests
- **Graceful Error Handling:** Comprehensive error handling with appropriate HTTP status codes
- **Environment-Based Configuration:** Flexible configuration using environment variables
- **Clean Architecture:**  Well-organized code structure following Go best practices