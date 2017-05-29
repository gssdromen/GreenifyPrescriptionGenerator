package main

import "io/ioutil"
import "fmt"
import "strings"

func main() {
	fmt.Println(`<prescriptions xmlns="http://greenify.github.io/schemas/prescription/v3">`)
	dat, _ := ioutil.ReadFile("a.txt")
	list := strings.Split(string(dat), " ")

	for _, v := range list {
		if v != "" {
			array := strings.Split(v, "/")
			packageName := array[0]
			className := array[1]
			format := `<prescription type="service" package="%s" class="%s" sender="any"/>
			`
			fmt.Printf(format, packageName, className)
		}
	}
	fmt.Println(`</prescriptions>`)
}
