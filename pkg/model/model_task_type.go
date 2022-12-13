/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// TaskType : 任务类型（部分任务待废弃）具体请参考
type TaskType string

// List of TaskType
const (
	TaskType_TASK_TYPE_REPORT_AGENCY_ADVERTISER_HOURLY                             TaskType = "TASK_TYPE_REPORT_AGENCY_ADVERTISER_HOURLY"
	TaskType_TASK_TYPE_TSA_IMEI                                                    TaskType = "TASK_TYPE_TSA_IMEI"
	TaskType_TASK_TYPE_CREATIVE_PICTURE_REPORT                                     TaskType = "TASK_TYPE_CREATIVE_PICTURE_REPORT"
	TaskType_TASK_TYPE_AD_HOURLY_REPORT                                            TaskType = "TASK_TYPE_AD_HOURLY_REPORT"
	TaskType_TASK_TYPE_ADGROUP_HOURLY_REPORT                                       TaskType = "TASK_TYPE_ADGROUP_HOURLY_REPORT"
	TaskType_TASK_TYPE_WECHAT_MOMENTS_ADGROUP_HOURLY_REPORT                        TaskType = "TASK_TYPE_WECHAT_MOMENTS_ADGROUP_HOURLY_REPORT"
	TaskType_TASK_TYPE_ADGROUP_DAILY_REPORT                                        TaskType = "TASK_TYPE_ADGROUP_DAILY_REPORT"
	TaskType_TASK_TYPE_INDUSTRY_CREATIVE_TEMPLATE_REPORT                           TaskType = "TASK_TYPE_INDUSTRY_CREATIVE_TEMPLATE_REPORT"
	TaskType_TASK_TYPE_WECHAT_ADGROUP_HOURLY_REPORT                                TaskType = "TASK_TYPE_WECHAT_ADGROUP_HOURLY_REPORT"
	TaskType_TASK_TYPE_WECHAT_ADGROUP_DAILY_REPORT                                 TaskType = "TASK_TYPE_WECHAT_ADGROUP_DAILY_REPORT"
	TaskType_TASK_TYPE_WECHAT_AD_HOURLY_REPORT                                     TaskType = "TASK_TYPE_WECHAT_AD_HOURLY_REPORT"
	TaskType_TASK_TYPE_WECHAT_AD_DAILY_REPORT                                      TaskType = "TASK_TYPE_WECHAT_AD_DAILY_REPORT"
	TaskType_TASK_TYPE_WECHAT_ADVERTISING_DATA                                     TaskType = "TASK_TYPE_WECHAT_ADVERTISING_DATA"
	TaskType_TASK_TYPE_POI_HOURLY_REPORT                                           TaskType = "TASK_TYPE_POI_HOURLY_REPORT"
	TaskType_TASK_TYPE_WECHAT_POI_HOURLY_REPORT                                    TaskType = "TASK_TYPE_WECHAT_POI_HOURLY_REPORT"
	TaskType_TASK_TYPE_ENCRYPTED_CUSTOM_IMEI                                       TaskType = "TASK_TYPE_ENCRYPTED_CUSTOM_IMEI"
	TaskType_TASK_TYPE_ENCRYPTED_TENCENT_IMEI                                      TaskType = "TASK_TYPE_ENCRYPTED_TENCENT_IMEI"
	TaskType_TASK_TYPE_CLEAR_TENCENT_IMEI                                          TaskType = "TASK_TYPE_CLEAR_TENCENT_IMEI"
	TaskType_TASK_TYPE_DEVICE_INFO                                                 TaskType = "TASK_TYPE_DEVICE_INFO"
	TaskType_TASK_TYPE_REBATE_DETAIL                                               TaskType = "TASK_TYPE_REBATE_DETAIL"
	TaskType_TASK_TYPE_DEVICE_DETAIL_CPC                                           TaskType = "TASK_TYPE_DEVICE_DETAIL_CPC"
	TaskType_TASK_TYPE_OM_ADVERTISING_INFO                                         TaskType = "TASK_TYPE_OM_ADVERTISING_INFO"
	TaskType_TASK_TYPE_CREATE_ANDROID_CHANNEL_PACKAGE                              TaskType = "TASK_TYPE_CREATE_ANDROID_CHANNEL_PACKAGE"
	TaskType_TASK_TYPE_UPDATE_ANDROID_CHANNEL_PACKAGE                              TaskType = "TASK_TYPE_UPDATE_ANDROID_CHANNEL_PACKAGE"
	TaskType_TASK_TYPE_CREATE_ANDROID_UNION_CHANNEL_PACKAGE                        TaskType = "TASK_TYPE_CREATE_ANDROID_UNION_CHANNEL_PACKAGE"
	TaskType_TASK_TYPE_UPDATE_ANDROID_UNION_CHANNEL_PACKAGE                        TaskType = "TASK_TYPE_UPDATE_ANDROID_UNION_CHANNEL_PACKAGE"
	TaskType_TASK_TYPE_UNION_POSITION_REPORT                                       TaskType = "TASK_TYPE_UNION_POSITION_REPORT"
	TaskType_TASK_TYPE_UPDATE_ANDROID_UNION_CHANNEL_PACKAGE_BY_URL                 TaskType = "TASK_TYPE_UPDATE_ANDROID_UNION_CHANNEL_PACKAGE_BY_URL"
	TaskType_TASK_TYPE_PRODUCT_DAILY_DATA                                          TaskType = "TASK_TYPE_PRODUCT_DAILY_DATA"
	TaskType_TASK_TYPE_PRODUCT_ADGROUP_DATA                                        TaskType = "TASK_TYPE_PRODUCT_ADGROUP_DATA"
	TaskType_TASK_TYPE_PRODUCT_AD_DATA                                             TaskType = "TASK_TYPE_PRODUCT_AD_DATA"
	TaskType_TASK_TYPE_DOWNLOAD_OPERATION_LOG                                      TaskType = "TASK_TYPE_DOWNLOAD_OPERATION_LOG"
	TaskType_TASK_TYPE_CAMPAIGN_DATA                                               TaskType = "TASK_TYPE_CAMPAIGN_DATA"
	TaskType_TASK_TYPE_REVIEW_ELEMENT_PREREVIEW_RESULT                             TaskType = "TASK_TYPE_REVIEW_ELEMENT_PREREVIEW_RESULT"
	TaskType_TASK_TYPE_INTEGRATED_DATA                                             TaskType = "TASK_TYPE_INTEGRATED_DATA"
	TaskType_TASK_TYPE_AD_EDITOR_DOWNLOAD                                          TaskType = "TASK_TYPE_AD_EDITOR_DOWNLOAD"
	TaskType_TASK_TYPE_SEARCH_KEYWORD_RECOMMEND_DATA                               TaskType = "TASK_TYPE_SEARCH_KEYWORD_RECOMMEND_DATA"
	TaskType_TASK_TYPE_SEARCH_BIDWORD_REPORT_DATA                                  TaskType = "TASK_TYPE_SEARCH_BIDWORD_REPORT_DATA"
	TaskType_TASK_TYPE_SEARCH_QUERYWORD_REPORT_DATA                                TaskType = "TASK_TYPE_SEARCH_QUERYWORD_REPORT_DATA"
	TaskType_TASK_TYPE_BATCH_IMPORT                                                TaskType = "TASK_TYPE_BATCH_IMPORT"
	TaskType_TASK_TYPE_WX_SEARCH_KEYWORD_DATA                                      TaskType = "TASK_TYPE_WX_SEARCH_KEYWORD_DATA"
	TaskType_TASK_TYPE_SEARCH_BIDWORD_DOWNLOAD                                     TaskType = "TASK_TYPE_SEARCH_BIDWORD_DOWNLOAD"
	TaskType_TASK_TYPE_UPDATE_ACCOUNT_BIDWORD_PRICE                                TaskType = "TASK_TYPE_UPDATE_ACCOUNT_BIDWORD_PRICE"
	TaskType_TASK_TYPE_UPDATE_ACCOUNT_BIDWORD_MATCH_TYPE                           TaskType = "TASK_TYPE_UPDATE_ACCOUNT_BIDWORD_MATCH_TYPE"
	TaskType_TASK_TYPE_BATCH_REQUEST                                               TaskType = "TASK_TYPE_BATCH_REQUEST"
	TaskType_TASK_TYPE_CLICK_FORWARDED_DATA                                        TaskType = "TASK_TYPE_CLICK_FORWARDED_DATA"
	TaskType_TASK_TYPE_ADP_OFFLINE_RPT                                             TaskType = "TASK_TYPE_ADP_OFFLINE_RPT"
	TaskType_TASK_TYPE_AGENCY_OFFLINE_RPT                                          TaskType = "TASK_TYPE_AGENCY_OFFLINE_RPT"
	TaskType_TASK_TYPE_TRANSCODE_TWENTY_EIGHT_TASK                                 TaskType = "TASK_TYPE_TRANSCODE_TWENTY_EIGHT_TASK"
	TaskType_TASK_TYPE_MULTI_ACCOUNT_INTEGRATED_DATA                               TaskType = "TASK_TYPE_MULTI_ACCOUNT_INTEGRATED_DATA"
	TaskType_TASK_TYPE_LANDING_PAGE_RPT                                            TaskType = "TASK_TYPE_LANDING_PAGE_RPT"
	TaskType_TASK_TYPE_UPDATE_UNION_POSITION_PACKAGE                               TaskType = "TASK_TYPE_UPDATE_UNION_POSITION_PACKAGE"
	TaskType_TASK_TYPE_UPDATE_EXCLUDE_UNION_POSITION_PACKAGE                       TaskType = "TASK_TYPE_UPDATE_EXCLUDE_UNION_POSITION_PACKAGE"
	TaskType_TASK_TYPE_UPDATE_TARGETING_ID                                         TaskType = "TASK_TYPE_UPDATE_TARGETING_ID"
	TaskType_TASK_TYPE_UPDATE_BID_STRATEGY                                         TaskType = "TASK_TYPE_UPDATE_BID_STRATEGY"
	TaskType_TASK_TYPE_UPDATE_DEEP_CONVERSION_BEHAVIOR_BID                         TaskType = "TASK_TYPE_UPDATE_DEEP_CONVERSION_BEHAVIOR_BID"
	TaskType_TASK_TYPE_CREATE_SCHEDULED_UPDATE_CAMPAIGN_DAILY_BUDGET               TaskType = "TASK_TYPE_CREATE_SCHEDULED_UPDATE_CAMPAIGN_DAILY_BUDGET"
	TaskType_TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADGROUP_DAILY_BUDGET                TaskType = "TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADGROUP_DAILY_BUDGET"
	TaskType_TASK_TYPE_DELETE_SCHEDULED_TASK                                       TaskType = "TASK_TYPE_DELETE_SCHEDULED_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_APP_ANDROID_CHANNEL_PACKAGE_ID               TaskType = "TASK_TYPE_UPDATE_ADGROUP_APP_ANDROID_CHANNEL_PACKAGE_ID"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_SPEED_MODE                                  TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_SPEED_MODE"
	TaskType_TASK_TYPE_DELETE_CAMPAIGN                                             TaskType = "TASK_TYPE_DELETE_CAMPAIGN"
	TaskType_TASK_TYPE_DELETE_ADGROUP                                              TaskType = "TASK_TYPE_DELETE_ADGROUP"
	TaskType_TASK_TYPE_DELETE_AD                                                   TaskType = "TASK_TYPE_DELETE_AD"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_RATE                   TaskType = "TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_RATE"
	TaskType_TASK_TYPE_UPDATE_ADCREATIVE_DEEP_LINK_URL                             TaskType = "TASK_TYPE_UPDATE_ADCREATIVE_DEEP_LINK_URL"
	TaskType_TASK_TYPE_TARGETINGS_SHARE                                            TaskType = "TASK_TYPE_TARGETINGS_SHARE"
	TaskType_TASK_TYPE_ADMASTER_CAMPAIGN_ADDSTRATEGY                               TaskType = "TASK_TYPE_ADMASTER_CAMPAIGN_ADDSTRATEGY"
	TaskType_TASK_TYPE_ADMASTER_ADGROUP_ADDSTRATEGY                                TaskType = "TASK_TYPE_ADMASTER_ADGROUP_ADDSTRATEGY"
	TaskType_TASK_TYPE_ADMASTER_ADVERTISER_ADDSTRATEGY                             TaskType = "TASK_TYPE_ADMASTER_ADVERTISER_ADDSTRATEGY"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_CONFIGURED_STATUS                           TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_CONFIGURED_STATUS"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_DAILY_BUDGET                                TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_DAILY_BUDGET"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_CONFIGURED_STATUS                            TaskType = "TASK_TYPE_UPDATE_ADGROUP_CONFIGURED_STATUS"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DAILY_BUDGET                                 TaskType = "TASK_TYPE_UPDATE_ADGROUP_DAILY_BUDGET"
	TaskType_TASK_TYPE_UPDATE_AD_CONFIGURED_STATUS                                 TaskType = "TASK_TYPE_UPDATE_AD_CONFIGURED_STATUS"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DATE                                         TaskType = "TASK_TYPE_UPDATE_ADGROUP_DATE"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_BID_AMOUNT                                   TaskType = "TASK_TYPE_UPDATE_ADGROUP_BID_AMOUNT"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_AUTO_ACQUISITION                             TaskType = "TASK_TYPE_UPDATE_ADGROUP_AUTO_ACQUISITION"
	TaskType_TASK_TYPE_UPDATE_ADCREATIVE_LANDING_PAGE                              TaskType = "TASK_TYPE_UPDATE_ADCREATIVE_LANDING_PAGE"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_BIND_RTA_POLICY                              TaskType = "TASK_TYPE_UPDATE_ADGROUP_BIND_RTA_POLICY"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DERIVE_CONF                                  TaskType = "TASK_TYPE_UPDATE_ADGROUP_DERIVE_CONF"
	TaskType_TASK_TYPE_UPDATE_ADVERTISER_DAILY_BUDGET                              TaskType = "TASK_TYPE_UPDATE_ADVERTISER_DAILY_BUDGET"
	TaskType_TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADVERTISER_DAILY_BUDGET             TaskType = "TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADVERTISER_DAILY_BUDGET"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_TIME                                         TaskType = "TASK_TYPE_UPDATE_ADGROUP_TIME"
	TaskType_TASK_TYPE_COPY_ORDER                                                  TaskType = "TASK_TYPE_COPY_ORDER"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_ADVANCED_RATE          TaskType = "TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_ADVANCED_RATE"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_TOTAL_BUDGET                                TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_TOTAL_BUDGET"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_FLOW_OPTIMIZATION_ENABLED                    TaskType = "TASK_TYPE_UPDATE_ADGROUP_FLOW_OPTIMIZATION_ENABLED"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_BIND_RTA_POLICY                             TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_BIND_RTA_POLICY"
	TaskType_TASK_TYPE_UPDATE_BIDWORD                                              TaskType = "TASK_TYPE_UPDATE_BIDWORD"
	TaskType_TASK_TYPE_DELETE_BIDWORD                                              TaskType = "TASK_TYPE_DELETE_BIDWORD"
	TaskType_TASK_TYPE_ADD_CAMPAIGN_NEGATIVE_WORD                                  TaskType = "TASK_TYPE_ADD_CAMPAIGN_NEGATIVE_WORD"
	TaskType_TASK_TYPE_ADD_ADGROUP_NEGATIVE_WORD                                   TaskType = "TASK_TYPE_ADD_ADGROUP_NEGATIVE_WORD"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_NEGATIVE_WORD                               TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_NEGATIVE_WORD"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_NEGATIVE_WORD                                TaskType = "TASK_TYPE_UPDATE_ADGROUP_NEGATIVE_WORD"
	TaskType_TASK_TYPE_SYNC_ANDROID_CHANNEL_PACKAGE_DATA                           TaskType = "TASK_TYPE_SYNC_ANDROID_CHANNEL_PACKAGE_DATA"
	TaskType_TASK_TYPE_COPY_ORDER_DELETE                                           TaskType = "TASK_TYPE_COPY_ORDER_DELETE"
	TaskType_TASK_TYPE_INVOICE_CONSUME                                             TaskType = "TASK_TYPE_INVOICE_CONSUME"
	TaskType_TASK_TYPE_AD_STATUS_CALC                                              TaskType = "TASK_TYPE_AD_STATUS_CALC"
	TaskType_TASK_TYPE_AUDIT_INVOICE_CONSUME                                       TaskType = "TASK_TYPE_AUDIT_INVOICE_CONSUME"
	TaskType_TASK_TYPE_AUTO_DERIVED_CREATIVE                                       TaskType = "TASK_TYPE_AUTO_DERIVED_CREATIVE"
	TaskType_TASK_TYPE_USER_PROJECT_CREATE_ORDER                                   TaskType = "TASK_TYPE_USER_PROJECT_CREATE_ORDER"
	TaskType_TASK_TYPE_ACCOUNT_PROJECT_CREATE_ORDER                                TaskType = "TASK_TYPE_ACCOUNT_PROJECT_CREATE_ORDER"
	TaskType_TASK_TYPE_ACCOUNT_PROJECT_CREATE_ORDER_SUB_TASK                       TaskType = "TASK_TYPE_ACCOUNT_PROJECT_CREATE_ORDER_SUB_TASK"
	TaskType_TASK_TYPE_REVIEW_PROCESS_CONSUME                                      TaskType = "TASK_TYPE_REVIEW_PROCESS_CONSUME"
	TaskType_TASK_TYPE_DC_BATCH_PROCESS_FINISHED                                   TaskType = "TASK_TYPE_DC_BATCH_PROCESS_FINISHED"
	TaskType_TASK_TYPE_SCHEDULED_UPDATE_ADVERTISER_DAILY_BUDGET                    TaskType = "TASK_TYPE_SCHEDULED_UPDATE_ADVERTISER_DAILY_BUDGET"
	TaskType_TASK_TYPE_SCHEDULED_UPDATE_CAMPAIGN_DAILY_BUDGET                      TaskType = "TASK_TYPE_SCHEDULED_UPDATE_CAMPAIGN_DAILY_BUDGET"
	TaskType_TASK_TYPE_SCHEDULED_UPDATE_ADGROUP_DAILY_BUDGET                       TaskType = "TASK_TYPE_SCHEDULED_UPDATE_ADGROUP_DAILY_BUDGET"
	TaskType_TASK_TYPE_CRON_CUSTOMER_REPORT_PUSH                                   TaskType = "TASK_TYPE_CRON_CUSTOMER_REPORT_PUSH"
	TaskType_TASK_TYPE_DELAY_CUSTOMER_REPORT_PUSH                                  TaskType = "TASK_TYPE_DELAY_CUSTOMER_REPORT_PUSH"
	TaskType_TASK_TYPE_DELAY_CUSTOMER_MESSAGE_PUSH                                 TaskType = "TASK_TYPE_DELAY_CUSTOMER_MESSAGE_PUSH"
	TaskType_TASK_TYPE_CHECK_EXPIRED                                               TaskType = "TASK_TYPE_CHECK_EXPIRED"
	TaskType_TASK_TYPE_UPDATE_UNION_POSITION_PACKAGE_SUB_TASK                      TaskType = "TASK_TYPE_UPDATE_UNION_POSITION_PACKAGE_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_EXCLUDE_UNION_POSITION_PACKAGE_SUB_TASK              TaskType = "TASK_TYPE_UPDATE_EXCLUDE_UNION_POSITION_PACKAGE_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_TARGETING_ID_SUB_TASK                                TaskType = "TASK_TYPE_UPDATE_TARGETING_ID_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_BID_STRATEGY_SUB_TASK                                TaskType = "TASK_TYPE_UPDATE_BID_STRATEGY_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_DEEP_CONVERSION_BEHAVIOR_BID_SUB_TASK                TaskType = "TASK_TYPE_UPDATE_DEEP_CONVERSION_BEHAVIOR_BID_SUB_TASK"
	TaskType_TASK_TYPE_CREATE_SCHEDULED_UPDATE_CAMPAIGN_DAILY_BUDGET_SUB_TASK      TaskType = "TASK_TYPE_CREATE_SCHEDULED_UPDATE_CAMPAIGN_DAILY_BUDGET_SUB_TASK"
	TaskType_TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADGROUP_DAILY_BUDGET_SUB_TASK       TaskType = "TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADGROUP_DAILY_BUDGET_SUB_TASK"
	TaskType_TASK_TYPE_DELETE_SCHEDULED_TASK_SUB_TASK                              TaskType = "TASK_TYPE_DELETE_SCHEDULED_TASK_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_APP_ANDROID_CHANNEL_PACKAGE_ID_SUB_TASK      TaskType = "TASK_TYPE_UPDATE_ADGROUP_APP_ANDROID_CHANNEL_PACKAGE_ID_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_SPEED_MODE_SUB_TASK                         TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_SPEED_MODE_SUB_TASK"
	TaskType_TASK_TYPE_DELETE_CAMPAIGN_SUB_TASK                                    TaskType = "TASK_TYPE_DELETE_CAMPAIGN_SUB_TASK"
	TaskType_TASK_TYPE_DELETE_ADGROUP_SUB_TASK                                     TaskType = "TASK_TYPE_DELETE_ADGROUP_SUB_TASK"
	TaskType_TASK_TYPE_DELETE_AD_SUB_TASK                                          TaskType = "TASK_TYPE_DELETE_AD_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_RATE_SUB_TASK          TaskType = "TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_RATE_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADCREATIVE_DEEP_LINK_URL_SUB_TASK                    TaskType = "TASK_TYPE_UPDATE_ADCREATIVE_DEEP_LINK_URL_SUB_TASK"
	TaskType_TASK_TYPE_TARGETINGS_SHARE_SUB_TASK                                   TaskType = "TASK_TYPE_TARGETINGS_SHARE_SUB_TASK"
	TaskType_TASK_TYPE_ADMASTER_CAMPAIGN_ADDSTRATEGY_SUB_TASK                      TaskType = "TASK_TYPE_ADMASTER_CAMPAIGN_ADDSTRATEGY_SUB_TASK"
	TaskType_TASK_TYPE_ADMASTER_ADGROUP_ADDSTRATEGY_SUB_TASK                       TaskType = "TASK_TYPE_ADMASTER_ADGROUP_ADDSTRATEGY_SUB_TASK"
	TaskType_TASK_TYPE_ADMASTER_ADVERTISER_ADDSTRATEGY_SUB_TASK                    TaskType = "TASK_TYPE_ADMASTER_ADVERTISER_ADDSTRATEGY_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_CONFIGURED_STATUS_SUB_TASK                  TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_CONFIGURED_STATUS_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_DAILY_BUDGET_SUB_TASK                       TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_DAILY_BUDGET_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_CONFIGURED_STATUS_SUB_TASK                   TaskType = "TASK_TYPE_UPDATE_ADGROUP_CONFIGURED_STATUS_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DAILY_BUDGET_SUB_TASK                        TaskType = "TASK_TYPE_UPDATE_ADGROUP_DAILY_BUDGET_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_AD_CONFIGURED_STATUS_SUB_TASK                        TaskType = "TASK_TYPE_UPDATE_AD_CONFIGURED_STATUS_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DATE_SUB_TASK                                TaskType = "TASK_TYPE_UPDATE_ADGROUP_DATE_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_BID_AMOUNT_SUB_TASK                          TaskType = "TASK_TYPE_UPDATE_ADGROUP_BID_AMOUNT_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_AUTO_ACQUISITION_SUB_TASK                    TaskType = "TASK_TYPE_UPDATE_ADGROUP_AUTO_ACQUISITION_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADCREATIVE_LANDING_PAGE_SUB_TASK                     TaskType = "TASK_TYPE_UPDATE_ADCREATIVE_LANDING_PAGE_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_BIND_RTA_POLICY_SUB_TASK                     TaskType = "TASK_TYPE_UPDATE_ADGROUP_BIND_RTA_POLICY_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DERIVE_CONF_SUB_TASK                         TaskType = "TASK_TYPE_UPDATE_ADGROUP_DERIVE_CONF_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADVERTISER_DAILY_BUDGET_SUB_TASK                     TaskType = "TASK_TYPE_UPDATE_ADVERTISER_DAILY_BUDGET_SUB_TASK"
	TaskType_TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADVERTISER_DAILY_BUDGET_SUB_TASK    TaskType = "TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADVERTISER_DAILY_BUDGET_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_TIME_SUB_TASK                                TaskType = "TASK_TYPE_UPDATE_ADGROUP_TIME_SUB_TASK"
	TaskType_TASK_TYPE_COPY_ORDER_SUB_TASK                                         TaskType = "TASK_TYPE_COPY_ORDER_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_ADVANCED_RATE_SUB_TASK TaskType = "TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_ADVANCED_RATE_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_TOTAL_BUDGET_SUB_TASK                       TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_TOTAL_BUDGET_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_FLOW_OPTIMIZATION_ENABLED_SUB_TASK           TaskType = "TASK_TYPE_UPDATE_ADGROUP_FLOW_OPTIMIZATION_ENABLED_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_BIND_RTA_POLICY_SUB_TASK                    TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_BIND_RTA_POLICY_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_BIDWORD_SUB_TASK                                     TaskType = "TASK_TYPE_UPDATE_BIDWORD_SUB_TASK"
	TaskType_TASK_TYPE_DELETE_BIDWORD_SUB_TASK                                     TaskType = "TASK_TYPE_DELETE_BIDWORD_SUB_TASK"
	TaskType_TASK_TYPE_ADD_CAMPAIGN_NEGATIVE_WORD_SUB_TASK                         TaskType = "TASK_TYPE_ADD_CAMPAIGN_NEGATIVE_WORD_SUB_TASK"
	TaskType_TASK_TYPE_ADD_ADGROUP_NEGATIVE_WORD_SUB_TASK                          TaskType = "TASK_TYPE_ADD_ADGROUP_NEGATIVE_WORD_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_CAMPAIGN_NEGATIVE_WORD_SUB_TASK                      TaskType = "TASK_TYPE_UPDATE_CAMPAIGN_NEGATIVE_WORD_SUB_TASK"
	TaskType_TASK_TYPE_UPDATE_ADGROUP_NEGATIVE_WORD_SUB_TASK                       TaskType = "TASK_TYPE_UPDATE_ADGROUP_NEGATIVE_WORD_SUB_TASK"
	TaskType_TASK_TYPE_REFRESH_CONTENT_TOKEN                                       TaskType = "TASK_TYPE_REFRESH_CONTENT_TOKEN"
	TaskType_TASK_TYPE_REFRESH_CONTENT_CONTENT                                     TaskType = "TASK_TYPE_REFRESH_CONTENT_CONTENT"
	TaskType_TASK_TYPE_WXGAME_GAUSS_CREATIVE                                       TaskType = "TASK_TYPE_WXGAME_GAUSS_CREATIVE"
	TaskType_TASK_TYPE_UPDATE_COVER_IMAGE_DATA                                     TaskType = "TASK_TYPE_UPDATE_COVER_IMAGE_DATA"
	TaskType_TASK_TYPE_PROCESS_USER_PAGE_OBJECT                                    TaskType = "TASK_TYPE_PROCESS_USER_PAGE_OBJECT"
	TaskType_TASK_TYPE_DC_BATCH_CREATE_ADCREATIVE                                  TaskType = "TASK_TYPE_DC_BATCH_CREATE_ADCREATIVE"
	TaskType_TASK_TYPE_DC_BATCH_UPDATE_ADCREATIVE                                  TaskType = "TASK_TYPE_DC_BATCH_UPDATE_ADCREATIVE"
	TaskType_TASK_TYPE_UPDATE_CREATIVE_WARNSTATUS                                  TaskType = "TASK_TYPE_UPDATE_CREATIVE_WARNSTATUS"
	TaskType_TASK_TYPE_MULTI_ACCOUNT_REPORT_DATA                                   TaskType = "TASK_TYPE_MULTI_ACCOUNT_REPORT_DATA"
	TaskType_UNKNOWN                                                               TaskType = "UNKNOWN"
)
