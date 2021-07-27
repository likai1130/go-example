package e

import "net/http"

var MsgFlags = map[string](int){

	ScodeOK:        http.StatusOK,       //正常 		200
	StatusAccepted: http.StatusAccepted, //请求成功无数据	202

	InvalidParams:      http.StatusBadRequest,   //参数异常 	400
	IllegalRequest:     http.StatusBadRequest,   //非法请求	400
	StatusUnauthorized: http.StatusUnauthorized, //认证失败	401
	StatusForbidden:    http.StatusForbidden,    //无权限，权限失效 403
	StatusNotFound:     http.StatusBadRequest,   //资源不存在	404
	StatusConflict:     http.StatusConflict,     //资源已存在	409

	InternalError:            http.StatusInternalServerError, //内部错误 	500
	StatusBadGateway:         http.StatusBadGateway,          //网关异常   502
	StatusServiceUnavailable: http.StatusServiceUnavailable,  //服务异常	503
	StatusGatewayTimeout:     http.StatusGatewayTimeout,      //超时		504

}

