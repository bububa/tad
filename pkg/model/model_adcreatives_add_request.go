/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdcreativesAddRequest struct {
	CampaignId                    *int64                        `json:"campaign_id,omitempty"`
	AdcreativeName                *string                       `json:"adcreative_name,omitempty"`
	PromotedObjectType            PromotedObjectType            `json:"promoted_object_type,omitempty"`
	PageType                      DestinationType               `json:"page_type,omitempty"`
	AutomaticSiteEnabled          *bool                         `json:"automatic_site_enabled,omitempty"`
	SiteSet                       *[]string                     `json:"site_set,omitempty"`
	PageSpec                      *PageSpec                     `json:"page_spec,omitempty"`
	LinkPageType                  LinkPageType                  `json:"link_page_type,omitempty"`
	LinkNameType                  LinkNameTypeMp                `json:"link_name_type,omitempty"`
	LinkPageSpec                  *LinkPageSpec                 `json:"link_page_spec,omitempty"`
	ConversionDataType            ConversionDataType            `json:"conversion_data_type,omitempty"`
	ConversionTargetType          ConversionTargetType          `json:"conversion_target_type,omitempty"`
	QqMiniGameTrackingQueryString *string                       `json:"qq_mini_game_tracking_query_string,omitempty"`
	DeepLinkUrl                   *string                       `json:"deep_link_url,omitempty"`
	AndroidDeepLinkAppId          *string                       `json:"android_deep_link_app_id,omitempty"`
	IosDeepLinkAppId              *string                       `json:"ios_deep_link_app_id,omitempty"`
	UniversalLinkUrl              *string                       `json:"universal_link_url,omitempty"`
	PromotedObjectId              *string                       `json:"promoted_object_id,omitempty"`
	ProfileId                     *int64                        `json:"profile_id,omitempty"`
	ShareContentSpec              *ShareContentSpec             `json:"share_content_spec,omitempty"`
	DynamicAdcreativeSpec         *DynamicAdcreativeSpec        `json:"dynamic_adcreative_spec,omitempty"`
	ComponentId                   *int64                        `json:"component_id,omitempty"`
	Category                      *[]int64                      `json:"category,omitempty"`
	Label                         *[]string                     `json:"label,omitempty"`
	UnionMarketSwitch             *bool                         `json:"union_market_switch,omitempty"`
	PlayablePageMaterialId        *string                       `json:"playable_page_material_id,omitempty"`
	VideoEndPage                  *VideoEndPageSpec             `json:"video_end_page,omitempty"`
	FeedsVideoCommentSwitch       *bool                         `json:"feeds_video_comment_switch,omitempty"`
	WebviewUrl                    *string                       `json:"webview_url,omitempty"`
	SimpleCanvasSubType           SimpleCanvasSubType           `json:"simple_canvas_sub_type,omitempty"`
	FloatingZone                  *FloatingZone                 `json:"floating_zone,omitempty"`
	MarketingPendantImageId       *string                       `json:"marketing_pendant_image_id,omitempty"`
	CountdownSwitch               *bool                         `json:"countdown_switch,omitempty"`
	PageTrackUrl                  *string                       `json:"page_track_url,omitempty"`
	AccountId                     *int64                        `json:"account_id,omitempty"`
	AdcreativeTemplateId          *int64                        `json:"adcreative_template_id,omitempty"`
	AdcreativeElements            *AdcreativeCreativeElementsMp `json:"adcreative_elements,omitempty"`
}
