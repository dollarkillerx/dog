package dog

import (
	"fmt"

	"github.com/dollarkillerx/processes"
	"github.com/spf13/cobra"
)

func init() {
	dockerCmd.AddCommand(dockerClean)

	// main
	rootCmd.AddCommand(dockerCmd)
}

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker Tools",
	Long:  "Docker Tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dockerClean = &cobra.Command{
	Use:   "clean",
	Short: "rm none",
	Long:  "1.stop none container 2.rm none container 3.rm none image",
	Run: func(cmd *cobra.Command, args []string) {
		command, _ := processes.RunCommand(`docker stop $(docker ps -a | grep "Exited" | awk '{print $1 }')`)
		fmt.Println("Stop None Container")
		fmt.Println(command)

		command, _ = processes.RunCommand(`docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')`)
		fmt.Println("Rm None Container")
		fmt.Println(command)

		command, _ = processes.RunCommand(`docker rmi $(docker images | grep "none" | awk '{print $3}')`)
		fmt.Println("Rm None Image")
		fmt.Println(command)

		fmt.Println("docker clean success")
	},
}
