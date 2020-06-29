package secrets

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func CreateSecret(description, name, secret string) {

	session := session.Must(session.NewSession())

	svc := secretsmanager.New(session, aws.NewConfig().WithRegion(os.Getenv("aws_region")))

	input := &secretsmanager.CreateSecretInput{
		Description:  aws.String(description),
		Name:         aws.String(name),
		SecretString: aws.String(secret),
	}

	result, err := svc.CreateSecret(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}
