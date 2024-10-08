# TSA Innovation Lab - Practical Code Challenge

## Synopsis
This project involves creating an HTTP server (API) that accepts and persists contact data into a relational database. The server is developed in Go using the gin-gonic framework.

## Guidelines

### API Development
- The API should be implemented in Go using the gin-gonic framework.

### Data Validation
- Only valid data must be accepted by the HTTP endpoint.

### Phone Number Format
- The phone number must be persisted in E.164 format and should always be an Australian number.

### Security
- Assume that the outer architecture will handle security; it's not required for this challenge.

### Documentation
- Provide clear instructions on how to call the API.

## Example Contact Data
Here are some examples of the data the API should handle:

```json
{
  "full_name": "Alex Bell",
  "email": "alex@bell-labs.com",
  "phone_numbers": ["03 8578 6688", "1800728069"]
}
```

```json
{
  "full_name": "fredrik IDESTAM",
  "phone_numbers": ["+6139888998"]
}
```

```json
{
  "full_name": "radia perlman",
  "email": "rperl001@mit.edu",
  "phone_numbers": ["(03) 9333 7119", "0488445688", "+61488224568"]
}
```

## Outcome

### Artefacts
Provide necessary artefacts to deploy this into a production environment. This could include a ZIP file or a GitHub/GitLab repository.

### Documentation
Detail any assumptions, tests, design or architecture thoughts, comments, or decisions made relating to requirements or implementation.

### Incomplete Items
Articulate any items considered incomplete and propose how they could be completed in the future.

## Approach Recommendations

### Completeness
The solution does not need to be fully complete. Provide details on what is incomplete and how it could be improved.

### Simplicity
Keep the solution as simple and clear as possible.

### Technologies
Use any comfortable technologies for the database, but use Golang & Gin for the backend.

## API Documentation

### Endpoints

#### POST /contacts
**Description:** Create a new contact.

**Request Body:**

```json
{
  "full_name": "string",
  "email": "string (optional)",
  "phone_numbers": ["string"]
}
```

**Response:**

- `201 Created`: When the contact is successfully created.
- `400 Bad Request`: When the request data is invalid.

**Example Request:**

```sh
curl -X POST http://localhost:8080/contacts \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Alex Bell",
    "email": "alex@bell-labs.com",
    "phone_numbers": ["03 8578 6688", "1800728069"]
  }'
```

#### GET /contacts
**Description:** Gets all contact.

**Response:**

- `200 OK`: Fetched list of contacts

**Example Request:**

```sh
curl -X GET http://localhost:8080/contacts \
  -H "Content-Type: application/json" 
```

#### GET /contacts/:id
**Description:** Gets contact infoby id.

**Response:**

- `200 OK`: Fetched contact info

**Example Request:**

```sh
curl -X GET http://localhost:8080/contacts/1 \
  -H "Content-Type: application/json" 
```


#### PUT /contacts/:id
**Description:** Updates contact info

**Request Body:**

```json
{
  "full_name": "string",
  "email": "string (optional)",
  "phone_numbers": ["string"]
}
```

**Response:**

- `200 Ok`: When the contact is successfully updated.
- `400 Bad Request`: When the request data is invalid.

**Example Request:**

```sh
curl -X PUT http://localhost:8080/contacts/1 \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Alex Bell",
    "email": "alex@bell-labs.com",
    "phone_numbers": ["03 8578 6688", "1800728069"]
  }'
```

#### DELETE /contacts/:id
**Description:** Deletes a contact.

**Response:**

- `200 OK`: When the contact is successfully created.
- `400 Bad Request`: When the request data is invalid.

**Example Request:**

```sh
curl -X DELETE http://localhost:8080/contacts/1 \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Alex Bell",
    "email": "alex@bell-labs.com",
    "phone_numbers": ["03 8578 6688", "1800728069"]
  }'
```

## Assumptions
- Data Validation: Implementations should handle invalid data gracefully.
- Database: Used postgresql
- Docker: used docker for containerization of database
- Added Bruno API collection (this is like postman, visit https://www.usebruno.com/ for more info)

## Design Thoughts
- Used regexp at first for phone number validation and normalizing into E.164 format.
- Deprecated regexp and used github.com/ttacon/libphonenumber package for validation. I thought it would be better to use the library as its more reliable its a fork of google's phone number validator.
- Used a provider to easily switch between regexp and libphonenumber if needed
- Created a service layer that handles database operations separate from controller methods


## Tests
    ```sh
    go test ./... -v
    ```
Test is included in the services folder

## Future Improvements
- Error Handling: Improve error handling and validation messages.
- Authentication: Add authentication and authorization if needed.
- CI/CD Deployment - Fix Docker setup and Procfile for Heroku distribution
- Documentation - Implement API documentation via swaggo/swag package

## Getting Started

1. **Clone the repository:**

    ```sh
    git clone <repository_url>
    ```

2. **Navigate to the project directory:**

    ```sh
    cd <project_directory>
    ```

3. **Install dependencies:**

    ```sh
    go mod tidy
    ```

4. **Run the server:**

    ```sh
    go run main.go
    ```

5. **Use the API as described above.**

---