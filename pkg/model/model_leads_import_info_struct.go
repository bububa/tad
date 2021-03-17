/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 导入的线索信息结构
type LeadsImportInfoStruct struct {
	OuterLeadsId             *string             `json:"outer_leads_id,omitempty"`
	LeadsType                LeadCluesLeadsType  `json:"leads_type,omitempty"`
	LeadsUserId              *string             `json:"leads_user_id,omitempty"`
	LeadsUserType            LeadsUserType       `json:"leads_user_type,omitempty"`
	LeadsUserWechatAppid     *string             `json:"leads_user_wechat_appid,omitempty"`
	LeadsTel                 *string             `json:"leads_tel,omitempty"`
	LeadsQq                  *int64              `json:"leads_qq,omitempty"`
	LeadsWechat              *string             `json:"leads_wechat,omitempty"`
	LeadsName                *string             `json:"leads_name,omitempty"`
	LeadsGender              LeadCluesGenderType `json:"leads_gender,omitempty"`
	LeadsEmail               *string             `json:"leads_email,omitempty"`
	LeadsArea                *string             `json:"leads_area,omitempty"`
	Bundle                   *string             `json:"bundle,omitempty"`
	ClickId                  *string             `json:"click_id,omitempty"`
	OuterLeadsConvertType    *string             `json:"outer_leads_convert_type,omitempty"`
	OuterLeadsIneffectReason *string             `json:"outer_leads_ineffect_reason,omitempty"`
}