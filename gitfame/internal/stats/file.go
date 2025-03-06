package stats

import (
	"gitlab.com/slon/shad-go/gitfame/internal/git"
)

type fileStats struct {
	Lines   map[string]int
	Commits map[string]map[string]bool
}

func getFileStats(repoPath, file, revision string, useCommitter bool) *fileStats {
	stats := &fileStats{
		Lines:   make(map[string]int),
		Commits: make(map[string]map[string]bool),
	}

	if git.IsFileEmpty(repoPath, file, revision) {
		commitHash := git.GetLastCommitForFile(repoPath, file, revision)
		author := git.GetAuthor(repoPath, commitHash)

		stats.Lines[author] = 0
		stats.Commits[author] = make(map[string]bool)
		stats.Commits[author][commitHash] = true

		return stats
	}
	blameLines := git.GetLineStats(repoPath, file, revision)

	for _, line := range blameLines {
		author := line.Author
		if useCommitter {
			author = line.Committer
		}

		stats.Lines[author]++

		if _, ok := stats.Commits[author]; !ok {
			stats.Commits[author] = make(map[string]bool)
		}
		stats.Commits[author][line.CommitHash] = true
	}

	return stats
}
