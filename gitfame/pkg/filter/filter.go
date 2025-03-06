package filter

func AtOnce(files, extensions, languages, exclude, restrict []string) []string {
	filterMap := getMapping()

	files = Extensions(files, extensions)
	files = Languages(files, languages, filterMap)

	files = ExcludeFiles(files, exclude)
	files = RestrictFiles(files, restrict)

	return files
}
