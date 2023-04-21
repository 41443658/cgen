package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type ApiConfig struct {
	ApiName   string
	ApiExport string
}

// genApiCmd represents the genApi command
var genApiCmd = &cobra.Command{
	Use:     "api",
	Short:   "Generate API entry for project",
	Long:    ``,
	Example: `cgen api games`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Example)
			os.Exit(2)
		}

		conf := new(ApiConfig)
		conf.ApiName = args[0] + "Api"
		conf.ApiExport = ToCamelInitCase(args[0], true)
		genFile("template/api.go.tmpl", filepath.Join("api", args[0]+".go"), conf)

		_, _ = cmd.OutOrStdout().Write([]byte("api create successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(genApiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
