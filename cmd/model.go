package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// genModelCmd represents the genModel command
var genModelCmd = &cobra.Command{
	Use:     "model",
	Short:   "Create Model for project",
	Long:    ``,
	Example: `cgen model -d "root:password@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local" -t bundle`,
	Run: func(cmd *cobra.Command, args []string) {
		var stderr bytes.Buffer
		dsn, _ := cmd.Flags().GetString("dsn")
		table, _ := cmd.Flags().GetString("table")

		outputDir, outputName := "./model/", cmd.Flag("table").Value.String()+".go"
		cmdd := exec.Command("gentool", "-dsn", dsn, "-tables", table, "-db", "mysql", "-onlyModel", "-outFile", outputName, "-outPath", outputDir)
		cmdd.Stderr = &stderr
		if err := cmdd.Run(); err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			os.Exit(2)
		}
		//rename model file
		err := os.Rename(outputDir+cmd.Flag("table").Value.String()+".gen.go", outputDir+outputName)
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			os.Exit(2)
		}
		_, _ = cmd.OutOrStdout().Write([]byte("model create successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(genModelCmd)

	genModelCmd.Flags().StringP("dsn", "d", "", "dsn info for the database connection")
	genModelCmd.Flags().StringP("table", "t", "", "database table name")
	_ = genModelCmd.MarkFlagRequired("dsn")
	_ = genModelCmd.MarkFlagRequired("table")
}
