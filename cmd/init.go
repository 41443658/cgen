package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
)

type InitConfig struct {
	ProjectDir string
	ModuleName string
	GoVersion  string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init the cortex internal framework scaffold",
	Long: `The cortex-scaffold init command creates a new gin application. will install the Gin framework and gentool
Example:`,
	Example: `cgen init blog
cgen init /var/wwww/cortex-blog
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Long, cmd.Example)
			os.Exit(2)
		}
		projectDir, err := filepath.Abs(args[0])
		if err != nil {
			panic(err)
		}

		conf := new(InitConfig)
		conf.ProjectDir, conf.ModuleName = projectDir, filepath.Base(projectDir)

		reg1 := regexp.MustCompile(`go(\d+.\d+).\d+`)
		conf.GoVersion = reg1.FindStringSubmatch(runtime.Version())[1]

		initScaffold(conf)
	},
}

func initScaffold(cnf *InitConfig) {
	//1. create the package
	fmt.Println("start to init the project structure")
	packages := []string{"conf", "api", "dao", "model", "router"}
	for _, v := range packages {
		dir := filepath.Join(cnf.ProjectDir, v)
		_ = os.MkdirAll(dir, 00755)
	}

	genFile("template/main.go.tmpl", filepath.Join(cnf.ProjectDir, "main.go"), cnf)
	genFile("template/router.go.tmpl", filepath.Join(cnf.ProjectDir, "router/router.go"), cnf)
	genFile("template/go.mod.tmpl", filepath.Join(cnf.ProjectDir, "go.mod"), cnf)

	fmt.Println("create the structure successfully!")
	fmt.Println("start to install the gin framework and other dependency")
	//execute the `go mod tidy`
	cmdd := exec.Command("go", "mod", "tidy")
	cmdd.Dir = cnf.ProjectDir
	if err := cmdd.Run(); err != nil {
		panic(err)
	}
	fmt.Println("gin framework and other dependency install successfully!")
	fmt.Println("start to install gentool")
	//install gentool
	installcmd := exec.Command("go", "install", "gorm.io/gen/tools/gentool@latest")
	if err := installcmd.Run(); err != nil {
		panic(err)
	}
	fmt.Println("project init done")
}

func init() {
	rootCmd.AddCommand(initCmd)

	//initCmd.MarkFlagDirname()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
