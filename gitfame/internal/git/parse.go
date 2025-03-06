package git

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

type blameLine struct {
	CommitHash string
	Author     string
	Committer  string
}

func parseBlameOutput(output []byte) []blameLine {
	lines := bytes.Split(output, []byte{'\n'})
	var blameLines []blameLine
	var current blameLine

	for i := 0; i < len(lines); i++ {
		line := string(lines[i])
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "author "):
			current.Author = strings.TrimPrefix(line, "author ")
		case strings.HasPrefix(line, "committer "):
			current.Committer = strings.TrimPrefix(line, "committer ")
		case strings.HasPrefix(line, "filename "):
			blameLines = append(blameLines, current)
			current = blameLine{}
		default:
			if strings.Contains(line, " ") {
				parts := strings.Split(line, " ")
				if len(parts) >= 3 {
					_, err := strconv.Atoi(parts[1])
					if err != nil {
						continue
					}
					_, err = strconv.Atoi(parts[2])
					if err != nil {
						continue
					}

					current.CommitHash = parts[0]
				}
			}
		}
	}

	return blameLines
}

func GetLineStats(repoPath, file, revision string) []blameLine {
	log.Printf("Collecting blame statistics from file '%s'\n", file)
	defer func() {
		log.Printf("-- Collecting blame statistics from file '%s' - done\n\n", file)
	}()

	output := getGitBlame(repoPath, file, revision)
	return parseBlameOutput(output)
}
