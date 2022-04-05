package metrics

import "time"

// 统计

type Metrics struct {
	ResponseTimes map[string][]float64
	TimeStamps    map[string][]float64
	//executor
}

// 记录响应时间
func (m *Metrics) RecordResponseTime(apiName string, responseTime float64) {
	if m.ResponseTimes == nil {
		m.ResponseTimes = map[string][]float64{}
	}
	m.ResponseTimes[apiName] = append(m.ResponseTimes[apiName], responseTime)
}

// 记录时间戳
func (m *Metrics) RecordTimestamp(apiName string, timestamp float64) {
	if m.TimeStamps == nil {
		m.TimeStamps = map[string][]float64{}
	}
	m.TimeStamps[apiName] = append(m.TimeStamps[apiName], timestamp)
}

func (m *Metrics) StartRepeatedReport(period int, unit time.Duration) {
	stats := map[string]map[string]float64{}
	for apiName, apiRespTimes := range m.ResponseTimes {
		stats[apiName]["max"] = m.max(apiRespTimes)
		stats[apiName]["avg"] = m.avg(apiRespTimes)
	}
	for apiName, apiTimestamps := range m.TimeStamps {
		stats[apiName]["count"] = float64(len(apiTimestamps))
	}

}

// 最大值
func (m *Metrics) max(dateset []float64) float64 {
	return 0

}

// 平均值
func (m *Metrics) avg(dateset []float64) float64 {
	return 0
}
