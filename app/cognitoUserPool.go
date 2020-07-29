package app

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/request"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func (cli *App) getCognitoUserPoolPaginated(params *cognito.ListUsersInput) (*cognito.ListUsersOutput, error) {
	var users []*cognito.UserType

	ctx := context.Background()

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

	return output, p.Err()
}
