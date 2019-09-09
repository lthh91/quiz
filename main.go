package main
import(
    "bufio"
    "flag"
    "fmt"
    "encoding/csv"
    "io"
    "io/ioutil"
    "os"
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
    // pdf viewer's name is "Document Viewer"
    var csv_file = flag.String("csv_file", "problems.csv",
                               "Path to the data csv file")
    var time_limit = flag.Int("time_limit", 30, "Time allowed")
    flag.Parse()
    reader := bufio.NewReader(os.Stdin)
    var correct int
    timer := time.NewTimer(time.Second*time_limit)
    fmt.Print("Press Enter to start...")
    _, _ := reader.ReadString('\n')
    go func() {
        <-timer.c

    questions, answers := read_csv(*csv_file)
    for i := range questions {
        fmt.Print(questions[i],":")
        answer, _ := reader.ReadString('\n')
        if strings.Compare(answer, answers[i]) == 0 {
            correct ++;
        }
    }
    fmt.Printf("%d/%d\n", correct, len(questions))
}
