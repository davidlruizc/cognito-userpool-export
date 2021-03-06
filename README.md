# Export Cognito User Pool

Allows to export in a JSON file all the User Pool data coming from AWS Cognito. This happens paginating all the data stored in that service and concatenating the results in a JSON file.

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

Available commands:

```sh
go run main.go export --region <aws-region> --poolid <user-pool-id> --clientid <app-client-id>
```

To see the available commands con `export`

```sh
go run main.go export --help
```

export `.json`file in the root path where you cloned the project.

## TODO

- [x] Flag for userpool id and client id.
- [ ] Command to export a `json` or `csv` file.
- [ ] Add location path to export the file generated.
- [ ] Allow add custom `aws` credentials.

## Getting started with go

Initialize the project with dockergo allows to get the package manager of the project.

```sh
go mod init dockergo
```

[source blog](https://medium.com/@alemarcha/primera-aplicaci%C3%B3n-go-usando-docker-6b4618833073)

## Result expected

A JSON file in the root of the project called `user-pool-output.json`.
