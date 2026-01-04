// Copyright 2025 Henghua Zhang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
	// BeginSeqNo (7) = EndSeqNo (16). If request is for all messages after a particular message, EndSeqNo (16) = 0
	// (representing infinity).
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

	// TagRawData - Unformatted raw data, can include bitmaps, word processor documents, etc.
	TagRawData

	// TagPossResend - Indicates that message may contain information that has been sent under another sequence number.
	TagPossResend

	// TagEncryptMethod - Method of encryption.
	TagEncryptMethod

	// TagStopPx - Price per share
	TagStopPx

	// TagExDestination - Execution destination as defined by institution when order is entered.
	TagExDestination

	// 101
	_

	// TagCxlRejReason - Code to identify reason for cancel rejection.
	TagCxlRejReason

	// TagOrdRejReason - Code to identify reason for order rejection.
	TagOrdRejReason

	// TagIOIQualifier - Code to qualify IOI use.
	TagIOIQualifier

	// TagWaveNo - Identifier to aid in the management of multiple lists derived from a single, master list.
	TagWaveNo

	// TagIssuer - Company name of security issuer (e.g. International Business Machines)
	TagIssuer

	// TagSecurityDesc - Security description.
	TagSecurityDesc

	// TagHeartBtInt - Heartbeat <35=0> interval (seconds)
	TagHeartBtInt

	// TagClientID - Firm identifier used in third party-transactions (should not be a substitute for OnBehalfOfCompID
	// (115)/DeliverToCompID (128)).
	TagClientID

	// TagMinQty - Minimum quantity of an order to be executed.
	TagMinQty

	// TagMaxFloor - Maximum number of shares within an order to be shown on the exchange floor at any given time.
	TagMaxFloor

	// TagTestReqID - Identifier included in Test Request <35=1> message to be returned in resulting Heartbeat <35=0>
	TagTestReqID

	// TagReportToExch - Identifies party of trade responsible for exchange reporting.
	TagReportToExch

	// TagLocateReqd - Indicates whether the broker is to locate the stock in conjunction with a short sell order.
	TagLocateReqd

	// TagOnBehalfOfCompID - Assigned value used to identify firm originating message if the message was delivered by a
	// third party i.e. the third party firm identifier would be delivered in the SenderCompID (49) field and the firm
	// originating the message in this field.
	TagOnBehalfOfCompID

	// TagOnBehalfOfSubID - Assigned value used to identify specific message originator (i.e. trader) if the message was
	// delivered by a third party
	TagOnBehalfOfSubID

	// TagQuoteID - Unique identifier for quote
	TagQuoteID

	// TagNetMoney - Total amount due as the result of the transaction (e.g. for Buy order - principal + commission +
	// fees) reported in currency of execution.
	TagNetMoney

	// TagSettlCurrAmt - Total amount due expressed in settlement currency (includes the effect of the forex
	// transaction)
	TagSettlCurrAmt

	// TagSettlCurrency - Currency code of settlement denomination.
	TagSettlCurrency

	// TagForexReq - Indicates request for forex accommodation trade to be executed along with security transaction.
	TagForexReq

	// TagOrigSendingTime - Original time of message transmission (always expressed in UTC (Universal Time Coordinated,
	// also known as 'GMT') when transmitting orders as the result of a resend request.
	TagOrigSendingTime

	// TagGapFillFlag - Indicates that the Sequence Reset <35=4> message is replacing administrative or application
	// messages which will not be resent.
	TagGapFillFlag

	// TagNoExecs - No of execution repeating group entries to follow.
	TagNoExecs

	// TagCxlType - No longer used. Included here for reference to prior versions.
	TagCxlType

	// TagExpireTime - Time/Date of order expiration (always expressed in UTC (Universal Time Coordinated, also known as
	// 'GMT')
	TagExpireTime

	// TagDKReason - Reason for execution rejection.
	TagDKReason

	// TagDeliverToCompID - Assigned value used to identify the firm targeted to receive the message if the message is
	// delivered by a third party i.e. the third party firm identifier would be delivered in the TargetCompID (56) field
	// and the ultimate receiver firm ID in this field.
	TagDeliverToCompID

	// TagDeliverToSubID - Assigned value used to identify specific message recipient (i.e. trader) if the message is
	// delivered by a third party
	TagDeliverToSubID

	// TagIOINaturalFlag - Indicates that IOI is the result of an existing agency order or a facilitation position
	// resulting from an agency order, not from principal trading or order solicitation activity.
	TagIOINaturalFlag

	// TagQuoteReqID - Unique identifier for quote request
	TagQuoteReqID

	// TagBidPx - Bid price/rate
	TagBidPx

	// TagOfferPx - Offer price/rate
	TagOfferPx

	// TagBidSize - Quantity of bid
	TagBidSize

	// TagOfferSize - Quantity of offer
	TagOfferSize

	// TagNoMiscFees - Number of repeating groups of miscellaneous fees
	TagNoMiscFees

	// TagMiscFeeAmt - Miscellaneous fee value
	TagMiscFeeAmt

	// TagMiscFeeCurr - Currency (15) of miscellaneous fee
	TagMiscFeeCurr

	// TagMiscFeeType - Indicates type of miscellaneous fee.
	TagMiscFeeType

	// TagPrevClosePx - Previous closing price of security.
	TagPrevClosePx

	// TagResetSeqNumFlag - Indicates that the both sides of the FIX session should reset sequence numbers.
	TagResetSeqNumFlag

	// TagSenderLocationID - Assigned value used to identify specific message originator's location (i.e. geographic
	// location and/or desk, trader)
	TagSenderLocationID

	// TagTargetLocationID - Assigned value used to identify specific message destination's location (i.e. geographic
	// location and/or desk, trader)
	TagTargetLocationID

	// TagOnBehalfOfLocationID - Assigned value used to identify specific message originator's location (i.e. geographic
	// location and/or desk, trader) if the message was delivered by a third party
	TagOnBehalfOfLocationID

	// TagDeliverToLocationID - Assigned value used to identify specific message recipient's location (i.e. geographic
	// location and/or desk, trader) if the message was delivered by a third party
	TagDeliverToLocationID

	// TagNoRelatedSym - Specifies the number of repeating symbols specified.
	TagNoRelatedSym

	// TagSubject - The subject of an Email <35=C> message
	TagSubject

	// TagHeadline - The headline of a News <35=B> message
	TagHeadline

	// TagURLLink - A URL (Uniform Resource Locator) link to additional information (i.e.
	// https://en.wikipedia.org/wiki/URL)
	TagURLLink

	// TagExecType - Describes the specific Execution Report <35=8> (i.e. Pending Cancel) while OrdStatus (39) will
	// always identify the current order status (i.e. Partially Filled)
	TagExecType

	// TagLeavesQty - Amount of shares open for further execution. If the OrdStatus (39) is Canceled, DoneForTheDay,
	// Expired, Calculated, or Rejected (in which case the order is no longer active) then LeavesQty (151) could be 0,
	// otherwise LeavesQty (151) = OrderQty (38) - CumQty (14).
	TagLeavesQty

	// TagCashOrderQty - Specifies the approximate order quantity desired in total monetary units vs. as a number of
	// shares. The broker would be responsible for converting and calculating a share quantity (OrderQty (38)) based
	// upon this amount to be used for the actual order and subsequent messages.
	TagCashOrderQty

	// TagAllocAvgPx - AvgPx (6) for a specific AllocAccount (79)
	TagAllocAvgPx

	// TagAllocNetMoney - NetMoney (118) for a specific AllocAccount (79)
	TagAllocNetMoney

	// TagSettlCurrFxRate - Foreign exchange rate used to compute SettlCurrAmt (119) from Currency (15) to SettlCurrency
	// (120)
	TagSettlCurrFxRate

	// TagSettlCurrFxRateCalc - Specifies whether SettlCurrFxRate (155) should be multiplied or divided.
	TagSettlCurrFxRateCalc

	// TagNumDaysInterest - Number of Days of Interest for convertible bonds and fixed income
	TagNumDaysInterest

	// TagAccruedInterestRate - Accrued Interest Rate for convertible bonds and fixed income
	TagAccruedInterestRate

	// TagAccruedInterestAmt - Amount of Accrued Interest for convertible bonds and fixed income
	TagAccruedInterestAmt

	// TagSettlInstMode - Indicates mode used for Settlement Instructions <35=T>
	TagSettlInstMode

	// TagAllocText - Free format text related to a specific AllocAccount (79).
	TagAllocText

	// TagSettlInstID - Unique identifier for Settlement Instructions <35=T> message.
	TagSettlInstID

	// TagSettlInstTransType - Settlement Instructions <35=T> message transaction type
	TagSettlInstTransType

	// TagEmailThreadID - Unique identifier for an email thread (new and chain of replies)
	TagEmailThreadID

	// TagSettlInstSource - Indicates source of Settlement Instructions <35=T>
	TagSettlInstSource

	// TagSettlLocation - Identifies Settlement Depository or Country (421) Code (ISITC spec)
	TagSettlLocation

	// TagSecurityType - Indicates type of security (ISITC spec)
	TagSecurityType

	// TagEffectiveTime - Time the details within the message should take effect (always expressed in UTC (Universal
	// Time Coordinated, also known as 'GMT')
	TagEffectiveTime

	// TagStandInstDbType - Identifies the Standing Instruction database used
	TagStandInstDbType

	// TagStandInstDbName - Name of the Standing Instruction database represented with StandInstDbType (169) (i.e. the
	// Global Custodian's name).
	TagStandInstDbName

	// TagStandInstDbID - Unique identifier used on the Standing Instructions database for the Standing Instructions to
	// be referenced.
	TagStandInstDbID

	// TagSettlDeliveryType - Identifies type of settlement
	TagSettlDeliveryType

	// TagSettlDepositoryCode - Broker's account code at the depository (i.e. CEDEL ID for CEDEL, FINS for DTC, or
	// Euroclear ID for Euroclear) if SettlLocation (166) is a depository
	TagSettlDepositoryCode

	// TagSettlBrkrCode - BIC (Bank Identification Code-Swift managed) code of the broker involved (i.e. for
	// multi-company brokerage firms)
	TagSettlBrkrCode

	// TagSettlInstCode - BIC (Bank Identification Code-Swift managed) code of the institution involved (i.e. for
	// multi-company institution firms)
	TagSettlInstCode

	// TagSecuritySettlAgentName - Name of SettlInstSource (165)'s local agent bank if SettlLocation (166) is not a
	// depository
	TagSecuritySettlAgentName

	// TagSecuritySettlAgentCode - BIC (Bank Identification Code-Swift managed) code of the SettlInstSource (165)'s
	// local agent bank if SettlLocation (166) is not a depository
	TagSecuritySettlAgentCode

	// TagSecuritySettlAgentAcctNum - SettlInstSource (165)'s account number at local agent bank if SettlLocation (166)
	// is not a depository
	TagSecuritySettlAgentAcctNum

	// TagSecuritySettlAgentAcctName - Name of SettlInstSource (165)'s account at local agent bank if SettlLocation
	// (166) is not a depository
	TagSecuritySettlAgentAcctName

	// TagSecuritySettlAgentContactName - Name of contact at local agent bank for SettlInstSource (165)'s account if
	// SettlLocation (166) is not a depository
	TagSecuritySettlAgentContactName

	// TagSecuritySettlAgentContactPhone - Phone number for contact at local agent bank if SettlLocation (166) is not a
	// depository
	TagSecuritySettlAgentContactPhone

	// TagCashSettlAgentName - Name of SettlInstSource (165)'s local agent bank if SettlDeliveryType (172)=Free
	TagCashSettlAgentName

	// TagCashSettlAgentCode - BIC (Bank Identification Code-Swift managed) code of the SettlInstSource (165)'s local
	// agent bank if SettlDeliveryType (172)=Free
	TagCashSettlAgentCode

	// TagCashSettlAgentAcctNum - SettlInstSource (165)'s account number at local agent bank if SettlDeliveryType
	// (172)=Free
	TagCashSettlAgentAcctNum

	// TagCashSettlAgentAcctName - Name of SettlInstSource (165)'s account at local agent bank if SettlDeliveryType
	// (172)=Free
	TagCashSettlAgentAcctName

	// TagCashSettlAgentContactName - Name of contact at local agent bank for SettlInstSource (165)'s account if
	// SettlDeliveryType (172)=Free
	TagCashSettlAgentContactName

	// TagCashSettlAgentContactPhone - Phone number for contact at local agent bank for SettlInstSource (165)'s account
	// if SettlDeliveryType (172)=Free
	TagCashSettlAgentContactPhone

	// TagBidSpotRate - Bid F/X spot rate.
	TagBidSpotRate

	// TagBidForwardPoints - Bid F/X forward points added to spot rate. May be a negative value.
	TagBidForwardPoints

	// TagOfferSpotRate - Offer F/X spot rate.
	TagOfferSpotRate

	// TagOfferForwardPoints - Offer F/X forward points added to spot rate. May be a negative value.
	TagOfferForwardPoints

	// TagOrderQty2 - OrderQty (38) of the future part of a F/X swap order.
	TagOrderQty2

	// TagFutSettDate2 - FutSettDate (64) of the future part of a F/X swap order.
	TagFutSettDate2

	// TagLastSpotRate - F/X spot rate.
	TagLastSpotRate

	// TagLastForwardPoints - F/X forward points added to LastSpotRate (194). May be a negative value.
	TagLastForwardPoints

	// TagAllocLinkID - Can be used to link two different Allocation <35=J> messages (each with unique AllocID (70))
	// together, i.e. for F/X 'Netting' or 'Swaps'. Should be unique.
	TagAllocLinkID

	// TagAllocLinkType - Identifies the type of Allocation <35=J> linkage when AllocLinkID (196) is used.
	TagAllocLinkType

	// TagSecondaryOrderID - Assigned by the party which accepts the order. Can be used to provide the OrderID (37) used
	// by an exchange or executing system.
	TagSecondaryOrderID

	// TagNoIOIQualifiers - Number of repeating groups of IOIQualifiers.
	TagNoIOIQualifiers

	// TagMaturityMonthYear - Month and Year of the maturity for SecurityType (167)=FUT or SecurityType (167)=OPT.
	// Required if MaturityDay (205) is specified.
	TagMaturityMonthYear

	// TagPutOrCall - Indicates whether an Option is for a put or call.
	TagPutOrCall

	// TagStrikePrice - Strike Price for an Option.
	TagStrikePrice

	// TagCoveredOrUncovered - Used for options
	TagCoveredOrUncovered

	// TagCustomerOrFirm - Used for options when delivering the order to an execution system/exchange to specify if the
	// order is for a customer or the firm placing the order itself.
	TagCustomerOrFirm

	// TagMaturityDay - Day of month used in conjunction with MaturityMonthYear (200) to specify the maturity date for
	// SecurityType (167)=FUT or SecurityType (167)=OPT.
	TagMaturityDay

	// TagOptAttribute - Can be used for SecurityType (167)=OPT to identify a particular security.
	TagOptAttribute

	// TagSecurityExchange - Market used to help identify a security.
	TagSecurityExchange

	// TagNotifyBrokerOfCredit - Indicates whether details should be communicated to BrokerOfCredit (92) (i.e. step-in
	// broker).
	TagNotifyBrokerOfCredit

	// TagAllocHandlInst - Indicates how the receiver (i.e. third party) of Allocation <35=J> message should
	// handle/process the account details.
	TagAllocHandlInst

	// TagMaxShow - Maximum number of shares within an order to be shown to other customers (i.e. sent via an IOI).
	TagMaxShow

	// TagPegDifference - Amount (signed) added to the price of the peg for a pegged order.
	TagPegDifference

	// TagXmlDataLen - Length of the XmlData (213) data block.
	TagXmlDataLen

	// TagXmlData - Actual XML data stream (e.g. FIXML). See approriate XML reference (e.g. FIXML). Note: may contain
	// embedded SOH characters.
	TagXmlData

	// TagSettlInstRefID - Reference identifier for the SettlInstID (162) with Cancel and Replace SettlInstTransType
	// (163) transaction types.
	TagSettlInstRefID

	// TagNoRoutingIDs - Number of repeating groups of RoutingID (217) and RoutingType (216) values.
	TagNoRoutingIDs

	// TagRoutingType - Indicates the type of RoutingID (217) specified.
	TagRoutingType

	// TagRoutingID - Assigned value used to identify a specific routing destination.
	TagRoutingID

	// TagSpreadToBenchmark - For Fixed Income. Basis points relative to a benchmark. To be expressed as "count of basis
	// points" (vs. an absolute value). E.g. High Grade Corporate Bonds may express price as basis points relative to
	// benchmark (the Benchmark (219) field). Note: Basis points can be negative.
	TagSpreadToBenchmark

	// TagBenchmark - For Fixed Income. Identifies the benchmark (e.g. used in conjunction with the SpreadToBenchmark
	// (218) field).
	TagBenchmark

	// TagCouponRate - For Fixed Income. Coupon rate of the bond. Will be zero for step-up bonds.
	TagCouponRate = iota + 4

	// TagContractMultiplier - Specifies the ratio or multiply factor to convert from contracts to shares (e.g. 1.0,
	// 100, 1000, etc.). Applicable For Fixed Income, Convertible Bonds, Derivatives, etc. Note: If used, quantities
	// should be expressed in the "nominal" (e.g. contracts vs. shares) amount.
	TagContractMultiplier = iota + 11

	// TagMDReqID - Unique identifier for Market Data Request <35=V>
	TagMDReqID = iota + 41

	// TagSubscriptionRequestType - Subscription Request Type
	TagSubscriptionRequestType

	// TagMarketDepth - Depth of market for Book Snapshot
	TagMarketDepth

	// TagMDUpdateType - Specifies the type of Market Data update.
	TagMDUpdateType

	// TagAggregatedBook - Specifies whether book entries should be aggregated.
	TagAggregatedBook

	// TagNoMDEntryTypes - Number of MDEntryType (269) fields requested.
	TagNoMDEntryTypes

	// TagNoMDEntries - Number of entries in Market Data message.
	TagNoMDEntries

	// TagMDEntryType - Type Market Data entry.
	TagMDEntryType

	// TagMDEntryPx - Price of the Market Data Entry.
	TagMDEntryPx

	// TagMDEntrySize - Number of shares represented by the Market Data Entry.
	TagMDEntrySize

	// TagMDEntryDate - Date of Market Data Entry.
	TagMDEntryDate

	// TagMDEntryTime - Time of Market Data Entry.
	TagMDEntryTime

	// TagTickDirection - Direction of the "tick".
	TagTickDirection

	// TagMDMkt - Market posting quote / trade.
	TagMDMkt

	// TagQuoteCondition - Space-delimited list of conditions describing a quote.
	TagQuoteCondition

	// TagTradeCondition - Space-delimited list of conditions describing a trade
	TagTradeCondition

	// TagMDEntryID - Unique Market Data Entry identifier.
	TagMDEntryID

	// TagMDUpdateAction - Type of Market Data update action.
	TagMDUpdateAction

	// TagMDEntryRefID - Refers to a previous MDEntryID (278).
	TagMDEntryRefID

	// TagMDReqRejReason - Reason for the rejection of aMarket Data Request <35=V>.
	TagMDReqRejReason

	// TagMDEntryOriginator - Originator of a Market Data Entry
	TagMDEntryOriginator

	// TagLocationID - Identification of a Market Maker's location
	TagLocationID

	// TagDeskID - Identification of a Market Maker's desk
	TagDeskID

	// TagDeleteReason - Reason for deletion.
	TagDeleteReason

	// TagOpenCloseSettleFlag - Flag that identifies a price.
	TagOpenCloseSettleFlag

	// TagSellerDays - Specifies the number of days that may elapse before delivery of the security
	TagSellerDays

	// TagMDEntryBuyer - Buying party in a trade
	TagMDEntryBuyer

	// TagMDEntrySeller - Selling party in a trade
	TagMDEntrySeller

	// TagMDEntryPositionNo - Display position of a bid or offer, numbered from most competitive to least competitive,
	// per market side, beginning with 1.
	TagMDEntryPositionNo

	// TagFinancialStatus - Identifies a firm's financial status.
	TagFinancialStatus

	// TagCorporateAction - Identifies the type of Corporate Action.
	TagCorporateAction

	// TagDefBidSize - Default Bid Size.
	TagDefBidSize

	// TagDefOfferSize - Default Offer Size.
	TagDefOfferSize

	// TagNoQuoteEntries - The number of quote entries for a QuoteSet.
	TagNoQuoteEntries

	// TagNoQuoteSets - The number of sets of quotes in the message.
	TagNoQuoteSets

	// TagQuoteAckStatus - Identifies the status of the quote acknowledgement.
	TagQuoteAckStatus

	// TagQuoteCancelType - Identifies the type of quote cancel.
	TagQuoteCancelType

	// TagQuoteEntryID - Uniquely identifies the quote as part of a QuoteSet.
	TagQuoteEntryID

	// TagQuoteRejectReason - Reason Quote was rejected.
	TagQuoteRejectReason

	// TagQuoteResponseLevel - Level of Response requested from receiver of quote messages.
	TagQuoteResponseLevel

	// TagQuoteSetID - Unique id for the Quote Set.
	TagQuoteSetID

	// TagQuoteRequestType - Indicates the type of Quote Request <35=R> being generated
	TagQuoteRequestType

	// TagTotQuoteEntries - Total number of quotes for the quote set across all messages. Should be the sum of all
	// NoQuoteEntries (295) in each message that has repeating quotes that are part of the same quote set.
	TagTotQuoteEntries

	// TagUnderlyingIDSource - Underlying security's IDSource.
	TagUnderlyingIDSource

	// TagUnderlyingIssuer - Underlying security's Issuer.
	TagUnderlyingIssuer

	// TagUnderlyingSecurityDesc - Underlying security's SecurityDesc.
	TagUnderlyingSecurityDesc

	// TagUnderlyingSecurityExchange - Underlying security's SecurityExchange. Can be used to identify the underlying
	// security.
	TagUnderlyingSecurityExchange

	// TagUnderlyingSecurityID - Underlying security's SecurityID.
	TagUnderlyingSecurityID

	// TagUnderlyingSecurityType - Underlying security's SecurityType.
	TagUnderlyingSecurityType

	// TagUnderlyingSymbol - Underlying security's Symbol.
	TagUnderlyingSymbol

	// TagUnderlyingSymbolSfx - Underlying security's SymbolSfx.
	TagUnderlyingSymbolSfx

	// TagUnderlyingMaturityMonthYear - Underlying security's MaturityMonthYear. Required if UnderlyingMaturityDay (314)
	// is specified.
	TagUnderlyingMaturityMonthYear

	// TagUnderlyingMaturityDay - Underlying security's MaturityDay.
	TagUnderlyingMaturityDay

	// TagUnderlyingPutOrCall - Underlying security's PutOrCall.
	TagUnderlyingPutOrCall

	// TagUnderlyingStrikePrice - Underlying security's StrikePrice.
	TagUnderlyingStrikePrice

	// TagUnderlyingOptAttribute - Underlying security's OptAttribute.
	TagUnderlyingOptAttribute

	// TagUnderlyingCurrency - Underlying security's Currency.
	TagUnderlyingCurrency

	// TagRatioQty - Quantity of a particular leg in the security.
	TagRatioQty

	// TagSecurityReqID - Unique ID of a Security Definition Request <35=c>.
	TagSecurityReqID

	// TagSecurityRequestType - Type of Security Definition Request <35=c>.
	TagSecurityRequestType

	// TagSecurityResponseID - Unique ID of a Security Definition <35=d> message.
	TagSecurityResponseID

	// TagSecurityResponseType - Type of Security Definition <35=d> message response.
	TagSecurityResponseType

	// TagSecurityStatusReqID - Unique ID of a Security Status Request <35=e> message.
	TagSecurityStatusReqID

	// TagUnsolicitedIndicator - Indicates whether message is being sent as a result of a subscription request or not.
	TagUnsolicitedIndicator

	// TagSecurityTradingStatus - Identifies the trading status applicable to the transaction.
	TagSecurityTradingStatus

	// TagHaltReason - Denotes the reason for the Opening Delay or Trading Halt.
	TagHaltReason

	// TagInViewOfCommon - Indicates whether the halt was due to Common Stock trading being halted.
	TagInViewOfCommon

	// TagDueToRelated - Indicates whether the halt was due to the Related Security being halted.
	TagDueToRelated

	// TagBuyVolume - Number of shares bought.
	TagBuyVolume

	// TagSellVolume - Number of shares sold.
	TagSellVolume

	// TagHighPx - Represents an indication of the high end of the price range for a security prior to the open or
	// reopen
	TagHighPx

	// TagLowPx - Represents an indication of the low end of the price range for a security prior to the open or reopen
	TagLowPx

	// TagAdjustment - Identifies the type of adjustment.
	TagAdjustment

	// TagTradSesReqID - Unique ID of a Trading Session Status <35=h> message.
	TagTradSesReqID

	// TagTradingSessionID - Identifier for Trading Session
	//
	// Can be used to represent a specific market trading session (e.g. 'PRE-OPEN", "CROSS_2", "AFTER-HOURS",
	// "TOSTNET1", "TOSTNET2", etc.).
	//
	// Values should be bi-laterally agreed to between counterparties.
	TagTradingSessionID

	// TagContraTrader - Identifies the trader (e.g. "badge number") of the ContraBroker (375).
	TagContraTrader

	// TagTradSesMethod - Method of trading
	TagTradSesMethod

	// TagTradSesMode - Trading Session Mode
	TagTradSesMode

	// TagTradSesStatus - State of the trading session.
	TagTradSesStatus

	// TagTradSesStartTime - Starting time of the trading session
	TagTradSesStartTime

	// TagTradSesOpenTime - Time of the opening of the trading session
	TagTradSesOpenTime

	// TagTradSesPreCloseTime - Time of the pre-closed of the trading session
	TagTradSesPreCloseTime

	// TagTradSesCloseTime - Closing time of the trading session
	TagTradSesCloseTime

	// TagTradSesEndTime - End time of the trading session
	TagTradSesEndTime

	// TagNumberOfOrders - Number of orders in the market.
	TagNumberOfOrders

	// TagMessageEncoding - Type of message encoding (non-ASCII (non-English) characters) used in a message's 'Encoded'
	// fields.
	TagMessageEncoding

	// TagEncodedIssuerLen - Byte length of encoded (non-ASCII characters) EncodedIssuer (349) field.
	TagEncodedIssuerLen

	// TagEncodedIssuer - Encoded (non-ASCII characters) representation of the Issuer (106) field in the encoded format
	// specified via the MessageEncoding (347) field. If used, the ASCII (English) representation should also be
	// specified in the Issuer (106) field.
	TagEncodedIssuer

	// TagEncodedSecurityDescLen - Byte length of encoded (non-ASCII characters) EncodedSecurityDesc (351) field.
	TagEncodedSecurityDescLen

	// TagEncodedSecurityDesc - Encoded (non-ASCII characters) representation of the SecurityDesc (107) field in the
	// encoded format specified via the MessageEncoding (347) field. If used, the ASCII (English) representation should
	// also be specified in the SecurityDesc (107) field.
	TagEncodedSecurityDesc

	// TagEncodedListExecInstLen - Byte length of encoded (non-ASCII characters) EncodedListExecInst (353) field.
	TagEncodedListExecInstLen

	// TagEncodedListExecInst - Encoded (non-ASCII characters) representation of the ListExecInst (69) field in the
	// encoded format specified via the MessageEncoding (347) field. If used, the ASCII (English) representation should
	// also be specified in the ListExecInst (69) field.
	TagEncodedListExecInst

	// TagEncodedTextLen - Byte length of encoded (non-ASCII characters) EncodedText (355) field.
	TagEncodedTextLen

	// TagEncodedText - Encoded (non-ASCII characters) representation of the Text (58) field in the encoded format
	// specified via the MessageEncoding (347) field. If used, the ASCII (English) representation should also be
	// specified in the Text (58) field.
	TagEncodedText

	// TagEncodedSubjectLen - Byte length of encoded (non-ASCII characters) EncodedSubject (357) field.
	TagEncodedSubjectLen

	// TagEncodedSubject - Encoded (non-ASCII characters) representation of the Subject (147) field in the encoded
	// format specified via the MessageEncoding (347) field. If used, the ASCII (English) representation should also be
	// specified in the Subject (147) field.
	TagEncodedSubject

	// TagEncodedHeadlineLen - Byte length of encoded (non-ASCII characters) EncodedHeadline (359) field.
	TagEncodedHeadlineLen

	// TagEncodedHeadline - Encoded (non-ASCII characters) representation of the Headline (148) field in the encoded
	// format specified via the MessageEncoding (347) field. If used, the ASCII (English) representation should also be
	// specified in the Headline (148) field.
	TagEncodedHeadline

	// TagEncodedAllocTextLen - Byte length of encoded (non-ASCII characters) EncodedAllocText (361) field.
	TagEncodedAllocTextLen

	// TagEncodedAllocText - Encoded (non-ASCII characters) representation of the AllocText (161) field in the encoded
	// format specified via the MessageEncoding (347) field. If used, the ASCII (English) representation should also be
	// specified in the AllocText (161) field.
	TagEncodedAllocText

	// TagEncodedUnderlyingIssuerLen - Byte length of encoded (non-ASCII characters) EncodedUnderlyingIssuer (363)
	// field.
	TagEncodedUnderlyingIssuerLen

	// TagEncodedUnderlyingIssuer - Encoded (non-ASCII characters) representation of the UnderlyingIssuer (306) field in
	// the encoded format specified via the MessageEncoding (347) field. If used, the ASCII (English) representation
	// should also be specified in the UnderlyingIssuer (306) field.
	TagEncodedUnderlyingIssuer

	// TagEncodedUnderlyingSecurityDescLen - Byte length of encoded (non-ASCII characters) EncodedUnderlyingSecurityDesc
	// (365) field.
	TagEncodedUnderlyingSecurityDescLen

	// TagEncodedUnderlyingSecurityDesc - Encoded (non-ASCII characters) representation of the UnderlyingSecurityDesc
	// (307) field in the encoded format specified via the MessageEncoding (347) field. If used, the ASCII (English)
	// representation should also be specified in the UnderlyingSecurityeDesc field.
	TagEncodedUnderlyingSecurityDesc

	// TagAllocPrice - Executed price for an AllocAccount (79) entry used when using 'executed price' vs. 'average
	// price' allocations (e.g. Japan).
	TagAllocPrice

	// TagQuoteSetValidUntilTime - Indicates expiration time of this particular QuoteSet (always expressed in UTC
	// (Universal Time Coordinated, also known as 'GMT')
	TagQuoteSetValidUntilTime

	// TagQuoteEntryRejectReason - Reason Quote Entry was rejected: 1--9
	TagQuoteEntryRejectReason

	// TagLastMsgSeqNumProcessed - The last MsgSeqNum (34) value received and processed. Can be specified on every
	// message sent. Useful for detecting a backlog with a counterparty.
	TagLastMsgSeqNumProcessed

	// TagOnBehalfOfSendingTime - Used when a message is sent via a 'hub' or 'service bureau'. If A sends to Q (the hub)
	// who then sends to B via a separate FIX session, then when Q sends to B the value of this field should represent
	// the SendingTime (52) on the message A sent to Q. (always expressed in UTC (Universal Time Coordinated, also known
	// as 'GMT')
	TagOnBehalfOfSendingTime

	// TagRefTagID - The tag number of the FIX field being referenced.
	TagRefTagID

	// TagRefMsgType - The MsgType (35) of the FIX message being referenced.
	TagRefMsgType

	// TagSessionRejectReason - Code to identify reason for a session-level Reject <35=3> message.
	TagSessionRejectReason

	// TagBidRequestTransType - Identifies the Bid Request <35=k> message type.
	TagBidRequestTransType

	// TagContraBroker - Identifies contra broker. Standard NASD market-maker mnemonic is preferred.
	TagContraBroker

	// TagComplianceID - ID used to represent this transaction for compliance purposes (e.g. OATS reporting).
	TagComplianceID

	// TagSolicitedFlag - Indicates whether the order was solicited.
	TagSolicitedFlag

	// TagExecRestatementReason - Code to identify reason for an ExecutionRpt message sent with ExecType (150)=Restated
	// or used when communicating an unsolicited cancel.
	TagExecRestatementReason

	// TagBusinessRejectRefID - The value of the business-level 'ID' field on the message being referenced.
	TagBusinessRejectRefID

	// TagBusinessRejectReason - Code to identify reason for a Business Message Reject <35=j> message.
	TagBusinessRejectReason

	// TagGrossTradeAmt - Total amount traded (e.g. CumQty (14) * AvgPx (6)) expressed in units of currency.
	TagGrossTradeAmt

	// TagNoContraBrokers - The number of ContraBroker (375) entries.
	TagNoContraBrokers

	// TagMaxMessageSize - Maximum number of bytes supported for a single message.
	TagMaxMessageSize

	// TagNoMsgTypes - Number of MsgType (35) in repeating group.
	TagNoMsgTypes

	// TagMsgDirection - Specifies the direction of the message.
	TagMsgDirection

	// TagNoTradingSessions - Number of TradingSessionID (336) in repeating group.
	TagNoTradingSessions

	// TagTotalVolumeTraded - Total volume (quantity) traded.
	TagTotalVolumeTraded

	// TagDiscretionInst - Code to identify the price a DiscretionOffset (389) is related to and should be
	// mathematically added to.
	TagDiscretionInst

	// TagDiscretionOffset - Amount (signed) added to the 'related to' price specified via DiscretionInst (388).
	TagDiscretionOffset

	// TagBidID - Unique identifier for Bid Response <35=l> as assigned by broker. Uniqueness must be guaranteed within
	// a single trading day.
	TagBidID

	// TagClientBidID - Unique identifier for a Bid Request <35=k> as assigned by institution. Uniqueness must be
	// guaranteed within a single trading day.
	TagClientBidID

	// TagListName - Descriptive name for list order.
	TagListName

	// TagTotalNumSecurities - Total number of securities.
	TagTotalNumSecurities

	// TagBidType - Code to identify the type of Bid Request <35=k>.
	TagBidType

	// TagNumTickets - Total number of tickets.
	TagNumTickets

	// TagSideValue1 - Amounts in currency
	TagSideValue1

	// TagSideValue2 - Amounts in currency
	TagSideValue2

	// TagNoBidDescriptors - Number of BidDescriptor (400) entries.
	TagNoBidDescriptors

	// TagBidDescriptorType - Code to identify the type of BidDescriptor (400).
	TagBidDescriptorType

	// TagBidDescriptor - BidDescriptor value. Usage depends upon BidDescriptorType (399).
	//
	// If BidDescriptorType = 1
	//
	// Industrials etc. - Free text
	//
	// If BidDescriptorType = 2
	//
	// "FR" etc. - ISO Country (421) Codes
	//
	// If BidDescriptorType = 3
	//
	// FT100, FT250, STOX - Free text
	TagBidDescriptor

	// TagSideValueInd - Code to identify which "SideValue" the value refers to. SideValue1 (396) and SideValue2 (397)
	// are used as opposed to Buy or Sell so that the basket can be quoted either way as Buy or Sell.
	TagSideValueInd

	// TagLiquidityPctLow - Liquidity indicator or lower limit if TotalNumSecurities (393) > 1. Represented as a
	// percentage.
	TagLiquidityPctLow

	// TagLiquidityPctHigh - Upper liquidity indicator if TotalNumSecurities (393) > 1. Represented as a percentage.
	TagLiquidityPctHigh

	// TagLiquidityValue - Value between LiquidityPctLow (402) and LiquidityPctHigh (403) in Currency (15)
	TagLiquidityValue

	// TagEFPTrackingError - Eg Used in EFP trades 12% (EFP - Exchange for Physical ). Represented as a percentage.
	TagEFPTrackingError

	// TagFairValue - Used in EFP trades
	TagFairValue

	// TagOutsideIndexPct - Used in EFP trades. Represented as a percentage.
	TagOutsideIndexPct

	// TagValueOfFutures - Used in EFP trades
	TagValueOfFutures

	// TagLiquidityIndType - Code to identify the type of liquidity indicator.
	TagLiquidityIndType

	// TagWtAverageLiquidity - Overall weighted average liquidity expressed as a % of average daily volume. Represented
	// as a percentage.
	TagWtAverageLiquidity

	// TagExchangeForPhysical - Indicates whether to exchange for physical.
	TagExchangeForPhysical

	// TagOutMainCntryUIndex - Value of stocks in Currency (15)
	TagOutMainCntryUIndex

	// TagCrossPercent - Percentage of program that crosses in Currency (15). Represented as a percentage.
	TagCrossPercent

	// TagProgRptReqs - Code to identify the desired frequency of progress reports.
	TagProgRptReqs

	// TagProgPeriodInterval - Time in minutes between each List Status <35=N> report sent by SellSide. Zero means don't
	// send status.
	TagProgPeriodInterval

	// TagIncTaxInd - Code to represent whether value is net (inclusive of tax) or gross.
	TagIncTaxInd

	// TagNumBidders - Indicates the total number of bidders on the list
	TagNumBidders

	// TagTradeType - Code to represent the type of trade.
	TagTradeType

	// TagBasisPxType - Code to represent the basis price type.
	TagBasisPxType

	// TagNoBidComponents - Indicates the number of list entries.
	TagNoBidComponents

	// TagCountry - ISO Country Code in field.
	TagCountry

	// TagTotNoStrikes - Total number of strike price entries across all messages. Should be the sum of all NoStrikes
	// (428) in each message that has repeating strike price entries related to the same ListID (66). Used to support
	// fragmentation.
	TagTotNoStrikes

	// TagPriceType - Code to represent the price type.
	TagPriceType

	// TagDayOrderQty - For GT orders, the OrderQty (38) less all shares (adjusted for stock splits) that traded on
	// previous days. DayOrderQty (424) = OrderQty (38) - (CumQty (14) - DayCumQty (425))
	TagDayOrderQty

	// TagDayCumQty - The number of shares on a GT order that have traded today.
	TagDayCumQty

	// TagDayAvgPx - The average price of shares on a GT order that have traded today.
	TagDayAvgPx

	// TagGTBookingInst - Code to identify whether to book out executions on a part-filled GT order on the day of
	// execution or to accumulate.
	TagGTBookingInst

	// TagNoStrikes - Number of list strike price entries.
	TagNoStrikes

	// TagListStatusType - Code to represent the price type.
	TagListStatusType

	// TagNetGrossInd - Code to represent whether value is net (inclusive of tax) or gross.
	TagNetGrossInd

	// TagListOrderStatus - Code to represent the status of a list order.
	TagListOrderStatus

	// TagExpireDate - Date of order expiration (last day the order can trade), always expressed in terms of the local
	// market date. The time at which the order expires is determined by the local market's business practices
	TagExpireDate

	// TagListExecInstType - Identifies the type of ListExecInst (69).
	TagListExecInstType

	// TagCxlRejResponseTo - Identifies the type of request that a Order Cancel Reject <35=9> is in response to.
	TagCxlRejResponseTo

	// TagUnderlyingCouponRate - Underlying security's CouponRate.
	TagUnderlyingCouponRate

	// TagUnderlyingContractMultiplier - Underlying security's ContractMultiplier.
	TagUnderlyingContractMultiplier

	// TagContraTradeQty - Quantity traded with the ContraBroker (375).
	TagContraTradeQty

	// TagContraTradeTime - Identifes the time of the trade with the ContraBroker (375). (always expressed in UTC
	// (Universal Time Coordinated, also known as 'GMT')
	TagContraTradeTime

	// TagClearingFirm - Firm that will clear the trade. Used if different from the executing firm.
	TagClearingFirm

	// TagClearingAccount - Supplemental accounting information forward to clearing house/firm.
	TagClearingAccount

	// TagLiquidityNumSecurities - Number of Securites between LiquidityPctLow (402) and LiquidityPctHigh (403) in
	// Currency (15).
	TagLiquidityNumSecurities

	// TagMultiLegReportingType - Used to indicate what an Execution Report <35=8> represents (e.g. used with multi-leg
	// securities, such as option strategies, spreads, etc.).
	TagMultiLegReportingType

	// TagStrikeTime - The time at which current market prices are used to determine the value of a basket.
	TagStrikeTime

	// TagListStatusText - Free format text string related to List Status <35=N>.
	TagListStatusText

	// TagEncodedListStatusTextLen - Byte length of encoded (non-ASCII characters) EncodedListStatusText (446) field.
	TagEncodedListStatusTextLen

	// TagEncodedListStatusText - Encoded (non-ASCII characters) representation of the ListStatusText (444) field in the
	// encoded format specified via the MessageEncoding (347) field. If used, the ASCII (English) representation should
	// also be specified in the ListStatusText (444) field.
	TagEncodedListStatusText
)
