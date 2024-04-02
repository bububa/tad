/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// MarketingTargetType : 推广内容资产类型
type MarketingTargetType string

// List of MarketingTargetType
const (
	MarketingTargetType_UNKNOWN                          MarketingTargetType = "MARKETING_TARGET_TYPE_UNKNOWN"
	MarketingTargetType_APP_ANDROID                      MarketingTargetType = "MARKETING_TARGET_TYPE_APP_ANDROID"
	MarketingTargetType_APP_IOS                          MarketingTargetType = "MARKETING_TARGET_TYPE_APP_IOS"
	MarketingTargetType_WECHAT_OFFICIAL_ACCOUNT          MarketingTargetType = "MARKETING_TARGET_TYPE_WECHAT_OFFICIAL_ACCOUNT"
	MarketingTargetType_PRODUCT                          MarketingTargetType = "MARKETING_TARGET_TYPE_PRODUCT"
	MarketingTargetType_TRAFFIC                          MarketingTargetType = "MARKETING_TARGET_TYPE_TRAFFIC"
	MarketingTargetType_HOUSE_PROPERTY                   MarketingTargetType = "MARKETING_TARGET_TYPE_HOUSE_PROPERTY"
	MarketingTargetType_LOCAL_STORE                      MarketingTargetType = "MARKETING_TARGET_TYPE_LOCAL_STORE"
	MarketingTargetType_WECHAT_MINI_GAME                 MarketingTargetType = "MARKETING_TARGET_TYPE_WECHAT_MINI_GAME"
	MarketingTargetType_CONSUMER_PRODUCT                 MarketingTargetType = "MARKETING_TARGET_TYPE_CONSUMER_PRODUCT"
	MarketingTargetType_WECHAT_CHANNELS                  MarketingTargetType = "MARKETING_TARGET_TYPE_WECHAT_CHANNELS"
	MarketingTargetType_WECHAT_CHANNELS_LIVE             MarketingTargetType = "MARKETING_TARGET_TYPE_WECHAT_CHANNELS_LIVE"
	MarketingTargetType_WECHAT_CHANNELS_LIVE_RESERVATION MarketingTargetType = "MARKETING_TARGET_TYPE_WECHAT_CHANNELS_LIVE_RESERVATION"
	MarketingTargetType_MINI_PROGRAM_WECHAT              MarketingTargetType = "MARKETING_TARGET_TYPE_MINI_PROGRAM_WECHAT"
	MarketingTargetType_APP_QUICK_APP                    MarketingTargetType = "MARKETING_TARGET_TYPE_APP_QUICK_APP"
	MarketingTargetType_CONSUME_MEDICAL                  MarketingTargetType = "MARKETING_TARGET_TYPE_CONSUME_MEDICAL"
	MarketingTargetType_COMPREHENSIVE_HOUSEKEEPING       MarketingTargetType = "MARKETING_TARGET_TYPE_COMPREHENSIVE_HOUSEKEEPING"
	MarketingTargetType_FICTION                          MarketingTargetType = "MARKETING_TARGET_TYPE_FICTION"
	MarketingTargetType_SHORT_DRAMA                      MarketingTargetType = "MARKETING_TARGET_TYPE_SHORT_DRAMA"
	MarketingTargetType_AUDIOVISUAL_ENTERTAINMENT        MarketingTargetType = "MARKETING_TARGET_TYPE_AUDIOVISUAL_ENTERTAINMENT"
	MarketingTargetType_BEAUTY_AND_PERSONAL_CARE         MarketingTargetType = "MARKETING_TARGET_TYPE_BEAUTY_AND_PERSONAL_CARE"
	MarketingTargetType_WEDDING_AND_PORTRAIT_PHOTOGRAPHY MarketingTargetType = "MARKETING_TARGET_TYPE_WEDDING_AND_PORTRAIT_PHOTOGRAPHY"
	MarketingTargetType_FRANCHISE_BRAND                  MarketingTargetType = "MARKETING_TARGET_TYPE_FRANCHISE_BRAND"
	MarketingTargetType_ENTERPRISE_SERVICES              MarketingTargetType = "MARKETING_TARGET_TYPE_ENTERPRISE_SERVICES"
	MarketingTargetType_EXHIBITION_BOOTH_DESIGN          MarketingTargetType = "MARKETING_TARGET_TYPE_EXHIBITION_BOOTH_DESIGN"
	MarketingTargetType_INSURANCE                        MarketingTargetType = "MARKETING_TARGET_TYPE_INSURANCE"
	MarketingTargetType_BANK                             MarketingTargetType = "MARKETING_TARGET_TYPE_BANK"
	MarketingTargetType_CREDIT                           MarketingTargetType = "MARKETING_TARGET_TYPE_CREDIT"
	MarketingTargetType_INVESTMENT_CONSULTING            MarketingTargetType = "MARKETING_TARGET_TYPE_INVESTMENT_CONSULTING"
	MarketingTargetType_REAL_ESTATE                      MarketingTargetType = "MARKETING_TARGET_TYPE_REAL_ESTATE"
	MarketingTargetType_TELECOMMUNICATIONS_OPERATOR      MarketingTargetType = "MARKETING_TARGET_TYPE_TELECOMMUNICATIONS_OPERATOR"
	MarketingTargetType_TOURIST_ATTRACTIONS_TICKETS      MarketingTargetType = "MARKETING_TARGET_TYPE_TOURIST_ATTRACTIONS_TICKETS"
	MarketingTargetType_RENOVATION_SERVICES              MarketingTargetType = "MARKETING_TARGET_TYPE_RENOVATION_SERVICES"
	MarketingTargetType_FURNITURE_AND_BUILDING_MATERIALS MarketingTargetType = "MARKETING_TARGET_TYPE_FURNITURE_AND_BUILDING_MATERIALS"
	MarketingTargetType_EXHIBITION_SALES                 MarketingTargetType = "MARKETING_TARGET_TYPE_EXHIBITION_SALES"
	MarketingTargetType_MEDICINE_INDUSTRY_COMMERCIAL     MarketingTargetType = "MARKETING_TARGET_TYPE_MEDICINE_INDUSTRY_COMMERCIAL"
	MarketingTargetType_FINANCE                          MarketingTargetType = "MARKETING_TARGET_TYPE_FINANCE"
	MarketingTargetType_LOCAL_STORE_PACKAGE              MarketingTargetType = "MARKETING_TARGET_TYPE_LOCAL_STORE_PACKAGE"
	MarketingTargetType_CATERING_AND_LEISURE             MarketingTargetType = "MARKETING_TARGET_TYPE_CATERING_AND_LEISURE"
	MarketingTargetType_CHAIN_RESTAURANT                 MarketingTargetType = "MARKETING_TARGET_TYPE_CHAIN_RESTAURANT"
	MarketingTargetType_COMMODITY_SET                    MarketingTargetType = "MARKETING_TARGET_TYPE_COMMODITY_SET"
	MarketingTargetType_TOURIST_TRAVEL_ROUTE             MarketingTargetType = "MARKETING_TARGET_TYPE_TOURIST_TRAVEL_ROUTE"
	MarketingTargetType_TOURIST_CRUISE_LINE              MarketingTargetType = "MARKETING_TARGET_TYPE_TOURIST_CRUISE_LINE"
	MarketingTargetType_TOURIST_HOTEL_SERVICE            MarketingTargetType = "MARKETING_TARGET_TYPE_TOURIST_HOTEL_SERVICE"
	MarketingTargetType_TOURIST_AIRLINE_TICKETS          MarketingTargetType = "MARKETING_TARGET_TYPE_TOURIST_AIRLINE_TICKETS"
	MarketingTargetType_LOCAL_STORE_COMBINE_WITH_PRODUCT MarketingTargetType = "MARKETING_TARGET_TYPE_LOCAL_STORE_COMBINE_WITH_PRODUCT"
	MarketingTargetType_ACTIVITY                         MarketingTargetType = "MARKETING_TARGET_TYPE_ACTIVITY"
	MarketingTargetType_STORE                            MarketingTargetType = "MARKETING_TARGET_TYPE_STORE"
	MarketingTargetType_MINI_GAME_QQ                     MarketingTargetType = "MARKETING_TARGET_TYPE_MINI_GAME_QQ"
	MarketingTargetType_APP_GAME_ANDROID                 MarketingTargetType = "MARKETING_TARGET_TYPE_APP_GAME_ANDROID"
	MarketingTargetType_APP_GAME_IOS                     MarketingTargetType = "MARKETING_TARGET_TYPE_APP_GAME_IOS"
	MarketingTargetType_PC_GAME                          MarketingTargetType = "MARKETING_TARGET_TYPE_PC_GAME"
	MarketingTargetType_WECHAT_WORK                      MarketingTargetType = "MARKETING_TARGET_TYPE_WECHAT_WORK"
	MarketingTargetType_LIVE_STREAM_ROOM                 MarketingTargetType = "MARKETING_TARGET_TYPE_LIVE_STREAM_ROOM"
	MarketingTargetType_PERSONAL_STORE                   MarketingTargetType = "MARKETING_TARGET_TYPE_PERSONAL_STORE"
	MarketingTargetType_PLATFORM_CHANNEL                 MarketingTargetType = "MARKETING_TARGET_TYPE_PLATFORM_CHANNEL"
	MarketingTargetType_TWO_WHEEL_VEHICLE                MarketingTargetType = "MARKETING_TARGET_TYPE_TWO_WHEEL_VEHICLE"
)
