package cmd

import (
	"dockergo/app"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
)

// exportCmd represents the export pool command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export a JSON of cognito userpool",
	Long:  `export allows you to generate a JSON file with all of user pools signed on your AWS cognito user pool.`,
	Run: func(cmd *cobra.Command, args []string) {
		exportPoolCognito()
		fmt.Println("Your user pool has been exported successfully as user-pool-output.json")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func exportPoolCognito() {
	// dot env initial config.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	region := os.Getenv("AWS_REGION")

	// creating aws session
	conf := &aws.Config{Region: aws.String(region)}
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

	users := cli.CognitoUserPoolPaginated().Users

	// export data in JSON file
	rankingsJSON, _ := json.Marshal(users)
	err = ioutil.WriteFile("user-pool-output.json", rankingsJSON, 0644)
}
