package consoleReporter

import (
	"Algorithms/serverCountv2/aggregator"
	"Algorithms/serverCountv2/metrics/storage"
	"code.geesunn.com/common/gocommon/util/jsonUtil"
	"fmt"
	"time"
)

type ConsoleReporter struct {
	metricStorage storage.MetricsStorage
	// executor  定时执行
}

func (c *ConsoleReporter) ConsoleReporter(metricsStorage storage.MetricsStorage) {
	c.metricStorage = metricsStorage
}
func (c *ConsoleReporter) StartRepeatedReport(periodInSeconds, durationInSeconds int) {
	durationInMillis := durationInSeconds * 1000
	endTimeInMills := int(time.Now().Unix())
	startTimeInMillis := endTimeInMills - durationInMillis
	requestInfos := c.metricStorage.GetRequestMapInfos(startTimeInMillis, endTimeInMills)
	stats := map[string]aggregator.RequestStat{}
	aggre := &aggregator.Aggregator{}
	for apiName, requestInfoList := range requestInfos {
		stats[apiName] = aggre.Aggregate(requestInfoList, durationInMillis)
	}
	statsStr, err := jsonUtil.MarshalToString(stats)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(statsStr)
}

type EmailSender struct {
}
type EmailReporter struct {
	metricsStorage storage.MetricsStorage
	emailSender    EmailSender
	toAddress      []string
}

func (e *EmailReporter) EmailReporter(metricsStorage storage.MetricsStorage, sender EmailSender) {
	e.metricsStorage = metricsStorage
	e.emailSender = sender
}

func (e *EmailReporter) AddToAddress(address string) {
	e.toAddress = append(e.toAddress, address)
}
func (e *EmailReporter)StartDeilyReport()  {

}
