package cmd

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var makeCommandCmd = &cobra.Command{
	Use:                   "make:command [name]",
	Short:                 "Создать команду.",
	Long:                  `Создать команду.`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		originalName, err := lo.Nth(args, 0)
		if err != nil {
			return errors.New("название команды не указано")
		}
		snakeName := lo.SnakeCase(originalName)
		camelName := lo.CamelCase(originalName)

		wd, err := os.Getwd()
		if err != nil {
			return errors.New(err.Error())
		}

		commandFileName := fmt.Sprintf("%s/cmd/%s.go", wd, snakeName)

		if _, err := os.Stat(commandFileName); err == nil {
			return errors.New(fmt.Sprintf("файл %s уже существует", commandFileName))
		}

		file, err := os.Create(commandFileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		commandContent := strings.ReplaceAll(TEMPLATE, "{{ .CmdName }}", originalName)
		commandContent = strings.ReplaceAll(commandContent, "{{ .CmdVarName }}", camelName)
		_, err = file.WriteString(commandContent)
		if err != nil {
			return errors.New(err.Error())
		}

		fmt.Printf("%s created at %s\n", originalName, commandFileName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(makeCommandCmd)
}

const TEMPLATE = `package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// {{ .CmdVarName }}Cmd represents the {{ .CmdVarName }} command
var {{ .CmdVarName }}Cmd = &cobra.Command{
	Use:   "{{ .CmdName }}",
	Short: "A brief description of your command",
	Long: ` + "`" + `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.` + "`" + `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("{{ .CmdVarName }} called")
	},
}

func init() {
	rootCmd.AddCommand({{ .CmdVarName }}Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// {{ .CmdVarName }}Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// {{ .CmdVarName }}Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
`
