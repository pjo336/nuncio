package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const defaultConfigFileName = ".nuncio"

var configFileLocation string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate the base inventory file for nuncio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		if _, err := os.Stat(configLocation()); !os.IsNotExist(err) {
			// .nuncio file already exists, error
			fmt.Println(".nuncio config file already exists, cannot init")
			os.Exit(1)
		}
		_, err := os.Create(configLocation())
		errorExit(err)
		fmt.Printf("Base config file for nuncio generated at %s\n", configLocation())
	},
}

// configLocation locates the configuration file path at the users homedir/.nuncio
func configLocation() string {
	if configFileLocation == "" {
		home, err := homedir.Dir()
		errorExit(err)
		return fmt.Sprintf("%s/%s", home, configFileLocation)
	}
	return configFileLocation
}

// errorExit handles exiting the command if an error occurs, printing what the error was
func errorExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	initCmd.Flags().StringVarP(&configFileLocation, "config", "c", "", "Set where the base configuration file is located")
	RootCmd.AddCommand(initCmd)
}
