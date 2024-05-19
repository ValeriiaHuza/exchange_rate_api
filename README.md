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

### The app will start running at http://localhost:8080

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
     

#### Request example
   
<img width="919" alt="Знімок екрана 2024-05-19 о 12 48 19" src="https://github.com/ValeriiaHuza/exchange_rate_api/assets/66443473/283adc47-fd54-4cfe-b7ab-fd4fc7418483">


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

#### Successful subscribe example

<img width="741" alt="Знімок екрана 2024-05-19 о 12 46 13" src="https://github.com/ValeriiaHuza/exchange_rate_api/assets/66443473/44d50332-7221-4f58-bf19-3b3e226bcfcc">

#### Invalid format example

<img width="947" alt="Знімок екрана 2024-05-19 о 12 46 39" src="https://github.com/ValeriiaHuza/exchange_rate_api/assets/66443473/3b280299-8cb1-4589-ae2b-f62afc1c942e">

#### Email exists in database example

<img width="935" alt="Знімок екрана 2024-05-19 о 12 47 08" src="https://github.com/ValeriiaHuza/exchange_rate_api/assets/66443473/1c6d0571-3f97-4134-b8fb-be0903728305">
