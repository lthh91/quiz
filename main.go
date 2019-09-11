package main
import(
    "bufio"
    "flag"
    "fmt"
    "encoding/csv"
    "io"
    "io/ioutil"
    "os"
    "time"
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
    var csv_file = flag.String("csv_file", "problems.csv",
                               "Path to the data csv file")
    var time_limit = flag.Int("time_limit", 30, "Time allowed")
    flag.Parse()
    reader := bufio.NewReader(os.Stdin)
    var correct int
    questions, answers := read_csv(*csv_file)
    //timer := time.NewTimer(time.Duration(*time_limit)*time.Second)
    fmt.Print("Press Enter to start...")
    fmt.Scanln()
    go func() {
        for i := range questions {
            fmt.Print(questions[i],":")
            answer, _ := reader.ReadString('\n')
            if strings.Compare(answer, answers[i]) == 0 {
                correct ++;
            }
        }
    }()
    <-time.After(time.Second * time.Duration(*time_limit))
    fmt.Printf("\n%d/%d\n", correct, len(questions))
}
