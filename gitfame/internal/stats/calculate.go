package stats

import (
	"gitlab.com/slon/shad-go/gitfame/internal/cli"
	"gitlab.com/slon/shad-go/gitfame/internal/git"
	"gitlab.com/slon/shad-go/gitfame/pkg/filter"
	"log"
	"sync"
)

type authorStats struct {
	Name    string `json:"name"`
	Lines   int    `json:"lines"`
	Commits int    `json:"commits"`
	Files   int    `json:"files"`
}

type authorStatsHelper struct {
	Name    string
	Lines   int32
	Commits *safeMapBool
	Files   int32
}

func CalculateStats(config *cli.Config) []authorStats {
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

	statsMap := newSafeMapHelper()
	var waitGroup = sync.WaitGroup{}
	for index, file := range filteredFiles {
		waitGroup.Add(1)

		go func(id int) {
			log.Printf("Goroutine %d is processing file %s\n", id, file)
			fileStatsPtr := getFileStats(config.Repository, file, config.Revision, config.UseCommitter)

			for author, lines := range fileStatsPtr.Lines {
				statsMap.Create(author)

				statsMap.IncreaseLines(author, lines)
				statsMap.IncreaseFiles(author)
				for commitHash := range fileStatsPtr.Commits[author] {
					statsMap.LogCommits(author, commitHash)
				}
			}

			log.Printf("-- Goroutine %d finished processing file %s\n", id, file)
			waitGroup.Done()
		}(index + 1)
	}
	waitGroup.Wait()

	stats := make([]authorStats, 0, len(statsMap.data))
	for _, stat := range statsMap.data {
		stats = append(stats, authorStats{
			Name:    stat.Name,
			Lines:   int(stat.Lines),
			Commits: stat.Commits.Len(),
			Files:   int(stat.Files),
		})
	}

	sortStats(stats, config.OrderBy)
	return stats
}
