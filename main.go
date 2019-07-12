/*
  Create a program that will read in a quiz provided via a CSV file (more details below)
  and will then give the quiz to a user keeping track of how many questions they
  get right and how many they get incorrect. Regardless of whether the answer is
  correct or wrong the next question should be asked immediately afterwards.
  The CSV file should default to problems.csv (example shown below), but the user
  should be able to customize the filename via a flag.
  The CSV file will be in a format like below, where the first column is a question
  and the second column in the same row is the answer to that question.
*/

package main

import (
  "encoding/csv"
  "flag"
  "fmt"
  "strconv"
  // "strings"
  // "time"
  "io"
  "log"
  "os"
)

var linesLength string // -flag
const FileLineLength = 100

func init() {
  const (
    defaultGopher = "problems.csv"
    usage         = "a csv file in the format of (question,answer)"
  )

  flag.StringVar(&linesLength, "csv", defaultGopher, usage)
  flag.Parse()
}

func main() {
  linesLength, err := strconv.Atoi(linesLength)
  if err != nil {
    fmt.Println(err)
    return
  }
  if linesLength > FileLineLength {
    fmt.Println("Use flag less than <", linesLength)
    return
  }
  fmt.Printf("Quiz has %d question(s)\n\n", linesLength)
  readCsv(linesLength)
}

func readCsv(lines int) {
  file, err := os.Open("problems.csv") // For read access.
  defer os.Exit(1)

  if err != nil {
    fmt.Println(err)
  }

  r := csv.NewReader(file)

  i := 0
  var trueCount = make([]int, 0, 3) // new slice
  for i < lines {
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }

    fmt.Printf("%s = ", record[0])

    var userAnswer int
    fmt.Scan(&userAnswer)

    answer, err := strconv.Atoi(record[1])
    if err != nil {
      fmt.Println(err)
      return
    }

    if answer == userAnswer {
      trueCount = append(trueCount, 1)
    }

    i++
  }
  result(trueCount, i)
}

func result(rightAnswers []int, allQuestions int) {
  fmt.Println("\nAll questions = ", allQuestions)
  fmt.Println("Right:", len(rightAnswers))
  fmt.Println("Result:", len(rightAnswers) > allQuestions/2)
}
