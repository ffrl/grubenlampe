package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gl",
	Short: "FFRL Grubenlampe",
	Long:  "Grubenlampe (gl) is a client for the Grubenlampe tunnel as a service service",
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff here
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("user", "u", "", "Username")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Password")
	rootCmd.PersistentFlags().StringP("server", "s", "[::1]:20170", "Server")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
