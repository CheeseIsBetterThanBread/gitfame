package git

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func IsFileEmpty(repoPath, file, revision string) bool {
	cmd := exec.Command("git", "-C", repoPath, "show", fmt.Sprintf("%s:%s", revision, file))
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf(
			"\tCould not show file '%s' from repo '%s' on commit '%s': %v",
			file,
			repoPath,
			revision,
			err,
		)
	}
	return len(out) == 0
}

func GetLastCommitForFile(repoPath, file, revision string) string {
	cmd := exec.Command("git", "-C", repoPath, "log", "-1", "--pretty=format:%H", revision, "--", file)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf(
			"\tCould not get last commit for file '%s' from repo '%s': %v",
			file,
			repoPath,
			err,
		)
	}

	return strings.TrimSpace(string(out))
}

func GetAuthor(repoPath, commitHash string) string {
	cmd := exec.Command("git", "-C", repoPath, "show", "-s", "--pretty=format:%an", commitHash)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf(
			"\tCould not get author for commit '%s' from repo '%s': %v",
			commitHash,
			repoPath,
			err,
		)
	}

	return strings.TrimSpace(string(out))
}
