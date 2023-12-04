package main

import (
	// "bytes"
	"bytes"
	"fmt"
	"strings"

	// "strings"
	"regexp"
)

func main() {
	// use_string_library()
	// use_bytes_library()
	use_regexp_library()
}

func use_bytes_library() {
	price := "€100"
	fmt.Println("String Prefix:", strings.HasPrefix(price, "€"))
	fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price), []byte{226, 130}))
}

func use_string_library() {
	product := "Kayak"
	fmt.Println("Product:", product)
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
	description := "A boat for sailing"
	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.Title(description))
	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))
}

func use_regexp_library() {
	pattern, compileErr := regexp.Compile("[A-z]oat")
	description := "A boat for one person"
	// match, err := regexp.MatchString("[A-z]oat", description)
	// if err == nil {
	// 	fmt.Println("Match:", match)
	// } else {
	// 	fmt.Println("Error:", err)
	// }
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}
}
