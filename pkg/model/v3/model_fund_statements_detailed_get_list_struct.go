/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type FundStatementsDetailedGetListStruct struct {
	AccountId      *int64         `json:"account_id,omitempty"`
	FundType       AccountTypeMap `json:"fund_type,omitempty"`
	Balance        *int64         `json:"balance,omitempty"`
	Time           *int64         `json:"time,omitempty"`
	ExternalBillNo *string        `json:"external_bill_no,omitempty"`
	TradeTypeExt   TradeTypeExt   `json:"trade_type_ext,omitempty"`
	Amount         *int64         `json:"amount,omitempty"`
	Description    *string        `json:"description,omitempty"`
}