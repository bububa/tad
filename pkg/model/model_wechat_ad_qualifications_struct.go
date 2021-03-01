/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 微信广告资质信息
type WechatAdQualificationsStruct struct {
	QualificationId     *int64              `json:"qualification_id,omitempty"`
	QualificationName   *string             `json:"qualification_name,omitempty"`
	ImageUrl            *string             `json:"image_url,omitempty"`
	ExpiredDate         *string             `json:"expired_date,omitempty"`
	QualificationStatus QualificationStatus `json:"qualification_status,omitempty"`
	RejectMessage       *string             `json:"reject_message,omitempty"`
}
