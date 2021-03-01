/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 自定义裁剪后缩放信息
type CropCustomizedSpecAndResize struct {
	CropWidth    *int64 `json:"crop_width,omitempty"`
	CropHeight   *int64 `json:"crop_height,omitempty"`
	AxisX        *int64 `json:"axis_x,omitempty"`
	AxisY        *int64 `json:"axis_y,omitempty"`
	ResizeWidth  *int64 `json:"resize_width,omitempty"`
	ResizeHeight *int64 `json:"resize_height,omitempty"`
}
