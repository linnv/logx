package logx

type logStruct struct {
	Line     int    `json:"line"`
	File     string `json:"file"`
	Datetime string `json:"time"`
	Content  string `json:"log"`
}
