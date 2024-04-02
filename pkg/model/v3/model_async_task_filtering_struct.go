/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 过滤条件
type AsyncTaskFilteringStruct struct {
	Field    *string        `json:"field,omitempty"`
	Operator FilterOperator `json:"operator,omitempty"`
	Values   *[]string      `json:"values,omitempty"`
}
