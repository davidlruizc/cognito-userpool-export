package app

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// CognitoUserPoolPaginated user pool pagination
func (cli *App) CognitoUserPoolPaginated() (*cognitoidentityprovider.ListUsersOutput, error) {
	var users []*cognito.UserType

	ctx := context.Background()

	params := &cognito.ListUsersInput{
		UserPoolId: &cli.UserPoolID,
		Limit:      aws.Int64(40),
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

	if len(users) > 0 {
		return output, nil
	}

	return nil, errors.New("Error: Must be an error on your user pool id or app client id values")
}
