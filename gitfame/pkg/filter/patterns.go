package filter

import (
	"log"
	"path/filepath"
)

func MatchesAnyPattern(file string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, err := filepath.Match(pattern, file)
		if err != nil {
			log.Fatalf("\tIn matching pattern '%s' on file '%s': %v", pattern, file, err)
		}
		if matched {
			return true
		}
	}
	return false
}

func ExcludeFiles(files []string, excludePatterns []string) []string {
	log.Print("Filtering patterns\n")
	defer func() {
		log.Printf("-- Filtering patterns - done\n\n")
	}()

	if len(excludePatterns) == 0 {
		log.Print("No patterns excluded from filter\n")
		return files
	}

	filtered := make([]string, 0)

	for _, file := range files {
		exclude := MatchesAnyPattern(file, excludePatterns)
		if !exclude {
			filtered = append(filtered, file)
		}
	}

	return filtered
}

func RestrictFiles(files []string, restrictPatterns []string) []string {
	log.Print("Restricting paths\n")
	defer func() {
		log.Printf("-- Restricting paths - done\n\n")
	}()

	if len(restrictPatterns) == 0 {
		log.Print("No patterns restricted to files\n")
		return files
	}

	filtered := make([]string, 0)

	for _, file := range files {
		include := MatchesAnyPattern(file, restrictPatterns)
		if include {
			filtered = append(filtered, file)
		}
	}

	return filtered
}
