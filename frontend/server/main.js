const express = require('express')
const cookieParser = require('cookie-parser')

const loginMiddleware = require('./loginMiddleware')
const db = require('./db')
const k8s = require('./k8s')

const app = express()

app.use(express.json())
app.use(cookieParser())

app.get('/question', async (req, res) => {
  const category = req.query.category
  const questions = await db.getQuestions(category)
  res.json(questions)
})

app.post('/submission', loginMiddleware.loggedIn, async (req, res) => {
  const user = req.user
  console.log('Body', req.body)
  console.log('User', user)
  // Write submission ID and the rest to Postgres
  const submissionID = await db.createSubmission(user.email, req.body.questionID, req.body.text, req.body.language)

  // Use go daddy lib to create a pod
  await k8s.createPod(submissionID)

  res.json({
    id: submissionID
  })
})

app.get('/', loginMiddleware.loggedIn, (req, res) => {
  res.end('Hello')
})

app.get('/submission/:submissionID/outcome', loginMiddleware.loggedIn, async (req, res) => {
  let submission = req.params.submissionID
  console.log("Submission id is", submission)
  const outcomeResponse = await db.getOutcome(submission)
  console.log("Outcome response from DB", outcomeResponse)
  res.json(outcomeResponse)
})

loginMiddleware.setupAuthRoutes(app)

app.listen(80, async () => {
  await db.connect()
  console.log('Server Started')
})

// app.listen(8080, async () => {
//   // await db.connect()
//   console.log('Server Started')
// })