package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

// App holds internals for auth flow.
type App struct {
	CognitoClient *cognito.CognitoIdentityProvider
	UserPoolID    string
	AppClientID   string
}

func main() {
	// dot env initial config.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	// env variables declaration
	UserPoolID := os.Getenv("USER_POOL_ID")
	//AppClientID := os.Getenv("APP_CLIENT_ID")
	// fmt.Println(os.Getenv("AWS_ACCESS_KEY_ID"))

	// config aws region
	conf := &aws.Config{Region: aws.String("us-east-1")}

	// init session
	// sessInit, err := session.NewSession(conf)
	//if err != nil {
	//fmt.Println("Failed to create session: ", err)
	//return
	//}

	sess := session.Must(session.NewSession(conf))

	// test new credentials
	creds, err := stscreds.NewCredentials(sess, "arn:aws:iam::043735794078:user/daragon")
	if err != nil {
		log.Fatal("NOPE")
	}
	fmt.Println(creds)

	// New cognito instance declaration
	svc := cognito.New(sess)

	params := &cognito.ListUsersInput{
		UserPoolId: aws.String(UserPoolID),
		AttributesToGet: []*string{
			aws.String("email"),
		},
	}

	// List users coming from User Pool
	resp, err := svc.ListUsers(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)

	// echo request
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
