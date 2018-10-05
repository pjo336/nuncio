package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// VERSION is set during build
	VERSION string
)

var cfgFile string
var Env string

var RootCmd = &cobra.Command{
	Use:   "nuncio",
	Short: "Nuncio microservice development orchestration",
	Long:  "Nuncio facilitates running many docker based microservices in network without docker compose",
}

func Execute(version string) {
	VERSION = version
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Find home directory.
	// home, err := homedir.Dir()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// f, err := ioutil.ReadFile(initFile)
}
