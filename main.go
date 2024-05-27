/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/mathis-deconchat/pg-gen-fake/cmd"
)

var Logger *log.Logger


func main() {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: false,
		Level: log.InfoLevel,
		TimeFormat: "15:04:05",
	})
	cmd.Execute()
}
