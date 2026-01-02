package fields

const BeginStringFIX42 = "FIX.4.2"

// OpenClose
const (
	Close = "C"
	Open  = "O"
)

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

// OrdType
const (
	OrdTypeMarket                = "1" // Market
	OrdTypeLimit                 = "2" // Limit
	OrdTypeStop                  = "3" // Stop
	OrdTypeStopLimit             = "4" // Stop limit
	OrdTypeMarketOnClose         = "5" // Market on close
	OrdTypeWithOrWithout         = "6" // With or without
	OrdTypeLimitOrBetter         = "7" // Limit or better
	OrdTypeLimitWithOrWithout    = "8" // Limit with or without
	OrdTypeOnBasis               = "9" // On basis
	OrdTypeOnClose               = "A" // On close
	OrdTypeLimitOnClose          = "B" // Limit on close
	OrdTypeForexMarket           = "C" // Forex - Market
	OrdTypePreviouslyQuoted      = "D" // Previously quoted
	OrdTypePreviouslyIndicated   = "E" // Previously indicated
	OrdTypeForexLimit            = "F" // Forex - Limit
	OrdTypeForexSwap             = "G" // Forex - Swap
	OrdTypeForexPreviouslyQuoted = "H" // Forex - Previously Quoted
	OrdTypeFunari                = "I" // Funari (Limit Day Order with unexecuted portion handled as Market On Close. e.g. Japan)
	OrdTypePegged                = "P" // Pegged
)

// PossDupFlag
const (
	PossDupFlagNo  = "N" // Original transmission
	PossDupFlagYes = "Y" // Possible duplicate
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

// TimeInForce
const (
	TimeInForceDay = "0" // Day
	TimeInForceGTC = "1" // Good Till Cancel (GTC)
	TimeInForceOPG = "2" // At the Opening (OPG)
	TimeInForceIOC = "3" // Immediate or Cancel (IOC)
	TimeInForceFOK = "4" // Fill or Kill (FOK)
	TimeInForceGTX = "5" // Good Till Crossing (GTX)
	TimeInForceGTD = "6" // Good Till Date
)

// Urgency
const (
	UrgencyNormal     = "0" // Normal
	UrgencyFlash      = "1" // Flash
	UrgencyBackground = "2" // Background
)

// SettlmntTyp
const (
	SettlmntTypRegular       = "0" // Regular
	SettlmntTypCash          = "1" // Cash
	SettlmntTypNextDay       = "2" // Next Day
	SettlmntTypT2            = "3" // T+2
	SettlmntTypT3            = "4" // T+3
	SettlmntTypT4            = "5" // T+4
	SettlmntTypFuture        = "6" // Future
	SettlmntTypWhenIssued    = "7" // When Issued
	SettlmntTypSellersOption = "8" // Sellers Option
	SettlmntTypT5            = "9" // T+5
)

// AllocTransType
const (
	AllocTransTypeNew                          = "0" // New
	AllocTransTypeReplace                      = "1" // Replace
	AllocTransTypeCancel                       = "2" // Cancel
	AllocTransTypePreliminary                  = "3" // Preliminary (without MiscFees and NetMoney)
	AllocTransTypeCalculated                   = "4" // Calculated (includes MiscFees and NetMoney)
	AllocTransTypeCalculatedWithoutPreliminary = "5" // Calculated without Preliminary (sent unsolicited by broker, includes MiscFees and NetMoney)
)

// ProcessCode
const (
	ProcessCodeRegular           = "0" // regular
	ProcessCodeSoftDollar        = "1" // soft dollar
	ProcessCodeStepIn            = "2" // step-in
	ProcessCodeStepOut           = "3" // step-out
	ProcessCodeSoftDollarStepIn  = "4" // soft-dollar step-in
	ProcessCodeSoftDollarStepOut = "5" // soft-dollar step-out
	ProcessCodePlanSponsor       = "6" // plan sponsor
)

// AllocStatus
const (
	AllocStatusAccepted      = "0" // accepted (successfully processed)
	AllocStatusRejected      = "1" // rejected
	AllocStatusPartialAccept = "2" // partial accept
	AllocStatusReceived      = "3" // received (received, not yet processed)
)

// AllocRejCode
const (
	AllocRejCodeUnknownAccounts                = "0" // unknown account(s)
	AllocRejCodeIncorrectQuantity              = "1" // incorrect quantity
	AllocRejCodeIncorrectAveragePrice          = "2" // incorrect average price
	AllocRejCodeUnknownExecutingBrokerMnemonic = "3" // unknown executing broker mnemonic
	AllocRejCodeCommissionDifference           = "4" // commission difference
	AllocRejCodeUnknownOrderID                 = "5" // unknown OrderID
	AllocRejCodeUnknownListID                  = "6" // unknown ListID
	AllocRejCodeOther                          = "7" // other
)

// EmailType
const (
	EmailTypeNew        = "0" // New
	EmailTypeReply      = "1" // Reply
	EmailTypeAdminReply = "2" // Admin Reply
)
