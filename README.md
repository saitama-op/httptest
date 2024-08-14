# Go HTTP Server Example

This repository contains a simple Go HTTP server application designed to run on port 80. The server has two main endpoints:

## Endpoints

### `/`

- **Description**: This endpoint serves a responsive HTML page.
- **Content**: The page displays the private IP address of the server along with a Bootstrap-styled layout. The HTML template includes a title with the server's IP address and a basic three-column layout.

### `/health`

- **Description**: This endpoint performs a health check.
- **Response**: Returns "Success" to indicate that the server is running properly.

## Code Overview

### Main Function

- Sets up HTTP handlers for the root path and the `/health` endpoint.
- Retrieves the private IP address of the server and stores it in a global variable.
- Starts the HTTP server on port 80.

### Health Check Handler

- Provides a simple "Success" message for health checks.

### IP Retrieval Function

- Retrieves and prints the server's private IP address.
- Filters to return only IPv4 addresses that are private.

### Root Handler

- Serves an HTML page with the server's IP address embedded in the title and content.
- Uses Bootstrap for styling to create a responsive page layout.

## How to Run

1. Clone this repository to your local machine.
2. Navigate to the directory containing the `main.go` file.
3. Run the application using `go run main.go`.
4. Access the application in your web browser at `http://localhost`.

## Dependencies

- Go 1.16 or later
- Bootstrap (for styling the HTML page)

## License

This project is licensed under the MIT License. See the LICENSE file for details.
