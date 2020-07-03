package connect

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	awsS3Region = "ap-south-1"
)

//ConnectAWS func ..
func ConnectAWS() *session.Session {
	sesn, err := session.NewSession(
		&aws.Config{Region: aws.String(awsS3Region)},
	)
	if err != nil {
		fmt.Println("Erorr while retrieving from session", err)
	}
	return sesn
}
