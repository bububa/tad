/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// AuditStatus : 商品审核状态
type AuditStatus string

// List of AuditStatus
const (
	AuditStatus_AUDITSTATUS_NORMAL             AuditStatus = "AUDITSTATUS_NORMAL"
	AuditStatus_AUDITSTATUS_PENDING            AuditStatus = "AUDITSTATUS_PENDING"
	AuditStatus_PRODUCT_STATUS_PENDING         AuditStatus = "PRODUCT_STATUS_PENDING"
	AuditStatus_PRODUCT_STATUS_AUDITING        AuditStatus = "PRODUCT_STATUS_AUDITING"
	AuditStatus_PRODUCT_STATUS_PASS            AuditStatus = "PRODUCT_STATUS_PASS"
	AuditStatus_PRODUCT_STATUS_REJECTED        AuditStatus = "PRODUCT_STATUS_REJECTED"
	AuditStatus_AUDITSTATUS_DELAY              AuditStatus = "AUDITSTATUS_DELAY"
	AuditStatus_AUDITSTATUS_DENIED             AuditStatus = "AUDITSTATUS_DENIED"
	AuditStatus_RESOURCE_AUDIT_STATUS_UNKNOWN  AuditStatus = "RESOURCE_AUDIT_STATUS_UNKNOWN"
	AuditStatus_RESOURCE_AUDIT_STATUS_APPROVED AuditStatus = "RESOURCE_AUDIT_STATUS_APPROVED"
	AuditStatus_RESOURCE_AUDIT_STATUS_PENDING  AuditStatus = "RESOURCE_AUDIT_STATUS_PENDING"
	AuditStatus_RESOURCE_AUDIT_STATUS_REJECTED AuditStatus = "RESOURCE_AUDIT_STATUS_REJECTED"
)
