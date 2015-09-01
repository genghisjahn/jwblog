package main

import "fmt"

func ProcessText(text string, items []TextLengthItem) string {
	firstParts := []string{"This text is %v characters in length.", "Text is %v characters long.", "Text is %v characters in length."}
	result_template := "%v  %v"
	textLength := len(text) + 2
	for _, firstPart := range firstParts {
		for _, item := range items {
			length_sentence := fmt.Sprintf(firstPart, item.Text)
			if textLength+len(length_sentence) == item.Value {
				return fmt.Sprintf(result_template, text, length_sentence)
			}

		}
	}
	return "No match found."
}
