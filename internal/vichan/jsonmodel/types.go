package jsonmodel

type Thread struct {
	Posts []Message
}

type Message struct {
	No             int    `json:"no"`
	Title          string `json:"sub"`
	Comment        string `json:"com"`
	Author         string `json:"name"`
	trip           string
	Time           int64 `json:"time"`
	omitted_posts  int
	omitted_images int
	sticky         int
	locked         int
	cyclical       string
	last_modified  int
	tn_h           int
	tn_w           int
	H              int
	W              int
	Fsize          uint64 `json:"fsize"`
	Filename       string `json:"filename"`
	Ext            string `json:"ext"`
	Tim            string `json:"tim"`
	md5            string
	Resto          int     `json:"resto"`
	ExtraFiles     []Extra `json:"extra_files"`
}

type Extra struct {
	tn_h     int
	tn_w     int
	H        int
	W        int
	Fsize    uint64 `json:"fsize"`    // in bytes
	Filename string `json:"filename"` // Original filename, without extension
	Ext      string `json:"ext"`      // eg: ".png"
	Tim      string `json:"tim"`      // Uploaded filename, without extension
	md5      string
}
