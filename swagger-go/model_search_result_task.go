/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SearchResultTask struct {
	TotalHits int64 `json:"totalHits,omitempty"`
	Results []Task `json:"results,omitempty"`
}