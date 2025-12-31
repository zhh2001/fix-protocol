package fields

// Side
const (
	SideBuy             = "1" // Buy
	SideSell            = "2" // Sell
	SideBuyMinus        = "3" // Buy minus
	SideSellPlus        = "4" // Sell plus
	SideSellShort       = "5" // Sell short
	SideSellShortExempt = "6" // Sell short exempt
	SideUndisclosed     = "7" // Undisclosed (valid for IOI and List Order messages only)
	SideCross           = "8" // Cross (orders where counterparty is an exchange, valid for all messages except IOIs)
	SideCrossShort      = "9" // Cross short
)
