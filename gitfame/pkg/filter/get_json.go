package filter

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type fileType struct {
	Name       string
	Type       string
	Extensions []string
}

func getMapping() map[string][]string {
	file, err := os.Open("../../configs/language_extensions.json")
	if err != nil {
		log.Fatalf("\tError opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("\tError reading file: %v", err)
	}

	var fileTypes []fileType
	err = json.Unmarshal(byteValue, &fileTypes)
	if err != nil {
		log.Fatalf("\tError unmarshalling JSON: %v", err)
	}

	extensionsMap := make(map[string][]string)
	for _, ft := range fileTypes {
		extensionsMap[strings.ToLower(ft.Name)] = ft.Extensions
	}

	return extensionsMap
}
