package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "zhihu-golang-web",
	Short: "zhihu-golang-web sample",
	Long:  "zhihu-golang-web sample",
	// RunE:         server,
	SilenceUsage: true,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
