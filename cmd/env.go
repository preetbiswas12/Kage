package cmd

import (
	"os"
	"strings"

	"github.com/preetbiswas12/Kage/color"
	"github.com/preetbiswas12/Kage/config"
	"github.com/preetbiswas12/Kage/constant"
	"github.com/preetbiswas12/Kage/style"
	"github.com/preetbiswas12/Kage/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
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
	Long: `Display all environment variables that can be used to configure Kage.

Environment variables provide an alternative to configuration files and are
especially useful in containerized environments or CI/CD pipelines.

Each variable controls a specific setting like download directory, default source,
output format, etc.

FLAGS:
  -s, --set-only     Show only variables that are currently set
  -u, --unset-only   Show only variables that are not set

EXAMPLES:
  # Show all available environment variables
  kage env
  
  # Show only set variables
  kage env --set-only
  
  # Show only unset variables
  kage env --unset-only
  
  # Use environment variables
  export KAGE_DOWNLOAD_DIR="$HOME/Manga"
  export KAGE_DEFAULT_SOURCE="Mangadex"
  kage inline -q "Death Note" -d`,
	Run: func(cmd *cobra.Command, args []string) {
		setOnly := lo.Must(cmd.Flags().GetBool("set-only"))
		unsetOnly := lo.Must(cmd.Flags().GetBool("unset-only"))

		config.EnvExposed = append(config.EnvExposed, where.EnvConfigPath)
		slices.Sort(config.EnvExposed)
		for _, env := range config.EnvExposed {
			if env != where.EnvConfigPath {
				env = strings.ToUpper(constant.Kage + "_" + config.EnvKeyReplacer.Replace(env))
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
