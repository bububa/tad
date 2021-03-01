/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AgencyPeerTransferAddResponseData struct {
	Amount         *int64         `json:"amount,omitempty"`
	ExternalBillNo *string        `json:"external_bill_no,omitempty"`
	IsRepeated     Boolean        `json:"is_repeated,omitempty"`
	FundType       AccountTypeMap `json:"fund_type,omitempty"`
	Time           *int64         `json:"time,omitempty"`
}
