/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type ProductCatalogsReportsGetListStruct struct {
	Date                      *string  `json:"date,omitempty"`
	ProductCatalogId          *int64   `json:"product_catalog_id,omitempty"`
	AdgroupId                 *int64   `json:"adgroup_id,omitempty"`
	ProductOuterId            *string  `json:"product_outer_id,omitempty"`
	FirstCategoryId           *int64   `json:"first_category_id,omitempty"`
	SecondCategoryId          *int64   `json:"second_category_id,omitempty"`
	ThirdCategoryId           *int64   `json:"third_category_id,omitempty"`
	ViewCount                 *int64   `json:"view_count,omitempty"`
	ValidClickCount           *int64   `json:"valid_click_count,omitempty"`
	Ctr                       *float64 `json:"ctr,omitempty"`
	Cpc                       *int64   `json:"cpc,omitempty"`
	ThousandDisplayPrice      *int64   `json:"thousand_display_price,omitempty"`
	Cost                      *int64   `json:"cost,omitempty"`
	ActivatedCount            *int64   `json:"activated_count,omitempty"`
	ActivatedCost             *int64   `json:"activated_cost,omitempty"`
	ActivatedRate             *float64 `json:"activated_rate,omitempty"`
	WebKeyPageViewCost        *int64   `json:"web_key_page_view_cost,omitempty"`
	WebCommodityPageViewCount *int64   `json:"web_commodity_page_view_count,omitempty"`
	WebCommodityPageViewCost  *int64   `json:"web_commodity_page_view_cost,omitempty"`
	WebRegisterCount          *int64   `json:"web_register_count,omitempty"`
	PagePhoneCallDirectCount  *int64   `json:"page_phone_call_direct_count,omitempty"`
	OwnPageNavigationCount    *int64   `json:"own_page_navigation_count,omitempty"`
	OwnPageNavigationCost     *int64   `json:"own_page_navigation_cost,omitempty"`
	WebApplicationCount       *int64   `json:"web_application_count,omitempty"`
	WebApplicationCost        *int64   `json:"web_application_cost,omitempty"`
	WebOrderCount             *int64   `json:"web_order_count,omitempty"`
	WebOrderRate              *float64 `json:"web_order_rate,omitempty"`
	AppOrderRate              *float64 `json:"app_order_rate,omitempty"`
	WebOrderCost              *int64   `json:"web_order_cost,omitempty"`
	WebCheckoutAmount         *int64   `json:"web_checkout_amount,omitempty"`
	WebCheckoutCount          *int64   `json:"web_checkout_count,omitempty"`
	WebCheckoutCost           *int64   `json:"web_checkout_cost,omitempty"`
	DownloadRate              *float64 `json:"download_rate,omitempty"`
	DownloadCost              *int64   `json:"download_cost,omitempty"`
	InstallCost               *int64   `json:"install_cost,omitempty"`
	ClickActivatedRate        *float64 `json:"click_activated_rate,omitempty"`
	RetentionCount            *int64   `json:"retention_count,omitempty"`
	RetentionRate             *float64 `json:"retention_rate,omitempty"`
	RetentionCost             *int64   `json:"retention_cost,omitempty"`
	AppKeyPageViewCount       *int64   `json:"app_key_page_view_count,omitempty"`
	WebKeyPageViewCount       *int64   `json:"web_key_page_view_count,omitempty"`
	AppCommodityPageViewCount *int64   `json:"app_commodity_page_view_count,omitempty"`
	AppCommodityPageViewRate  *float64 `json:"app_commodity_page_view_rate,omitempty"`
	WebCommodityPageViewRate  *float64 `json:"web_commodity_page_view_rate,omitempty"`
	AppCommodityPageViewCost  *int64   `json:"app_commodity_page_view_cost,omitempty"`
	AppRegisterCount          *int64   `json:"app_register_count,omitempty"`
	AppRegisterCost           *int64   `json:"app_register_cost,omitempty"`
	AppApplicationCount       *int64   `json:"app_application_count,omitempty"`
	AppApplicationCost        *int64   `json:"app_application_cost,omitempty"`
	AppOrderCount             *int64   `json:"app_order_count,omitempty"`
	AppOrderCost              *int64   `json:"app_order_cost,omitempty"`
	FollowCost                *int64   `json:"follow_cost,omitempty"`
	ForwardCost               *int64   `json:"forward_cost,omitempty"`
	ReadCost                  *int64   `json:"read_cost,omitempty"`
	PraiseCount               *int64   `json:"praise_count,omitempty"`
	PraiseCost                *int64   `json:"praise_cost,omitempty"`
	CommentCount              *int64   `json:"comment_count,omitempty"`
	LikeOrComment             *int64   `json:"like_or_comment,omitempty"`
	CommentCost               *int64   `json:"comment_cost,omitempty"`
	AppCheckoutRate           *float64 `json:"app_checkout_rate,omitempty"`
	AppRegisterRate           *float64 `json:"app_register_rate,omitempty"`
	Impression                *int64   `json:"impression,omitempty"`
	Click                     *int64   `json:"click,omitempty"`
	Download                  *int64   `json:"download,omitempty"`
	Follow                    *int64   `json:"follow,omitempty"`
	Activation                *int64   `json:"activation,omitempty"`
	Share                     *int64   `json:"share,omitempty"`
	Read                      *int64   `json:"read,omitempty"`
	AppPaymentCount           *int64   `json:"app_payment_count,omitempty"`
	Reservation               *int64   `json:"reservation,omitempty"`
	AppInstallation           *int64   `json:"app_installation,omitempty"`
	AppPaymentAmount          *int64   `json:"app_payment_amount,omitempty"`
	AppAddToCarCount          *int64   `json:"app_add_to_car_count,omitempty"`
	AppAddToCarCost           *int64   `json:"app_add_to_car_cost,omitempty"`
}
