package apollo

// Path 路径信息
type Path struct {
	AppID     string
	NameSpace string
	Key       string
	Index     uint64 // CAS
}
