/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 素材信息
type MaterialStruct struct {
	Type_        MaterialTypeEnum `json:"type,omitempty"`
	MediaId      *string          `json:"media_id,omitempty"`
	MaterialName *string          `json:"material_name,omitempty"`
}
