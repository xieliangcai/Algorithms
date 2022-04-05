package userController

import (
	"Algorithms/serverCount/metrics"
	"time"
)

type UserController struct {
	M *metrics.Metrics
}

func NewUserController() *UserController {
	return &UserController{
		M: &metrics.Metrics{},
	}
}
func (this *UserController) Controller() {
	this.M.StartRepeatedReport(60, time.Second)
}

type User struct {

}

// 统计用户注册的响应时间
func (this *UserController) register(userVo User) {
	startTimeStamp := time.Now().Unix()
	this.M.RecordTimestamp("register", float64(startTimeStamp))
	// TODO 注册业务逻辑
	respTime := time.Now().Unix()
	this.M.RecordResponseTime("register", float64(respTime-startTimeStamp))
}

// 统计用户的访问次数
func (this *UserController) UserVologin(telephone string, password string) {
	startTimeStamp := time.Now().Unix()
	this.M.RecordTimestamp("login", float64(startTimeStamp))
	// TODO 登录业务逻辑
	respTime := time.Now().Unix()
	this.M.RecordResponseTime("login", float64(respTime-startTimeStamp))
}
