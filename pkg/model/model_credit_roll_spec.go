/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 固定信用金更多信息
type CreditRollSpec struct {
	LimitAmount *int64 `json:"limit_amount,omitempty"`
	UsedAmount  *int64 `json:"used_amount,omitempty"`
	UsableAmout *int64 `json:"usable_amout,omitempty"`
}
