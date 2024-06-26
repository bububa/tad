/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告渠道包任务处理信息
type ChannelPackageInfo struct {
	AndroidAppId     *int64            `json:"android_app_id,omitempty"`
	PackageName      *string           `json:"package_name,omitempty"`
	Status           ChannelTaskStatus `json:"status,omitempty"`
	ErrorCode        ChannelTaskError  `json:"error_code,omitempty"`
	CreatedTime      *int64            `json:"created_time,omitempty"`
	LastModifiedTime *int64            `json:"last_modified_time,omitempty"`
	ChannelPackageId *string           `json:"channel_package_id,omitempty"`
}
