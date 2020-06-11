package manage

import (
	"encoding/json"
	"fmt"
	"log"
)

// TagMap ads dasd
type TagMap map[string]string

// TagMap ads dasd
type Associations map[string]string

type vpc struct {
	Vpcid      string
	Tags       []TagMap
	CidrBlocks []Associations
}

type VpcList struct {
	vpcs []vpc
}

func VpcDescribe() {

	svc := AwsEC2SessionHelper()

	result, err := svc.DescribeVpcs(nil)

	if err != nil {
		log.Println(err)
	}

	// fmt.Println(result)

	for _, vpcValue := range result.Vpcs {

		var vpcMap vpc

		// Get VPC id
		vpcMap.Vpcid = *vpcValue.VpcId

		// Get Tags
		/////////////////////////////////////////////////////////
		/////////////////////////////////////////////////////////

		var tagList []TagMap
		tag := make(TagMap)

		for _, values := range vpcValue.Tags {

			tag[*values.Key] = *values.Value

		}

		tagList = append(tagList, tag)

		vpcMap.Tags = tagList

		/////////////////////////////////////////////////////////
		/////////////////////////////////////////////////////////

		// Get CidrBloc within AssociationId

		var AssociationsList []Associations
		cidrB := make(map[string]string)

		for _, AssociationId := range vpcValue.CidrBlockAssociationSet {

			cidrB[*AssociationId.AssociationId] = *AssociationId.CidrBlock

		}

		AssociationsList = append(AssociationsList, cidrB)

		vpcMap.CidrBlocks = AssociationsList

		vpcMapJSON, err := json.Marshal(vpcMap)

		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(vpcMapJSON))

	}
}

func SgDescribe() {

	svc := AwsEC2SessionHelper()

	result, err := svc.DescribeSecurityGroups(nil)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

}

func SubnetDescribe() {

	svc := AwsEC2SessionHelper()

	result, err := svc.DescribeSubnets(nil)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

}
