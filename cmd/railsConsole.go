package cmd

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
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
				Values: []*string{aws.String(os.Getenv("EC2_SERVER_NAME"))},
			},
		}
		ec2List.GetEC2ByFilter()

		sshInfo, err := getTargeCID(ec2List.EC2s)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Please execute this command [ ssh -t %s 'docker exec -it %s bundle exec rails console' ]\n", sshInfo.Host, sshInfo.Result)
	},
}

func getTargeCID(ec2s []EC2Info) (SSHInfo, error) {
	checkCommand := fmt.Sprintf("docker ps | grep ecs-%s- | wc -l", os.Getenv("ECS_SERVICE_NAME"))
	getCommand := fmt.Sprintf("docker ps | grep ecs-%s- | head -1 | cut -f1 -d' '", os.Getenv("ECS_SERVICE_NAME"))
	for _, ec2 := range ec2s {
		sshInfo := SSHInfo{}
		sshInfo.User = "ec2-user"
		sshInfo.Host = ec2.PrivateIpAddress
		log.Println(checkCommand)
		err := sshInfo.execRemoteCommand(checkCommand)
		if err != nil {
			log.Fatal(err)
		}
		sshInfo.Result = strings.TrimRight(sshInfo.Result, "\n")
		if sshInfo.Result != "0" {
			log.Println(getCommand)
			err = sshInfo.execRemoteCommand(getCommand)
			sshInfo.Result = strings.TrimRight(sshInfo.Result, "\n")
			return sshInfo, nil
		}
	}
	return SSHInfo{}, errors.New("There is no target container.")
}
