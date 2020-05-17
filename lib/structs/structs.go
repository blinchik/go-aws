package structs

import "time"

type DescribeVolumesOutput struct {

	// The NextToken value to include in a future DescribeVolumes request. When
	// the results of a DescribeVolumes request exceed MaxResults, this value can
	// be used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the volumes.
	Volumes []*Volume `locationName:"volumeSet" locationNameList:"item" type:"list"`
	// contains filtered or unexported fields
}

type Volume struct {

	// Information about the volume attachments.
	Attachments []*VolumeAttachment `locationName:"attachmentSet" locationNameList:"item" type:"list"`

	// The Availability Zone for the volume.
	AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

	// The time stamp when volume creation was initiated.
	CreateTime *time.Time `locationName:"createTime" type:"timestamp"`

	// Indicates whether the volume is encrypted.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// Indicates whether the volume was created using fast snapshot restore.
	FastRestored *bool `locationName:"fastRestored" type:"boolean"`

	// The number of I/O operations per second (IOPS) that the volume supports.
	// For Provisioned IOPS SSD volumes, this represents the number of IOPS that
	// are provisioned for the volume. For General Purpose SSD volumes, this represents
	// the baseline performance of the volume and the rate at which the volume accumulates
	// I/O credits for bursting. For more information, see Amazon EBS Volume Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Constraints: Range is 100-16,000 IOPS for gp2 volumes and 100 to 64,000IOPS
	// for io1 volumes, in most Regions. The maximum IOPS for io1 of 64,000 is guaranteed
	// only on Nitro-based instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).
	// Other instance families guarantee performance up to 32,000 IOPS.
	//
	// Condition: This parameter is required for requests to create io1 volumes;
	// it is not used in requests to create gp2, st1, sc1, or standard volumes.
	Iops *int64 `locationName:"iops" type:"integer"`

	// The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS)
	// customer master key (CMK) that was used to protect the volume encryption
	// key for the volume.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// Indicates whether Amazon EBS Multi-Attach is enabled.
	MultiAttachEnabled *bool `locationName:"multiAttachEnabled" type:"boolean"`

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string `locationName:"outpostArn" type:"string"`

	// The size of the volume, in GiBs.
	Size *int64 `locationName:"size" type:"integer"`

	// The snapshot from which the volume was created, if applicable.
	SnapshotId *string `locationName:"snapshotId" type:"string"`

	// The volume state.
	State *string `locationName:"status" type:"string" enum:"VolumeState"`

	// Any tags assigned to the volume.
	Tags []*Tag `locationName:"tagSet" locationNameList:"item" type:"list"`

	// The ID of the volume.
	VolumeId *string `locationName:"volumeId" type:"string"`

	// The volume type. This can be gp2 for General Purpose SSD, io1 for Provisioned
	// IOPS SSD, st1 for Throughput Optimized HDD, sc1 for Cold HDD, or standard
	// for Magnetic volumes.
	VolumeType *string `locationName:"volumeType" type:"string" enum:"VolumeType"`
	// contains filtered or unexported fields
}

type VolumeAttachment struct {

	// The time stamp when the attachment initiated.
	AttachTime *time.Time `locationName:"attachTime" type:"timestamp"`

	// Indicates whether the EBS volume is deleted on instance termination.
	DeleteOnTermination *bool `locationName:"deleteOnTermination" type:"boolean"`

	// The device name.
	Device *string `locationName:"device" type:"string"`

	// The ID of the instance.
	InstanceId *string `locationName:"instanceId" type:"string"`

	// The attachment state of the volume.
	State *string `locationName:"status" type:"string" enum:"VolumeAttachmentState"`

	// The ID of the volume.
	VolumeId *string `locationName:"volumeId" type:"string"`
	// contains filtered or unexported fields
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
