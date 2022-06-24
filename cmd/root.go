/*
Copyright © 2022 Moeid Heidari

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//=================================================================================================================================================================================
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "Gitea batch repo aggregator programm",
	Long: `                                                        
	.d8888b.  8888888 888888b.   8888888b.  8888888888        d8888 
	d88P  Y88b   888   888  "88b  888   Y88b 888              d88888 
	888    888   888   888  .88P  888    888 888             d88P888 
	888          888   8888888K.  888   d88P 8888888        d88P 888 
	888  88888   888   888  "Y88b 8888888P"  888           d88P  888 
	888    888   888   888    888 888 T88b   888          d88P   888 
	Y88b  d88P   888   888   d88P 888  T88b  888         d8888888888 
	 "Y8888P88 8888888 8888888P"  888   T88b 8888888888 d88P     888
	
	
    
Gibrea is a batch aggregator program that aggregates a bunch of git repositories from different sources and migrate them to a git repository system`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

//=================================================================================================================================================================================
var CmdStart = &cobra.Command{
	Use:   "start",
	Short: "starts the gibrea programm",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		config, _ := cmd.Flags().GetString("config")

		fmt.Println(config)

	},
}

//=================================================================================================================================================================================
//=================================================================================================================================================================================
//=================================================================================================================================================================================
//=================================================================================================================================================================================
//=================================================================================================================================================================================
//=================================================================================================================================================================================
//=================================================================================================================================================================================
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(CmdStart)
	CmdStart.PersistentFlags().String("config", "", "name of the user")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

//=================================================================================================================================================================================
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.main.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//=================================================================================================================================================================================
