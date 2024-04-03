/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件字段结构
type ComponentMetadataValueValid struct {
	Required        *bool                           `json:"required,omitempty"`
	Type_           ComponentMetadataFieldValidType `json:"type,omitempty"`
	FileFormat      *string                         `json:"file_format,omitempty"`
	FileSizeKbLimit *int64                          `json:"file_size_kb_limit,omitempty"`
	Height          *int64                          `json:"height,omitempty"`
	Width           *int64                          `json:"width,omitempty"`
	MaxLength       *int64                          `json:"max_length,omitempty"`
	MinLength       *int64                          `json:"min_length,omitempty"`
	Pattern         *string                         `json:"pattern,omitempty"`
	MinEmojiOccurs  *int64                          `json:"min_emoji_occurs,omitempty"`
	MaxEmojiOccurs  *int64                          `json:"max_emoji_occurs,omitempty"`
	RatioHeight     *int64                          `json:"ratio_height,omitempty"`
	RatioWidth      *int64                          `json:"ratio_width,omitempty"`
	MinDuration     *int64                          `json:"min_duration,omitempty"`
	MaxDuration     *int64                          `json:"max_duration,omitempty"`
	MinHeight       *int64                          `json:"min_height,omitempty"`
	MinWidth        *int64                          `json:"min_width,omitempty"`
}