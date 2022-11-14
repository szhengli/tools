package tools

import "fmt"

func Echo(name string) string {
	fmt.Println(name)
	return name + "processed"
}
