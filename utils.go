package main

func setAsFirstString(lst []string, elem string) []string {
	lst = removeStringFromArray(lst, elem)
	return append([]string{elem}, lst...)
}

func removeStringFromArray(lst []string, elem string) []string {
	for {
		if len(lst) == 0 {
			return lst
		}

		for index, value := range lst {
			if value == elem {
				lst = append(lst[:index], lst[index+1:]...)
				break
			}

			if len(lst) == index+1 {
				return lst
			}
		}
	}
}
