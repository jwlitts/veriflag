package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateInRange(typename string) string {

	inrange := `func InRangeGEN(start GEN, end GEN) func(x GEN) bool {
		return func(x GEN) bool {
			return x > start && x < end
		}
	}
	`
	return strings.ReplaceAll(inrange, "GEN", typename)
}
func GenerateInList(typename string) string {

	inrange := `func InListGEN(list []GEN) func(x GEN) bool {
		return func(x GEN) bool {
			for _, elem := range list {
				if elem == x {
					return true
				}
			}
			return false
		}
	}
	`
	return strings.ReplaceAll(inrange, "GEN", typename)
}

func main() {
	fl, err := os.OpenFile("conditions.go", os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fl.Close()
	fl.WriteString("package main\n")
	typenames := []string{"int", "int32", "int64", "string", "float32", "float64", "uint", "uint32", "uint64"}
	for _, typename := range typenames {

		fl.WriteString(GenerateInList(typename))
		fl.WriteString(GenerateInRange(typename))
	}
}
