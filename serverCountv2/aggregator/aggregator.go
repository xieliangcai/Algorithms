package aggregator

import "Algorithms/serverCountv2/metrics"

type Aggregator struct {
}

func (a *Aggregator) Aggregate(requestInfos []*metrics.RequestInfo, durationInMillis int) RequestStat {
	maxRespTime := int(999)
	minRespTime := int(0)
	avgRespTime := int(-1)
	p999RespTime := int(-1)
	p99RespTime := int(-1)
	sumRespTime := 0
	count := 0
	for _, requestInfo := range requestInfos {
		count++
		respTime := requestInfo.GetResponseTime()
		if maxRespTime < respTime {
			maxRespTime = respTime
		}
		if minRespTime > respTime {
			minRespTime = respTime
		}

		sumRespTime += respTime
	}

	if count != 0 {
		avgRespTime = sumRespTime / count
	}

	tps := count / durationInMillis * 1000
	idx999 := int(float64(count) * 0.999)
	idx99 := int(float64(count) * 0.99)
	if count != 0 {
		p99RespTime = requestInfos[idx99].GetResponseTime()
		p999RespTime = requestInfos[idx999].GetResponseTime()
	}

	requestStat := RequestStat{
		maxResponseTime:  maxRespTime,
		minResponseTime:  minRespTime,
		avgResponseTime:  avgRespTime,
		p99ResponseTime:  p99RespTime,
		p999ResponseTime: p999RespTime,
		count:            count,
		tps:              tps,
	}
	return requestStat

}

type RequestStat struct {
	maxResponseTime  int
	minResponseTime  int
	avgResponseTime  int
	p999ResponseTime int
	p99ResponseTime  int
	count            int
	tps              int
}
