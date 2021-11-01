/*
 * Demo Server
 *
 * This is the API descriptor for demo server
 *
 * API version: 1.0.0
 * Contact: bensong2017@hotmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

type ApiResponseData struct {

	Items []string `json:"items,omitempty"`

	AreAllEncrypted bool `json:"areAllEncrypted,omitempty"`
}
