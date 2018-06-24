package cmd

import (
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
			sshInfo := SSHInfo{}
			sshInfo.User = "ec2-user"
			sshInfo.Host = ec2.PrivateIpAddress
			sshInfo.KeyPath = ""
			err := sshInfo.execRemoteCommand("docker ps | grep AAAA | wc -l")
			if err != nil {
				log.Fatal(err)
			}
			if sshInfo.Result != "0" {
				err = sshInfo.execRemoteCommand("docker ps | grep XXXXXX")
				log.Println(sshInfo.Result)
			}
		}
	},
}
