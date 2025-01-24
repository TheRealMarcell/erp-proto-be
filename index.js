const express = require('express')
const morgan = require('morgan')
const app = express()

// middleware
app.use(express.json())
app.use(morgan('tiny'))

app.get("", (request, response) => {
  response.send('<h1>ERP API</h1>')
})

app.listen(3001, () => {
  console.log(`Server running on port 3001`)
})