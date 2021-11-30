package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if !(len(os.Args) == 2 && len(os.Args[1]) >= 10 && os.Args[1][:10] == "--reverse=") {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}
	temp, err := readLines("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	file, err2 := readLines(os.Args[1][10:])
	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Println(reverse(temp, file))
}

// This function reads the font file
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// This function reverses ascii art
func reverse(temp []string, file []string) string {

	result := ""
	startPos := 0
	endPos := 0
	temp = temp[1:]

	fileLines := len(file) / 8

	for i := 0; i < fileLines; i++ {
		fileSpace := []int{}
		//Check how long is file
		for j := 0; j < len(file[i]); j++ {
			fileSpace = append(fileSpace, j)
		}
		for k := 0; k < len(fileSpace); k++ {
			for l := 0; l < 8; l++ {
				if l == 7 {
					endPos = k + 1
					flag := false
					//Check every letter in template
					for m := 0; m < len(temp); m += 9 {
						for line := 0; line < 8; line++ {
							//Break if letter doesn't match, go to next letter in template
							if file[line][startPos:endPos] != temp[m+line] {
								break
							}
							//If matches, add correct letter to result
							if line == 7 {
								result += string(rune(m/9 + 32))
								flag = true
							}
						}
						//Flag controls the location of letter position, if it's true it goes to the next letter
						if flag {
							startPos = endPos
							break
						}
					}
					break
				}
			}
		}
	}
	return result
}
