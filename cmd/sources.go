package cmd

import (
	"fmt"
	"github.com/preetbiswas12/Kage/color"
	"github.com/preetbiswas12/Kage/constant"
	"github.com/preetbiswas12/Kage/key"
	"github.com/preetbiswas12/Kage/tui"
	"github.com/preetbiswas12/Kage/util"
	"github.com/spf13/viper"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/preetbiswas12/Kage/filesystem"
	"github.com/preetbiswas12/Kage/icon"
	"github.com/preetbiswas12/Kage/provider"
	"github.com/preetbiswas12/Kage/style"
	"github.com/preetbiswas12/Kage/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sourcesCmd)
}

var sourcesCmd = &cobra.Command{
	Use:   "sources",
	Short: "Manage manga sources",
	Long:  `Install, remove, generate and list manga sources (scrapers) for downloading manga from various websites.`,
}

func init() {
	sourcesCmd.AddCommand(sourcesListCmd)

	sourcesListCmd.Flags().BoolP("raw", "r", false, "do not print headers")
	sourcesListCmd.Flags().BoolP("custom", "c", false, "show only custom sources")
	sourcesListCmd.Flags().BoolP("builtin", "b", false, "show only builtin sources")

	sourcesListCmd.MarkFlagsMutuallyExclusive("custom", "builtin")
	sourcesListCmd.SetOut(os.Stdout)
}

var sourcesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available manga sources",
	Long:  `Display all built-in and custom manga sources that can be used to search and download manga.`,
	Run: func(cmd *cobra.Command, args []string) {
		printHeader := !lo.Must(cmd.Flags().GetBool("raw"))
		headerStyle := style.New().Foreground(color.HiBlue).Bold(true).Render
		h := func(s string) {
			if printHeader {
				cmd.Println(headerStyle(s))
			}
		}

		printBuiltin := func() {
			h("Builtin:")
			for _, p := range provider.Builtins() {
				cmd.Println(p.Name)
			}
		}

		printCustom := func() {
			h("Custom:")
			for _, p := range provider.Customs() {
				cmd.Println(p.Name)
			}
		}

		switch {
		case lo.Must(cmd.Flags().GetBool("builtin")):
			printBuiltin()
		case lo.Must(cmd.Flags().GetBool("custom")):
			printCustom()
		default:
			printBuiltin()
			if printHeader {
				cmd.Println()
			}
			printCustom()
		}
	},
}

func init() {
	sourcesCmd.AddCommand(sourcesRemoveCmd)

	sourcesRemoveCmd.Flags().StringArrayP("name", "n", []string{}, "name of the source to remove")
	lo.Must0(sourcesRemoveCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		sources, err := filesystem.Api().ReadDir(where.Sources())
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		return lo.FilterMap(sources, func(item os.FileInfo, _ int) (string, bool) {
			name := item.Name()
			if !strings.HasSuffix(name, provider.CustomProviderExtension) {
				return "", false
			}

			return util.FileStem(filepath.Base(name)), true
		}), cobra.ShellCompDirectiveNoFileComp
	}))
}

var sourcesRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove custom manga sources",
	Long:  `Remove one or more custom manga sources (scrapers) from your installation.`,
	Example: `  # Remove a single source
  mangal sources remove -n mangakakalot

  # Remove multiple sources
  mangal sources remove -n source1 -n source2`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range lo.Must(cmd.Flags().GetStringArray("name")) {
			path := filepath.Join(where.Sources(), name+provider.CustomProviderExtension)
			handleErr(filesystem.Api().Remove(path))
			fmt.Printf("%s successfully removed %s\n", icon.Get(icon.Success), style.Fg(color.Yellow)(name))
		}
	},
}

func init() {
	sourcesCmd.AddCommand(sourcesInstallCmd)
}

var sourcesInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Browse and install custom scrapers",
	Long: `Browse and install custom scrapers from official GitHub repo.
https://github.com/preetbiswas12/Kage-scrapers`,
	Run: func(cmd *cobra.Command, args []string) {
		handleErr(tui.Run(&tui.Options{Install: true}))
	},
}

func init() {
	sourcesCmd.AddCommand(sourcesGenCmd)

	sourcesGenCmd.Flags().StringP("name", "n", "", "name of the source")
	sourcesGenCmd.Flags().StringP("url", "u", "", "url of the website")

	lo.Must0(sourcesGenCmd.MarkFlagRequired("name"))
	lo.Must0(sourcesGenCmd.MarkFlagRequired("url"))
}

var sourcesGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new Lua scraper template",
	Long: `Generate a template for creating a custom manga scraper in Lua.

This command creates a new Lua file with boilerplate code and comments
to help you build your own manga source scraper. The generated template
includes function stubs for searching, listing chapters, and fetching pages.`,
	Example: `  # Generate a scraper for a specific site
  mangal sources gen -n "My Site" -u https://example.com

  # The generated file will be saved to your sources directory
  mangal sources gen -n mangasite -u https://mangasite.com`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetOut(os.Stdout)

		author := viper.GetString(key.GenAuthor)
		if author == "" {
			usr, err := user.Current()
			if err == nil {
				author = usr.Username
			} else {
				author = "Anonymous"
			}
		}

		s := struct {
			Name            string
			URL             string
			SearchMangaFn   string
			MangaChaptersFn string
			ChapterPagesFn  string
			Author          string
		}{
			Name:            lo.Must(cmd.Flags().GetString("name")),
			URL:             lo.Must(cmd.Flags().GetString("url")),
			SearchMangaFn:   constant.SearchMangaFn,
			MangaChaptersFn: constant.MangaChaptersFn,
			ChapterPagesFn:  constant.ChapterPagesFn,
			Author:          author,
		}

		funcMap := template.FuncMap{
			"repeat": strings.Repeat,
			"plus":   func(a, b int) int { return a + b },
			"max":    util.Max[int],
		}

		tmpl, err := template.New("source").Funcs(funcMap).Parse(constant.SourceTemplate)
		handleErr(err)

		target := filepath.Join(where.Sources(), util.SanitizeFilename(s.Name)+".lua")
		f, err := filesystem.Api().Create(target)
		handleErr(err)

		defer util.Ignore(f.Close)

		err = tmpl.Execute(f, s)
		handleErr(err)

		cmd.Println(target)
	},
}
