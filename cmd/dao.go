package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type DaoConfig struct {
	DaoName   string
	DaoExport string
}

// genDaoCmd represents the genDao command
var genDaoCmd = &cobra.Command{
	Use:     "dao",
	Short:   "Create Dao for project",
	Long:    ``,
	Example: `cgen dao games`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Example)
			os.Exit(2)
		}

		conf := new(DaoConfig)
		conf.DaoName = args[0] + "Dao"
		conf.DaoExport = ToCamelInitCase(args[0], true)
		genFile("template/dao.go.tmpl", filepath.Join("dao", args[0]+".go"), conf)

		_, _ = cmd.OutOrStdout().Write([]byte("dao create successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(genDaoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genDaoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genDaoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
