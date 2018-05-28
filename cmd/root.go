package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "kafkaproducer",
	Short: "kafka data producer",
	Long:  "read data from log file, reformat the data and put them into kafka",
	// RunE:         server,
	SilenceUsage: true,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
