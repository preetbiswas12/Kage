package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/invopop/jsonschema"
	"github.com/preetbiswas12/Kage/anilist"
	"github.com/preetbiswas12/Kage/converter"
	"github.com/preetbiswas12/Kage/filesystem"
	"github.com/preetbiswas12/Kage/inline"
	"github.com/preetbiswas12/Kage/key"
	"github.com/preetbiswas12/Kage/provider"
	"github.com/preetbiswas12/Kage/query"
	"github.com/preetbiswas12/Kage/source"
	"github.com/preetbiswas12/Kage/update"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func init() {
	rootCmd.AddCommand(inlineCmd)

	inlineCmd.Flags().StringP("query", "q", "", "query to search for")
	inlineCmd.Flags().StringP("manga", "m", "", "manga selector")
	inlineCmd.Flags().StringP("chapters", "c", "", "chapter selector")
	inlineCmd.Flags().BoolP("download", "d", false, "download chapters")
	inlineCmd.Flags().BoolP("json", "j", false, "JSON output")
	inlineCmd.Flags().BoolP("populate-pages", "p", false, "Populate chapters pages")
	inlineCmd.Flags().BoolP("fetch-metadata", "f", false, "Populate manga metadata")
	inlineCmd.Flags().BoolP("include-anilist-manga", "a", false, "Include anilist manga in the output")
	lo.Must0(viper.BindPFlag(key.MetadataFetchAnilist, inlineCmd.Flags().Lookup("fetch-metadata")))

	inlineCmd.Flags().StringP("output", "o", "", "output file")

	lo.Must0(inlineCmd.MarkFlagRequired("query"))
	inlineCmd.MarkFlagsMutuallyExclusive("download", "json")
	inlineCmd.MarkFlagsMutuallyExclusive("include-anilist-manga", "download")

	inlineCmd.RegisterFlagCompletionFunc("query", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return query.SuggestMany(toComplete), cobra.ShellCompDirectiveNoFileComp
	})
}

var inlineCmd = &cobra.Command{
	Use:   "inline",
	Short: "Run in inline mode for scripting",
	Long: `Launch in inline mode for scripting and automation.

This mode allows you to search, select, and download manga programmatically
using command-line arguments, making it perfect for scripts and integrations.

Manga selectors:
  first - first manga in the list
  last - last manga in the list
  [number] - select manga by index (starting from 0)

Chapter selectors:
  first - first chapter in the list
  last - last chapter in the list
  all - all chapters in the list
  [number] - select chapter by index (starting from 0)
  [from]-[to] - select chapters by range
  @[substring]@ - select chapters by name substring

When using the json flag, manga selector can be omitted to select all mangas.`,

	Example: `  # Search and download first manga's first chapter
  mangal inline -q "one piece" -m first -c first -d

  # Get JSON output for all search results
  mangal inline -q "naruto" -j

  # Download chapters 1-10 from first manga
  mangal inline -q "bleach" -m first -c 1-10 -d

More examples: https://github.com/preetbiswas12/Kage/wiki/Inline-mode`,
	PreRun: func(cmd *cobra.Command, args []string) {
		json, _ := cmd.Flags().GetBool("json")

		if !json {
			lo.Must0(cmd.MarkFlagRequired("manga"))
		}

		if lo.Must(cmd.Flags().GetBool("populate-pages")) {
			lo.Must0(cmd.MarkFlagRequired("json"))
		}

		if _, err := converter.Get(viper.GetString(key.FormatsUse)); err != nil {
			handleErr(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			sources []source.Source
			err     error
		)

		for _, name := range viper.GetStringSlice(key.DownloaderDefaultSources) {
			if name == "" {
				handleErr(errors.New("source not set"))
			}

			p, ok := provider.Get(name)
			if !ok {
				handleErr(fmt.Errorf("source not found: %s", name))
			}

			src, err := p.CreateSource()
			handleErr(err)

			sources = append(sources, src)
		}

		query := lo.Must(cmd.Flags().GetString("query"))

		output := lo.Must(cmd.Flags().GetString("output"))
		var writer io.Writer
		if output != "" {
			writer, err = filesystem.Api().Create(output)
			handleErr(err)
		} else {
			writer = os.Stdout
		}

		mangaFlag := lo.Must(cmd.Flags().GetString("manga"))
		mangaPicker := mo.None[inline.MangaPicker]()
		if mangaFlag != "" {
			fn, err := inline.ParseMangaPicker(query, lo.Must(cmd.Flags().GetString("manga")))
			handleErr(err)
			mangaPicker = mo.Some(fn)
		}

		chapterFlag := lo.Must(cmd.Flags().GetString("chapters"))
		chapterFilter := mo.None[inline.ChaptersFilter]()
		if chapterFlag != "" {
			fn, err := inline.ParseChaptersFilter(chapterFlag)
			handleErr(err)
			chapterFilter = mo.Some(fn)
		}

		options := &inline.Options{
			Sources:             sources,
			Download:            lo.Must(cmd.Flags().GetBool("download")),
			Json:                lo.Must(cmd.Flags().GetBool("json")),
			Query:               query,
			PopulatePages:       lo.Must(cmd.Flags().GetBool("populate-pages")),
			IncludeAnilistManga: lo.Must(cmd.Flags().GetBool("include-anilist-manga")),
			MangaPicker:         mangaPicker,
			ChaptersFilter:      chapterFilter,
			Out:                 writer,
		}

		handleErr(inline.Run(options))
	},
}

func init() {
	inlineCmd.AddCommand(inlineAnilistCmd)
}

var inlineAnilistCmd = &cobra.Command{
	Use:   "anilist",
	Short: "Manage Anilist manga integration",
	Long:  `Search, bind, and manage manga metadata from Anilist for better organization and tracking.`,
}

func init() {
	inlineAnilistCmd.AddCommand(inlineAnilistSearchCmd)

	inlineAnilistSearchCmd.Flags().StringP("name", "n", "", "manga name to search")
	inlineAnilistSearchCmd.Flags().IntP("id", "i", 0, "anilist manga id")

	inlineAnilistSearchCmd.MarkFlagsMutuallyExclusive("name", "id")
}

var inlineAnilistSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for manga on Anilist",
	Long:  `Search Anilist for manga by name or retrieve specific manga by ID. Returns results in JSON format.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().Changed("name") && !cmd.Flags().Changed("id") {
			handleErr(errors.New("name or id flag is required"))
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		mangaName := lo.Must(cmd.Flags().GetString("name"))
		mangaId := lo.Must(cmd.Flags().GetInt("id"))

		var toEncode any

		if mangaName != "" {
			mangas, err := anilist.SearchByName(mangaName)
			handleErr(err)
			toEncode = mangas
		} else {
			manga, err := anilist.GetByID(mangaId)
			handleErr(err)
			toEncode = manga
		}

		handleErr(json.NewEncoder(os.Stdout).Encode(toEncode))
	},
}

func init() {
	inlineAnilistCmd.AddCommand(inlineAnilistGetCmd)

	inlineAnilistGetCmd.Flags().StringP("name", "n", "", "manga name to get the bind for")
	lo.Must0(inlineAnilistGetCmd.MarkFlagRequired("name"))
}

var inlineAnilistGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Anilist binding for a manga",
	Long:  `Retrieve the Anilist manga that is bound to a local manga name. Returns the binding in JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			m   *anilist.Manga
			err error
		)

		name := lo.Must(cmd.Flags().GetString("name"))
		m, err = anilist.FindClosest(name)

		if err != nil {
			m, err = anilist.FindClosest(name)
			handleErr(err)
		}

		handleErr(json.NewEncoder(os.Stdout).Encode(m))
	},
}

func init() {
	inlineAnilistCmd.AddCommand(inlineAnilistBindCmd)

	inlineAnilistBindCmd.Flags().StringP("name", "n", "", "manga name")
	inlineAnilistBindCmd.Flags().IntP("id", "i", 0, "anilist manga id")

	lo.Must0(inlineAnilistBindCmd.MarkFlagRequired("name"))
	lo.Must0(inlineAnilistBindCmd.MarkFlagRequired("id"))

	inlineAnilistBindCmd.MarkFlagsRequiredTogether("name", "id")
}

var inlineAnilistBindCmd = &cobra.Command{
	Use:   "set",
	Short: "Bind a manga to an Anilist entry",
	Long:  `Create a binding between a local manga name and an Anilist manga ID for metadata enrichment and tracking.`,
	Run: func(cmd *cobra.Command, args []string) {
		anilistManga, err := anilist.GetByID(lo.Must(cmd.Flags().GetInt("id")))
		handleErr(err)

		mangaName := lo.Must(cmd.Flags().GetString("name"))

		handleErr(anilist.SetRelation(mangaName, anilistManga))
	},
}

func init() {
	inlineAnilistCmd.AddCommand(inlineAnilistUpdateCmd)

	inlineAnilistUpdateCmd.Flags().StringP("path", "p", "", "path to the manga")
	lo.Must0(inlineAnilistUpdateCmd.MarkFlagRequired("path"))
}

var inlineAnilistUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update manga metadata from Anilist",
	Long:  `Refresh manga metadata according to its current Anilist binding. Updates cover art, description, and other information.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := lo.Must(cmd.Flags().GetString("path"))
		handleErr(update.Metadata(path))
	},
}

func init() {
	inlineCmd.AddCommand(inlineSchemaCmd)

	inlineSchemaCmd.Flags().BoolP("anilist", "a", false, "generate anilist search output schema")
}

var inlineSchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Generate JSON schemas for inline mode output",
	Long:  `Generate JSON schemas that describe the structure of inline mode JSON outputs. Useful for validation and IDE support.`,
	Run: func(cmd *cobra.Command, args []string) {
		reflector := new(jsonschema.Reflector)
		reflector.Anonymous = true
		reflector.Namer = func(t reflect.Type) string {
			name := t.Name()
			switch strings.ToLower(name) {
			case "manga", "chapter", "page", "date", "output":
				return filepath.Base(t.PkgPath()) + "." + name
			}

			return name
		}

		var schema *jsonschema.Schema

		switch {
		case lo.Must(cmd.Flags().GetBool("anilist")):
			schema = reflector.Reflect([]*anilist.Manga{})
		default:
			schema = reflector.Reflect(&inline.Output{})
		}

		handleErr(json.NewEncoder(os.Stdout).Encode(schema))
	},
}
