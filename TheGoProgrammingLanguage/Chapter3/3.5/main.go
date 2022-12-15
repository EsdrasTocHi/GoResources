package main

import "fmt"

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := range s {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}

	return false
}

func main() {
	rawString := `with this type of string
	
	the content is taken
	
	\n literally
	
	Raw strings are convenient way to write
	templates, regular expressions and text with multiple lines`

	fmt.Println(rawString)
	fmt.Println(HasPrefix(rawString, "with"))      // true
	fmt.Println(HasSuffix(rawString, "lines"))     // true
	fmt.Println(Contains(rawString, "convenient")) // true
}
