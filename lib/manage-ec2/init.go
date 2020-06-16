package manage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var svc *ec2.EC2

var home, err = os.UserHomeDir()

var awsCred = os.Getenv("GO_AWS_CRED")

var awsRegion = os.Getenv("aws_region")

func AwsEC2SessionHelper() (svc *ec2.EC2) {

	var sess *session.Session

	creds := credentials.NewSharedCredentials(filepath.FromSlash(fmt.Sprintf("%s/.aws/credentials", home)), awsCred)

	_, err := creds.Get()

	if err != nil {

		sess, err = session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
		})

		if err != nil {

			log.Fatal(err)
		}

		// Create an EC2 service client.
		svc = ec2.New(sess)
		return svc

	}

	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: creds,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Create an EC2 service client.
	svc = ec2.New(sess)
	return svc

}

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	svc = AwsEC2SessionHelper()

}
