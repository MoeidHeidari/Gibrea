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
		assService, _ := cmd.Flags().GetBool("as-service")
		fmt.Println(config, assService)

	},
}

//=================================================================================================================================================================================
var CmdStop = &cobra.Command{
	Use:   "stop",
	Short: "stops the program with state persisted",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goodbye:(")

	},
}

//=================================================================================================================================================================================
var CmdStatus = &cobra.Command{
	Use:   "status",
	Short: "shows the current status of the programm",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("current status")

	},
}

//=================================================================================================================================================================================
var CmdPause = &cobra.Command{
	Use:   "pause",
	Short: "pauses the prrogramm execution temorarily with last state persisted",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("program paused")

	},
}

//=================================================================================================================================================================================
var CmdLogs = &cobra.Command{
	Use:   "logs",
	Short: "shows the the logs of running program",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		printLogFilePath, _ := cmd.Flags().GetString("print")
		fmt.Println(printLogFilePath)

	},
}

//=================================================================================================================================================================================
var CmdResum = &cobra.Command{
	Use:   "resum",
	Short: "reads the last state of the program and resumes the execution",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("program resumed")

	},
}

//=================================================================================================================================================================================
//=================================================================================================================================================================================
//=================================================================================================================================================================================
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(CmdStart)
	CmdStart.PersistentFlags().String("config", "", "path to the configuration file")
	CmdStart.PersistentFlags().Bool("as-service", false, "runs the program as a background service")
	rootCmd.AddCommand(CmdStop)
	rootCmd.AddCommand(CmdStatus)
	rootCmd.AddCommand(CmdLogs)
	CmdLogs.PersistentFlags().String("print", "", "path to the log file you want to print the logs inside")
	rootCmd.AddCommand(CmdPause)
	rootCmd.AddCommand(CmdResum)
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
