/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 跳转落地页信息
type LandingPageStructure struct {
	PageSpec            *PageSpec           `json:"page_spec,omitempty"`
	PageType            DestinationType     `json:"page_type,omitempty"`
	LinkNameType        LinkUrlLinkNameType `json:"link_name_type,omitempty"`
	LandingPagePlatform LandingPagePlatform `json:"landing_page_platform,omitempty"`
	LandingPageInfo     *LandingPageInfo    `json:"landing_page_info,omitempty"`
}