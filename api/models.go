package handlers

type DSTResponse struct {
	DST           float64 `json:"dst"`
	StormStrength string  `json:"stormStrength"`
}

type BzResponse struct {
	Bz            float64 `json:"bz"`
	StormStrength string  `json:"stormStrength"`
}
