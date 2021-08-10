/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 拉取创意结构
type AdCreativeGetSpec struct {
	AdcreativeId       *int64                          `json:"adcreative_id,omitempty"`
	AdcreativeName     *string                         `json:"adcreative_name,omitempty"`
	AdcreativeElements *DpAdcreativeCreativeElementsMp `json:"adcreative_elements,omitempty"`
	PageType           DestinationType                 `json:"page_type,omitempty"`
	PageSpec           *DpPageSpec                     `json:"page_spec,omitempty"`
	ProfileId          *int64                          `json:"profile_id,omitempty"`
}
