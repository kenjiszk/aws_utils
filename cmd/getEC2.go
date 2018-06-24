package cmd

import (
	"fmt"
	"log"
	//"strconv"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(getEC2)
	RootCmd.PersistentFlags().StringP("name", "a", "", "EC2 name for filtering.")
}

var getEC2 = &cobra.Command{
	Use:   "getEC2",
	Short: "get EC2 list with filter.",
	Long:  "get EC2 list with filter.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := awsEnvValidator(); err != nil {
			log.Fatal(err)
		}
		ec2List := EC2List{}
		ec2List.GetEC2ByFilter()
		for _, ec2 := range ec2List.EC2s {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\n", ec2.Name, ec2.InstanceId, ec2.InstanceType, ec2.PrivateIpAddress, ec2.PublicIpAddress, ec2.VPC, ec2.State)
		}
	},
}

