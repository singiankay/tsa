TSA Innovation Lab - Practical Code Challenge
Synopsis
This project involves creating an HTTP server (API) that accepts and persists contact data into a relational database. The server is developed in Go using the gin-gonic framework.

Guidelines
API Development: The API should be implemented in Go using the gin-gonic framework.
Data Validation: Only valid data must be accepted by the HTTP endpoint.
Phone Number Format: The phone number must be persisted in E.164 format and should always be an Australian number.
Security: Assume that the outer architecture will handle security; it's not required for this challenge.
Documentation: Provide clear instructions on how to call the API.
Example Contact Data
Here are some examples of the data the API should handle:

{
  "full_name": "Alex Bell",
  "email": "alex@bell-labs.com",
  "phone_numbers": ["03 8578 6688", "1800728069"]
}
{
  "full_name": "fredrik IDESTAM",
  "phone_numbers": ["+6139888998"]
}
{
  "full_name": "radia perlman",
  "email": "rperl001@mit.edu",
  "phone_numbers": ["(03) 9333 7119", "0488445688", "+61488224568"]
}

Outcome
Artefacts: Provide necessary artefacts to deploy this into a production environment. This could include a ZIP file or a GitHub/GitLab repository.
Documentation: Detail any assumptions, tests, design or architecture thoughts, comments, or decisions made relating to requirements or implementation.
Incomplete Items: Articulate any items considered incomplete and propose how they could be completed in the future.

Approach Recommendations
Completeness: The solution does not need to be fully complete. Provide details on what is incomplete and how it could be improved.
Simplicity: Keep the solution as simple and clear as possible.
Technologies: Use any comfortable technologies for the database, but use Golang & Gin for the backend.

API Documentation
Endpoints
POST /contacts
Description: Create a new contact.

Request Body:
{
  "full_name": "string",
  "email": "string (optional)",
  "phone_numbers": ["string"]
}

Response:
201 Created: When the contact is successfully created.
400 Bad Request: When the request data is invalid.

Example Request:
curl -X POST http://localhost:8080/contacts \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Alex Bell",
    "email": "alex@bell-labs.com",
    "phone_numbers": ["03 8578 6688", "1800728069"]
  }'

Assumptions
Data Validation: Implementations should handle invalid data gracefully.
Database: Use a relational database of your choice to persist the data.
Tests
Include any relevant tests to ensure the API behaves as expected.

Future Improvements
Error Handling: Improve error handling and validation messages.
Authentication: Add authentication and authorization if needed.
Performance: Optimize performance for larger datasets.

Getting Started
Clone the repository:
    git clone <repository_url>

Navigate to the project directory:
    cd <project_directory>

Install dependencies:
    go mod tidy

Run the server:
    go run main.go

Use the API as described above.

