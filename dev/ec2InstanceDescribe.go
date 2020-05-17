
type DescribeInstancesOutput struct {
	NextToken *string `locationName:"nextToken" type:"string"`
	Reservations []*Reservation `locationName:"reservationSet" locationNameList:"item" type:"list"`
}

type Reservation struct {
	Groups []*GroupIdentifier `locationName:"groupSet" locationNameList:"item" type:"list"`
	Instances []*Instance `locationName:"instancesSet" locationNameList:"item" type:"list"`
	OwnerId *string `locationName:"ownerId" type:"string"`
	RequesterId *string `locationName:"requesterId" type:"string"`
	ReservationId *string `locationName:"reservationId" type:"string"`
}

type InstanceBlockDeviceMapping struct {
	DeviceName *string `locationName:"deviceName" type:"string"`
	Ebs *EbsInstanceBlockDevice `locationName:"ebs" type:"structure"`
}

type EbsInstanceBlockDevice struct {
	AttachTime *time.Time `locationName:"attachTime" type:"timestamp"`
	DeleteOnTermination *bool `locationName:"deleteOnTermination" type:"boolean"`
	Status *string `locationName:"status" type:"string" enum:"AttachmentStatus"`
	VolumeId *string `locationName:"volumeId" type:"string"`
}

type CapacityReservationSpecificationResponse struct {
	CapacityReservationPreference *string `locationName:"capacityReservationPreference" type:"string" enum:"CapacityReservationPreference"`
	CapacityReservationTarget *CapacityReservationTargetResponse `locationName:"capacityReservationTarget" type:"structure"`
}

type CapacityReservationTargetResponse struct {
	CapacityReservationId *string `locationName:"capacityReservationId" type:"string"`
}

type ElasticGpuAssociation struct {
	ElasticGpuAssociationId *string `locationName:"elasticGpuAssociationId" type:"string"`
	ElasticGpuAssociationState *string `locationName:"elasticGpuAssociationState" type:"string"`
	ElasticGpuAssociationTime *string `locationName:"elasticGpuAssociationTime" type:"string"`
	ElasticGpuId *string `locationName:"elasticGpuId" type:"string"`
}

type ElasticInferenceAcceleratorAssociation struct {
	ElasticInferenceAcceleratorArn *string `locationName:"elasticInferenceAcceleratorArn" type:"string"`
	ElasticInferenceAcceleratorAssociationId *string `locationName:"elasticInferenceAcceleratorAssociationId" type:"string"`
	ElasticInferenceAcceleratorAssociationState *string `locationName:"elasticInferenceAcceleratorAssociationState" type:"string"`
	ElasticInferenceAcceleratorAssociationTime *time.Time `locationName:"elasticInferenceAcceleratorAssociationTime" type:"timestamp"`
}

type LicenseConfiguration struct {
	LicenseConfigurationArn *string `locationName:"licenseConfigurationArn" type:"string"`
}

type InstanceNetworkInterface struct {
	Association *InstanceNetworkInterfaceAssociation `locationName:"association" type:"structure"`
	Attachment *InstanceNetworkInterfaceAttachment `locationName:"attachment" type:"structure"`
	Description *string `locationName:"description" type:"string"`
	Groups []*GroupIdentifier `locationName:"groupSet" locationNameList:"item" type:"list"`
	Ipv6Addresses []*InstanceIpv6Address `locationName:"ipv6AddressesSet" locationNameList:"item" type:"list"`
	MacAddress *string `locationName:"macAddress" type:"string"`
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string"`
	OwnerId *string `locationName:"ownerId" type:"string"`
	PrivateDnsName *string `locationName:"privateDnsName" type:"string"`
	PrivateIpAddress *string `locationName:"privateIpAddress" type:"string"`
	PrivateIpAddresses []*InstancePrivateIpAddress `locationName:"privateIpAddressesSet" locationNameList:"item" type:"list"`
	SourceDestCheck *bool `locationName:"sourceDestCheck" type:"boolean"`
	Status *string `locationName:"status" type:"string" enum:"NetworkInterfaceStatus"`
	SubnetId *string `locationName:"subnetId" type:"string"`
	VpcId *string `locationName:"vpcId" type:"string"`
}


type InstanceNetworkInterfaceAttachment struct {
	AttachTime *time.Time `locationName:"attachTime" type:"timestamp"`
	AttachmentId *string `locationName:"attachmentId" type:"string"`
	DeleteOnTermination *bool `locationName:"deleteOnTermination" type:"boolean"`
	DeviceIndex *int64 `locationName:"deviceIndex" type:"integer"`
	Status *string `locationName:"status" type:"string" enum:"AttachmentStatus"`
}


type InstanceNetworkInterfaceAssociation struct {
	IpOwnerId *string `locationName:"ipOwnerId" type:"string"`
	PublicDnsName *string `locationName:"publicDnsName" type:"string"`
	PublicIp *string `locationName:"publicIp" type:"string"`
}

type InstancePrivateIpAddress struct {
	Association *InstanceNetworkInterfaceAssociation `locationName:"association" type:"structure"`
	Primary *bool `locationName:"primary" type:"boolean"`
	PrivateDnsName *string `locationName:"privateDnsName" type:"string"`
	PrivateIpAddress *string `locationName:"privateIpAddress" type:"string"`
}


type ProductCode struct {
	ProductCodeId *string `locationName:"productCode" type:"string"`
	ProductCodeType *string `locationName:"type" type:"string" enum:"ProductCodeValues"`
}

type GroupIdentifier struct {
	GroupId *string `locationName:"groupId" type:"string"`
	GroupName *string `locationName:"groupName" type:"string"`
}

type InstanceIpv6Address struct {
	Ipv6Address *string `locationName:"ipv6Address" type:"string"`
}

type Instance struct {
	AmiLaunchIndex *int64 `locationName:"amiLaunchIndex" type:"integer"`
	Architecture *string `locationName:"architecture" type:"string" enum:"ArchitectureValues"`
	BlockDeviceMappings []*InstanceBlockDeviceMapping `locationName:"blockDeviceMapping" locationNameList:"item" type:"list"`
	CapacityReservationId *string `locationName:"capacityReservationId" type:"string"`
	CapacityReservationSpecification *CapacityReservationSpecificationResponse `locationName:"capacityReservationSpecification" type:"structure"`
	ClientToken *string `locationName:"clientToken" type:"string"`
	CpuOptions *CpuOptions `locationName:"cpuOptions" type:"structure"`
	EbsOptimized *bool `locationName:"ebsOptimized" type:"boolean"`
	ElasticGpuAssociations []*ElasticGpuAssociation `locationName:"elasticGpuAssociationSet" locationNameList:"item" type:"list"`
	ElasticInferenceAcceleratorAssociations []*ElasticInferenceAcceleratorAssociation `locationName:"elasticInferenceAcceleratorAssociationSet" locationNameList:"item" type:"list"`
	EnaSupport *bool `locationName:"enaSupport" type:"boolean"`
	HibernationOptions *HibernationOptions `locationName:"hibernationOptions" type:"structure"`
	Hypervisor *string `locationName:"hypervisor" type:"string" enum:"HypervisorType"`
	IamInstanceProfile *IamInstanceProfile `locationName:"iamInstanceProfile" type:"structure"`
	ImageId *string `locationName:"imageId" type:"string"`
	InstanceId *string `locationName:"instanceId" type:"string"`
	InstanceLifecycle *string `locationName:"instanceLifecycle" type:"string" enum:"InstanceLifecycleType"`
	InstanceType *string `locationName:"instanceType" type:"string" enum:"InstanceType"`
	KernelId *string `locationName:"kernelId" type:"string"`
	KeyName *string `locationName:"keyName" type:"string"`
	LaunchTime *time.Time `locationName:"launchTime" type:"timestamp"`
	Licenses []*LicenseConfiguration `locationName:"licenseSet" locationNameList:"item" type:"list"`
	Monitoring *Monitoring `locationName:"monitoring" type:"structure"`
	NetworkInterfaces []*InstanceNetworkInterface `locationName:"networkInterfaceSet" locationNameList:"item" type:"list"`
	Placement *Placement `locationName:"placement" type:"structure"`
	Platform *string `locationName:"platform" type:"string" enum:"PlatformValues"`
	PrivateDnsName *string `locationName:"privateDnsName" type:"string"`
	PrivateIpAddress *string `locationName:"privateIpAddress" type:"string"`
	ProductCodes []*ProductCode `locationName:"productCodes" locationNameList:"item" type:"list"`
	PublicDnsName *string `locationName:"dnsName" type:"string"`
	PublicIpAddress *string `locationName:"ipAddress" type:"string"`
	RamdiskId *string `locationName:"ramdiskId" type:"string"`
	RootDeviceName *string `locationName:"rootDeviceName" type:"string"`
	RootDeviceType *string `locationName:"rootDeviceType" type:"string" enum:"DeviceType"`
	SecurityGroups []*GroupIdentifier `locationName:"groupSet" locationNameList:"item" type:"list"`
	SourceDestCheck *bool `locationName:"sourceDestCheck" type:"boolean"`
	SpotInstanceRequestId *string `locationName:"spotInstanceRequestId" type:"string"`
	SriovNetSupport *string `locationName:"sriovNetSupport" type:"string"`
	State *InstanceState `locationName:"instanceState" type:"structure"`
	StateReason *StateReason `locationName:"stateReason" type:"structure"`
	StateTransitionReason *string `locationName:"reason" type:"string"`
	SubnetId *string `locationName:"subnetId" type:"string"`
	Tags []*Tag `locationName:"tagSet" locationNameList:"item" type:"list"`
	VirtualizationType *string `locationName:"virtualizationType" type:"string" enum:"VirtualizationType"`
	VpcId *string `locationName:"vpcId" type:"string"`
}

type Tag struct {

    // The key of the tag.
    //
    // Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode
    // characters. May not begin with aws:.
    Key *string `locationName:"key" type:"string"`

    // The value of the tag.
    //
    // Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode
    // characters.
    Value *string `locationName:"value" type:"string"`
    // contains filtered or unexported fields
}

type StateReason struct {

    // The reason code for the state change.
    Code *string `locationName:"code" type:"string"`

    // The message for the state change.
    //
    //    * Server.InsufficientInstanceCapacity: There was insufficient capacity
    //    available to satisfy the launch request.
    //
    //    * Server.InternalError: An internal error caused the instance to terminate
    //    during launch.
    //
    //    * Server.ScheduledStop: The instance was stopped due to a scheduled retirement.
    //
    //    * Server.SpotInstanceShutdown: The instance was stopped because the number
    //    of Spot requests with a maximum price equal to or higher than the Spot
    //    price exceeded available capacity or because of an increase in the Spot
    //    price.
    //
    //    * Server.SpotInstanceTermination: The instance was terminated because
    //    the number of Spot requests with a maximum price equal to or higher than
    //    the Spot price exceeded available capacity or because of an increase in
    //    the Spot price.
    //
    //    * Client.InstanceInitiatedShutdown: The instance was shut down using the
    //    shutdown -h command from the instance.
    //
    //    * Client.InstanceTerminated: The instance was terminated or rebooted during
    //    AMI creation.
    //
    //    * Client.InternalError: A client error caused the instance to terminate
    //    during launch.
    //
    //    * Client.InvalidSnapshot.NotFound: The specified snapshot was not found.
    //
    //    * Client.UserInitiatedHibernate: Hibernation was initiated on the instance.
    //
    //    * Client.UserInitiatedShutdown: The instance was shut down using the Amazon
    //    EC2 API.
    //
    //    * Client.VolumeLimitExceeded: The limit on the number of EBS volumes or
    //    total storage was exceeded. Decrease usage or request an increase in your
    //    account limits.
    Message *string `locationName:"message" type:"string"`
    // contains filtered or unexported fields
}

type InstanceState struct {

    // The state of the instance as a 16-bit unsigned integer.
    //
    // The high byte is all of the bits between 2^8 and (2^16)-1, which equals decimal
    // values between 256 and 65,535. These numerical values are used for internal
    // purposes and should be ignored.
    //
    // The low byte is all of the bits between 2^0 and (2^8)-1, which equals decimal
    // values between 0 and 255.
    //
    // The valid values for instance-state-code will all be in the range of the
    // low byte and they are:
    //
    //    * 0 : pending
    //
    //    * 16 : running
    //
    //    * 32 : shutting-down
    //
    //    * 48 : terminated
    //
    //    * 64 : stopping
    //
    //    * 80 : stopped
    //
    // You can ignore the high byte value by zeroing out all of the bits above 2^8
    // or 256 in decimal.
    Code *int64 `locationName:"code" type:"integer"`

    // The current state of the instance.
    Name *string `locationName:"name" type:"string" enum:"InstanceStateName"`
    // contains filtered or unexported fields
}

type Placement struct {

    // The affinity setting for the instance on the Dedicated Host. This parameter
    // is not supported for the ImportInstance command.
    Affinity *string `locationName:"affinity" type:"string"`

    // The Availability Zone of the instance.
    //
    // If not specified, an Availability Zone will be automatically chosen for you
    // based on the load balancing criteria for the Region.
    AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

    // The name of the placement group the instance is in.
    GroupName *string `locationName:"groupName" type:"string"`

    // The ID of the Dedicated Host on which the instance resides. This parameter
    // is not supported for the ImportInstance command.
    HostId *string `locationName:"hostId" type:"string"`

    // The ARN of the host resource group in which to launch the instances. If you
    // specify a host resource group ARN, omit the Tenancy parameter or set it to
    // host.
    HostResourceGroupArn *string `locationName:"hostResourceGroupArn" type:"string"`

    // The number of the partition the instance is in. Valid only if the placement
    // group strategy is set to partition.
    PartitionNumber *int64 `locationName:"partitionNumber" type:"integer"`

    // Reserved for future use.
    SpreadDomain *string `locationName:"spreadDomain" type:"string"`

    // The tenancy of the instance (if the instance is running in a VPC). An instance
    // with a tenancy of dedicated runs on single-tenant hardware. The host tenancy
    // is not supported for the ImportInstance command.
    Tenancy *string `locationName:"tenancy" type:"string" enum:"Tenancy"`
    // contains filtered or unexported fields
}
type CpuOptions struct {

    // The number of CPU cores for the instance.
    CoreCount *int64 `locationName:"coreCount" type:"integer"`

    // The number of threads per CPU core.
    ThreadsPerCore *int64 `locationName:"threadsPerCore" type:"integer"`
    // contains filtered or unexported fields
}

type HibernationOptions struct {

    // If this parameter is set to true, your instance is enabled for hibernation;
    // otherwise, it is not enabled for hibernation.
    Configured *bool `locationName:"configured" type:"boolean"`
    // contains filtered or unexported fields
}

type IamInstanceProfile struct {

    // The Amazon Resource Name (ARN) of the instance profile.
    Arn *string `locationName:"arn" type:"string"`

    // The ID of the instance profile.
    Id *string `locationName:"id" type:"string"`
    // contains filtered or unexported fields
}

type Monitoring struct {

    // Indicates whether detailed monitoring is enabled. Otherwise, basic monitoring
    // is enabled.
    State *string `locationName:"state" type:"string" enum:"MonitoringState"`
    // contains filtered or unexported fields
}
	  
func main() {
		sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	// Create an EC2 service client.
	svc := ec2.New(sess)

	result, err := svc.DescribeInstances(nil)
	if err != nil {
	 fmt.Println("Error", err)
	} else {
	 fmt.Println("Success",reflect.TypeOf(result) )
	}


	var data DescribeInstancesOutput

	err = json.NewDecoder(result).Decode(&data)
	

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

	
}

