package main

// program configuration
type Conf struct {
	ServerPort int      // server port
	ClientPort int      // client port
	HostIp     string   // host ip address
	Role       int      // role definition
	Paths      []string // file paths
	BatchSize  int64    // data batch size
}

// file header struct
type Header struct {
	Uuid          string `json:"uuid"`           // current send uuid
	TotalSize     int64  `json:"total_size"`     // total size of file
	ContentLength int64  `json:"content_length"` // current length of data
	ContentType   string `json:"content_type"`   // file content-type
	Offset        int64  `json:"offset"`         // current offset
	Filename      string `json:"filename"`       // filename
	LastStatus    bool   `json:"last_status"`    // last transfer status
}

// ack info to assure that data has sent successfully
type Ack struct {
	Status bool   `json:"status"`
	Offset int64  `json:"offset"`
	Uuid   string `json:"uuid"`
}

// constants
const (
	Host = iota
	Client
)
