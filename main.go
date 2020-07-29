package main

import (
	"context"
	"dockergo/app"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	var users []*cognito.UserType

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
	cli := app.App{
		CognitoClient: cognito.New(sess),
		UserPoolID:    os.Getenv("USER_POOL_ID"),
		AppClientID:   os.Getenv("APP_CLIENT_ID"),
	}

	params := &cognito.ListUsersInput{
		UserPoolId: &cli.UserPoolID,
		AttributesToGet: []*string{
			aws.String("email"),
		},
		Limit: aws.Int64(40),
	}

	ctx := context.Background()

	// request pagination
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			req, _ := cli.CognitoClient.ListUsersRequest(params)
			req.SetContext(ctx)
			return req, nil
		},
	}

	// pagination iterator and append the array with results
	for p.Next() {
		page := p.Page().(*cognito.ListUsersOutput)
		for _, obj := range page.Users {
			users = append(users, obj)
		}
	}

	fmt.Println(len(users))

	// echo request
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})

	// declaring the port
	e.Logger.Fatal(e.Start(":1323"))
}
