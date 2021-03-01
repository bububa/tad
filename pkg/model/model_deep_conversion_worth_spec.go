/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// oCPC/oCPM 优化 ROI 配置
type DeepConversionWorthSpec struct {
	Goal        DeepConversionWorthGoal `json:"goal,omitempty"`
	ExpectedRoi *float64                `json:"expected_roi,omitempty"`
}
