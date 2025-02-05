package options

import (
	"fmt"
	"os"
	"strconv"
)

var (
	AWS_AMI                  = "AWS_AMI"
	AWS_DISK_SIZE            = "AWS_DISK_SIZE"
	AWS_INSTANCE_TYPE        = "AWS_INSTANCE_TYPE"
	AWS_REGION               = "AWS_REGION"
	AWS_SECURITY_GROUP_ID    = "AWS_SECURITY_GROUP_ID"
	AWS_SUBNET_ID            = "AWS_SUBNET_ID"
	AWS_VPC_ID               = "AWS_VPC_ID"
	AWS_INSTANCE_PROFILE_ARN = "AWS_INSTANCE_PROFILE_ARN"
)

type Options struct {
	DiskImage          string
	DiskSizeGB         int
	MachineFolder      string
	MachineID          string
	MachineType        string
	VpcID              string
	SubnetID           string
	SecurityGroupID    string
	InstanceProfileArn string
	Zone               string
}

func FromEnv(init bool) (*Options, error) {
	retOptions := &Options{}

	var err error

	retOptions.MachineType, err = fromEnvOrError(AWS_INSTANCE_TYPE)
	if err != nil {
		return nil, err
	}

	diskSizeGB, err := fromEnvOrError(AWS_DISK_SIZE)
	if err != nil {
		return nil, err
	}

	retOptions.DiskSizeGB, err = strconv.Atoi(diskSizeGB)
	if err != nil {
		return nil, err
	}

	retOptions.DiskImage = os.Getenv(AWS_AMI)
	retOptions.SecurityGroupID = os.Getenv(AWS_SECURITY_GROUP_ID)
	retOptions.SubnetID = os.Getenv(AWS_SUBNET_ID)
	retOptions.VpcID = os.Getenv(AWS_VPC_ID)
	retOptions.InstanceProfileArn = os.Getenv(AWS_INSTANCE_PROFILE_ARN)
	retOptions.Zone = os.Getenv(AWS_REGION)

	// Return eraly if we're just doing init
	if init {
		return retOptions, nil
	}

	retOptions.MachineID, err = fromEnvOrError("MACHINE_ID")
	if err != nil {
		return nil, err
	}
	// prefix with devpod-
	retOptions.MachineID = "devpod-" + retOptions.MachineID

	retOptions.MachineFolder, err = fromEnvOrError("MACHINE_FOLDER")
	if err != nil {
		return nil, err
	}

	return retOptions, nil
}

func fromEnvOrError(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf(
			"couldn't find option %s in environment, please make sure %s is defined",
			name,
			name,
		)
	}

	return val, nil
}
