package qweather

/*
单位
和风天气默认采用公制单位，例如：公里、摄氏度等，如需要采用英制单位，可以通过添加对应的参数进行设置。

单位参数

单位	API请求参数	iOS常量属性	Android常量属性
公制单位	        m	        UNIT_TYPE_M	METRIC
英制单位	        i	        UNIT_TYPE_I	IMPERIAL

单位列表

	> 有一些数据项不存在英制单位的，统一使用公制单位。

数据项	    公制单位	        英制单位
温度	        摄氏度	        华氏度
风速	        公里/小时 km/h	英里/小时 mile/h
能见度	    公里 km	        英里 mile
大气压强	    百帕 hPa	        百帕 hPa
降水量	    毫米 mm	        毫米 mm
PM2.5	    微克/立方米 μg/m3	微克/立方米 μg/m3
PM10	    微克/立方米 μg/m3	微克/立方米 μg/m3
O3	        微克/立方米 μg/m3	微克/立方米 μg/m3
SO2	        微克/立方米 μg/m3	微克/立方米 μg/m3
CO	        毫克/立方米 mg/m3	毫克/立方米 mg/m3
NO2	        微克/立方米 μg/m3	微克/立方米 μg/m3
*/

type UnitType int8

func (u UnitType) String() string {
	switch u {
	case METRIC:
		return "m"
	case IMPERIAL:
		return "i"
	default:
		return "m" //  默认公制
	}
}

const (
	_        UnitType = iota
	METRIC            // 公制单位
	IMPERIAL          // 英制单位
)
