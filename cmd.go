package dog

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Dog いぬさま",
	Short: "Your good helper for development 你的开发好帮手",
	Long: `
▓█████▄  ▒█████    ▄████    
▒██▀ ██▌▒██▒  ██▒ ██▒ ▀█▒   
░██   █▌▒██░  ██▒▒██░▄▄▄░   
░▓█▄   ▌▒██   ██░░▓█  ██▓   
░▒████▓ ░ ████▓▒░░▒▓███▀▒   
 ▒▒▓  ▒ ░ ▒░▒░▒░  ░▒   ▒    
 ░ ▒  ▒   ░ ▒ ▒░   ░   ░    
 ░ ░  ░ ░ ░ ░ ▒  ░ ░   ░    
   ░        ░ ░        ░    
 ░                        
https://github.com/dollarkillerx/dog`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
