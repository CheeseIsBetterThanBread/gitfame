package stats

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func PrintStats(stats []authorStats, format string) {
	log.Printf("Displaying statistics\n")
	defer func() {
		log.Printf("-- Displaying statistics - done\n\n")
	}()

	var err error
	switch format {
	case "csv":
		w := csv.NewWriter(os.Stdout)
		defer w.Flush()

		if err = w.Write([]string{"Name", "Lines", "Commits", "Files"}); err != nil {
			log.Fatalf("\tError writing record to csv: %v\n", err)
		}

		for _, stat := range stats {
			if err = w.Write([]string{
				stat.Name,
				fmt.Sprintf("%d", stat.Lines),
				fmt.Sprintf("%d", stat.Commits),
				fmt.Sprintf("%d", stat.Files),
			}); err != nil {
				log.Fatalf("\tError writing record to csv: %v\n", err)
			}
		}

	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err = enc.Encode(stats); err != nil {
			log.Fatalf("\tError writing record to json: %v\n", err)
		}

	case "json-lines":
		enc := json.NewEncoder(os.Stdout)
		for _, stat := range stats {
			if err = enc.Encode(stat); err != nil {
				log.Fatalf("\tError writing record to json-lines: %v\n", err)
			}
		}
	case "tabular":
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		if _, err = fmt.Fprintln(w, "Name\tLines\tCommits\tFiles"); err != nil {
			log.Fatalf("\tError writing record to tabwriter: %v\n", err)
		}

		for _, stat := range stats {
			if _, err = fmt.Fprintf(w, "%s\t%d\t%d\t%d\n",
				stat.Name,
				stat.Lines,
				stat.Commits,
				stat.Files,
			); err != nil {
				log.Fatalf("\tError writing record to tabwriter: %v\n", err)
			}
		}

		if err = w.Flush(); err != nil {
			log.Fatalf("\tError writing record to tabwriter: %v\n", err)
		}
	default:
		log.Fatalf("\tUnknown format %s\n", format)
	}
}
