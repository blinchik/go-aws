package manage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"

	// "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func DescribeAllVols() {

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
