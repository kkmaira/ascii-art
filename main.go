package main

import (
	"errors"
	"fmt"
	"log"
	fs1 "mkassymk/ascii-art-reverse/functions"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var Dict map[byte][]string

func main() {
	banner, inputOs, output, align, reverse, err := fs1.GetArgs(os.Args)
	if err != nil {
		log.Println(err)
		return
	}
	if reverse != "" {
		err = fs1.Reverse(reverse)
		if err != nil {
			log.Println(err)
		}
		return
	}
	if banner == "" {
		banner = "standard"
	}

	input, err := InputBase(inputOs)
	if err != nil {
		log.Println(err)
		return
	}
	for i := range input {
		input[i] = strings.Join(strings.Fields(input[i]), " ")
	}
	Dict, err = GetASCII(banner)
	if err != nil {
		log.Println(err)
		return
	}

	ASCIILetters := MakeASCII(input)
	if err = IsValid(input); err != nil {
		log.Print(err)
		return
	}
	result := ApplyAlign(ASCIILetters, align, input)
	if strings.HasSuffix(os.Args[1], "\\n") && !HasOnlyNewLines(input) {
		result += "\n"
	}
	if output != "" {
		if strings.HasSuffix(output, ".txt") {
			err := os.WriteFile(output, []byte(result), 0o644)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("Wrong format of a result file")
			return
		}
		return
	}
	fmt.Print(result)
}
func IsValid(input []string) error {
	for i := range input {
		if GetTerminalWidth() < getInputLine(input[i]) {
			return errors.New("the input is too long")
		}
	}
	return nil
}
func HasOnlyNewLines(Input []string) bool {
	for i := range Input {
		if len(Input[i]) > 0 {
			return false
		}
	}
	return true
}
func GetSpaces(num int) [][][]string {
	spaceStr := ""
	for num != 0 {
		spaceStr += " "
		num--
	}
	temp := []string{}
	temp2 := [][]string{}
	for i := 0; i < 8; i++ {
		temp = append(temp, spaceStr)
	}
	return append([][][]string{}, append(temp2, temp))
}
func InputWidth(letters [][]string) (num int) {
	for i := range letters {
		if len(letters[i]) == 1 {
			continue
		}
		num += len(letters[i][7])
	}
	return num
}
func InputWidthForJustify(sentence [][][]string) (num int) {
	if len(sentence[0]) > 2 {
		for i := range sentence[0] {
			num += len([]byte(sentence[0][i][5]))
		}
	}
	return num
}
func ApplyAlign(res [][][]string, align string, input []string) string {
	var temp string
	for i, sentence := range res {
		switch align {
		case "right":
			for index := 0; index < 8; index++ {
				spaceLen := GetTerminalWidth() - InputWidth(sentence) - 1
				var s string
				for spaceLen != 0 {
					s += " "
					spaceLen--
				}
				spaceLen = 0
				temp += s
				for i := range sentence {
					temp += sentence[i][index]
				}
				temp += "\n"
			}
		case "center":
			for index := 0; index < 8; index++ {
				spaceLen := (GetTerminalWidth() - InputWidth(sentence)) / 2
				var s string
				for spaceLen != 0 {
					s += " "
					spaceLen--
				}
				spaceLen = 0
				temp += s
				for i := range sentence {
					temp += sentence[i][index]
				}
				temp += "\n"
			}
		case "justify":
			temp += ApplyMiniJustify(sentence, input[i])
		default:
			for index := 0; index < 8; index++ {
				stop := false
				for i := range sentence {
					if sentence[i] == nil {
						temp += "\n"
						stop = true
						break
					} else {
						temp += sentence[i][index]
					}
				}
				if stop {
					break
				}
				temp += "\n"
			}
		}
	}
	return temp
}
func getInputLine(s string) int {
	return InputWidthForJustify(MakeASCII([]string{(s)})) + 1
}
func ApplyMiniJustify(sentence [][]string, input string) string {
	var temp string
	if input == "" {
		return temp + "\n"
	}
	if strings.Count(input, " ") == 0 {
		for index := 0; index < 8; index++ {
			for i := range sentence {
				temp += sentence[i][index]
			}
			temp += "\n"
		}
		return temp
	}
	for index := 0; index < 8; index++ {
		first := true
		var spaceLen int
		var bonusSpaces int
		var s string
		for i := range input {
			temp += sentence[i][index]
			if input[i] == 32 {
				if first {
					first = false
					spaceLen += (GetTerminalWidth() - getInputLine(input)) / strings.Count(input, " ")
					bonusSpaces = getBonusSpaces(input, spaceLen)
					for spaceLen > 0 {
						s += " "
						spaceLen--
					}
				}
				for bonusSpaces != 0 {
					bonusSpaces--
					temp += " "
				}
				temp += s
			}
		}
		temp += "\n"
	}
	return temp
}
func getBonusSpaces(s string, spaceLen int) int {
	terminalWidth := GetTerminalWidth()
	inputWidth := getInputLine(s)
	countOfSpaces := strings.Count(s, " ")
	return (terminalWidth - (inputWidth + spaceLen*countOfSpaces))
}
func MakeASCII(input []string) [][][]string {

	var res [][][]string
	for i := range input {
		sentence := [][]string{}
		if input[i] == "" {
			res = append(res, [][]string{nil})
			continue
		}
		for j := range input[i] {
			sentence = append(sentence, Dict[input[i][j]])
		}
		res = append(res, sentence)
	}
	return res
}
func InputBase(arg string) ([]string, error) {
	ArrayElement := ""
	InputBase := []string{}
	for i := 0; i < len(arg); i++ {
		if arg[i] == 10 {
			InputBase = append(InputBase, ArrayElement)
			ArrayElement = ""
			continue
		}
		if i+1 < len(arg) && arg[i] == '\\' && arg[i+1] == 'n' {
			j := i
			countOfSlashes := 0
			for j >= 0 && arg[j] == '\\' {
				countOfSlashes++
				j--
			}
			if countOfSlashes%2 != 0 {
				if ArrayElement == "" {
					InputBase = append(InputBase, ArrayElement)
					i++
				} else {
					InputBase = append(InputBase, ArrayElement)
					ArrayElement = ""
					i++
				}
			}
		} else {
			ArrayElement += string(arg[i])
			if i == len(arg)-1 {
				InputBase = append(InputBase, ArrayElement)
			}
		}
	}
	return InputBase, nil
}
func GetASCII(font string) (map[byte][]string, error) {
	dict := make(map[byte][]string)
	f, err := os.ReadFile("fonts/" + font + ".txt")
	if err != nil {
		return nil, err
	}
	content := strings.Split(strings.ReplaceAll(string(f[1:]), "\r", ""), "\n\n")
	letter := byte(32)
	for i := range content {
		dict[letter] = strings.Split(content[i], "\n")
		letter++
	}
	return dict, nil
}
func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, _ := cmd.Output()
	width, _ := strconv.Atoi(strings.Fields(string(output))[1])
	return width
}
