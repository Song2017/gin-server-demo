# Nomad Envoy Utils Application
export APP_TIMEOUT='5'
export SECURITY_CA_KEY='test'
export GIN_MODE='debug' # debug test release

# Platform configuration
export APP_OPERATIONS='AreAllEncrypted,EncryptBatch,DecryptBatch'

# DB
export REDIS_SERVER_HOST='reids_server:6379'
export REDIS_SERVER_PASSWORD='test'
export REDIS_SERVER_HOST_PILOT='reids_server2:6379'
export REDIS_SERVER_PASSWORD_PILOT='test'


