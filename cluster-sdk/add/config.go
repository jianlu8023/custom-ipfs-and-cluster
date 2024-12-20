package add

type Info struct {
	FileName             string `json:"fileName"`
	FilePath             string `json:"filePath"`
	ExpireAt             string `json:"expireAt"`
	ReplicationFactorMin int    `json:"replicationFactorMin"`
	ReplicationFactorMax int    `json:"replicationFactorMax"`
}
