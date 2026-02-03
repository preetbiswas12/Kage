package cmd

import (
	"github.com/metafates/mangal/provider/custom"
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
	Long: `Run Lua 5.1 scripts for testing custom scrapers or using mangal as a Lua interpreter.

This command is useful for debugging custom manga sources before installation
or for running standalone Lua scripts with mangal's built-in libraries.`,
	Args:    cobra.ExactArgs(1),
	Example: `  # Run a custom scraper for testing
  mangal run ./my-scraper.lua

  # Run with lenient mode (skip function warnings)
  mangal run -l ./test.lua`,
	Run: func(cmd *cobra.Command, args []string) {
		sourcePath := args[0]

		// LoadSource runs file when it's loaded
		_, err := custom.LoadSource(sourcePath, !lo.Must(cmd.Flags().GetBool("lenient")))
		handleErr(err)
	},
}
