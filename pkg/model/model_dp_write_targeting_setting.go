/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 定向详细设置
type DpWriteTargetingSetting struct {
	Age    *[]AgeStruct `json:"age,omitempty"`
	Gender *[]string    `json:"gender,omitempty"`
}
