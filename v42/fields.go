package fields

type Tag int

const (
	// TagAccount - Account mnemonic as agreed between broker and institution.
	TagAccount Tag = iota + 1

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
)
