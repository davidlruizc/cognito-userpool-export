# Export Cognito User Pool

Allows to export in a JSON file all the User Pool data coming from AWS Cognito. This happens paginating all the data stored in that service and concatenating the results in a JSON file.

## Getting started with go

Initialize the project with dockergo allows to get the package manager of the project.

```sh
go mod init dockergo
```

[source blog](https://medium.com/@alemarcha/primera-aplicaci%C3%B3n-go-usando-docker-6b4618833073)

## Usage

Clone the project

```sh
https://github.com/davidlruizc/cognito-userpool-export.git
```

Add `.env` file with the follow data

```sh
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
USER_POOL_ID=
APP_CLIENT_ID=
AWS_REGION=
```

Then you can run

```sh
go run main.go
```

## Result expected

A JSON file in the root of the project called `user-pool-output.json`.
