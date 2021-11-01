package apiserver

type NomadLog struct {
	Platform    string  `json:"platform"`
	Latency     float64 `json:"latency"`
	TimingStart string  `json:"timeBegin"`
	TimingEnd   string  `json:"timeEnd"`
	Label       string  `json:"label"`
	RemoteAddr  string  `json:"remoteAddr"`
	Input       string  `json:"input"`
	Response    string  `json:"response"`
}
