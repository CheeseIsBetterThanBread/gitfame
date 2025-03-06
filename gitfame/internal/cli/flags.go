package cli

import (
	"github.com/spf13/cobra"
	"log"
)

func SetupFlags(cmd *cobra.Command) {
	log.Print("Setting up flags\n")
	defer func() {
		log.Printf("-- Setting up flags - done\n\n")
	}()

	cmd.Flags().StringP("repository", "r", ".", "Path to the Git repository")
	cmd.Flags().StringP("revision", "R", "HEAD", "Commit revision")
	cmd.Flags().StringP("order-by", "o", "lines", "Sort by lines, commits, or files")
	cmd.Flags().BoolP("use-committer", "c", false, "Use committer instead of author")
	cmd.Flags().StringP("format", "f", "tabular", "Output format: tabular, csv, json, json-lines")
	cmd.Flags().StringSliceP("extensions", "e", []string{}, "File extensions to include")
	cmd.Flags().StringSliceP("languages", "l", []string{}, "Languages to include")
	cmd.Flags().StringSliceP("exclude", "x", []string{}, "Glob patterns to exclude")
	cmd.Flags().StringSliceP("restrict-to", "t", []string{}, "Glob patterns to restrict to")
}

type Config struct {
	Repository   string
	Revision     string
	OrderBy      string
	UseCommitter bool
	Format       string
	Extensions   []string
	Languages    []string
	Exclude      []string
	RestrictTo   []string
}

func ParseFlags(cmd *cobra.Command) *Config {
	log.Print("Parsing flags\n")
	defer func() {
		log.Printf("-- Parsing flags - done\n\n")
	}()

	var err error
	repo, err := cmd.Flags().GetString("repository")
	if err != nil {
		log.Fatalf("\tCouldn't parse repository flag: %v\n", err)
	}

	revision, err := cmd.Flags().GetString("revision")
	if err != nil {
		log.Fatalf("\tCouldn't parse revision flag: %v\n", err)
	}

	orderBy, err := cmd.Flags().GetString("order-by")
	if err != nil {
		log.Fatalf("\tCouldn't parse order-by flag: %v\n", err)
	}

	useCommitter, err := cmd.Flags().GetBool("use-committer")
	if err != nil {
		log.Fatalf("\tCouldn't parse use-commiter flag: %v\n", err)
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		log.Fatalf("\tCouldn't parse format flag: %v\n", err)
	}

	extensions, err := cmd.Flags().GetStringSlice("extensions")
	if err != nil {
		log.Fatalf("\tCouldn't parse extensions flag: %v\n", err)
	}

	languages, err := cmd.Flags().GetStringSlice("languages")
	if err != nil {
		log.Fatalf("\tCouldn't parse languages flag: %v\n", err)
	}

	exclude, err := cmd.Flags().GetStringSlice("exclude")
	if err != nil {
		log.Fatalf("\tCouldn't parse exclude flag: %v\n", err)
	}

	restrictTo, err := cmd.Flags().GetStringSlice("restrict-to")
	if err != nil {
		log.Fatalf("\tCouldn't parse restrict-to flag: %v\n", err)
	}

	return &Config{
		Repository:   repo,
		Revision:     revision,
		OrderBy:      orderBy,
		UseCommitter: useCommitter,
		Format:       format,
		Extensions:   extensions,
		Languages:    languages,
		Exclude:      exclude,
		RestrictTo:   restrictTo,
	}
}
