package cmd

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/kardianos/osext"
	"github.com/spf13/cobra"
	"github.com/asdfsx/zhihu-golang-web/helpers"
	"github.com/asdfsx/zhihu-golang-web/server"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of KafkaProducer",
	Long:  "All software has versions. This is KafkaProducer's.",
	RunE: func(cmd *cobra.Command, args []string) error {
		printServerVersion()
		return nil
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func printServerVersion() {
	if server.BuildDate == "" {
		setBuildDate() // set the build date from executable's mdate
	} else {
		formatBuildDate() // format the compile time
	}
	if server.CommitHash == "" {
		jww.FEEDBACK.Printf("server v%s %s/%s BuildDate: %s\n", helpers.ServerVersion(), runtime.GOOS, runtime.GOARCH, server.BuildDate)
	} else {
		jww.FEEDBACK.Printf("server v%s-%s %s/%s BuildDate: %s\n", helpers.ServerVersion(), strings.ToUpper(server.CommitHash), runtime.GOOS, runtime.GOARCH, server.BuildDate)
	}
}

func setBuildDate() {
	fname, _ := osext.Executable()
	dir, err := filepath.Abs(filepath.Dir(fname))
	if err != nil {
		jww.ERROR.Println(err)
		return
	}
	fi, err := os.Lstat(filepath.Join(dir, filepath.Base(fname)))
	if err != nil {
		jww.ERROR.Println(err)
		return
	}
	t := fi.ModTime()
	server.BuildDate = t.Format(time.RFC3339)
}

func formatBuildDate() {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", server.BuildDate)
	server.BuildDate = t.Format(time.RFC3339)
}
