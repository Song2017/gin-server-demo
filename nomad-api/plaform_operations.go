package apiserver

import model "apiserver/v1/nomad-model"

// Operation : operation name, AreAllEncrypted, EncryptBatch and DecryptBatch
// List of Operation
const (
	ARE_ALL_ENCRYPTED string = "AreAllEncrypted"
	ENCRYPT_BATCH     string = "EncryptBatch"
	DECRYPT_BATCH     string = "DecryptBatch"
)

type PlatformCypher struct {
	model.Cypher
}
type PlatformAble interface {
	GetToken() (bool, string)

	DecryptCypher(resp *model.ApiResponse)
	EncryptCypher(resp *model.ApiResponse)
	VerifyCypher(resp *model.ApiResponse)
}

// Interface integrity Check
var _ PlatformAble = (*PlatformCypher)(nil)
