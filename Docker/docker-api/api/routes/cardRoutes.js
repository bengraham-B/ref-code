const express = require('express')
const router = express.Router()

const {getCards} = require('../controller/cards')

router.get("/", getCards)

module.exports = router