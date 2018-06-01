package cmd

import (
	"github.com/spf13/cobra"
	"github.com/asdfsx/zhihu-golang-web/migrate"
)

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
		err = migrateFunc(cmd, args)
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

func migrateFunc(cmd *cobra.Command, args []string) error {
	err := migrate.ConnectDB(config.Database.Type, config.Database.User, config.Database.Passwd,
		config.Database.Host, config.Database.Port, config.Database.DBName, config.Database.TablePrefix)
	if err != nil{
		return err
	}
	defer migrate.Close()
	migrate.Migrate()
	return nil
}

