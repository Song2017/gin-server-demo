/*
 * Nomad Envoy Utils
 *
 * This is the API descriptor for the Nomad Platform and Customs utils. Developed by `Samarkand Global <https://global>`_
 *
 * API version: 1.0.0
 * Contact: paul@global
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

import (
	model "apiserver/v1/nomad-model"
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// BatchOperatecypher - cypher
func BatchOperatecypher(c *gin.Context) {
	// get body data
	var cypher_body model.Cypher
	var bodyBytes []byte

	bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	json.Unmarshal(bodyBytes, &cypher_body)

	resp := model.ApiResponse{
		Code:    200,
		Data:    model.ApiResponseData{Items: []string{}, AreAllEncrypted: false},
		Message: "OK",
	}

	// get platform
	platform := cypher_body.Platform
	var platform_able PlatformAble
	switch platform {
	case cypher_body.Platform:
		platform_able = &PlatformCypher{cypher_body}
	default:
		platform_able = &PlatformCypher{cypher_body}
	}

	// execute platform functions
	switch cypher_body.Operation {
	case DECRYPT_BATCH:
		platform_able.DecryptCypher(&resp)
	case ENCRYPT_BATCH:
		platform_able.EncryptCypher(&resp)
	case ARE_ALL_ENCRYPTED:
		platform_able.VerifyCypher(&resp)
	default:
		platform_able.DecryptCypher(&resp)
	}

	c.JSON(int(resp.Code), resp)
}
