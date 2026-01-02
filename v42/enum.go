package fields

const BeginStringFIX42 = "FIX.4.2"

// AdvSide
const (
	AdvSideBuy   = "B" // Buy
	AdvSideSell  = "S" // Sell
	AdvSideTrade = "T" // Trade
	AdvSideCross = "X" // Cross
)

// AdvTransType
const (
	AdvTransTypeCancel  = "C" // Cancel
	AdvTransTypeNew     = "N" // New
	AdvTransTypeReplace = "R" // Replace
)

// CommType
const (
	CommTypePerShare   = "1" // Per share
	CommTypePercentage = "2" // Percentage
	CommTypeAbsolute   = "3" // Absolute
)

// ExecInst
const (
	ExecInstStayOnOfferside            = "0" // Stay on offerside
	ExecInstNotHeld                    = "1" // Not held
	ExecInstWork                       = "2" // Work
	ExecInstGoAlong                    = "3" // Go along
	ExecInstOverTheDay                 = "4" // Over the day
	ExecInstHeld                       = "5" // Held
	ExecInstParticipateDontInitiate    = "6" // Participate don't initiate
	ExecInstStrictScale                = "7" // Strict scale
	ExecInstTryToScale                 = "8" // Try to scale
	ExecInstStayOnBidside              = "9" // Stay on bidside
	ExecInstNoCross                    = "A" // No cross (cross is forbidden)
	ExecInstOKToCross                  = "B" // OK to cross
	ExecInstCallFirst                  = "C" // Call first
	ExecInstPercentOfVolume            = "D" // Percent of volume (indicates that the sender does not want to be all the volume on the floor vs. a specific percentage)
	ExecInstDNI                        = "E" // Do not increase - DNI
	ExecInstDNR                        = "F" // Do not reduce - DNR
	ExecInstAON                        = "G" // All or none - AON
	ExecInstInstitutionsOnly           = "I" // Institutions only
	ExecInstLastPeg                    = "L" // Last peg (last sale)
	ExecInstMidPricePeg                = "M" // Mid-price peg (midprice of inside quote)
	ExecInstNonNegotiable              = "N" // Non-negotiable
	ExecInstOpeningPeg                 = "O" // Opening peg
	ExecInstMarketPeg                  = "P" // Market peg
	ExecInstPrimaryPeg                 = "R" // Primary peg (primary market - buy at bid/sell at offer)
	ExecInstSuspend                    = "S" // Suspend
	ExecInstFixedPeg                   = "T" // Fixed Peg to Local best bid or offer at time of order
	ExecInstCustomerDisplayInstruction = "U" // Customer Display Instruction (Rule11Ac1-1/4)
	ExecInstNetting                    = "V" // Netting (for Forex)
	ExecInstPegToVWAP                  = "W" // Peg to VWAP
)

// ExecTransType
const (
	ExecTransTypeNew     = "0" // New
	ExecTransTypeCancel  = "1" // Cancel
	ExecTransTypeCorrect = "2" // Correct
	ExecTransTypeStatus  = "3" // Status
)

// HandlInst
const (
	HandlInstAutoExecOrderPrivate = "1" // Automated execution order, private, no Broker intervention
	HandlInstAutoExecOrderPublic  = "2" // Automated execution order, public, Broker intervention OK
	HandlInstManualOrder          = "3" // Manual order, best execution
)

// IDSource
const (
	IDSourceCUSIP       = "1" // CUSIP
	IDSourceSEDOL       = "2" // SEDOL
	IDSourceQUIK        = "3" // QUIK
	IDSourceISIN        = "4" // ISIN number
	IDSourceRIC         = "5" // RIC code
	IDSourceISOCurrency = "6" // ISO Currency Code
	IDSourceISOCountry  = "7" // ISO Country Code
	IDSourceExchange    = "8" // Exchange Symbol
	IDSourceCTA         = "9" // Consolidated Tape Association (CTA) Symbol (SIAC CTS/CQS line format)
)

// IOIQltyInd
const (
	IOIQltyIndHigh   = "H" // High
	IOIQltyIndLow    = "L" // Low
	IOIQltyIndMedium = "M" // Medium
)

// IOIShares
const (
	IOISharesLarge  = "L" // Large
	IOISharesMedium = "M" // Medium
	IOISharesSmall  = "S" // Small
)

// IOITransType
const (
	IOITransTypeCancel  = "C" // Cancel
	IOITransTypeNew     = "N" // New
	IOITransTypeReplace = "R" // Replace
)

// LastCapacity
const (
	LastCapacityAgent            = "1" // Agent
	LastCapacityCrossAsAgent     = "2" // Cross as agent
	LastCapacityCrossAsPrincipal = "3" // Cross as principal
	LastCapacityPrincipal        = "4" // Principal
)

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
