package git

import (
	"log"
	"os/exec"
	"strings"
)

func GetGitFiles(repoPath, revision string) []string {
	files := make([]string, 0)
	log.Print("Collecting files\n")
	defer func() {
		if len(files) > 0 {
			log.Printf("Found files:\n")
			for _, file := range files {
				log.Printf("'%s'\n", file)
			}
		} else {
			log.Printf("No files were found\n")
		}

		log.Printf("-- Collecting files - done\n\n")
	}()

	cmd := exec.Command("git", "-C", repoPath, "ls-tree", "-r", revision, "--name-only")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("\tError getting git files from '%s': %v\n", repoPath, err)
	}

	result := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, file := range result {
		if file != "" {
			files = append(files, file)
		}
	}

	return files
}

func getGitBlame(repoPath, file, revision string) []byte {
	cmd := exec.Command("git", "-C", repoPath, "blame", "--line-porcelain", revision, "--", file)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf(
			"\tError getting git blame from file '%s' in repo '%s' on commit '%s': %v\n",
			file,
			repoPath,
			revision,
			err,
		)
	}
	return out
}
