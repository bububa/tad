/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 链路模板节点结构
type ConversionLinkNodeStruct struct {
	ConversionLinkNodeId     *int64                       `json:"conversion_link_node_id,omitempty"`
	ConversionLinkNodeName   *string                      `json:"conversion_link_node_name,omitempty"`
	ConversionLinkNodeIndex  *int64                       `json:"conversion_link_node_index,omitempty"`
	ConversionLinkActionType ActionType                   `json:"conversion_link_action_type,omitempty"`
	CarrierId                *int64                       `json:"carrier_id,omitempty"`
	CarrierName              *string                      `json:"carrier_name,omitempty"`
	DataSource               ConversionLinkNodeDataSource `json:"data_source,omitempty"`
	OgId                     *[]int64                     `json:"og_id,omitempty"`
	OgList                   *[]OgStruct                  `json:"og_list,omitempty"`
	RoiOgId                  *[]int64                     `json:"roi_og_id,omitempty"`
	RoiOgList                *[]RoiOgStruct               `json:"roi_og_list,omitempty"`
	ConversionInfo           *[]ConversionInfoStruct      `json:"conversion_info,omitempty"`
}