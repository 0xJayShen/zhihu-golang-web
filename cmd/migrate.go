package cmd

import "github.com/spf13/cobra"

// versionCmd represents the version command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "create database & tables",
	Long:  "Create database & tables that application used.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := readConfig()
		if err != nil{
			return err
		}
		err = migrate(cmd, args)
		if err != nil{
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")
}

func migrate(cmd *cobra.Command, args []string) error {
    return nil
}

