package geo

type SearchResponse struct {
	Code     string `json:"code"`
	Location []struct {
		Name      string `json:"name"`
		Id        string `json:"id"`
		Lat       string `json:"lat"`
		Lon       string `json:"lon"`
		Adm2      string `json:"adm2"`
		Adm1      string `json:"adm1"`
		Country   string `json:"country"`
		Tz        string `json:"tz"`
		UtcOffset string `json:"utcOffset"`
		IsDst     string `json:"isDst"`
		Type      string `json:"type"`
		Rank      string `json:"rank"`
		FxLink    string `json:"fxLink"`
	} `json:"location"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}

type HitResponse struct {
	Code        string `json:"code"`
	TopCityList []struct {
		Name      string `json:"name"`
		Id        string `json:"id"`
		Lat       string `json:"lat"`
		Lon       string `json:"lon"`
		Adm2      string `json:"adm2"`
		Adm1      string `json:"adm1"`
		Country   string `json:"country"`
		Tz        string `json:"tz"`
		UtcOffset string `json:"utcOffset"`
		IsDst     string `json:"isDst"`
		Type      string `json:"type"`
		Rank      string `json:"rank"`
		FxLink    string `json:"fxLink"`
	} `json:"topCityList"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}

type POIResponse struct {
	Code string `json:"code"`
	Poi  []struct {
		Name      string `json:"name"`
		Id        string `json:"id"`
		Lat       string `json:"lat"`
		Lon       string `json:"lon"`
		Adm2      string `json:"adm2"`
		Adm1      string `json:"adm1"`
		Country   string `json:"country"`
		Tz        string `json:"tz"`
		UtcOffset string `json:"utcOffset"`
		IsDst     string `json:"isDst"`
		Type      string `json:"type"`
		Rank      string `json:"rank"`
		FxLink    string `json:"fxLink"`
	} `json:"poi"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}
