package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"os"
)

type EC2Info struct {
	InstanceId       string
	InstanceType     string
	PrivateIpAddress string
	PublicIpAddress  string
	Name             string
	VPC              string
	State            string
}

type EC2List struct {
	EC2s   []EC2Info
	Filter []*ec2.Filter
}

func (e *EC2List) GetEC2ByFilter() {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})

	nextToken := aws.String("")
	keepLoop := true
	maxResults := int64(5)

	for keepLoop {
		params := &ec2.DescribeInstancesInput{
			/* Filters: []*ec2.Filter{
				{
					Name: aws.String("vpc-id"),
					Values: []*string{
						aws.String(vpcId),
					},
				},
				{
					Name: aws.String("instance-state-name"),
					Values: []*string{
						aws.String("running"),
						aws.String("stopping"),
						aws.String("stopped"),
					},
				},
			},*/
			MaxResults: aws.Int64(maxResults),
			NextToken:  nextToken,
		}
		resp, err := svc.DescribeInstances(params)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		// log.Println(resp)
		resultInstanceNum := int64(0)
		for _, resv := range resp.Reservations {
			for _, ins := range resv.Instances {
				resultInstanceNum++
				ec2Info := EC2Info{}
				ec2Info.InstanceId = *ins.InstanceId
				ec2Info.InstanceType = *ins.InstanceType
				ec2Info.PrivateIpAddress = *ins.PrivateIpAddress
				ec2Info.PublicIpAddress = *ins.PublicIpAddress
				ec2Info.State = *ins.State.Name
				ec2Info.VPC = *ins.VpcId
				ec2Info.Name = getEC2Name(ins.Tags)
				e.EC2s = append(e.EC2s, ec2Info)
			}
		}
		if resultInstanceNum == maxResults {
			nextToken = resp.NextToken
		} else {
			keepLoop = false
		}
	}
}

func getEC2Name(tags []*ec2.Tag) string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}
	return "No-Name"
}
