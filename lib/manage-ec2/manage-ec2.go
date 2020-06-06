package manage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/aws/aws-sdk-go/aws"
)

var home, err = os.UserHomeDir()

var awsCred = os.Getenv("GO_AWS_CRED")

var awsRegion = os.Getenv("aws_region")

func AwsEC2SessionHelper() (svc *ec2.EC2) {

	var sess *session.Session

	creds := credentials.NewSharedCredentials(filepath.FromSlash(fmt.Sprintf("%s/.aws/credentials", home)), awsCred)

	_, err := creds.Get()

	if err != nil {
		log.Println("skipping ~/.aws/credentials file")

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

type SummaryEC2 struct {
	InstanceId       []*string `locationName:"instanceId" type:"string"`
	PrivateIpAddress []*string `locationName:"privateIpAddress" type:"string"`
	PublicIp         []*string `locationName:"publicIp" type:"string"`
	TagValue         []*string `locationName:"value" type:"string"`
}

func StopEC2(InstanceID []*string) {

	svc := AwsEC2SessionHelper()

	input := &ec2.StopInstancesInput{
		InstanceIds: InstanceID,
		DryRun:      aws.Bool(false),
	}

	result, err := svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)

	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.StopInstances(input)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Success", result.StoppingInstances)
		}
	} else {
		fmt.Println("Error", result.StoppingInstances)
	}

}

func StartEC2(InstanceID []*string) {

	svc := AwsEC2SessionHelper()

	input := &ec2.StartInstancesInput{
		InstanceIds: InstanceID,
		DryRun:      aws.Bool(false),
	}

	result, err := svc.StartInstances(input)
	if err != nil {
		fmt.Println(err)
	}

	inputStatus := &ec2.DescribeInstanceStatusInput{
		InstanceIds: InstanceID,
	}

	svc.WaitUntilInstanceStatusOk(inputStatus)

	fmt.Println(result)
}

func DescribeByOperationTag(value string) (summary SummaryEC2) {

	svc := AwsEC2SessionHelper()

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:OperationGroup"),
				Values: []*string{
					aws.String(value),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)

	if err != nil {

		fmt.Println("Error", err)

	} else {

		for i := 0; i < len(result.Reservations); i++ {

			if *result.Reservations[i].Instances[0].State.Name != "terminated" {

				summary.InstanceId = append(summary.InstanceId, result.Reservations[i].Instances[0].InstanceId)
				summary.PrivateIpAddress = append(summary.PrivateIpAddress, result.Reservations[i].Instances[0].PrivateIpAddress)

				AssociationStruct := result.Reservations[i].Instances[0].String()
				ec2IsPub := strings.Contains(AssociationStruct, "Association")

				if ec2IsPub {

					summary.PublicIp = append(summary.PublicIp, result.Reservations[i].Instances[0].NetworkInterfaces[0].Association.PublicIp)

				}

			}

		}

	}

	return summary

}

func DescribeAllRunning() (summary SummaryEC2) {

	svc := AwsEC2SessionHelper()

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)

	if err != nil {

		fmt.Println("Error", err)

	} else {

		for i := 0; i < len(result.Reservations); i++ {

			summary.InstanceId = append(summary.InstanceId, result.Reservations[i].Instances[0].InstanceId)
			summary.PrivateIpAddress = append(summary.PrivateIpAddress, result.Reservations[i].Instances[0].PrivateIpAddress)

			AssociationStruct := result.Reservations[i].Instances[0].String()
			ec2IsPub := strings.Contains(AssociationStruct, "Association")

			if ec2IsPub {

				summary.PublicIp = append(summary.PublicIp, result.Reservations[i].Instances[0].NetworkInterfaces[0].Association.PublicIp)

			} else {

				summary.PublicIp = append(summary.PublicIp, nil)

			}

		}

	}

	return summary

}

func DescribeByGeneralTag(tag string, value string) (summary SummaryEC2) {

	svc := AwsEC2SessionHelper()

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(fmt.Sprintf("tag:%s", tag)),
				Values: []*string{
					aws.String(value),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)

	if err != nil {

		fmt.Println("Error", err)

	} else {

		for i := 0; i < len(result.Reservations); i++ {

			if *result.Reservations[i].Instances[0].State.Name != "terminated" {

				summary.InstanceId = append(summary.InstanceId, result.Reservations[i].Instances[0].InstanceId)
				summary.PrivateIpAddress = append(summary.PrivateIpAddress, result.Reservations[i].Instances[0].PrivateIpAddress)

				AssociationStruct := result.Reservations[i].Instances[0].String()
				ec2IsPub := strings.Contains(AssociationStruct, "Association")

				if ec2IsPub {

					summary.PublicIp = append(summary.PublicIp, result.Reservations[i].Instances[0].NetworkInterfaces[0].Association.PublicIp)

				}

			}

		}

	}

	return summary

}

func DescribeAllMentionedTag(tag string) (summary SummaryEC2) {

	svc := AwsEC2SessionHelper()

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(fmt.Sprintf("tag:%s", tag)),
				Values: []*string{
					aws.String("*"),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)

	if err != nil {

		fmt.Println("Error", err)

	} else {

		for i := 0; i < len(result.Reservations); i++ {

			if *result.Reservations[i].Instances[0].State.Name != "terminated" {

				summary.InstanceId = append(summary.InstanceId, result.Reservations[i].Instances[0].InstanceId)
				summary.PrivateIpAddress = append(summary.PrivateIpAddress, result.Reservations[i].Instances[0].PrivateIpAddress)

				Tags := result.Reservations[i].Instances[0].Tags
				for _, v := range Tags {
					if *v.Key == tag {
						summary.TagValue = append(summary.TagValue, v.Value)
					}
				}

				AssociationStruct := result.Reservations[i].Instances[0].String()
				ec2IsPub := strings.Contains(AssociationStruct, "Association")

				if ec2IsPub {

					summary.PublicIp = append(summary.PublicIp, result.Reservations[i].Instances[0].NetworkInterfaces[0].Association.PublicIp)

				} else {

					summary.PublicIp = append(summary.PublicIp, nil)

				}

			}

		}

	}

	return summary

}

func DescribeAllVols() {

	svc := AwsEC2SessionHelper()
	input := &ec2.DescribeVolumesInput{}

	result, err := svc.DescribeVolumes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	for _, Attachments := range result.Volumes {

		fmt.Println("=================================")
		for _, tag := range Attachments.Tags {
			fmt.Println(*tag.Value)
		}

		fmt.Println(*Attachments.VolumeId)

	}

}

func DetachVol() {

	svc := ec2.New(session.New())
	input := &ec2.DetachVolumeInput{
		VolumeId: aws.String("vol-1234567890abcdef0"),
	}

	result, err := svc.DetachVolume(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func DeleteVol() {
	svc := ec2.New(session.New())
	input := &ec2.DeleteVolumeInput{
		VolumeId: aws.String("vol-049df61146c4d7901"),
	}

	result, err := svc.DeleteVolume(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}
