package metrics


type RedisMetricStorage struct {
}

func (s *RedisMetricStorage) SaverRequestInfos(requestInfo *RequestInfo){
	// TODO 业务代码

}

func (s *RedisMetricStorage)GetRequestInfos(apiname string, startTimeMillis, endTimeMillis int) []*RequestInfo {
	//  TODO 业务代码
	return nil

}

func (s *RedisMetricStorage)GetRequestMapInfos(startTimeMillis, endTimeMillis int) map[string][]*RequestInfo{
	//  TODO 业务代码
	return nil
}
