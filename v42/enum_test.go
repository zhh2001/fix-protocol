package fields

import "testing"

// TestSideEnumValues verifies that Side enum values comply with the official FIX specification
func TestSideEnumValues(t *testing.T) {
	// Define test cases: enum name → actual value → expected value
	testCases := []struct {
		name     string
		actual   string
		expected string
	}{
		{"BeginStringFIX42", BeginStringFIX42, "FIX.4.2"},

		{"AdvSideBuy", AdvSideBuy, "B"},
		{"AdvSideSell", AdvSideSell, "S"},
		{"AdvSideTrade", AdvSideTrade, "T"},
		{"AdvSideCross", AdvSideCross, "X"},

		{"AdvTransTypeCancel", AdvTransTypeCancel, "C"},
		{"AdvTransTypeNew", AdvTransTypeNew, "N"},
		{"AdvTransTypeReplace", AdvTransTypeReplace, "R"},

		{"CommTypePerShare", CommTypePerShare, "1"},
		{"CommTypePercentage", CommTypePercentage, "2"},
		{"CommTypeAbsolute", CommTypeAbsolute, "3"},

		{"ExecInstStayOnOfferside", ExecInstStayOnOfferside, "0"},
		{"ExecInstNotHeld", ExecInstNotHeld, "1"},
		{"ExecInstWork", ExecInstWork, "2"},
		{"ExecInstGoAlong", ExecInstGoAlong, "3"},
		{"ExecInstOverTheDay", ExecInstOverTheDay, "4"},
		{"ExecInstHeld", ExecInstHeld, "5"},
		{"ExecInstParticipateDontInitiate", ExecInstParticipateDontInitiate, "6"},
		{"ExecInstStrictScale", ExecInstStrictScale, "7"},
		{"ExecInstTryToScale", ExecInstTryToScale, "8"},
		{"ExecInstStayOnBidside", ExecInstStayOnBidside, "9"},
		{"ExecInstNoCross", ExecInstNoCross, "A"},
		{"ExecInstOKToCross", ExecInstOKToCross, "B"},
		{"ExecInstCallFirst", ExecInstCallFirst, "C"},
		{"ExecInstPercentOfVolume", ExecInstPercentOfVolume, "D"},
		{"ExecInstDNI", ExecInstDNI, "E"},
		{"ExecInstDNR", ExecInstDNR, "F"},
		{"ExecInstAON", ExecInstAON, "G"},
		{"ExecInstInstitutionsOnly", ExecInstInstitutionsOnly, "I"},
		{"ExecInstLastPeg", ExecInstLastPeg, "L"},
		{"ExecInstMidPricePeg", ExecInstMidPricePeg, "M"},
		{"ExecInstNonNegotiable", ExecInstNonNegotiable, "N"},
		{"ExecInstOpeningPeg", ExecInstOpeningPeg, "O"},
		{"ExecInstMarketPeg", ExecInstMarketPeg, "P"},
		{"ExecInstPrimaryPeg", ExecInstPrimaryPeg, "R"},
		{"ExecInstSuspend", ExecInstSuspend, "S"},
		{"ExecInstFixedPeg", ExecInstFixedPeg, "T"},
		{"ExecInstCustomerDisplayInstruction", ExecInstCustomerDisplayInstruction, "U"},
		{"ExecInstNetting", ExecInstNetting, "V"},
		{"ExecInstPegToVWAP", ExecInstPegToVWAP, "W"},

		{"ExecTransTypeNew", ExecTransTypeNew, "0"},
		{"ExecTransTypeCancel", ExecTransTypeCancel, "1"},
		{"ExecTransTypeCorrect", ExecTransTypeCorrect, "2"},
		{"ExecTransTypeStatus", ExecTransTypeStatus, "3"},

		{"HandlInstAutoExecOrderPrivate", HandlInstAutoExecOrderPrivate, "1"},
		{"HandlInstAutoExecOrderPublic", HandlInstAutoExecOrderPublic, "2"},
		{"HandlInstManualOrder", HandlInstManualOrder, "3"},

		{"IDSourceCUSIP", IDSourceCUSIP, "1"},
		{"IDSourceSEDOL", IDSourceSEDOL, "2"},
		{"IDSourceQUIK", IDSourceQUIK, "3"},
		{"IDSourceISIN", IDSourceISIN, "4"},
		{"IDSourceRIC", IDSourceRIC, "5"},
		{"IDSourceISOCurrency", IDSourceISOCurrency, "6"},
		{"IDSourceISOCountry", IDSourceISOCountry, "7"},
		{"IDSourceExchange", IDSourceExchange, "8"},
		{"IDSourceCTA", IDSourceCTA, "9"},

		{"IOIQltyIndHigh", IOIQltyIndHigh, "H"},
		{"IOIQltyIndLow", IOIQltyIndLow, "L"},
		{"IOIQltyIndMedium", IOIQltyIndMedium, "M"},

		{"IOISharesLarge", IOISharesLarge, "L"},
		{"IOISharesMedium", IOISharesMedium, "M"},
		{"IOISharesSmall", IOISharesSmall, "S"},

		{"IOITransTypeCancel", IOITransTypeCancel, "C"},
		{"IOITransTypeNew", IOITransTypeNew, "N"},
		{"IOITransTypeReplace", IOITransTypeReplace, "R"},

		{"LastCapacityAgent", LastCapacityAgent, "1"},
		{"LastCapacityCrossAsAgent", LastCapacityCrossAsAgent, "2"},
		{"LastCapacityCrossAsPrincipal", LastCapacityCrossAsPrincipal, "3"},
		{"LastCapacityPrincipal", LastCapacityPrincipal, "4"},

		{"MsgTypeAdvertisement", MsgTypeAdvertisement, "7"},
		{"MsgTypeAllocation", MsgTypeAllocation, "J"},
		{"MsgTypeAllocationAck", MsgTypeAllocationAck, "P"},
		{"MsgTypeBidRequest", MsgTypeBidRequest, "k"},
		{"MsgTypeBidResponse", MsgTypeBidResponse, "l"},
		{"MsgTypeBusinessMessageReject", MsgTypeBusinessMessageReject, "j"},
		{"MsgTypeDontKnowTrade", MsgTypeDontKnowTrade, "Q"},
		{"MsgTypeEmail", MsgTypeEmail, "C"},
		{"MsgTypeExecutionReport", MsgTypeExecutionReport, "8"},
		{"MsgTypeHeartbeat", MsgTypeHeartbeat, "0"},
		{"MsgTypeIndicationOfInterest", MsgTypeIndicationOfInterest, "6"},
		{"MsgTypeListCancelRequest", MsgTypeListCancelRequest, "K"},
		{"MsgTypeListExecute", MsgTypeListExecute, "L"},
		{"MsgTypeListStatus", MsgTypeListStatus, "N"},
		{"MsgTypeListStatusRequest", MsgTypeListStatusRequest, "M"},
		{"MsgTypeListStrikePrice", MsgTypeListStrikePrice, "m"},
		{"MsgTypeLogon", MsgTypeLogon, "A"},
		{"MsgTypeLogout", MsgTypeLogout, "5"},
		{"MsgTypeMarketDataIncrementalRefresh", MsgTypeMarketDataIncrementalRefresh, "X"},
		{"MsgTypeMarketDataRequest", MsgTypeMarketDataRequest, "V"},
		{"MsgTypeMarketDataRequestReject", MsgTypeMarketDataRequestReject, "Y"},
		{"MsgTypeMarketDataSnapshotFullRefresh", MsgTypeMarketDataSnapshotFullRefresh, "W"},
		{"MsgTypeMassQuote", MsgTypeMassQuote, "i"},
		{"MsgTypeNewOrderList", MsgTypeNewOrderList, "E"},
		{"MsgTypeNews", MsgTypeNews, "B"},
		{"MsgTypeOrderCancelReject", MsgTypeOrderCancelReject, "9"},
		{"MsgTypeOrderCancelRequest", MsgTypeOrderCancelRequest, "F"},
		{"MsgTypeOrderCancelOrReplaceRequest", MsgTypeOrderCancelOrReplaceRequest, "G"},
		{"MsgTypeOrderSingle", MsgTypeOrderSingle, "D"},
		{"MsgTypeOrderStatusRequest", MsgTypeOrderStatusRequest, "H"},
		{"MsgTypeQuote", MsgTypeQuote, "S"},
		{"MsgTypeQuoteAcknowledgement", MsgTypeQuoteAcknowledgement, "b"},
		{"MsgTypeQuoteCancel", MsgTypeQuoteCancel, "Z"},
		{"MsgTypeQuoteRequest", MsgTypeQuoteRequest, "R"},
		{"MsgTypeQuoteStatusRequest", MsgTypeQuoteStatusRequest, "a"},
		{"MsgTypeReject", MsgTypeReject, "3"},
		{"MsgTypeResendRequest", MsgTypeResendRequest, "2"},
		{"MsgTypeSecurityDefinition", MsgTypeSecurityDefinition, "d"},
		{"MsgTypeSecurityDefinitionRequest", MsgTypeSecurityDefinitionRequest, "c"},
		{"MsgTypeSecurityStatus", MsgTypeSecurityStatus, "f"},
		{"MsgTypeSecurityStatusRequest", MsgTypeSecurityStatusRequest, "e"},
		{"MsgTypeSequenceReset", MsgTypeSequenceReset, "4"},
		{"MsgTypeSettlementInstructions", MsgTypeSettlementInstructions, "T"},
		{"MsgTypeTestRequest", MsgTypeTestRequest, "1"},
		{"MsgTypeTradingSessionStatus", MsgTypeTradingSessionStatus, "h"},
		{"MsgTypeTradingSessionStatusRequest", MsgTypeTradingSessionStatusRequest, "g"},

		{"OrdStatusNew", OrdStatusNew, "0"},
		{"OrdStatusPartiallyFilled", OrdStatusPartiallyFilled, "1"},
		{"OrdStatusFilled", OrdStatusFilled, "2"},
		{"OrdStatusDoneForDay", OrdStatusDoneForDay, "3"},
		{"OrdStatusCanceled", OrdStatusCanceled, "4"},
		{"OrdStatusReplaced", OrdStatusReplaced, "5"},
		{"OrdStatusPendingCancel", OrdStatusPendingCancel, "6"},
		{"OrdStatusStopped", OrdStatusStopped, "7"},
		{"OrdStatusRejected", OrdStatusRejected, "8"},
		{"OrdStatusSuspended", OrdStatusSuspended, "9"},
		{"OrdStatusPendingNew", OrdStatusPendingNew, "A"},
		{"OrdStatusCalculated", OrdStatusCalculated, "B"},
		{"OrdStatusExpired", OrdStatusExpired, "C"},
		{"OrdStatusAcceptedForBidding", OrdStatusAcceptedForBidding, "D"},
		{"OrdStatusPendingReplace", OrdStatusPendingReplace, "E"},

		{"SideBuy", SideBuy, "1"},
		{"SideSell", SideSell, "2"},
		{"SideBuyMinus", SideBuyMinus, "3"},
		{"SideSellPlus", SideSellPlus, "4"},
		{"SideSellShort", SideSellShort, "5"},
		{"SideSellShortExempt", SideSellShortExempt, "6"},
		{"SideUndisclosed", SideUndisclosed, "7"},
		{"SideCross", SideCross, "8"},
		{"SideCrossShort", SideCrossShort, "9"},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf("Enum %s has an incorrect value: actual %q, expected %q", tc.name, tc.actual, tc.expected)
			}
		})
	}
}
