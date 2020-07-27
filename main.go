package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
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

	// creating aws session
	conf := &aws.Config{Region: aws.String("us-east-1")}
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}

	// App definition
	cli := App{
		CognitoClient: cognito.New(sess),
		UserPoolID:    os.Getenv("USER_POOL_ID"),
		AppClientID:   os.Getenv("APP_CLIENT_ID"),
	}

	params := &cognito.ListUsersInput{
		UserPoolId: &cli.UserPoolID,
		AttributesToGet: []*string{
			aws.String("email"),
		},
	}

	// List users coming from User Pool
	resp, err := cli.CognitoClient.ListUsers(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(len(resp.Users))

	// echo request
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// declaring the port
	e.Logger.Fatal(e.Start(":1323"))
}
