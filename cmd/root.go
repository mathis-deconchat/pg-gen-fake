/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/charmbracelet/log"

)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pg-gen-fake",
	Aliases: []string{"pg-gen-fake", "pgf"},
	Short: "A tool to generate fake data for an entire database, with integrity ",
	Long: `
--------------------------------------------------------------------------------------------------
	######    #####           #####   #######  #     #         #######     #     #    #  #######
	#     #  #     #         #     #  #        ##    #         #          # #    #   #   #      
	#     #  #               #        #        # #   #         #         #   #   #  #    #      
	######   #  ####  #####  #  ####  #####    #  #  #  #####  #####    #     #  ###     #####  
	#        #     #         #     #  #        #   # #         #        #######  #  #    #      
	#        #     #         #     #  #        #    ##         #        #     #  #   #   #      
	#         #####           #####   #######  #     #         #        #     #  #    #  #######
																		   ❇️ Mathis Deconchat 2024
--------------------------------------------------------------------------------------------------

pg-gen-fake is a tool to generate fake data for an entire database, with integrity.`,

	// Run: func(cmd *cobra.Command, args []string) { },
}
 var Logger *log.Logger



// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: false,
		Level: log.InfoLevel,
		TimeFormat: "15:04:05",
	})
	

	if viper.GetString("log_level") == "debug" {
	Logger.SetLevel(log.DebugLevel)
	}
	if viper.GetString("log_level") == "error" {
	Logger.SetLevel(log.ErrorLevel)
	}

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.BindPFlag("message", RootCmd.PersistentFlags().Lookup("message"))


}

func initConfig(){
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		Logger.Error("Error reading config file", err)
	}
	
}


