package stats

import (
	"log"
	"sort"
)

type triple struct {
	first  int
	second int
	third  int
}

func equal(lhs, rhs triple) bool {
	return lhs.first == rhs.first && lhs.second == rhs.second && lhs.third == rhs.third
}

func less(lhs, rhs triple) bool {
	if lhs.first != rhs.first {
		return lhs.first < rhs.first
	}

	if lhs.second != rhs.second {
		return lhs.second < rhs.second
	}

	return lhs.third < rhs.third
}

func more(lhs, rhs triple) bool {
	return less(rhs, lhs)
}

// (lines, commits, files) by default
func sortStats(stats []authorStats, orderBy string) {
	switch orderBy {
	case "commits":
		sort.Slice(stats, func(i, j int) bool {
			lhs := triple{stats[i].Commits, stats[i].Lines, stats[i].Files}
			rhs := triple{stats[j].Commits, stats[j].Lines, stats[j].Files}

			if equal(lhs, rhs) {
				return stats[i].Name < stats[j].Name
			}
			return more(lhs, rhs)
		})

	case "files":
		sort.Slice(stats, func(i, j int) bool {
			lhs := triple{stats[i].Files, stats[i].Lines, stats[i].Commits}
			rhs := triple{stats[j].Files, stats[j].Lines, stats[j].Commits}

			if equal(lhs, rhs) {
				return stats[i].Name < stats[j].Name
			}
			return more(lhs, rhs)
		})
	case "lines":
		sort.Slice(stats, func(i, j int) bool {
			lhs := triple{stats[i].Lines, stats[i].Commits, stats[i].Files}
			rhs := triple{stats[j].Lines, stats[j].Commits, stats[j].Files}

			if equal(lhs, rhs) {
				return stats[i].Name < stats[j].Name
			}
			return more(lhs, rhs)
		})
	default:
		log.Fatalf("\tInvalid sort order %s\n", orderBy)
	}
}
