package warning

type RealTimeResponse struct {
	Code       string     `json:"code"`       // 状态码
	UpdateTime string     `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/resource/glossary/#update-time
	FxLink     string     `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Warning    []struct { // 注意：如果查询的地区当前没有灾害预警信息，返回的warning字段为空。
		Id        string `json:"id"`        // 本条预警的唯一标识，可判断本条预警是否已经存在
		Sender    string `json:"sender"`    // 预警发布单位，**可能为空**
		PubTime   string `json:"pubTime"`   // 预警发布时间
		Title     string `json:"title"`     // 预警信息标题
		StartTime string `json:"startTime"` // 预警开始时间，**可能为空**
		EndTime   string `json:"endTime"`   // 预警结束时间 https://dev.qweather.com/docs/resource/warning-info/#expiry-time 可能为空
		Status    string `json:"status"`    // 预警信息的发布状态 https://dev.qweather.com/docs/resource/warning-info/#status
		// Level         string `json:"level"`         预警等级 https://dev.qweather.com/docs/resource/warning-info/#level-deprecated（已弃用），不要再使用这个字段，该字段已弃用，目前返回为空或未更新的值。请使用severity和severityColor代替
		Severity      string `json:"severity"`      // 预警严重等级 https://dev.qweather.com/docs/resource/warning-info/#severity
		SeverityColor string `json:"severityColor"` // 预警严重等级颜色 https://dev.qweather.com/docs/resource/warning-info/#severity-color **可能为空**
		Type          string `json:"type"`          // 预警类型ID https://dev.qweather.com/docs/resource/warning-info/#warning-type
		TypeName      string `json:"typeName"`      // 预警类型名称 https://dev.qweather.com/docs/resource/warning-info/#warning-type
		Urgency       string `json:"urgency"`       // 预警信息的紧迫程度，可能为空 https://dev.qweather.com/docs/resource/warning-info/#urgency
		Certainty     string `json:"certainty"`     // 预警信息的确定性，可能为空 https://dev.qweather.com/docs/resource/warning-info/#certainty
		Text          string `json:"text"`          // 预警详细文字描述
		Related       string `json:"related"`       // 与本条预警相关联的预警ID，当预警状态为cancel或update时返回。**可能为空**
	} `json:"warning"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}

type CityListResponse struct {
	Code           string `json:"code"`       // 状态码
	UpdateTime     string `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/resource/glossary/#update-time
	WarningLocList []struct {
		LocationId string `json:"locationId"` // 当前国家预警的LocationID
	} `json:"warningLocList"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}

type T struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Warning    []struct {
		Id            string `json:"id"`
		Sender        string `json:"sender"`
		PubTime       string `json:"pubTime"`
		Title         string `json:"title"`
		StartTime     string `json:"startTime"`
		EndTime       string `json:"endTime"`
		Status        string `json:"status"`
		Level         string `json:"level"`
		Severity      string `json:"severity"`
		SeverityColor string `json:"severityColor"`
		Type          string `json:"type"`
		TypeName      string `json:"typeName"`
		Urgency       string `json:"urgency"`
		Certainty     string `json:"certainty"`
		Text          string `json:"text"`
		Related       string `json:"related"`
	} `json:"warning"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}
