const cards = require("../cards.json")
const {cl, dateString} = require('goosefuncs')

const getCards = (req, res) => {
    
    res.status(200).json(cards)
}

module.exports = {getCards}