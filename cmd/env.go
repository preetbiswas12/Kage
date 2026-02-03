package cmd

import (
	"github.com/metafates/mangal/color"
	"github.com/metafates/mangal/config"
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/style"
	"github.com/metafates/mangal/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.Flags().BoolP("set-only", "s", false, "only show variables that are set")
	envCmd.Flags().BoolP("unset-only", "u", false, "only show variables that are unset")

	envCmd.MarkFlagsMutuallyExclusive("set-only", "unset-only")
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "List available environment variables",
	Long: `Display all environment variables that can be used to configure mangal.

Environment variables provide an alternative way to configure mangal
without modifying the configuration file. They are especially useful
in containerized environments or CI/CD pipelines.`,
	Example: `  # Show all available environment variables
  mangal env

  # Show only variables that are currently set
  mangal env --set-only

  # Show only unset variables
  mangal env --unset-only`,
	Run: func(cmd *cobra.Command, args []string) {
		setOnly := lo.Must(cmd.Flags().GetBool("set-only"))
		unsetOnly := lo.Must(cmd.Flags().GetBool("unset-only"))

		config.EnvExposed = append(config.EnvExposed, where.EnvConfigPath)
		slices.Sort(config.EnvExposed)
		for _, env := range config.EnvExposed {
			if env != where.EnvConfigPath {
				env = strings.ToUpper(constant.Mangal + "_" + config.EnvKeyReplacer.Replace(env))
			}
			value := os.Getenv(env)
			present := value != ""

			if setOnly || unsetOnly {
				if !present && setOnly {
					continue
				}

				if present && unsetOnly {
					continue
				}
			}

			cmd.Print(style.New().Bold(true).Foreground(color.Purple).Render(env))
			cmd.Print("=")

			if present {
				cmd.Println(style.Fg(color.Green)(value))
			} else {
				cmd.Println(style.Fg(color.Red)("unset"))
			}
		}
	},
}
