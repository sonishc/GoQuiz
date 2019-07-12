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
  "io"
  "log"
  "os"
  "strconv"
)

const FileLineLength = 100

var quizSize int // -csv=

func init() {
  const (
    defaultGopher = FileLineLength
    usage         = "a csv file in the format of (question,answer)"
  )

  flag.IntVar(&quizSize, "csv", defaultGopher, usage)
  flag.Parse()
}

func main() {
  if validateLength() {
    return
  }
  readCsv(quizSize)
}

func validateLength() bool {
  status := quizSize > FileLineLength

  if status {
    fmt.Println("Use flag number less than <", quizSize)
  } else {
    fmt.Printf("Quiz has %d question(s)\n\n", quizSize)
  }
  return status
}

func readCsv(lines int) {
  file, err := os.Open("problems.csv") // For read access.
  defer os.Exit(1)

  if err != nil {
    fmt.Println(err)
  }

  questions := csv.NewReader(file)

  var trueAnsers = make([]int, 0) // new slice of True ansers

  i := 0
  for i < lines {
    expression, err := questions.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%s = ", expression[0])

    if checkUserAnswer(expression[1]) {
      trueAnsers = append(trueAnsers, 1)
    }
    i++
  }
  result(trueAnsers, i)
}

func checkUserAnswer(answer string) bool {
  answer1, err := strconv.Atoi(answer)
  if err != nil {
    log.Fatal(err)
  }

  var userAns int
  fmt.Scan(&userAns)

  return answer1 == userAns
}

func result(rightAnswers []int, allQuestions int) {
  fmt.Println("\nAll questions = ", allQuestions)
  fmt.Println("Right:", len(rightAnswers))
  fmt.Println("Result:", len(rightAnswers) > allQuestions/2)
}
