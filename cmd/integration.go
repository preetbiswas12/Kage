package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/preetbiswas12/Kage/icon"
	"github.com/preetbiswas12/Kage/integration/anilist"
	"github.com/preetbiswas12/Kage/key"
	"github.com/preetbiswas12/Kage/log"
	"github.com/preetbiswas12/Kage/open"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(integrationCmd)
	integrationCmd.AddCommand(integrationAnilistCmd)
	integrationAnilistCmd.Flags().BoolP("disable", "d", false, "Disable Anilist integration")
}

var integrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "Manage third-party service integrations",
	Long: `Configure and manage integrations with external services like Anilist.

Integrations allow Kage to sync your reading progress, fetch metadata,
and provide enhanced features through connection with online services.

Currently supported integrations:
  • Anilist - Track manga and sync reading progress

SUBCOMMANDS:
  integration anilist    Configure Anilist integration`,
}

var integrationAnilistCmd = &cobra.Command{
	Use:   "anilist",
	Short: "Configure Anilist manga tracking integration",
	Long: `Set up integration with Anilist for manga tracking and metadata synchronization.

Anilist integration enables:
  • Automatic progress tracking
  • Metadata fetching (title, cover, genres, etc.)
  • Synchronization with Anilist account
  • Enhanced manga information

FLAGS:
  -d, --disable    Disable Anilist integration

EXAMPLES:
  # Set up Anilist integration (interactive setup)
  kage integration anilist
  
  # Disable Anilist integration
  kage integration anilist --disable

For detailed setup instructions, visit:
https://github.com/preetbiswas12/Kage/wiki/Anilist-Integration`,
	Run: func(cmd *cobra.Command, args []string) {
		if lo.Must(cmd.Flags().GetBool("disable")) {
			viper.Set(key.AnilistEnable, false)
			viper.Set(key.AnilistCode, "")
			viper.Set(key.AnilistSecret, "")
			viper.Set(key.AnilistID, "")
			log.Info("Anilist integration disabled")
			handleErr(viper.WriteConfig())
		}

		if !viper.GetBool(key.AnilistEnable) {
			confirm := survey.Confirm{
				Message: "Anilist is disabled. Enable?",
				Default: false,
			}
			var response bool
			err := survey.AskOne(&confirm, &response)
			handleErr(err)

			if !response {
				return
			}

			viper.Set(key.AnilistEnable, response)
			err = viper.WriteConfig()
			if err != nil {
				switch err.(type) {
				case viper.ConfigFileNotFoundError:
					err = viper.SafeWriteConfig()
					handleErr(err)
				default:
					handleErr(err)
					log.Error(err)
				}
			}
		}

		if viper.GetString(key.AnilistID) == "" {
			input := survey.Input{
				Message: "Anilsit client ID is not set. Please enter it:",
				Help:    "",
			}
			var response string
			err := survey.AskOne(&input, &response)
			handleErr(err)

			if response == "" {
				return
			}

			viper.Set(key.AnilistID, response)
			err = viper.WriteConfig()
			handleErr(err)
		}

		if viper.GetString(key.AnilistSecret) == "" {
			input := survey.Input{
				Message: "Anilsit client secret is not set. Please enter it:",
				Help:    "",
			}
			var response string
			err := survey.AskOne(&input, &response)
			handleErr(err)

			if response == "" {
				return
			}

			viper.Set(key.AnilistSecret, response)
			err = viper.WriteConfig()
			handleErr(err)
		}

		if viper.GetString(key.AnilistCode) == "" {
			authURL := anilist.New().AuthURL()
			confirmOpenInBrowser := survey.Confirm{
				Message: "Open browser to authenticate with Anilist?",
				Default: false,
			}

			var openInBrowser bool
			err := survey.AskOne(&confirmOpenInBrowser, &openInBrowser)
			if err == nil && openInBrowser {
				err = open.Start(authURL)
			}

			if err != nil || !openInBrowser {
				fmt.Println("Please open the following URL in your browser:")
				fmt.Println(authURL)
			}

			input := survey.Input{
				Message: "Anilsit code is not set. Please copy it from the link and paste in here:",
				Help:    "",
			}

			var response string
			err = survey.AskOne(&input, &response)
			handleErr(err)

			if response == "" {
				return
			}

			viper.Set(key.AnilistCode, response)
			err = viper.WriteConfig()
			handleErr(err)
		}

		fmt.Printf("%s Anilist integration was set up\n", icon.Get(icon.Success))
	},
}
