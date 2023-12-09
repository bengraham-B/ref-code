const cards = require("../cards.json")

const getCards = (req, res) => {
    
    res.status(200).json(cards)
}

module.exports = {getCards}