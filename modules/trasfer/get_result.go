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
	Lattice  map[interface{}]interface{} `json:"lattice"`
	Lattice2 map[interface{}]interface{} `json:"lattice2"`
	Label    map[interface{}]interface{} `json:"label"`
}
