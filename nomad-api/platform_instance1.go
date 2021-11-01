package apiserver

import (
	model "apiserver/v1/nomad-model"
)

func (c PlatformCypher) GetToken() (bool, string) {
	token_store_id := "redis_table"

	return true, token_store_id
}

func (c PlatformCypher) DecryptCypher(resp *model.ApiResponse) {
	resp.Message = "DecryptCypher"
}

func (c PlatformCypher) EncryptCypher(resp *model.ApiResponse) {
	resp.Message = "EncryptCypher"
}

func (c PlatformCypher) VerifyCypher(resp *model.ApiResponse) {
	resp.Message = "VerifyCypher"
}
