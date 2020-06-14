package solution

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://hcclone.blob.core.windows.net/testcases/"
const baseDirectory = "./problem"

func DownloadTestCasesAndExpectedOutputs(submission *Submission) error {
	for i := 0; i < submission.NoOfTestCases; i++ {
		err := downloadFile(
			fmt.Sprintf("TestCase_%d_%d",
				submission.QuestionID,
				i))

		if err != nil {
			return err
		}

		err = downloadFile(
			fmt.Sprintf("Output_%d_%d",
				submission.QuestionID,
				i))

		if err != nil {
			return err
		}
	}

	return nil
}

func downloadFile(fileName string) error {
	url := fmt.Sprintf("%s%s", baseURL, fileName)
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	fileContents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(
		fmt.Sprintf("%s/%s", baseDirectory, fileName),
		fileContents,
		0644)
	return err
}
