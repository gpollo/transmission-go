package main

func setAsFirstString(lst []string, elem string) []string {
	removeFromStringArray(lst, elem)
	return append([]string{elem}, lst...)
}

func removeFromStringArray(lst []string, elem string) []string {
	for index, value := range lst {
		if value == elem {
			return append(lst[:index], lst[index+1:]...)
		}
	}
	return lst
}
