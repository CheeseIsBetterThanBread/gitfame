package stats

import (
	"gitlab.com/slon/shad-go/gitfame/internal/CLI"
	"gitlab.com/slon/shad-go/gitfame/internal/git"
	"gitlab.com/slon/shad-go/gitfame/pkg/filter"
	"log"
)

type authorStats struct {
	Name    string `json:"name"`
	Lines   int    `json:"lines"`
	Commits int    `json:"commits"`
	Files   int    `json:"files"`
}

type authorStatsHelper struct {
	Name    string
	Lines   int
	Commits map[string]bool
	Files   int
}

func CalculateStats(config *CLI.Config) []authorStats {
	log.Printf("Calculating statistics\n")
	defer func() {
		log.Printf("-- Calculating statistics - done\n\n")
	}()

	files := git.GetGitFiles(config.Repository, config.Revision)
	filteredFiles := filter.AtOnce(
		files,
		config.Extensions,
		config.Languages,
		config.Exclude,
		config.RestrictTo,
	)

	statsMap := make(map[string]*authorStatsHelper)
	for _, file := range filteredFiles {
		fileStatsPtr := getFileStats(config.Repository, file, config.Revision, config.UseCommitter)

		for author, lines := range fileStatsPtr.Lines {
			if _, ok := statsMap[author]; !ok {
				statsMap[author] = &authorStatsHelper{Name: author, Commits: make(map[string]bool)}
			}
			statsMap[author].Lines += lines
			statsMap[author].Files++

			for commitHash := range fileStatsPtr.Commits[author] {
				statsMap[author].Commits[commitHash] = true
			}
		}
	}

	stats := make([]authorStats, 0, len(statsMap))
	for _, stat := range statsMap {
		stats = append(stats, authorStats{
			Name:    stat.Name,
			Lines:   stat.Lines,
			Commits: len(stat.Commits),
			Files:   stat.Files,
		})
	}

	sortStats(stats, config.OrderBy)
	return stats
}
