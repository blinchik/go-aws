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

type sg struct {
	Sgid   string `json:"sg_id"`
	SgName string `json:"sg_name"`
	IpPerm []ipPerm
	Tags   []TagMap
}

type ipPerm struct {
	IpProtocol string
	FromPort   int64
	ToPort     int64
	IpRanges   []string
}

type sgList struct {
	sgs []sg
}

func VpcDescribe() []byte {

	var vpcOutputList VpcList

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

		vpcOutputList.vpcs = append(vpcOutputList.vpcs, vpcMap)

	}

	vpcOutput, err := json.Marshal(vpcOutputList.vpcs)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(vpcOutput))

	return vpcOutput

}

func SgDescribe() []byte {

	var sgOutputList sgList

	svc := AwsEC2SessionHelper()

	result, err := svc.DescribeSecurityGroups(nil)

	if err != nil {
		log.Println(err)
	}

	for _, sgValue := range result.SecurityGroups {

		var sgMap sg

		sgMap.Sgid = *sgValue.GroupId
		sgMap.SgName = *sgValue.GroupName

		var ipPermMap ipPerm

		for _, sgIpPerm := range sgValue.IpPermissions {

			if sgIpPerm.FromPort != nil {

				ipPermMap.FromPort = *sgIpPerm.FromPort

			}

			if sgIpPerm.IpProtocol != nil {

				ipPermMap.IpProtocol = *sgIpPerm.IpProtocol

			}

			if sgIpPerm.ToPort != nil {

				ipPermMap.ToPort = *sgIpPerm.ToPort

			}

			for _, ipRange := range sgIpPerm.IpRanges {

				ipPermMap.IpRanges = append(ipPermMap.IpRanges, *ipRange.CidrIp)

			}

			sgMap.IpPerm = append(sgMap.IpPerm, ipPermMap)

		}

		var tagList []TagMap
		tag := make(TagMap)

		for _, values := range sgValue.Tags {

			tag[*values.Key] = *values.Value

		}

		tagList = append(tagList, tag)

		sgMap.Tags = tagList

		sgOutputList.sgs = append(sgOutputList.sgs, sgMap)

	}

	sgOutput, err := json.Marshal(sgOutputList.sgs)

	if err != nil {
		log.Println(err)
	}

	// fmt.Println(result)

	fmt.Println(string(sgOutput))

	return sgOutput

}

func SubnetDescribe() {

	svc := AwsEC2SessionHelper()

	result, err := svc.DescribeSubnets(nil)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

}
