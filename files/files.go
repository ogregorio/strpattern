package files

import (
	"bufio"
	"fmt"
	"os"
)

func Read(path string) ([]string, error) {

	archive, err := os.Open(path) //open file

	if err != nil {
		return nil, err
	} //if error was found

	defer archive.Close() //archive will be closed

	// Read file with scanner
	var lines []string
	scanner := bufio.NewScanner(archive)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err() //return lines or error
}

func WriteMultiLine(lines []string, path string) error {

	archive, err := os.Create(path) //create archive to write

	if err != nil {
		return err
	} //if error was found

	defer archive.Close() //archive will be closed

	// Write file with Writer
	writer := bufio.NewWriter(archive)
	for _, lines := range lines {
		fmt.Fprintln(writer, lines)
	}

	return writer.Flush() //return lines or error
}

func WriteOneLine(content string, path string) error {

	archive, err := os.Create(path) //create archive to write

	if err != nil {
		return err
	} //if error was found

	defer archive.Close() //archive will be closed

	// Write file with Writer
	writer := bufio.NewWriter(archive)
	fmt.Fprintln(writer, content)

	return writer.Flush() //return lines or error
}
