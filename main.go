package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// func main() {
// 	// dot env initial config.
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file.")
// 	}

// 	region := os.Getenv("AWS_REGION")

// 	// creating aws session
// 	conf := &aws.Config{Region: aws.String(region)}
// 	sess, err := session.NewSession(conf)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// App definition
// 	cli := app.App{
// 		CognitoClient: cognito.New(sess),
// 		UserPoolID:    os.Getenv("USER_POOL_ID"),
// 		AppClientID:   os.Getenv("APP_CLIENT_ID"),
// 	}

// 	users := cli.CognitoUserPoolPaginated().Users

// 	// export data in JSON file
// 	rankingsJson, _ := json.Marshal(users)
// 	err = ioutil.WriteFile("user-pool-output.json", rankingsJson, 0644)

// 	// echo request
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, users)
// 	})

// 	// declaring the port
// 	e.Logger.Fatal(e.Start(":1323"))
// }

var rootCmd = &cobra.Command{
	Use:   "poolcog",
	Short: "Poolcog it's a cognito user pool export CLI",
	Long:  `A very useful way to export your users pool from AWS cognito user pool service.`,
	Run: func(cmd *cobra.Command, arg []string) {
		fmt.Println("TEST command")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
