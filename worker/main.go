package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/patnaikshekhar/worker/db"
	"github.com/patnaikshekhar/worker/solution"
)

func main() {
	// Steps

	// Connect to Postgres and get the submission
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	submissionID, err := strconv.Atoi(os.Getenv("SUBMISSION_ID"))
	if err != nil {
		panic(err)
	}

	sub, err := db.GetSubmission(database, submissionID)
	if err != nil {
		panic(err)
	}

	log.Printf("Submission is %+v", sub)

	// Based on code submitted create a file
	err = createFile(sub)
	if err != nil {
		panic(err)
	}

	// Based on question and language download all test cases
	err = solution.DownloadTestCasesAndExpectedOutputs(sub)
	if err != nil {
		panic(err)
	}

	// Execute Test Cases one by one and check results with expected output
	runResult, err := solution.Run(sub, baseDirectory)
	if err != nil {
		log.Printf("%s", err.Error())
		// panic(err)
	}
	log.Printf("Run Result is %v", runResult)

	// Write results to database
	err = db.UpdateSubmission(database, sub, runResult)
	if err != nil {
		panic(err)
	}
}

const baseDirectory = "./problem"

func createFile(sub *solution.Submission) error {
	extension := "py"

	switch sub.Language {
	case "Go":
		extension = ".go"
	case "Node":
		extension = ".js"
	}

	fileName := fmt.Sprintf("%s/solution.%s", baseDirectory, extension)

	err := ioutil.WriteFile(fileName, []byte(sub.Solution), 0644)
	if err != nil {
		return err
	}

	return nil
}
