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
    "sync/atomic"
    "math/rand"
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
    var shuffle = flag.Bool("shuffle", false, "If the questions should be shuffled")
    flag.Parse()
    reader := bufio.NewReader(os.Stdin)
    var correct int32
    questions, answers := read_csv(*csv_file)
    if *shuffle {
        rand.Shuffle(len(questions), func(i, j int) {
            questions[i], questions[j] = questions[j], questions[i]
            answers[i], answers[j] = answers[j], answers[i]
        })
    }
    fmt.Print("Press Enter to start...")
    fmt.Scanln()
    done := make(chan bool)
    go func() {
        for i := range questions {
            fmt.Print(questions[i],":")
            answer, _ := reader.ReadString('\n')
            if strings.Compare(strings.TrimSpace(answer), strings.TrimSpace(answers[i])) == 0 {
                atomic.AddInt32(&correct, 1)
            }
        }
        done <- true
    }()
    go func() {
        time.Sleep(time.Second * time.Duration(*time_limit))
        done <- true
    }()
    <- done
    fmt.Printf("\nFinished. Your score is %d/%d\n", correct, len(questions))
}
