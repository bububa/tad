/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告组小时报表查询条件
type TaskTypeWechatAdgroupHourlyReportSpec struct {
	Date      *string    `json:"date,omitempty"`
	HourRange *HourRange `json:"hour_range,omitempty"`
}
