package airqualityBeta

type NowPara struct {
	LocationID string
	Lang       string
	Station    bool
	Pollutant  bool
}

type StationPara struct {
	LocationID string
	Lang       string
}
