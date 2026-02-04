package cmd

import (
	"github.com/preetbiswas12/Kage/provider/custom"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("lenient", "l", false, "do not warn about missing functions")
}

var runCmd = &cobra.Command{
	Use:   "run [file]",
	Short: "Execute a Lua script file",
	Long: `Run Lua 5.1 scripts for testing custom scrapers or using Kage as a Lua interpreter.

This command is useful for:
  • Testing custom manga scrapers before installation
  • Debugging Lua scripts
  • Running standalone scripts with Kage's libraries
  • Developing new manga sources

FLAGS:
  -l, --lenient    Do not warn about missing functions

EXAMPLES:
  # Run a custom scraper for testing
  kage run ./my-scraper.lua
  
  # Run with lenient mode (skip function warnings)
  kage run -l ./test.lua
  
  # Test a new source before installation
  kage run ./sources/mysource.lua`,
	Run: func(cmd *cobra.Command, args []string) {
		sourcePath := args[0]

		// LoadSource runs file when it's loaded
		_, err := custom.LoadSource(sourcePath, !lo.Must(cmd.Flags().GetBool("lenient")))
		handleErr(err)
	},
}
