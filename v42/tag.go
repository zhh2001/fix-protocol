package fields

const (
	// TagAccount - Account mnemonic as agreed between broker and institution.
	TagAccount = iota + 1

	// TagAdvId - Unique identifier of Advertisement <35=7> message.
	TagAdvId

	// TagAdvRefID - Reference identifier used with CANCEL and REPLACE transaction types.
	TagAdvRefID

	// TagAdvSide - Broker's side of advertised trade
	TagAdvSide

	// TagAdvTransType - Identifies Advertisement <35=7> message transaction type
	TagAdvTransType

	// TagAvgPx - Calculated average price of all fills on this order.
	TagAvgPx

	// TagBeginSeqNo - Message sequence number of first message in range to be resent
	TagBeginSeqNo

	// TagBeginString - Identifies beginning of new message and protocol version. ALWAYS FIRST FIELD IN MESSAGE. (Always
	// unencrypted)
	TagBeginString

	// TagBodyLength - Message length, in bytes, forward to the CheckSum (10) field. ALWAYS SECOND FIELD IN MESSAGE.
	// (Always unencrypted)
	TagBodyLength

	// TagCheckSum - Three byte, simple checksum. ALWAYS LAST FIELD IN MESSAGE; i.e. serves, with the trailing , as the
	// end-of-message delimiter. Always defined as three characters. (Always unencrypted)
	TagCheckSum

	// TagClOrdID - Unique identifier for Order as assigned by institution (identified by SenderCompID (49) or
	// OnBehalfOfCompID (115) as appropriate). Uniqueness must be guaranteed within a single trading day. Firms,
	// particularly those which electronically submit multi-day orders, trade globally or throughout market close
	// periods,should ensure uniqueness across days, for example by embedding a date within the ClOrdID (11) field.
	TagClOrdID

	// TagCommission - Commission. Note if CommType (13) is percentage, Commission (12) of 5% should be represented as
	// .05.
	TagCommission

	// TagCommType - Commission (12) type
	TagCommType

	// TagCumQty - Total number of shares filled.
	TagCumQty

	// TagCurrency - Identifies currency used for price. Absence of this field is interpreted as the default for the
	// security. It is recommended that systems provide the currency value whenever possible.
	TagCurrency

	// TagEndSeqNo - Message sequence number of last message in range to be resent. If request is for a single message
	// BeginSeqNo (7) = EndSeqNo (16). If request is for all messages subsequent to a particular message, EndSeqNo (16)
	// = 0 (representing infinity).
	TagEndSeqNo

	// TagExecID - Unique identifier of execution message as assigned by broker (will be 0 (zero) for ExecTransType
	// (20)='3' (Status)).
	//
	// Uniqueness must be guaranteed within a single trading day or the life of a multi-day order. Firms which accept
	// multi-day orders should consider embedding a date within the ExecID (17) field to assure uniqueness across days.
	TagExecID

	// TagExecInst - Instructions for order handling on exchange trading floor. If more than one instruction is
	// applicable to an order, this field can contain multiple instructions separated by space.
	TagExecInst

	// TagExecRefID - Reference identifier used with Cancel and Correct transaction types.
	TagExecRefID

	// TagExecTransType - Identifies transaction type
	TagExecTransType

	// TagHandlInst - Instructions for order handling on Broker trading floor
	TagHandlInst

	// TagIDSource - Identifies class of alternative SecurityID (48)
	TagIDSource

	// TagIOIid - Unique identifier of Indication of Interest <35=6> message.
	TagIOIid

	// TagIOIOthSvc - No longer used as of FIX 4.2.
	TagIOIOthSvc

	// TagIOIQltyInd - Relative quality of indication
	TagIOIQltyInd

	// TagIOIRefID - Reference identifier used with CANCEL and REPLACE, transaction types.
	TagIOIRefID

	// TagIOIShares - Number of shares in numeric or relative size.
	TagIOIShares

	// TagIOITransType - Identifies IOI message transaction type
	TagIOITransType

	// TagLastCapacity - Broker capacity in order execution
	TagLastCapacity

	// TagLastMkt - Market of execution for last fill
	TagLastMkt

	// TagLastPx - Price of this (last) fill. Field not required for ExecTransType (20) ='3' (Status)
	TagLastPx

	// TagLastShares - Quantity of shares bought/sold on this (last) fill. Field not required for ExecTransType (20)
	// ='3' (Status)
	TagLastShares

	// TagLinesOfText - Identifies number of lines of text body
	TagLinesOfText

	// TagMsgSeqNum - Integer message sequence number.
	TagMsgSeqNum

	// TagMsgType - Defines message type. Always third field in message. Always unencrypted.
	//
	// The value is case-sensitive.
	//
	// "U" as the first character of the value (e.g. U1, U2) indicates that the message format is privately defined
	// between the sender and receiver.
	TagMsgType

	// TagNewSeqNo - New sequence number
	TagNewSeqNo

	// TagOrderID - Unique identifier for Order as assigned by broker. Uniqueness must be guaranteed within a single
	// trading day. Firms which accept multi-day orders should consider embedding a date within the OrderID (37) field
	// to assure uniqueness across days.
	TagOrderID

	// TagOrderQty - Number of shares ordered. This represents the number of shares for equities or based on normal
	// convention the number of contracts for options, futures, convertible bonds, etc.
	TagOrderQty

	// TagOrdStatus - Identifies current status of order.
	TagOrdStatus

	// TagOrdType - Order type.
	TagOrdType

	// TagOrigClOrdID - ClOrdID (11) of the previous order (NOT the initial order of the day) as assigned by the
	// institution, used to identify the previous order in cancel and cancel/replace requests.
	TagOrigClOrdID

	// TagOrigTime - Time of message origination (always expressed in UTC (Universal Time Coordinated, also known as
	// 'GMT'))
	TagOrigTime

	// TagPossDupFlag - Indicates possible retransmission of message with this sequence number
	TagPossDupFlag

	// TagPrice - Price per share
	TagPrice

	// TagRefSeqNum - Reference message sequence number
	TagRefSeqNum

	// TagRelatdSym - Symbol of issue related to story. Can be repeated within message to identify multiple companies.
	TagRelatdSym

	// TagRule80A - Note that the name of this field is changing to 'OrderCapacity' as Rule80A is a very US
	// market-specific term. Other world markets need to convey similar information, however, often a subset of the US
	// values.
	TagRule80A

	// TagSecurityID - CUSIP or other alternate security identifier
	TagSecurityID

	// TagSenderCompID - Assigned value used to identify firm sending message.
	TagSenderCompID

	// TagSenderSubID - Assigned value used to identify specific message originator (desk, trader, etc.)
	TagSenderSubID

	// TagSendingDate - No longer used. Included here for reference to prior versions.
	TagSendingDate

	// TagSendingTime - Time of message transmission (always expressed in UTC (Universal Time Coordinated, also known as
	// 'GMT')
	TagSendingTime

	// TagShares - Number of shares
	TagShares

	// TagSide - Side of order
	TagSide

	// TagSymbol - Ticker symbol
	TagSymbol

	// TagTargetCompID - Assigned value used to identify receiving firm.
	TagTargetCompID

	// TagTargetSubID - Assigned value used to identify specific individual or unit intended to receive message. 'ADMIN'
	// reserved for administrative messages not intended for a specific user.
	TagTargetSubID

	// TagText - Free format text string
	TagText

	// TagTimeInForce - Specifies how long the order remains in effect. Absence of this field is interpreted as DAY.
	TagTimeInForce

	// TagTransactTime - Time of execution/order creation (expressed in UTC (Universal Time Coordinated, also known as
	// 'GMT')
	TagTransactTime

	// TagUrgency - Urgency flag
	TagUrgency

	// TagValidUntilTime - Indicates expiration time of indication message (always expressed in UTC (Universal Time
	// Coordinated, also known as 'GMT')
	TagValidUntilTime

	// TagSettlmntTyp - Indicates order settlement period. Absence of this field is interpreted as Regular. Regular is
	// defined as the default settlement period for the particular security on the exchange of execution.
	TagSettlmntTyp

	// TagFutSettDate - Specific date of trade settlement (SettlementDate) in YYYYMMDD format. Required when SettlmntTyp
	// (63) = '6' (Future) or SettlmntTyp (63) = '8' (Sellers Option). (expressed in local time at place of settlement)
	TagFutSettDate
)
