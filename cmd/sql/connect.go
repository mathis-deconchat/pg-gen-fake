/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package sql

import (
	"fmt"

	"github.com/mathis-deconchat/pg-gen-fake/cmd"
	"github.com/spf13/cobra"
)

// sql/sqlCmd represents the sql/sql command
var sqlCmd = &cobra.Command{
	Use:   "connect",
	Aliases: []string{"conn", "c"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sql/sql called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(sqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sql/sqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sql/sqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
