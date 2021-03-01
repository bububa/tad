/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AsyncTasksAddRequest struct {
	AccountId *int64    `json:"account_id,omitempty"`
	TaskName  *string   `json:"task_name,omitempty"`
	TaskType  TaskType  `json:"task_type,omitempty"`
	TaskSpec  *TaskSpec `json:"task_spec,omitempty"`
}
