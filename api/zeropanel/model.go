package zeropanel

import "encoding/json"

// NodeInfoResponse is the response of node
type NodeInfoResponse struct {
	Class        int             `json:"node_class"`
	SpeedLimit   float64         `json:"node_speedlimit"`
	TrafficRate  float64         `json:"traffic_rate"`
	Type         int             `json:"node_type"`
	CustomConfig json.RawMessage `json:"custom_config"`
}

type CustomConfig struct {
	OffsetPortUser string          `json:"offset_port_user"`
	OffsetPortNode string          `json:"offset_port_node"`
	ServerSub      string          `json:"server_sub"`
	Host           string          `json:"host"`
	MuPort         string          `json:"mu_port"`
	MuEncryption   string          `json:"mu_encryption"`
	MuProtocol     string          `json:"mu_protocol"`
	MuObfs         string          `json:"mu_obfs"`
	MuSuffix       string          `json:"mu_suffix"`
	ServerPsk      string          `json:"server_psk"`
	V2Port         string          `json:"v2_port"`
	TLS            string          `json:"tls"`
	EnableVless    string          `json:"enable_vless"`
	Network        string          `json:"network"`
	Security       string          `json:"security"`
	Path           string          `json:"path"`
	VerifyCert     bool            `json:"verify_cert"`
	Obfs           string          `json:"obfs"`
	Header         json.RawMessage `json:"header"`
	TrojanPort     string          `json:"trojan_port"`
	AllowInsecure  string          `json:"allow_insecure"`
	Grpc           string          `json:"grpc"`
	Servicename    string          `json:"servicename"`
	Flow           string          `json:"flow"`
}

// UserResponse is the response of user
type UserResponse struct {
	ID          int     `json:"id"`
	Email       string  `json:"email"`
	Passwd      string  `json:"passwd"`
	Port        uint32  `json:"port"`
	SpeedLimit  float64 `json:"node_speedlimit"`
	DeviceLimit int     `json:"node_iplimit"`
	UUID        string  `json:"uuid"`
	AliveIP     int     `json:"alive_ip"`
}

// Response is the common response
type Response struct {
	Ret  uint            `json:"ret"`
	Data json.RawMessage `json:"data"`
}

// PostData is the data structure of post data
type PostData struct {
	Data interface{} `json:"data"`
}

// SystemLoad is the data structure of systemload
type SystemLoad struct {
	Uptime string `json:"uptime"`
	Load   string `json:"load"`
}

// OnlineUser is the data structure of online user
type OnlineUser struct {
	UID int    `json:"user_id"`
	IP  string `json:"ip"`
}

// UserTraffic is the data structure of traffic
type UserTraffic struct {
	UID      int   `json:"user_id"`
	Upload   int64 `json:"u"`
	Download int64 `json:"d"`
}

type RuleItem struct {
	ID      int    `json:"id"`
	Content string `json:"regex"`
}

type IllegalItem struct {
	ID  int `json:"list_id"`
	UID int `json:"user_id"`
}
