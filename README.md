
# Waysbooks (Backend)

Backend for waysbooks using RESTful API.


## Tech Stack

**Server:** Golang (Echo)

**Database:** MySQL

**Others:** JWT, Bycript, Cloudinary, S3 Bucket
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

#### SecretKey
`SECRET_KEY=<this-is-secret>`

#### Cloudinary API
`CLOUD_NAME=<cloudinary-cloud-name>`

`API_KEY=<cloudinary-api-key>`

`API_SECRET=<cloudinary-api-secret>`

#### Midtrans

`SERVER_KEY=<midtrans-server-key>`

#### Email

`EMAIL_SYSTEM=<your-gomail-email@gmail.com>`

`PASSWORD_SYSTEM=<gomail-application-password>`

#### Database
`DB_USER=<db-user>`

`DB_PASSWORD=<db-password>`

`DB_HOST=<db-host>`

`DB_PORT=<db-port>`

`DB_NAME=<db-name>`

#### AWS
`AWS_REGION=<aws-region-code>`

`AWS_ACCESS_KEY_ID=<your-aws-key-id>`

`AWS_SECRET_ACCESS_KEY=<your-aws-secret-key>`
## Run Locally

Clone the project

```bash
  git clone https://github.com/VindoKountur/backend-waysbeans-golang
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod tidy ; go mod download
```

Start the server

```bash
  go run main.go
```

