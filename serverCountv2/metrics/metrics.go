package metrics

import "Algorithms/serverCountv2/metrics/storage"

type MetricsCollector struct {
	metricStorage storage.MetricsStorage
}

// 依赖注入
func (m *MetricsCollector) MetricsCollector(metricsStorage *RedisMetricStorage) {
	m.metricStorage = metricsStorage
}

type RequestInfo struct {
	ApiName      string `json:"api_name"`
	ResponseTime int  `json:"response_time"`
	Timestamp    int64  `json:"timestamp"`
}

func (r *RequestInfo) GetResponseTime() int {
	return r.ResponseTime
}

func (m *MetricsCollector) RecordRequest(requestInfo *RequestInfo) {
	if (requestInfo == nil) || requestInfo.ApiName == "" {
		return
	}
	m.metricStorage.SaverRequestInfos(requestInfo)
}
