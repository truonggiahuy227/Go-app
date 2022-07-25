package util

import "fmt"

func RemoveElement(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func FindIndex(slice []string, item string) int {
	fmt.Println(item)
	for i, _ := range slice {
		fmt.Println(slice[i])
		if slice[i] == item {
			return i
		}
	}
	return -1
}

func SliceContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
