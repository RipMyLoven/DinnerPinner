package str

import "fmt"

func StrName(name1, name2 string) string {
	return fmt.Sprintf("%s %s", name1, name2)
}
func StrBool(name1, name2 string) bool {
	return name1 == name2
}
func StrInt(name1, name2 string) int {
	return len(name1) + len(name2)
}
