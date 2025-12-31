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
)
