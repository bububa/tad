/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	apierrors "github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:(?:application|text)/json)")
	xmlCheck  = regexp.MustCompile("(?i:(?:application|text)/xml)")
)

// APIClient manages communication with the Marketing API API v1.3
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	Cfg       *config.Configuration
	SDKConfig *config.SDKConfig
	common    service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AdDiagnosisApi *AdDiagnosisApiService

	AdLabelApi *AdLabelApiService

	AdParamApi *AdParamApiService

	AdcreativePreviewsApi *AdcreativePreviewsApiService

	AdcreativePreviewsQrcodeApi *AdcreativePreviewsQrcodeApiService

	AdcreativeTemplateApi *AdcreativeTemplateApiService

	AdcreativeTemplateDetailApi *AdcreativeTemplateDetailApiService

	AdcreativeTemplateListApi *AdcreativeTemplateListApiService

	AdcreativeTemplatePreviewApi *AdcreativeTemplatePreviewApiService

	AdcreativeTemplatePreviewsApi *AdcreativeTemplatePreviewsApiService

	AdcreativeTemplatesApi *AdcreativeTemplatesApiService

	AdcreativesApi *AdcreativesApiService

	AdcreativesRelatedCapabilityApi *AdcreativesRelatedCapabilityApiService

	AdgroupsApi *AdgroupsApiService

	AdsApi *AdsApiService

	AdvertiserApi *AdvertiserApiService

	AgencyInnerTransferApi *AgencyInnerTransferApiService

	AgencyRealtimeCostApi *AgencyRealtimeCostApiService

	AndroidChannelPackagesApi *AndroidChannelPackagesApiService

	AndroidUnionChannelPackagesApi *AndroidUnionChannelPackagesApiService

	AppAndroidChannelPackagesApi *AppAndroidChannelPackagesApiService

	AssetPermissionsApi *AssetPermissionsApiService

	AssetPrePermissionsApi *AssetPrePermissionsApiService

	AsyncReportFilesApi *AsyncReportFilesApiService

	AsyncReportsApi *AsyncReportsApiService

	AsyncTaskFilesApi *AsyncTaskFilesApiService

	AsyncTasksApi *AsyncTasksApiService

	AudienceGrantRelationsApi *AudienceGrantRelationsApiService

	AuthorizationApi *AuthorizationApiService

	BarrageApi *BarrageApiService

	BarrageRecommendApi *BarrageRecommendApiService

	BatchAsyncRequestSpecificationApi *BatchAsyncRequestSpecificationApiService

	BatchAsyncRequestsApi *BatchAsyncRequestsApiService

	BatchOperationApi *BatchOperationApiService

	BatchRequestsApi *BatchRequestsApiService

	BidSimulationApi *BidSimulationApiService

	BrandApi *BrandApiService

	BusinessManagerRelationsApi *BusinessManagerRelationsApiService

	BusinessMdmAccountRelationsApi *BusinessMdmAccountRelationsApiService

	CampaignsApi *CampaignsApiService

	CapabilitiesApi *CapabilitiesApiService

	ChannelsApi *ChannelsApiService

	ComplianceValidationApi *ComplianceValidationApiService

	ConversionsApi *ConversionsApiService

	CreativetoolsTextApi *CreativetoolsTextApiService

	CustomAudienceEstimationsApi *CustomAudienceEstimationsApiService

	CustomAudienceFilesApi *CustomAudienceFilesApiService

	CustomAudienceInsightsApi *CustomAudienceInsightsApiService

	CustomAudienceReportsApi *CustomAudienceReportsApiService

	CustomAudiencesApi *CustomAudiencesApiService

	CustomFeaturesApi *CustomFeaturesApiService

	CustomTagFilesApi *CustomTagFilesApiService

	CustomTagsApi *CustomTagsApiService

	DailyBalanceReportApi *DailyBalanceReportApiService

	DailyCostApi *DailyCostApiService

	DailyReportsApi *DailyReportsApiService

	DiagnosisApi *DiagnosisApiService

	DplabelAdLabelApi *DplabelAdLabelApiService

	DynamicAdImageTemplatesApi *DynamicAdImageTemplatesApiService

	DynamicAdImagesApi *DynamicAdImagesApiService

	DynamicAdTemplatesApi *DynamicAdTemplatesApiService

	DynamicAdVideoApi *DynamicAdVideoApiService

	DynamicAdVideoTemplatesApi *DynamicAdVideoTemplatesApiService

	DynamicCreativesApi *DynamicCreativesApiService

	EcommerceOrderApi *EcommerceOrderApiService

	EstimationApi *EstimationApiService

	FundStatementsDailyApi *FundStatementsDailyApiService

	FundStatementsDetailedApi *FundStatementsDetailedApiService

	FundTransferApi *FundTransferApiService

	FundsApi *FundsApiService

	HourlyReportsApi *HourlyReportsApiService

	ImageProcessingApi *ImageProcessingApiService

	ImagesApi *ImagesApiService

	LabelAudiencesApi *LabelAudiencesApiService

	LabelsApi *LabelsApiService

	LeadCluesApi *LeadCluesApiService

	LeadsFormApi *LeadsFormApiService

	LeadsFormListApi *LeadsFormListApiService

	LeadsInvalidPayApi *LeadsInvalidPayApiService

	LocalApi *LocalApiService

	LocalEndadsmanuallyApi *LocalEndadsmanuallyApiService

	LocalEstimatedamountApi *LocalEstimatedamountApiService

	LocalStoresApi *LocalStoresApiService

	LocalStoresAddressParsingResultApi *LocalStoresAddressParsingResultApiService

	LocalStoresCategoriesApi *LocalStoresCategoriesApiService

	LocalStoresSearchInfoApi *LocalStoresSearchInfoApiService

	LocalStoresWxpayMerchantsApi *LocalStoresWxpayMerchantsApiService

	OauthApi *OauthApiService

	OptimizationGoalPermissionsApi *OptimizationGoalPermissionsApiService

	OuterCluesApi *OuterCluesApiService

	OuterCluesContactApi *OuterCluesContactApiService

	PagesApi *PagesApiService

	PlayablePagesApi *PlayablePagesApiService

	ProductCatalogsApi *ProductCatalogsApiService

	ProductCatalogsReportsApi *ProductCatalogsReportsApiService

	ProductCategoriesListApi *ProductCategoriesListApiService

	ProductItemsApi *ProductItemsApiService

	ProductItemsDetailApi *ProductItemsDetailApiService

	ProductItemsVerticalsApi *ProductItemsVerticalsApiService

	ProductSeriesApi *ProductSeriesApiService

	ProductsSystemStatusApi *ProductsSystemStatusApiService

	ProfilesApi *ProfilesApiService

	PromotedObjectsApi *PromotedObjectsApiService

	PropertyFileSessionsApi *PropertyFileSessionsApiService

	PropertyFilesApi *PropertyFilesApiService

	PropertySetSchemasApi *PropertySetSchemasApiService

	PropertySetsApi *PropertySetsApiService

	QualificationsApi *QualificationsApiService

	RealtimeCostApi *RealtimeCostApiService

	ReportApi *ReportApiService

	ReviewElementPrereviewResultsApi *ReviewElementPrereviewResultsApiService

	SceneSpecTagsApi *SceneSpecTagsApiService

	ShopApi *ShopApiService

	SplitTestsApi *SplitTestsApiService

	SubcustomerTransferApi *SubcustomerTransferApiService

	SystemStatusApi *SystemStatusApiService

	TargetingTagReportsApi *TargetingTagReportsApiService

	TargetingTagsApi *TargetingTagsApiService

	TargetingTagsUvApi *TargetingTagsUvApiService

	TargetingsApi *TargetingsApiService

	TargetingsShareApi *TargetingsShareApiService

	TrackingReportsApi *TrackingReportsApiService

	UnionPositionPackagesApi *UnionPositionPackagesApiService

	UserActionSetReportsApi *UserActionSetReportsApiService

	UserActionSetsApi *UserActionSetsApiService

	UserActionsApi *UserActionsApiService

	UserPropertiesApi *UserPropertiesApiService

	UserPropertySetsApi *UserPropertySetsApiService

	VideomakerAutoadjustmentsApi *VideomakerAutoadjustmentsApiService

	VideomakerSubtitlesApi *VideomakerSubtitlesApiService

	VideomakerTasksApi *VideomakerTasksApiService

	VideomakerVideocapturesApi *VideomakerVideocapturesApiService

	VideosApi *VideosApiService

	WechatAdFollowersApi *WechatAdFollowersApiService

	WechatAdLabelsApi *WechatAdLabelsApiService

	WechatAdvertiserApi *WechatAdvertiserApiService

	WechatAdvertiserDetailApi *WechatAdvertiserDetailApiService

	WechatAdvertiserLocalBusinessApi *WechatAdvertiserLocalBusinessApiService

	WechatAdvertiserSpecificationApi *WechatAdvertiserSpecificationApiService

	WechatAgencyApi *WechatAgencyApiService

	WechatDailyCostApi *WechatDailyCostApiService

	WechatFundStatementsDetailedApi *WechatFundStatementsDetailedApiService

	WechatFundTransferApi *WechatFundTransferApiService

	WechatFundsApi *WechatFundsApiService

	WechatPagesApi *WechatPagesApiService

	WechatPagesCsgroupStatusApi *WechatPagesCsgroupStatusApiService

	WechatPagesCsgroupUserApi *WechatPagesCsgroupUserApiService

	WechatPagesCsgrouplistApi *WechatPagesCsgrouplistApiService

	WechatPagesCustomApi *WechatPagesCustomApiService

	WechatPagesGrantinfoApi *WechatPagesGrantinfoApiService

	WechatQualificationsApi *WechatQualificationsApiService

	WildcardsApi *WildcardsApiService

	XijingComplexTemplateApi *XijingComplexTemplateApiService

	XijingPageApi *XijingPageApiService

	XijingPageByComponentsApi *XijingPageByComponentsApiService

	XijingPageInteractiveApi *XijingPageInteractiveApiService

	XijingPageListApi *XijingPageListApiService

	XijingTemplateApi *XijingTemplateApiService

	XijingTemplateListApi *XijingTemplateListApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(sdkConfig *config.SDKConfig) *APIClient {
	cfg := sdkConfig.Configuration
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{}
	}

	c := &APIClient{}
	c.SDKConfig = sdkConfig
	c.Cfg = &cfg
	c.common.client = c

	// API Services
	c.AdDiagnosisApi = (*AdDiagnosisApiService)(&c.common)
	c.AdLabelApi = (*AdLabelApiService)(&c.common)
	c.AdParamApi = (*AdParamApiService)(&c.common)
	c.AdcreativePreviewsApi = (*AdcreativePreviewsApiService)(&c.common)
	c.AdcreativePreviewsQrcodeApi = (*AdcreativePreviewsQrcodeApiService)(&c.common)
	c.AdcreativeTemplateApi = (*AdcreativeTemplateApiService)(&c.common)
	c.AdcreativeTemplateDetailApi = (*AdcreativeTemplateDetailApiService)(&c.common)
	c.AdcreativeTemplateListApi = (*AdcreativeTemplateListApiService)(&c.common)
	c.AdcreativeTemplatePreviewApi = (*AdcreativeTemplatePreviewApiService)(&c.common)
	c.AdcreativeTemplatePreviewsApi = (*AdcreativeTemplatePreviewsApiService)(&c.common)
	c.AdcreativeTemplatesApi = (*AdcreativeTemplatesApiService)(&c.common)
	c.AdcreativesApi = (*AdcreativesApiService)(&c.common)
	c.AdcreativesRelatedCapabilityApi = (*AdcreativesRelatedCapabilityApiService)(&c.common)
	c.AdgroupsApi = (*AdgroupsApiService)(&c.common)
	c.AdsApi = (*AdsApiService)(&c.common)
	c.AdvertiserApi = (*AdvertiserApiService)(&c.common)
	c.AgencyInnerTransferApi = (*AgencyInnerTransferApiService)(&c.common)
	c.AgencyRealtimeCostApi = (*AgencyRealtimeCostApiService)(&c.common)
	c.AndroidChannelPackagesApi = (*AndroidChannelPackagesApiService)(&c.common)
	c.AndroidUnionChannelPackagesApi = (*AndroidUnionChannelPackagesApiService)(&c.common)
	c.AppAndroidChannelPackagesApi = (*AppAndroidChannelPackagesApiService)(&c.common)
	c.AssetPermissionsApi = (*AssetPermissionsApiService)(&c.common)
	c.AssetPrePermissionsApi = (*AssetPrePermissionsApiService)(&c.common)
	c.AsyncReportFilesApi = (*AsyncReportFilesApiService)(&c.common)
	c.AsyncReportsApi = (*AsyncReportsApiService)(&c.common)
	c.AsyncTaskFilesApi = (*AsyncTaskFilesApiService)(&c.common)
	c.AsyncTasksApi = (*AsyncTasksApiService)(&c.common)
	c.AudienceGrantRelationsApi = (*AudienceGrantRelationsApiService)(&c.common)
	c.AuthorizationApi = (*AuthorizationApiService)(&c.common)
	c.BarrageApi = (*BarrageApiService)(&c.common)
	c.BarrageRecommendApi = (*BarrageRecommendApiService)(&c.common)
	c.BatchAsyncRequestSpecificationApi = (*BatchAsyncRequestSpecificationApiService)(&c.common)
	c.BatchAsyncRequestsApi = (*BatchAsyncRequestsApiService)(&c.common)
	c.BatchOperationApi = (*BatchOperationApiService)(&c.common)
	c.BatchRequestsApi = (*BatchRequestsApiService)(&c.common)
	c.BidSimulationApi = (*BidSimulationApiService)(&c.common)
	c.BrandApi = (*BrandApiService)(&c.common)
	c.BusinessManagerRelationsApi = (*BusinessManagerRelationsApiService)(&c.common)
	c.BusinessMdmAccountRelationsApi = (*BusinessMdmAccountRelationsApiService)(&c.common)
	c.CampaignsApi = (*CampaignsApiService)(&c.common)
	c.CapabilitiesApi = (*CapabilitiesApiService)(&c.common)
	c.ChannelsApi = (*ChannelsApiService)(&c.common)
	c.ComplianceValidationApi = (*ComplianceValidationApiService)(&c.common)
	c.ConversionsApi = (*ConversionsApiService)(&c.common)
	c.CreativetoolsTextApi = (*CreativetoolsTextApiService)(&c.common)
	c.CustomAudienceEstimationsApi = (*CustomAudienceEstimationsApiService)(&c.common)
	c.CustomAudienceFilesApi = (*CustomAudienceFilesApiService)(&c.common)
	c.CustomAudienceInsightsApi = (*CustomAudienceInsightsApiService)(&c.common)
	c.CustomAudienceReportsApi = (*CustomAudienceReportsApiService)(&c.common)
	c.CustomAudiencesApi = (*CustomAudiencesApiService)(&c.common)
	c.CustomFeaturesApi = (*CustomFeaturesApiService)(&c.common)
	c.CustomTagFilesApi = (*CustomTagFilesApiService)(&c.common)
	c.CustomTagsApi = (*CustomTagsApiService)(&c.common)
	c.DailyBalanceReportApi = (*DailyBalanceReportApiService)(&c.common)
	c.DailyCostApi = (*DailyCostApiService)(&c.common)
	c.DailyReportsApi = (*DailyReportsApiService)(&c.common)
	c.DiagnosisApi = (*DiagnosisApiService)(&c.common)
	c.DplabelAdLabelApi = (*DplabelAdLabelApiService)(&c.common)
	c.DynamicAdImageTemplatesApi = (*DynamicAdImageTemplatesApiService)(&c.common)
	c.DynamicAdImagesApi = (*DynamicAdImagesApiService)(&c.common)
	c.DynamicAdTemplatesApi = (*DynamicAdTemplatesApiService)(&c.common)
	c.DynamicAdVideoApi = (*DynamicAdVideoApiService)(&c.common)
	c.DynamicAdVideoTemplatesApi = (*DynamicAdVideoTemplatesApiService)(&c.common)
	c.DynamicCreativesApi = (*DynamicCreativesApiService)(&c.common)
	c.EcommerceOrderApi = (*EcommerceOrderApiService)(&c.common)
	c.EstimationApi = (*EstimationApiService)(&c.common)
	c.FundStatementsDailyApi = (*FundStatementsDailyApiService)(&c.common)
	c.FundStatementsDetailedApi = (*FundStatementsDetailedApiService)(&c.common)
	c.FundTransferApi = (*FundTransferApiService)(&c.common)
	c.FundsApi = (*FundsApiService)(&c.common)
	c.HourlyReportsApi = (*HourlyReportsApiService)(&c.common)
	c.ImageProcessingApi = (*ImageProcessingApiService)(&c.common)
	c.ImagesApi = (*ImagesApiService)(&c.common)
	c.LabelAudiencesApi = (*LabelAudiencesApiService)(&c.common)
	c.LabelsApi = (*LabelsApiService)(&c.common)
	c.LeadCluesApi = (*LeadCluesApiService)(&c.common)
	c.LeadsFormApi = (*LeadsFormApiService)(&c.common)
	c.LeadsFormListApi = (*LeadsFormListApiService)(&c.common)
	c.LeadsInvalidPayApi = (*LeadsInvalidPayApiService)(&c.common)
	c.LocalApi = (*LocalApiService)(&c.common)
	c.LocalEndadsmanuallyApi = (*LocalEndadsmanuallyApiService)(&c.common)
	c.LocalEstimatedamountApi = (*LocalEstimatedamountApiService)(&c.common)
	c.LocalStoresApi = (*LocalStoresApiService)(&c.common)
	c.LocalStoresAddressParsingResultApi = (*LocalStoresAddressParsingResultApiService)(&c.common)
	c.LocalStoresCategoriesApi = (*LocalStoresCategoriesApiService)(&c.common)
	c.LocalStoresSearchInfoApi = (*LocalStoresSearchInfoApiService)(&c.common)
	c.LocalStoresWxpayMerchantsApi = (*LocalStoresWxpayMerchantsApiService)(&c.common)
	c.OauthApi = (*OauthApiService)(&c.common)
	c.OptimizationGoalPermissionsApi = (*OptimizationGoalPermissionsApiService)(&c.common)
	c.OuterCluesApi = (*OuterCluesApiService)(&c.common)
	c.OuterCluesContactApi = (*OuterCluesContactApiService)(&c.common)
	c.PagesApi = (*PagesApiService)(&c.common)
	c.PlayablePagesApi = (*PlayablePagesApiService)(&c.common)
	c.ProductCatalogsApi = (*ProductCatalogsApiService)(&c.common)
	c.ProductCatalogsReportsApi = (*ProductCatalogsReportsApiService)(&c.common)
	c.ProductCategoriesListApi = (*ProductCategoriesListApiService)(&c.common)
	c.ProductItemsApi = (*ProductItemsApiService)(&c.common)
	c.ProductItemsDetailApi = (*ProductItemsDetailApiService)(&c.common)
	c.ProductItemsVerticalsApi = (*ProductItemsVerticalsApiService)(&c.common)
	c.ProductSeriesApi = (*ProductSeriesApiService)(&c.common)
	c.ProductsSystemStatusApi = (*ProductsSystemStatusApiService)(&c.common)
	c.ProfilesApi = (*ProfilesApiService)(&c.common)
	c.PromotedObjectsApi = (*PromotedObjectsApiService)(&c.common)
	c.PropertyFileSessionsApi = (*PropertyFileSessionsApiService)(&c.common)
	c.PropertyFilesApi = (*PropertyFilesApiService)(&c.common)
	c.PropertySetSchemasApi = (*PropertySetSchemasApiService)(&c.common)
	c.PropertySetsApi = (*PropertySetsApiService)(&c.common)
	c.QualificationsApi = (*QualificationsApiService)(&c.common)
	c.RealtimeCostApi = (*RealtimeCostApiService)(&c.common)
	c.ReportApi = (*ReportApiService)(&c.common)
	c.ReviewElementPrereviewResultsApi = (*ReviewElementPrereviewResultsApiService)(&c.common)
	c.SceneSpecTagsApi = (*SceneSpecTagsApiService)(&c.common)
	c.ShopApi = (*ShopApiService)(&c.common)
	c.SplitTestsApi = (*SplitTestsApiService)(&c.common)
	c.SubcustomerTransferApi = (*SubcustomerTransferApiService)(&c.common)
	c.SystemStatusApi = (*SystemStatusApiService)(&c.common)
	c.TargetingTagReportsApi = (*TargetingTagReportsApiService)(&c.common)
	c.TargetingTagsApi = (*TargetingTagsApiService)(&c.common)
	c.TargetingTagsUvApi = (*TargetingTagsUvApiService)(&c.common)
	c.TargetingsApi = (*TargetingsApiService)(&c.common)
	c.TargetingsShareApi = (*TargetingsShareApiService)(&c.common)
	c.TrackingReportsApi = (*TrackingReportsApiService)(&c.common)
	c.UnionPositionPackagesApi = (*UnionPositionPackagesApiService)(&c.common)
	c.UserActionSetReportsApi = (*UserActionSetReportsApiService)(&c.common)
	c.UserActionSetsApi = (*UserActionSetsApiService)(&c.common)
	c.UserActionsApi = (*UserActionsApiService)(&c.common)
	c.UserPropertiesApi = (*UserPropertiesApiService)(&c.common)
	c.UserPropertySetsApi = (*UserPropertySetsApiService)(&c.common)
	c.VideomakerAutoadjustmentsApi = (*VideomakerAutoadjustmentsApiService)(&c.common)
	c.VideomakerSubtitlesApi = (*VideomakerSubtitlesApiService)(&c.common)
	c.VideomakerTasksApi = (*VideomakerTasksApiService)(&c.common)
	c.VideomakerVideocapturesApi = (*VideomakerVideocapturesApiService)(&c.common)
	c.VideosApi = (*VideosApiService)(&c.common)
	c.WechatAdFollowersApi = (*WechatAdFollowersApiService)(&c.common)
	c.WechatAdLabelsApi = (*WechatAdLabelsApiService)(&c.common)
	c.WechatAdvertiserApi = (*WechatAdvertiserApiService)(&c.common)
	c.WechatAdvertiserDetailApi = (*WechatAdvertiserDetailApiService)(&c.common)
	c.WechatAdvertiserLocalBusinessApi = (*WechatAdvertiserLocalBusinessApiService)(&c.common)
	c.WechatAdvertiserSpecificationApi = (*WechatAdvertiserSpecificationApiService)(&c.common)
	c.WechatAgencyApi = (*WechatAgencyApiService)(&c.common)
	c.WechatDailyCostApi = (*WechatDailyCostApiService)(&c.common)
	c.WechatFundStatementsDetailedApi = (*WechatFundStatementsDetailedApiService)(&c.common)
	c.WechatFundTransferApi = (*WechatFundTransferApiService)(&c.common)
	c.WechatFundsApi = (*WechatFundsApiService)(&c.common)
	c.WechatPagesApi = (*WechatPagesApiService)(&c.common)
	c.WechatPagesCsgroupStatusApi = (*WechatPagesCsgroupStatusApiService)(&c.common)
	c.WechatPagesCsgroupUserApi = (*WechatPagesCsgroupUserApiService)(&c.common)
	c.WechatPagesCsgrouplistApi = (*WechatPagesCsgrouplistApiService)(&c.common)
	c.WechatPagesCustomApi = (*WechatPagesCustomApiService)(&c.common)
	c.WechatPagesGrantinfoApi = (*WechatPagesGrantinfoApiService)(&c.common)
	c.WechatQualificationsApi = (*WechatQualificationsApiService)(&c.common)
	c.WildcardsApi = (*WildcardsApiService)(&c.common)
	c.XijingComplexTemplateApi = (*XijingComplexTemplateApiService)(&c.common)
	c.XijingPageApi = (*XijingPageApiService)(&c.common)
	c.XijingPageByComponentsApi = (*XijingPageByComponentsApiService)(&c.common)
	c.XijingPageInteractiveApi = (*XijingPageInteractiveApiService)(&c.common)
	c.XijingPageListApi = (*XijingPageListApiService)(&c.common)
	c.XijingTemplateApi = (*XijingTemplateApiService)(&c.common)
	c.XijingTemplateListApi = (*XijingTemplateListApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	case "multi":
		if jsonString, err := json.Marshal(obj); err == nil {
			return string(jsonString)
		}
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	if reflect.TypeOf(obj).Kind() == reflect.Struct {
		if jsonString, err := json.Marshal(obj); err == nil {
			return string(jsonString)
		}
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.Cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.Cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte,
	fileKey string) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(fileKey, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}
		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.Cfg.Host != "" {
		localVarRequest.Host = c.Cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.Cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(config.ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(config.ContextBasicAuth).(config.BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(config.ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.Cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		dec := json.NewDecoder(bytes.NewReader(b))
		if c.SDKConfig.IsStrictMode {
			dec.DisallowUnknownFields() // Force
		}
		if err := dec.Decode(v); err != nil {
			return apierrors.NewResponseStrictError(err.Error())
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
