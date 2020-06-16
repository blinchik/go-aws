package manage

import (
	"fmt"
	"log"

	// "path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	rsaKey "github.com/blinchik/go-utils/rsakey"

	"github.com/aws/aws-sdk-go/aws"
)

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

func ImportKey(name string) {

	rsaKey.SavePEMKey(fmt.Sprintf("%s/.ssh/%s.pem", rsaKey.Home, name), rsaKey.Key)
	pub := rsaKey.KeepPublicPEMKey(rsaKey.PublicKey)

	var keyInput ec2.ImportKeyPairInput

	awsName := aws.String(name)

	keyInput.KeyName = awsName
	keyInput.PublicKeyMaterial = pub

	// svc.CreateKeyPair(&keyInput)

	out, err := svc.ImportKeyPair(&keyInput)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(out)

}
