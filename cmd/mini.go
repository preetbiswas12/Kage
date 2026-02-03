package cmd

import (
	"github.com/preetbiswas12/Kage/converter"
	"github.com/preetbiswas12/Kage/key"
	"github.com/preetbiswas12/Kage/mini"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(miniCmd)

	miniCmd.Flags().BoolP("download", "d", false, "download mode")
	miniCmd.Flags().BoolP("continue", "c", false, "continue reading")

	miniCmd.MarkFlagsMutuallyExclusive("download", "continue")
}

var miniCmd = &cobra.Command{
	Use:   "mini",
	Short: "Launch in mini mode",
	Long: `Launch mangal in mini mode with a simplified, interactive interface.

Mini mode provides a streamlined experience similar to ani-cli,
with quick keyboard-based navigation and minimal visual clutter.
Perfect for terminal enthusiasts who prefer speed and efficiency.`,
	Example: `  # Launch mini mode
  mangal mini

  # Launch in download mode
  mangal mini -d

  # Continue reading from history
  mangal mini -c`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := converter.Get(viper.GetString(key.FormatsUse)); err != nil {
			handleErr(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		options := mini.Options{
			Download: lo.Must(cmd.Flags().GetBool("download")),
			Continue: lo.Must(cmd.Flags().GetBool("continue")),
		}
		err := mini.Run(&options)

		if err != nil && err.Error() != "interrupt" {
			handleErr(err)
		}
	},
}
