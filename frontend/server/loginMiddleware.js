const verifier = require('google-id-token-verifier')
const request = require('request-promise')

const config = require('./config.json')

const loggedIn = (req, res, next) => {
  // Check if cookie containing JWT exists
  const cookie = req.cookies['hr-login']
  console.log('Cookie is', cookie)
  if (cookie) {
    // Check validity of cookie
    verifier.verify(cookie, config.google_client_id, (err, tokenInfo) => {
      if (err) {
        console.log('Invalid cookie', err)
        res.redirect('/login')
      } else {
        req.user = tokenInfo
        next()
      }
    })
  } else {
    res.redirect('/login')
  }
}

const setupAuthRoutes = (app) => {
  app.get('/login', (req, res) => {
    const googleAuthURL = `${config.google_auth_uri}?scope=openid email&response_type=code&redirect_uri=${config.google_redirect_uri}&client_id=${config.google_client_id}`
    res.redirect(googleAuthURL)
  })

  app.get('/callback', async (req, res) => {
    if (req.query.code) {
      const resp = await request.post(config.google_token_uri, {
        form: {
          code: req.query.code,
          client_id: config.google_client_id,
          client_secret: config.google_client_secret,
          redirect_uri: config.google_redirect_uri,
          grant_type: 'authorization_code'
        }
      })

      const response = JSON.parse(resp)
      verifier.verify(response.id_token, config.google_client_id, (err, tokenInfo) => {
        if (err) {
          res.redirect('/login')
          return
        }
        console.log(tokenInfo)
        res.cookie('hr-login', response.id_token)
        res.redirect('/')
      })
    } else {
      res.end(req.params.error)
    }
  })
}

module.exports = {
  loggedIn,
  setupAuthRoutes
}