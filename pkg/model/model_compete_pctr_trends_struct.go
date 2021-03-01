/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 点击率趋势变化
type CompetePctrTrendsStruct struct {
	Score          *int64         `json:"score,omitempty"`
	SelfAvg        *[]PointStruct `json:"self_avg,omitempty"`
	SelfChoseAvg   *[]PointStruct `json:"self_chose_avg,omitempty"`
	WinnerChoseAvg *[]PointStruct `json:"winner_chose_avg,omitempty"`
	Conclusion     *string        `json:"conclusion,omitempty"`
}
