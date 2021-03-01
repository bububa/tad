/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type CustomTagsAddRequest struct {
	AccountId   *int64       `json:"account_id,omitempty"`
	ParentTagId *int64       `json:"parent_tag_id,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	TagCode     *string      `json:"tag_code,omitempty"`
	Platform    DataPlatform `json:"platform,omitempty"`
}
