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

// MsgType
const (
	MsgTypeAdvertisement                 = "7" // Advertisement
	MsgTypeAllocation                    = "J" // Allocation
	MsgTypeAllocationAck                 = "P" // Allocation Ack
	MsgTypeBidRequest                    = "k" // Bid Request
	MsgTypeBidResponse                   = "l" // Bid Response
	MsgTypeBusinessMessageReject         = "j" // Business Message Reject
	MsgTypeDontKnowTrade                 = "Q" // Don't Know Trade
	MsgTypeEmail                         = "C" // Email
	MsgTypeExecutionReport               = "8" // Execution Report
	MsgTypeHeartbeat                     = "0" // Heartbeat
	MsgTypeIndicationOfInterest          = "6" // Indication of Interest
	MsgTypeListCancelRequest             = "K" // List Cancel Request
	MsgTypeListExecute                   = "L" // List Execute
	MsgTypeListStatus                    = "N" // List Status
	MsgTypeListStatusRequest             = "M" // List Status Request
	MsgTypeListStrikePrice               = "m" // List Strike Price
	MsgTypeLogon                         = "A" // Logon
	MsgTypeLogout                        = "5" // Logout
	MsgTypeMarketDataIncrementalRefresh  = "X" // Market Data Incremental Refresh
	MsgTypeMarketDataRequest             = "V" // Market Data Request
	MsgTypeMarketDataRequestReject       = "Y" // Market Data Request Reject
	MsgTypeMarketDataSnapshotFullRefresh = "W" // Market Data Snapshot Full Refresh
	MsgTypeMassQuote                     = "i" // Mass Quote
	MsgTypeNewOrderList                  = "E" // New Order List
	MsgTypeNews                          = "B" // News
	MsgTypeOrderCancelReject             = "9" // Order Cancel Reject
	MsgTypeOrderCancelRequest            = "F" // Order Cancel Request
	MsgTypeOrderCancelOrReplaceRequest   = "G" // Order Cancel/Replace Request
	MsgTypeOrderSingle                   = "D" // Order Single
	MsgTypeOrderStatusRequest            = "H" // Order Status Request
	MsgTypeQuote                         = "S" // Quote
	MsgTypeQuoteAcknowledgement          = "b" // Quote Acknowledgement
	MsgTypeQuoteCancel                   = "Z" // Quote Cancel
	MsgTypeQuoteRequest                  = "R" // Quote Request
	MsgTypeQuoteStatusRequest            = "a" // Quote Status Request
	MsgTypeReject                        = "3" // Reject
	MsgTypeResendRequest                 = "2" // Resend Request
	MsgTypeSecurityDefinition            = "d" // Security Definition
	MsgTypeSecurityDefinitionRequest     = "c" // Security Definition Request
	MsgTypeSecurityStatus                = "f" // Security Status
	MsgTypeSecurityStatusRequest         = "e" // Security Status Request
	MsgTypeSequenceReset                 = "4" // Sequence Reset
	MsgTypeSettlementInstructions        = "T" // Settlement Instructions
	MsgTypeTestRequest                   = "1" // Test Request
	MsgTypeTradingSessionStatus          = "h" // Trading Session Status
	MsgTypeTradingSessionStatusRequest   = "g" // Trading Session Status Request
)

// OrdStatus
const (
	OrdStatusNew                = "0" // New
	OrdStatusPartiallyFilled    = "1" // Partially filled
	OrdStatusFilled             = "2" // Filled
	OrdStatusDoneForDay         = "3" // Done for day
	OrdStatusCanceled           = "4" // Canceled
	OrdStatusReplaced           = "5" // Replaced
	OrdStatusPendingCancel      = "6" // Pending Cancel (e.g. result of Order Cancel Request)
	OrdStatusStopped            = "7" // Stopped
	OrdStatusRejected           = "8" // Rejected
	OrdStatusSuspended          = "9" // Suspended
	OrdStatusPendingNew         = "A" // Pending New
	OrdStatusCalculated         = "B" // Calculated
	OrdStatusExpired            = "C" // Expired
	OrdStatusAcceptedForBidding = "D" // Accepted for bidding
	OrdStatusPendingReplace     = "E" // Pending Replace (e.g. result of Order Cancel/Replace Request)
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
