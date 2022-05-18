/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProgrammedTemplateGetResponseData struct {
	TemplateInfos *[]TemplateInfoStruct `json:"template_infos,omitempty"`
	PageInfo      *PageInfoStruct       `json:"page_info,omitempty"`
}
