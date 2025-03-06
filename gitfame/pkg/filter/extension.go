package filter

import (
	"log"
	"path/filepath"
	"strings"
)

func Extensions(files, extensions []string) []string {
	log.Print("Filtering extensions\n")
	defer func() {
		log.Printf("-- Filtering extensions - done\n\n")
	}()

	if len(extensions) == 0 {
		log.Print("No extensions found")
		return files
	}

	filterMap := map[string]bool{}
	for _, ext := range extensions {
		filterMap[ext] = true
	}

	filtered := make([]string, 0, len(files))
	for _, file := range files {
		ext := filepath.Ext(file)
		if filterMap[ext] {
			filtered = append(filtered, file)
		}
	}

	return filtered
}

func Languages(files, languages []string, languageExtensions map[string][]string) []string {
	log.Print("Filtering languages\n")
	defer func() {
		log.Printf("-- Filtering languages - done\n\n")
	}()

	if len(languages) == 0 {
		log.Print("No languages found")
		return files
	}

	filterMap := make(map[string]bool)
	for _, lang := range languages {
		language := strings.ToLower(lang)
		if extensions, ok := languageExtensions[language]; ok {
			for _, ext := range extensions {
				filterMap[ext] = true
			}
		} else {
			log.Printf("Warning: unknown language '%s'\n", lang)
		}
	}

	filtered := make([]string, 0, len(files))
	for _, file := range files {
		ext := filepath.Ext(file)
		if filterMap[ext] {
			filtered = append(filtered, file)
		}
	}

	return filtered
}
