package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/patnaikshekhar/worker/solution"
)

const (
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

func Connect() (*sql.DB, error) {

	host := os.Getenv("POSTGRES_HOST")
	password := os.Getenv("POSTGRES_PASSWORD")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetSubmission gets the submission from the database
func GetSubmission(db *sql.DB, submissionID int) (*solution.Submission, error) {
	sqlStatement := "SELECT submission.id, userId, qstnId, solution, language, status, no_of_test_cases FROM submission, question WHERE submission.id=$1 and submission.qstnId = question.id"

	row := db.QueryRow(sqlStatement, submissionID)

	var submission *solution.Submission = &solution.Submission{}

	if err := row.Scan(
		&submission.ID,
		&submission.UserID,
		&submission.QuestionID,
		&submission.Solution,
		&submission.Language,
		&submission.Status,
		&submission.NoOfTestCases,
	); err != nil {
		return nil, err
	}

	return submission, nil
}

// Update database to hold the outcome of tests made for a submission. Update status
// in the Submission table to "Success/Failed" and writes to a SubmissionOutcomes
// table.
func UpdateSubmission(db *sql.DB, sub *solution.Submission, outcome solution.RunOutcome) error {

	// Retrieve record from table and update status
	sqlStatement := "UPDATE submission SET status=$1 WHERE id = $2"

	_, err := db.Exec(sqlStatement, outcome.Status, sub.ID)

	if err != nil {
		return err
	}

	// Insert record into SubmissionOutcomes table

	insertOutcomeSqlStmt := "INSERT into submission_outcomes (test_case, expected_outcome, actual_outcome, submission_id) values ($1, $2, $3, $4)"

	for _, c := range outcome.FailedCases {

		_, err := db.Exec(insertOutcomeSqlStmt, c.Input, c.ExpectedOutput, c.ActualOutput, sub.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// CREATE TABLE submission_outcomes
// (
//   id serial PRIMARY KEY,
//   test_case text,
//   expected_outcome text,
//   actual_outcome text,
//   submission_id integer REFERENCES submission (id)
// );
