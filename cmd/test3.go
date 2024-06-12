package cmd

import (
	"github.com/gkwa/manyporter/test3"
	"github.com/spf13/cobra"
)

// test3Cmd represents the test3 command
var test3Cmd = &cobra.Command{
	Use:   "test3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		test3.Run()
	},
}

func init() {
	rootCmd.AddCommand(test3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
