package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structPractice/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display(userNote)
	err = userNote.SaveFile()

	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of the note: ")
	content := getUserInput("Enter the content of the note: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	text := strings.TrimSuffix(input, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
