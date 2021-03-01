/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 小程序落地页，mini_program_id和mini_program_path要同时填写
type DpMiniProgramSpec struct {
	MiniProgramId   *string `json:"mini_program_id,omitempty"`
	MiniProgramPath *string `json:"mini_program_path,omitempty"`
}
