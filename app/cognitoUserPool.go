package app

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// CognitoUserPoolPaginated user pool pagination
func (cli *App) CognitoUserPoolPaginated() *cognitoidentityprovider.ListUsersOutput {
	var users []*cognito.UserType

	ctx := context.Background()

	params := &cognito.ListUsersInput{
		UserPoolId: &cli.UserPoolID,
		AttributesToGet: []*string{
			aws.String("email"),
		},
		Limit: aws.Int64(40),
	}

	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			req, _ := cli.CognitoClient.ListUsersRequest(params)
			req.SetContext(ctx)
			return req, nil
		},
	}

	for p.Next() {
		page := p.Page().(*cognito.ListUsersOutput)
		for _, obj := range page.Users {
			users = append(users, obj)
		}
	}

	output := &cognito.ListUsersOutput{}
	output.SetUsers(users)

	return output
}
