package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type ServiceConfig struct {
	ServiceName   string
	ServiceExport string
}

// genServiceCmd represents the genService command
var genServiceCmd = &cobra.Command{
	Use:     "service",
	Short:   "Create Service for project",
	Long:    ``,
	Example: `cgen service bundle`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Example)
			os.Exit(2)
		}

		conf := new(ServiceConfig)
		conf.ServiceName = args[0] + "Service"
		conf.ServiceExport = ToCamelInitCase(args[0], true)
		genFile("template/service.go.tmpl", filepath.Join("service", args[0]+".go"), conf)

		_, _ = cmd.OutOrStdout().Write([]byte("service create successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(genServiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genServiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genServiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
