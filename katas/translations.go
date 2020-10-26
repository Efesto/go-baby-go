package main

import (
	"fmt"
	"os"
)

func main() {
	myFirstMap := map[string]string{
		"First":  "Erste",
		"Second": "Zweite",
		"Third":  "Dritte",
		"Fourth": "Vierte",
		"Fifth":  "Fünfte",
		"Sixth":  "Sechste",
	}

	for key, value := range myFirstMap {
		fmt.Println(key, "ist", value)
		fmt.Println(spellForInEnglish(value))
	}
}

func spellForInEnglish(term string) string {
	var spell string
	for _, c := range term {
		switch c {
		case 'w', 'W':
			spell += "v"
		case 'z', 'Z':
			spell += "ts"
		case 'V', 'v':
			spell += "f"
		default:
			spell += fmt.Sprintf("%c", c)
		}
	}
	return spell
}

func printingFiesta() {
	fmt.Printf("This prints on STDOUT on a single line\n")
	newLine()
	fmt.Println("This prints on STDOUT and leaves a line")
	newLine()
	fmt.Println(fmt.Sprint("This prints on a string"))
	newLine()
	fmt.Fprint(os.Stdout, "This ", "formats ", "something ", "in a string", "\n")
	newLine()
	fmt.Printf("this prints something like %v in standard format\n", [4]int{1, 2, 3, 4})
	newLine()
}
func newLine() {
	fmt.Println("---")
}
