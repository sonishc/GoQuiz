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
  "time"
)

const FileLineLength = 100

var quizSize int  // -csv=
var timeLimit int // -csv=

func init() {
  const (
    defaultLength = FileLineLength
    lengthUsage   = "a csv file in the format of (question,answer)"
    defaultTimer  = 5
    timerUsage    = "a time in seconds to answer the questions"
  )

  flag.IntVar(&quizSize, "csv", defaultLength, lengthUsage)
  flag.IntVar(&timeLimit, "limit", defaultTimer, timerUsage)
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

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  questions := csv.NewReader(file)

  var trueAnsers = make([]int, 0) // new slice of True ansers

  i := 0
  timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
questionsloop:
  for i < lines {
    expression, err := questions.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%s = ", expression[0])

    answerCh := make(chan int)
    go func() {
      var userAns int
      fmt.Scan(&userAns)
      answerCh <- userAns
    }()

    answerRight, err := strconv.Atoi(expression[1])
    if err != nil {
      log.Fatal(err)
    }

    select {
    case <-timer.C:
      break questionsloop
    case answer := <-answerCh:
      if answer == answerRight {
        trueAnsers = append(trueAnsers, 1)
      }
    }
    i++
  }
  result(trueAnsers, i)
}

func result(rightAnswers []int, allQuestions int) {
  fmt.Println("\nAll questions = ", allQuestions)
  fmt.Println("Right:", len(rightAnswers))
  isPassed(len(rightAnswers) > allQuestions/2)
}

func isPassed(summ bool) {
  if summ {
    fmt.Println("Result: Passed")
  } else {
    fmt.Println("Result: Failed")
  }
}
