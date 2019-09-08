package main
import(
	"fmt"
	"encoding/csv"
	"io"
	"io/ioutil"
	"strings"
)

func read_csv(csv_file string) ([]string, []string) {
	dat, _ := ioutil.ReadFile(csv_file)
	r := csv.NewReader(strings.NewReader(string(dat)))
	var questions []string
	var answers []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		questions = append(questions, record[0])
		answers = append(answers, record[1])
	}
	return questions, answers
}

func main() {
	csv_file := "problems.csv"
	questions, answers := read_csv(csv_file)
	fmt.Printf("%q\n", questions)
	fmt.Printf("%q\n", answers)
}