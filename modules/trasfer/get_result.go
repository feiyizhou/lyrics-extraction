package trasfer

type GetResultRequest struct {
	OrderId    string `json:"orderId"`
	ResultType string `json:"resultType"`
}

type GetResultResponse struct {
	Code     string              `json:"code"`
	DescInfo string              `json:"descInfo"`
	Content  getResultResContent `json:"content"`
}

type getResultResContent struct {
	OrderResult      string    `json:"orderResult"`
	OrderInfo        orderInfo `json:"orderInfo"`
	TaskEstimateTime int       `json:"taskEstimateTime"`
}

type orderInfo struct {
	FailType         int    `json:"failType"`
	Status           int    `json:"status"`
	OrderId          string `json:"orderId"`
	OriginalDuration int64  `json:"originalDuration"`
	RealDuration     int64  `json:"realDuration"`
}

type OrderResult struct {
	Lattice  []map[string]string      `json:"lattice"`
	Lattice2 []map[string]interface{} `json:"lattice2"`
	Label    label                    `json:"label"`
}

type Lattice struct {
	Json1Best Json1Best `json:"json_1best"`
}

type Json1Best struct {
	ST st `json:"st"`
}

type st struct {
	SC string `json:"sc"`
	PA string `json:"pa"`
	RT []rt   `json:"rt"`
	BG string `json:"bg"`
	RL string `json:"rl"`
	ED string `json:"ed"`
}

type rt struct {
	WS []ws `json:"ws"`
}

type ws struct {
	CW []cw `json:"cw"`
	WB int  `json:"wb"`
	WE int  `json:"we"`
}

type cw struct {
	W  string `json:"w"`
	WP string `json:"wp"`
	WC string `json:"wc"`
}

type label struct {
	TLTrack []rlTrack `json:"tl_track"`
}

type rlTrack struct {
	RL    string `json:"rl"`
	Track string `json:"track"`
}
