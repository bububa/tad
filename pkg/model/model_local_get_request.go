/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LocalGetRequest struct {
	AccountId int64              `json:"account_id,omitempty"`
	DateRange *DateRange         `json:"date_range,omitempty"`
	Filtering *[]FilteringStruct `json:"filtering,omitempty"`
	Level     DpApiReportLevel   `json:"level,omitempty"`
	Page      int64              `json:"page,omitempty"`
	PageSize  int64              `json:"page_size,omitempty"`
}
