# api_rest_go
Api of create products and comunication with database

Product API with Direct PostgreSQL Queries via HTTP Verbs

This is a straightforward approach to building a product API directly against a PostgreSQL database, employing HTTP verbs for data manipulation.

Technical Overview:

    API Architecture: Dockerized for containerized development with separation of concerns:

        Docker Container 1: API server

        Docker Container 2: PostgreSQL database

    Data Access: Directly using PostgreSQL queries within your API logic. You can employ libraries like pq for Go to interface with the database.

    Database Design: The database schema, "product" table, was defined and constructed using Dbeaver, a popular database management tool (https://dbeaver.io/download/).

API Routes:

    GET /products: Retrieves a list of all products

    POST /products: Creates a new product

    GET /products/{id}: Retrieves a specific product based on the ID provided

Project Setup:

    Install Docker: (https://www.docker.com/products/docker-desktop)

    Clone the Repository: Obtain the source code for the project from the repository.

    Build the Image: Execute the build command (usually docker build -t <image-name> . to create a container image.

    Run the API Server: Launch the Go application: go run main.go. The main.go file likely includes your server logic and API handlers.

    Test with Postman (or a similar HTTP client): Use a tool like Postman to interact with your API endpoints and test requests. 
