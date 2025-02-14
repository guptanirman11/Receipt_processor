# Project: Receipt Processor

## Business Requirement
A receipt processor application that calculates the number of points earned from a valid receipt based on specific constraints.

## System Requirements
### Functional Requirements
- A user interface with a text box to input JSON receipts.
- A button to calculate points and display earned points.
- The application should run in Docker containers using Golang for the backend.
- In-memory storage should be used, meaning past receipts do not persist.

### Non-Functional Requirements
- The application must be scalable and maintainable.
- Users should be able to retrieve points for a previously processed receipt by providing its receipt ID.

## High-Level Architecture
- The Processor service uses a **client** - **service** architecture.
- The **client** handles user input, displaying receipt processing results.
- The **server** receives receipt data, validates it, calculates points, and stores the results temporarily.
- **Nginx** serves static files and acts as a reverse proxy to route API calls to the backend running on port `8080`.
- The backend is built using the **Gin** framework.
- The receipt data is stored in an **in-memory map**.
- The service validates the receipt structure but assumes all required fields contain correct values.

## File Structure
```
Receipt_Process/
|── client/
|   |── index.html
|   |── style.css
|   |── script.js
|   |── receipt.json
|── server/
|   |── receipt_processor_service.go
|── go.mod
|── go.sum
|── README.md
```

## Run Instructions
### Docker Setup
To build and run the service, execute:
```sh
# Build Docker image
docker build -t receipt_service .

# Run container and expose required ports
docker run -p 80:80 -p 8080:8080 receipt_service
```

### Accessing the Application
- Open `http://localhost:80` in a web browser.
- Paste a valid JSON receipt.
- Click the **Submit Your Receipt** button to process it.
- If valid, the receipt ID is returned and auto-filled.
- Click **Calculate Your Points** to fetch the points earned.

## Server Details
- The backend is built using the **Gin** framework.
- The receipt data is stored in an **in-memory map**.
- The service validates the receipt structure but assumes all required fields contain correct values.

## Future Extensions
- Extend support for processing of multiple receipts. Currenlty only one receipt is handled by the service.
- Improve type validation for fields like `date`, `time`, and `price`.
- Implement persistent storage for better tracking of receipts.
- The requests can be made secure by implementing user sign in and use HTTPS for request/response.
- For scalability, we can parallelise the points calculation as well as utilize a queuing mechanism to store process/receipt api calls and implement asynchronously for multiple users.

## Assumptions
- The JSON input structure is correctly formatted.
- Required fields if exists then contains valid values.
- Time, date, and numerical fields are assumed to be valid and in the format mentioned in api.yml.

## References
1. https://gin-gonic.com/docs/quickstart/
2. https://stackoverflow.com/questions/10075304/nginx-fails-to-load-css-files
3. https://github.com/nginx/nginx/blob/master/conf/nginx.conf