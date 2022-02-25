package main

//ReqJob 截图的参数
type ReqJob struct {
	Url      string  `json:"url"`       //必填
	PxHeight int64   `json:"px_height"` //像素 可选
	PxWidth  int64   `json:"px_width"`  //像素 可选
	Quality  int64   `json:"quality"`   //图片质量 可选
	Sel      string  `json:"sel"`       //css 选择器 可选
	Timeout  int     `json:"timeout"`   //time out second
	Wait     int     `json:"wait"`      //等待时间  second
	Scale    float64 `json:"scale"`     //缩放比例 可选
}

//ResJob 截图的结果
type ResJob struct {
	Code int    `json:"code"` //200==ok
	Msg  string `json:"msg"`  //msg
	Uri  string `json:"uri"`  //uri
	Url  string `json:"url"`  //图片地址
}
