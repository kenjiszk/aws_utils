package cmd

import (
	"fmt"
	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(railsConsole)
}

var railsConsole = &cobra.Command{
	Use:   "railsConsole",
	Short: "execute rails console in target container.",
	Long:  "execute rails console in target container.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := awsEnvValidator(); err != nil {
			log.Fatal(err)
		}
		if err := railsConsoleValidator(); err != nil {
			log.Fatal(err)
		}
		ec2List := EC2List{}
		ec2List.Filter = []*ec2.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String(os.Getenv("ECS_SERVER_NAME"))},
			},
		}
		ec2List.GetEC2ByFilter()

		for _, ec2 := range ec2List.EC2s {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\n", ec2.Name, ec2.InstanceId, ec2.InstanceType, ec2.PrivateIpAddress, ec2.PublicIpAddress, ec2.VPC, ec2.State)
		}
	},
}
