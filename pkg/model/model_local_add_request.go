/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LocalAddRequest struct {
	Adgroup    *AdgroupCreateSpec    `json:"adgroup,omitempty"`
	Campaign   *CampaignCreateSpec   `json:"campaign,omitempty"`
	Adcreative *AdCreativeCreateSpec `json:"adcreative,omitempty"`
}
