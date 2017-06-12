package globals

type Config struct {
	Type DatabaseType `json:"type"`
	Host string `json:"host"`
	Database string `json:"database"`
}
