package fields

import "testing"

// TestTagValues verifies that FIX protocol constant values comply with the official specification
func TestTagValues(t *testing.T) {
	// Define test cases: constant name → actual value → expected value
	testCases := []struct {
		name     string
		actual   int
		expected int
	}{
		{"TagAccount", TagAccount, 1},
		{"TagAdvId", TagAdvId, 2},
		{"TagAdvRefID", TagAdvRefID, 3},
		{"TagAdvSide", TagAdvSide, 4},
		{"TagAdvTransType", TagAdvTransType, 5},
		{"TagAvgPx", TagAvgPx, 6},
		{"TagBeginSeqNo", TagBeginSeqNo, 7},
		{"TagBeginString", TagBeginString, 8},
		{"TagBodyLength", TagBodyLength, 9},
		{"TagCheckSum", TagCheckSum, 10},
		{"TagClOrdID", TagClOrdID, 11},
		{"TagCommission", TagCommission, 12},
		{"TagCommType", TagCommType, 13},
		{"TagCumQty", TagCumQty, 14},
		{"TagCurrency", TagCurrency, 15},
		{"TagEndSeqNo", TagEndSeqNo, 16},
		{"TagExecID", TagExecID, 17},
		{"TagExecInst", TagExecInst, 18},
		{"TagExecRefID", TagExecRefID, 19},
		{"TagExecTransType", TagExecTransType, 20},
		{"TagHandlInst", TagHandlInst, 21},
		{"TagIDSource", TagIDSource, 22},
		{"TagIOIid", TagIOIid, 23},
		{"TagIOIOthSvc", TagIOIOthSvc, 24},
		{"TagIOIQltyInd", TagIOIQltyInd, 25},
		{"TagIOIRefID", TagIOIRefID, 26},
		{"TagIOIShares", TagIOIShares, 27},
		{"TagIOITransType", TagIOITransType, 28},
		{"TagLastCapacity", TagLastCapacity, 29},
		{"TagLastMkt", TagLastMkt, 30},
		{"TagLastPx", TagLastPx, 31},
		{"TagLastShares", TagLastShares, 32},
		{"TagLinesOfText", TagLinesOfText, 33},
		{"TagMsgSeqNum", TagMsgSeqNum, 34},
		{"TagMsgType", TagMsgType, 35},
		{"TagNewSeqNo", TagNewSeqNo, 36},
		{"TagOrderID", TagOrderID, 37},
		{"TagOrderQty", TagOrderQty, 38},
		{"TagOrdStatus", TagOrdStatus, 39},
		{"TagOrdType", TagOrdType, 40},
		{"TagOrigClOrdID", TagOrigClOrdID, 41},
		{"TagOrigTime", TagOrigTime, 42},
		{"TagPossDupFlag", TagPossDupFlag, 43},
		{"TagPrice", TagPrice, 44},
		{"TagRefSeqNum", TagRefSeqNum, 45},
		{"TagRelatdSym", TagRelatdSym, 46},
		{"TagRule80A", TagRule80A, 47},
		{"TagSecurityID", TagSecurityID, 48},
		{"TagSenderCompID", TagSenderCompID, 49},
		{"TagSenderSubID", TagSenderSubID, 50},
		{"TagSendingDate", TagSendingDate, 51},
		{"TagSendingTime", TagSendingTime, 52},
		{"TagShares", TagShares, 53},
		{"TagSide", TagSide, 54},
		{"TagSymbol", TagSymbol, 55},
		{"TagTargetCompID", TagTargetCompID, 56},
		{"TagTargetSubID", TagTargetSubID, 57},
		{"TagText", TagText, 58},
		{"TagTimeInForce", TagTimeInForce, 59},
		{"TagTransactTime", TagTransactTime, 60},
		{"TagUrgency", TagUrgency, 61},
		{"TagValidUntilTime", TagValidUntilTime, 62},
		{"TagSettlmntTyp", TagSettlmntTyp, 63},
		{"TagFutSettDate", TagFutSettDate, 64},
		{"TagSymbolSfx", TagSymbolSfx, 65},
		{"TagListID", TagListID, 66},
		{"TagListSeqNo", TagListSeqNo, 67},
		{"TagTotNoOrders", TagTotNoOrders, 68},
		{"TagListExecInst", TagListExecInst, 69},
		{"TagAllocID", TagAllocID, 70},
		{"TagAllocTransType", TagAllocTransType, 71},
		{"TagRefAllocID", TagRefAllocID, 72},
		{"TagNoOrders", TagNoOrders, 73},
		{"TagAvgPrxPrecision", TagAvgPrxPrecision, 74},
		{"TagTradeDate", TagTradeDate, 75},
		{"TagExecBroker", TagExecBroker, 76},
		{"TagOpenClose", TagOpenClose, 77},
		{"TagNoAllocs", TagNoAllocs, 78},
		{"TagAllocAccount", TagAllocAccount, 79},
		{"TagAllocShares", TagAllocShares, 80},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf("Constant %s has an incorrect value: actual %d, expected %d", tc.name, tc.actual, tc.expected)
			}
		})
	}
}
