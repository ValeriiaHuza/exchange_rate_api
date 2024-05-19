# Exchange Rate API

This project is an Exchange Rate API built with Go, Gin, GORM, and Docker. It retrieves the current USD to UAH exchange rate from the National Bank of Ukraine, allows users to subscribe via email, and sends daily email updates.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Clone the Repository and go to the project directory

```sh
git clone https://github.com/ValeriiaHuza/exchange_rate_api.git

cd exchange_rate_api
```

### Create .env file

To run this project, you will need to add the following environment variables to your .env file

```sh
DB_HOST=db

DB_PORT=5432

DB_USERNAME=yourusername

DB_PASSWORD=yourpassword

DB_NAME=yourdatabase

MAIL_EMAIL=youremail@example.com

MAIL_PASSWORD=yourpassword`

```

### Build and run

Run the following command to build and start your Docker containers:

```sh
docker-compose up --build
```

## Rest API

#### Get exchange rate

Description: Get the current USD to UAH exchange rate

```http
  GET /rate
```
- Parameters: No parameters

- Response:
    - Status: 
         - 200 OK
         - 400 Bad request 
   

#### Subscribe an email

Description:  Subscribe an email to receive the current exchange rate

```http
  POST /subscribe
```

- Parameters :
    - email (string)
- Response:
    - Status: 
         - 200 OK
         - 400 Bad Request: Return if incorrect email format
         - 409 Conflict: Return if the email is already in the database
     - Body: Email added: "email" 
