package main

import "fmt"

func isOppoosite(dir1 string, dir2 string) bool {
	if dir1 == "NORTH" && dir2 == "SOUTH" || (dir1 == "SOUTH" && dir2 == "NORTH") || (dir1 == "EAST" && dir2 == "WEST") || (dir1 == "WEST" && dir2 == "EAST") {
		return true
	}

	return false
}

func DirReduc(arr []string) []string {
	var result []string

	for {
		var removed = false

		result = make([]string, 0)
		for _, val := range arr {
			if val != "" {
				result = append(result, val)
			}
		}

		arr = result

		for i := 0; i < len(arr); i++ {
			if i+1 < len(arr) {
				if arr[i] != "" && arr[i+1] != "" && isOppoosite(arr[i], arr[i+1]) {
					arr[i] = ""
					arr[i+1] = ""

					i = i + 2
					removed = true
				}
			}
		}

		if ! removed {
			break
		}
	}

	return result
}

func main() {
	var a = []string{"NORTH", "SOUTH", "EAST", "WEST"}
	fmt.Println(DirReduc(a))

	a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	fmt.Println(DirReduc(a))

	a = []string{"NORTH", "WEST", "SOUTH", "EAST"}
	fmt.Println(DirReduc(a))

	a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
	fmt.Println(DirReduc(a))
}
