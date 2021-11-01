package apiserver

type Cypher struct {
	StoreId   string   `json:"storeId"`
	Platform  string   `json:"Platform"`
	Operation string   `json:"operation"`
	Items     []string `json:"items"`
}

type AppSettings struct {
	AppTimeout     int                 `json:"AppTimeout"`
	AppOperations  []string            `json:"AppOperations"`
	AllStoreIDs    map[string][]string `json:"AllStoreIDs"`
	SecurityCaKey  string              `json:"SecurityCaKey"`
	GinMode        string              `json:"GinMode"`
	RedisHost      string              `json:"RedisHost"`
	RedisPass      string              `json:"RedisPass"`
	RedisHostPilot string              `json:"RedisHostPilot"`
	RedisPassPilot string              `json:"RedisPassPilot"`
}
