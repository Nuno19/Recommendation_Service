/*
 * Recommendation API
 *
 * This is a recommendation API using k-means Clustering
 *
 * API version: 1.0.0
 * Contact: capela.nuno@ua.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Item struct {

	// Identifier of the item
	Id string `json:"id"`

	// Data of the item
	Data string `json:"data"`

	//Which centroid is it nearest to
	BelongsTo int `json:"belongsTo,omitempty"`
}

type Items []Item
