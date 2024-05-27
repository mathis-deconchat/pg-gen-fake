package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os/exec"
)

var Hello string

// generatorCmd represents the generator command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen", "g"},
	Short:   "Generate data for an entire database with integrity",
	Long: `Respect Foreign Keys, Unique Constraints, indexes, and more.
	It also fakes data to be as close as possible to the real data.
	Like phone numbers, emails, names, etc. [WIP]`,
	Run: func(cmd *cobra.Command, args []string) {
		dockerFlag, _ := cmd.Flags().GetString("docker")
		viper.Set("use_docker", dockerFlag)

		// Save the updated config if needed
		err := viper.WriteConfig()
		if err != nil {
			fmt.Printf("Error writing config: %s\n", err)
		}

		checkFlag, _ := cmd.Flags().GetBool("check")
		if checkFlag {
			Logger.Info("Checking for different tools")
			if checkForDockerIsInstalled() {
				checkOs()
				checkIfDefaultCredentialsForPostgresAreInConfig()
				checkIfNetworkIsAccessible()
				checkIfGitIsInstalled()
				Logger.Info("All checks passed")

			}
		}

	},
}

func init() {
	RootCmd.AddCommand(generateCmd)

	// Add a flag to set use_docker value
	generateCmd.Flags().StringP("docker", "d", "true", "Set use_docker value")
	viper.BindPFlag("use_docker", generateCmd.Flags().Lookup("docker"))

	generateCmd.Flags().BoolP("check", "c", true, "Check for different tools")
	viper.BindPFlag("check", generateCmd.Flags().Lookup("check"))

	generateCmd.Flags().StringVarP(&Hello, "hello", "w", "Hello", "A help for foo")
}

func checkForDockerIsInstalled() bool {
	cmd := exec.Command("docker", "--version")
	err := cmd.Run()
	if err != nil {
		return false
	}

	Logger.Infof("🟢 Docker is installed")
	return true
}

func checkOs() {
	cmd := exec.Command("uname")
	_, err := cmd.Output()
	if err != nil {
		Logger.Errorf("🔴 Not linux...")
		return 
	}
	Logger.Infof("🟢 Linux is installed")

}

func checkIfDefaultCredentialsForPostgresAreInConfig() bool {
	if viper.GetString("postgres.user") == "postgres" && viper.GetString("postgres.password") == "postgres" {
		Logger.Infof("🟢 Postgres credentials set")
		return true
	}
	if len(viper.GetString("postgres_url")) > 0 {
		Logger.Infof("🟢 Postgres credentials set")
		return true
	}
	return false
}

func checkIfNetworkIsAccessible() bool {
	cmd := exec.Command("ping", "-c", "1", "8.8.8.8")
	err := cmd.Run()
	if err != nil {
		Logger.Errorf("🔴 Network is not accessible")
		return false
	}
	Logger.Infof("🟢 Network is accessible")
	return err == nil
}

func checkIfGitIsInstalled() bool {
	cmd := exec.Command("git", "--version")
	err := cmd.Run()
	if err != nil {
		Logger.Errorf("🔴 Git is not installed")
		return false
	}
	Logger.Infof("🟢 Git is installed")
	return true
}
