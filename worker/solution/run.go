package solution

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Run executes the command passing in inputs and compares outputs
func Run(submission *Submission, baseDirectory string) (RunOutcome, error) {

	result := RunOutcome{Status: "Passed", FailedCases: []Case{}}

	command := map[string]LanguageCommand{
		"python":     LanguageCommand{Command: "python3", Args: []string{"solution.py"}},
		"javascript": LanguageCommand{Command: "node", Args: []string{"solution.js"}},
		"go":         LanguageCommand{Command: "go", Args: []string{"run", "solution.go"}},
	}

	for i := 0; i < submission.NoOfTestCases; i++ {
		log.Printf(
			"Running command %s with args %s",
			command[submission.Language].Command,
			command[submission.Language].Args,
		)
		cmd := exec.Command(
			command[submission.Language].Command,
			command[submission.Language].Args...,
		)

		cmd.Dir = baseDirectory
		stdin, err := cmd.StdinPipe()
		if err != nil {
			result.Status = "Failed"
			return result, err
		}

		go func() {
			fileName := fmt.Sprintf("%s/TestCase_%d_%d", baseDirectory, submission.QuestionID, i)
			file, _ := os.OpenFile(fileName, os.O_RDONLY, 0755)
			if err != nil {
				result.Status = "Failed"
				return
			}
			defer file.Close()
			defer stdin.Close()
			io.Copy(stdin, file)
		}()

		out, _ := cmd.CombinedOutput()
		// if err != nil {
		// 	result.FailedCases = append(result.FailedCases, Case{string(inputFile), string(expectedOutput), string(out)})
		// 	result.Status = "Failed"
		// 	continue
		// }

		outputFileName := fmt.Sprintf("%s/Output_%d_%d", baseDirectory, submission.QuestionID, i)
		expectedOutput, err := ioutil.ReadFile(outputFileName)
		if err != nil {
			result.Status = "Failed"
			return result, err
		}

		// Compare the outputs
		if strings.TrimSpace(string(expectedOutput)) != strings.TrimSpace(string(out)) {
			result.Status = "Failed"
			fileName := fmt.Sprintf("%s/TestCase_%d_%d", baseDirectory, submission.QuestionID, i)
			inputFile, _ := ioutil.ReadFile(fileName)
			result.FailedCases = append(result.FailedCases, Case{string(inputFile), string(expectedOutput), string(out)})
		}

	}

	return result, nil
}

type RunOutcome struct {
	Status      string
	FailedCases []Case
}

type Case struct {
	Input          string
	ExpectedOutput string
	ActualOutput   string
}

type LanguageCommand struct {
	Command string
	Args    []string
}

// Submission Represents a code submission
type Submission struct {
	ID            int
	UserID        string
	QuestionID    int
	Solution      string
	Language      string
	Status        string
	NoOfTestCases int
}
