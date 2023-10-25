package vk

import "github.com/hootuu/tome/vn/value"

// BUSINESS KEYS
const (
	GMV         value.Key = "GMV"         // GrossMerchandise Volume
	Users       value.Key = "USERS"       // VN OR SP USER Count
	ActiveUsers value.Key = "AU"          // Active User Count
	Orders      value.Key = "ORDERS"      // Total Order Count
	OrderUsers  value.Key = "ORDER_USERS" // Total Order User Count
	PCT         value.Key = "PCT"         // Per Customer Transaction

	DailyGMV         value.Key = "D_GMV"
	DailyUsers       value.Key = "D_USERS"
	DailyActiveUsers value.Key = "D_AU"
	DailyOrders      value.Key = "D_ORDERS"
	DailyOrderUsers  value.Key = "D_ORDER_USERS"
	DailyPCT         value.Key = "D_PCT"

	WeeklyGMV         value.Key = "W_GMV"
	WeeklyUsers       value.Key = "W_USERS"
	WeeklyActiveUsers value.Key = "W_AU"
	WeeklyOrders      value.Key = "W_ORDERS"
	WeeklyOrderUsers  value.Key = "W_ORDER_USERS"
	WeeklyPCT         value.Key = "W_PCT"

	MonthlyGMV         value.Key = "M_GMV"
	MonthlyUsers       value.Key = "M_USERS"
	MonthlyActiveUsers value.Key = "M_AU"
	MonthlyOrders      value.Key = "M_ORDERS"
	MonthlyOrderUsers  value.Key = "M_ORDER_USERS"
	MonthlyPCT         value.Key = "M_PCT"

	QuarterlyGMV         value.Key = "Q_GMV"
	QuarterlyUsers       value.Key = "Q_USERS"
	QuarterlyActiveUsers value.Key = "Q_AU"
	QuarterlyOrders      value.Key = "Q_ORDERS"
	QuarterlyOrderUsers  value.Key = "Q_ORDER_USERS"
	QuarterlyPCT         value.Key = "Q_PCT"

	YearlyGMV         value.Key = "Y_GMV"
	YearlyUsers       value.Key = "Y_USERS"
	YearlyActiveUsers value.Key = "Y_AU"
	YearlyOrders      value.Key = "Y_ORDERS"
	YearlyOrderUsers  value.Key = "Y_ORDER_USERS"
	YearlyPCT         value.Key = "Y_PCT"
)
