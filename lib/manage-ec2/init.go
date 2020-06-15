package manage

import "github.com/aws/aws-sdk-go/service/ec2"

var svc *ec2.EC2

func init() {

	svc = AwsEC2SessionHelper()

}
