
# URL Shortener

This is a simple URL shortener application written in Go. It uses Redis as the database for storing the shortened URLs. The application provides an API to create short URLs and resolve them to the original long URLs.

## Features
Shorten long URLs to short URLs with custom or automatically generated IDs.
Store short URLs in Redis for efficient retrieval.
Resolve short URLs to the original long URLs.
Rate limiting for API requests.
Expiry time for short URLs.
Error handling for invalid URLs and rate limit exceeded.

## Technologies Used
Go: The programming language used to build the application.
Fiber: A fast and efficient web framework for Go.
Redis: An in-memory data structure store used as the database.
Docker: Containerization technology used for easy deployment.
Docker Compose: Tool for defining and running multi-container Docker applications.

## Getting Started
To get started with the URL shortener, follow the steps below:

Clone the repository: git clone https://github.com/your-username/url-shortener.git
Install Docker and Docker Compose.
Create an .env file in the root directory of the project and configure the environment variables. Here's an example:
makefile
Copy code
DB_ADDR=db:6379
DB_PASS=
APP_PORT=:3000
DOMAIN=localhost:3000
API_QUOTA=10
Run the application using Docker Compose: docker-compose up.
Access the application in your browser at http://localhost:3000.
API Endpoints
Shorten URL
URL: POST /api/v1
Request Body:
json
Copy code
{
  "url": "https://www.example.com",
  "customShort": "",
  "expiry": 0
}
Response:
json
Copy code
{
  "url": "https://www.example.com",
  "customShort": "",
  "expiry": 86400,
  "rateLimitRemaining": 10,
  "rateLimitReset": 10
}
Resolve URL
URL: GET /:url
Response:
If the URL is found, it redirects to the original long URL.
If the URL is not found, it returns a JSON response:
json
Copy code
{
  "error": "URL not found"
}

## Contributing
Contributions to the URL shortener are welcome! If you find any bugs or want to add new features, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.

## Acknowledgements
### Fiber: Fast and efficient Go web framework.
### Redis: In-memory data structure store used as the database. 
### Docker: Containerization platform used for easy deployment.
### Docker Compose: Tool for defining and running multi-container Docker applications.
