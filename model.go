package main

type ReqJob struct {
	Url      string `json:"url"`
	PxHeight int64  `json:"px_height"`
	PxWidth  int64  `json:"px_width"`
	Quality  int64  `json:"quality"`
	Sel      string `json:"sel"`     // css 选择器
	Timeout  int    `json:"timeout"` //等待second
	Wait     int    `json:"wait"`
	Format   string `json:"format"` //pdf or image
}

type ResJob struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Uri  string `json:"uri"`
	Url  string `json:"url"`
}
