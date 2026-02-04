package cmd

import (
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/preetbiswas12/Kage/color"
	"github.com/preetbiswas12/Kage/style"
	"github.com/preetbiswas12/Kage/version"
	"github.com/samber/lo"

	"github.com/preetbiswas12/Kage/constant"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.SetOut(os.Stdout)
	versionCmd.Flags().BoolP("short", "s", false, "print short version")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Long: `Display detailed version information including build details.

Shows:
  • Current version number
  • Git commit hash
  • Build date and time
  • Platform (OS/Architecture)
  • Available updates

FLAGS:
  -s, --short    Print only the version number (no additional info)

EXAMPLES:
  # Show full version information
  kage version
  
  # Show only version number
  kage version --short`,
	Run: func(cmd *cobra.Command, args []string) {
		if lo.Must(cmd.Flags().GetBool("short")) {
			cmd.Println(constant.Version)
			return
		}

		defer version.Notify()

		versionInfo := struct {
			Version  string
			OS       string
			Arch     string
			BuiltAt  string
			BuiltBy  string
			Revision string
			App      string
		}{
			Version:  constant.Version,
			App:      constant.Kage,
			OS:       runtime.GOOS,
			Arch:     runtime.GOARCH,
			BuiltAt:  strings.TrimSpace(constant.BuiltAt),
			BuiltBy:  constant.BuiltBy,
			Revision: constant.Revision,
		}

		t, err := template.New("version").Funcs(map[string]any{
			"faint":   style.Faint,
			"bold":    style.Bold,
			"magenta": style.Fg(color.Purple),
			"green":   style.Fg(color.Green),
			"repeat":  strings.Repeat,
			"concat": func(a, b string) string {
				return a + b
			},
		}).Parse(`{{ magenta "▇▇▇" }} {{ magenta .App }} 

  {{ faint "Version" }}         {{ bold .Version }}
  {{ faint "Git Commit" }}      {{ bold .Revision }} 
  {{ faint "Build Date" }}  	  {{ bold .BuiltAt }}
  {{ faint "Built By" }}        {{ bold .BuiltBy }}
  {{ faint "Platform" }}        {{ bold .OS }}/{{ bold .Arch }}
`)
		handleErr(err)
		handleErr(t.Execute(cmd.OutOrStdout(), versionInfo))
	},
}
