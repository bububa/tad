/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 基础图片组件元素<br>        图片大小: 不超过300KB<br>        图片尺寸: 宽度750, 高度不超过1536
type ImageSpec struct {
	ImageId       *string `json:"image_id,omitempty"`
	Width         *int64  `json:"width,omitempty"`
	Height        *int64  `json:"height,omitempty"`
	PaddingTop    *int64  `json:"padding_top,omitempty"`
	PaddingBottom *int64  `json:"padding_bottom,omitempty"`
}