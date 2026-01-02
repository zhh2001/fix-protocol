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

	// TagSymbolSfx - Additional information about the security (e.g. preferred, warrants, etc.).
	TagSymbolSfx

	// TagListID - Unique identifier for list as assigned by institution, used to associate multiple individual orders.
	// Uniqueness must be guaranteed within a single trading day. Firms which generate multi-day orders should consider
	// embedding a date within the ListID (66) field to assure uniqueness across days.
	TagListID

	// TagListSeqNo - Sequence of individual order within list (i.e. ListSeqNo (67) of TotNoOrders (68), 2 of 25, 3 of
	// 25, . . . )
	TagListSeqNo

	// TagTotNoOrders - Total number of list order entries across all messages. Should be the sum of all NoOrders (73)
	// in each message that has repeating list order entries related to the same ListID (66). Used to support
	// fragmentation.
	TagTotNoOrders

	// TagListExecInst - Free format text message containing list handling and execution instructions.
	TagListExecInst

	// TagAllocID - Unique identifier for Allocation <35=J> message.
	TagAllocID

	// TagAllocTransType - Identifies allocation transaction type
	TagAllocTransType

	// TagRefAllocID - Reference identifier to be used with Replace, Cancel, and Calculated AllocTransType (71)
	// messages.
	TagRefAllocID

	// TagNoOrders - Indicates number of orders to be combined for average pricing and allocation.
	TagNoOrders

	// TagAvgPrxPrecision - Indicates number of decimal places to be used for average pricing. Absence of this field
	// indicates that default precision arranged by the broker/institution is to be used.
	TagAvgPrxPrecision

	// TagTradeDate - Indicates date of trade referenced in this message in YYYYMMDD format. Absence of this field
	// indicates current day (expressed in local time at place of trade).
	TagTradeDate

	// TagExecBroker - Identifies executing / give-up broker. Standard NASD market-maker mnemonic is preferred.
	TagExecBroker

	// TagOpenClose - Indicates whether the resulting position after a trade should be an opening position or closing
	// position. Used for omnibus accounting - where accounts are held on a gross basis instead of being netted
	// together.
	TagOpenClose

	// TagNoAllocs - Number of repeating AllocAccount (79)/AllocPrice (366) entries.
	TagNoAllocs

	// TagAllocAccount - Sub-account mnemonic
	TagAllocAccount

	// TagAllocShares - Number of shares to be allocated to specific sub-account
	TagAllocShares

	// TagProcessCode - Processing code for sub-account. Absence of this field in AllocAccount (79) /
	// AllocPrice (366)/AllocShares (80) / ProcessCode (81) instance indicates regular trade.
	TagProcessCode

	// TagNoRpts - Total number of reports within series.
	TagNoRpts

	// TagRptSeq - Sequence number of message within report series.
	TagRptSeq

	// TagCxlQty - Total number of shares canceled for this order.
	TagCxlQty

	// TagNoDlvyInst - Number of delivery instruction fields to follow
	TagNoDlvyInst

	// TagDlvyInst - Free format text field to indicate delivery instructions
	TagDlvyInst

	// TagAllocStatus - Identifies status of allocation.
	TagAllocStatus

	// TagAllocRejCode - Identifies reason for rejection.
	TagAllocRejCode

	// TagSignature - Electronic signature
	TagSignature

	// TagSecureDataLen - Length of encrypted message
	TagSecureDataLen

	// TagSecureData - Actual encrypted data stream
	TagSecureData

	// TagBrokerOfCredit - Broker to receive trade credit.
	TagBrokerOfCredit

	// TagSignatureLength - Number of bytes in Signature (89) field.
	TagSignatureLength

	// TagEmailType - Email <35=C> message type.
	TagEmailType

	// TagRawDataLength - Number of bytes in raw data field.
	TagRawDataLength
)
