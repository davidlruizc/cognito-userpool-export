package cmd

import (
	"dockergo/app"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
)

var awsRegion string
var userPoolID string
var appClientID string

// exportCmd represents the export pool command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export a JSON of cognito userpool",
	Long:  `export allows you to generate a JSON file with all of user pools signed on your AWS cognito user pool.`,
	Run: func(cmd *cobra.Command, args []string) {
		if userPoolID != "" && appClientID != "" && awsRegion != "" {
			exportPoolCognito(userPoolID, appClientID, awsRegion)
		} else {
			fmt.Println("Sorry, you must specify the user pool id and the app client id like this:")
			fmt.Println("poolcognito export --region us-east-1 --poolid <user-pool-id> --clientid <app-client-id>")
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVar(&awsRegion, "region", "", "AWS region (required)")
	exportCmd.Flags().StringVar(&userPoolID, "poolid", "", "AWS cognito user pool id (required)")
	exportCmd.Flags().StringVar(&appClientID, "clientid", "", "AWS cognito app client id (required)")
}

func exportPoolCognito(userPoolID string, appClientID string, awsRegion string) {
	// dot env initial config.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	region := awsRegion

	// creating aws session
	conf := &aws.Config{Region: aws.String(region)}
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}

	// App definition
	cli := app.App{
		CognitoClient: cognito.New(sess),
		UserPoolID:    userPoolID,
		AppClientID:   appClientID,
	}

	users, err := cli.CognitoUserPoolPaginated()

	if err != nil {
		log.Fatal(err)
	}

	// export data in JSON file
	userPoolJSON, _ := json.Marshal(users.Users)
	err = ioutil.WriteFile("user-pool-output.json", userPoolJSON, 0644)
	fmt.Println("Your user pool has been exported successfully as user-pool-output.json")
}
