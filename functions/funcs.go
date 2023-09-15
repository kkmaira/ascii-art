package functions

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var Dict map[byte][]string

func IsValid(input []string) error {
	for i := range input {
		if GetTerminalWidth() < getInputLine(input[i]) {
			return errors.New("ERROR: the input is too long for this terminal window")
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

func InputWidth2(letters [][][]string) (num int) {
	for j := range letters {
		for i := range letters[j] {
			num += len([]byte(letters[j][i][7]))
		}
	}
	return num
}

func ApplyAlign(res [][][]string, align string, input []string) string {
	// for i := range input {
	// 	fmt.Println(i, input[i])
	// }
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
		case "justfiy":
			temp += ApplyMiniJustify(sentence, input[i])
		default:
			for index := 0; index < 8; index++ {
				for i := range sentence {
					temp += sentence[i][index]
				}
				temp += "\n"
			}
		}
	}

	return temp
}

func getInputLine(s string) int {
	return InputWidth2(MakeASCII([]string{(s)}))
}

func ApplyMiniJustify(sentence [][]string, input string) string {
	var temp string
	if input == "" {
		return temp + "\n"
	}
	amountOfSpaces := strings.Count(input, " ")
	if amountOfSpaces == 0 {
		for index := 0; index < 8; index++ {
			for i := range sentence {
				temp += sentence[i][index]
			}
			temp += "\n"
		}
		return temp
	}

	for index := 0; index < 8; index++ {
		for i := range input {
			temp += sentence[i][index]
			if input[i] == 32 {
				var s string
				spaceLen := (GetTerminalWidth() - getInputLine(input) - 2) / strings.Count(input, " ")
				for spaceLen != 0 {
					s += " "
					spaceLen--
				}
				temp += s
			}
		}
		temp += "\n"
	}

	return temp
}

func MakeASCII(input []string) [][][]string {
	var res [][][]string
	for i := range input {
		sentence := [][]string{}
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
