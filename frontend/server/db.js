const { Client } = require('pg')
const client = new Client({
  host: process.env.POSTGRES_HOST,
  password: process.env.POSTGRES_PASSWORD,
  database: 'postgres',
  user: 'postgres'
})

const connect = async () => {
  await client.connect()
}

const getQuestions = async (category) => {
  const res = await client.query('SELECT * FROM question WHERE category = $1', [category])
  return res.rows
}

const createSubmission = async (user, qstnId, solution, lang) => {
  const insertStatement = "INSERT INTO submission (userId, qstnId, solution, language, status) VALUES ($1, $2, $3, $4, $5) RETURNING id"
  const res = await client.query(insertStatement, [user, qstnId, solution, lang, 'Pending'])
  console.log("Submission to Postgres", res.rows[0])
  return res.rows[0].id

}

const getOutcome = async (submission) => {
  const res = await client.query('SELECT status FROM submission WHERE id = $1', [submission])
  if (res.rows[0].status != "Pending") {
    const outcomeRes = await client.query('SELECT * FROM submission_outcomes WHERE submission_id = $1', [submission])
    return {
      status: res.rows[0].status,
      outcomes: outcomeRes.rows
    }
  }
  return {
    status: "Pending",
    outcomes: []
  }
}

module.exports = {
  connect,
  getQuestions,
  createSubmission,
  getOutcome
}