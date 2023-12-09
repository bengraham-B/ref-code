const express = require('express')
const {dateString, cl} = require('goosefuncs')
const cors = require('cors')

const cardRoutes = require('./routes/cardRoutes')

const app = express()
app.use(cors())
app.use(express.json())

app.use('/api/cards', cardRoutes)

const port = 8029

app.listen(port, () =>{
    cl(`Server is running on PORT:${port} ${dateString(time=true)}`)
})