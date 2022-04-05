package storage

import "Algorithms/serverCountv2/metrics"

type MetricsStorage interface {
	SaverRequestInfos(requestInfo *metrics.RequestInfo)
	GetRequestInfos(apiname string, startTimeMillis, endTimeMillis int) []*metrics.RequestInfo
	GetRequestMapInfos(startTimeMillis, endTimeMillis int) map[string][]*metrics.RequestInfo
}

