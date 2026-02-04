package cmd

import (
	"fmt"

	"github.com/preetbiswas12/Kage/filesystem"
	"github.com/preetbiswas12/Kage/icon"
	"github.com/preetbiswas12/Kage/util"
	"github.com/preetbiswas12/Kage/where"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/spf13/cobra"
)

type clearTarget struct {
	name     string
	argLong  string
	argShort mo.Option[string]
	location func() string
}

var clearTargets = []clearTarget{
	{"cache directory", "cache", mo.Some("c"), where.Cache},
	{"history file", "history", mo.Some("s"), where.History},
	{"anilist binds", "anilist", mo.Some("a"), where.AnilistBinds},
	{"queries history", "queries", mo.Some("q"), where.Queries},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	for _, target := range clearTargets {
		help := fmt.Sprintf("clear %s", target.name)
		if target.argShort.IsPresent() {
			clearCmd.Flags().BoolP(target.argLong, target.argShort.MustGet(), false, help)
		} else {
			clearCmd.Flags().Bool(target.argLong, false, help)
		}
	}
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear cached files and data",
	Long: `Remove cached data, history, and other temporary files to free up space or reset state.

This command helps manage your local Kage data by removing:
  • Cache files (downloaded manga, images)
  • Reading history (chapter progress)
  • Anilist bindings (manga-to-Anilist mappings)
  • Queries history (search history)

FLAGS:
  -c, --cache     Clear downloaded manga cache
  -s, --history   Clear reading history
  -a, --anilist   Clear Anilist bindings
  -q, --queries   Clear search query history

EXAMPLES:
  # Clear all caches (use without flags)
  kage clear --cache --history --anilist --queries
  
  # Clear only cache
  kage clear --cache
  
  # Clear reading history
  kage clear --history`,
	Run: func(cmd *cobra.Command, args []string) {
		var anyCleared bool

		doClear := func(what string) bool {
			return lo.Must(cmd.Flags().GetBool(what))
		}

		for _, target := range clearTargets {
			if doClear(target.argLong) {
				anyCleared = true
				e := util.PrintErasable(fmt.Sprintf("%s Clearing %s...", icon.Get(icon.Progress), util.Capitalize(target.name)))
				_ = util.Delete(target.location())
				e()
				fmt.Printf("%s %s cleared\n", icon.Get(icon.Success), util.Capitalize(target.name))
				handleErr(filesystem.Api().RemoveAll(target.location()))
			}
		}

		if !anyCleared {
			handleErr(cmd.Help())
		}
	},
}
