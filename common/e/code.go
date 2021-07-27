package e

/**
 * 自定内内部错误好，统一异常信息处理
 */
const (
	ScodeOK        = "ScodeOK"
	StatusAccepted = "StatusAccepted"

	InvalidParams      = "InvalidParams"
	IllegalRequest     = "IllegalRequest"
	StatusUnauthorized = "StatusUnauthorized"
	StatusForbidden    = "StatusForbidden"
	StatusNotFound     = "StatusNotFound"
	StatusConflict     = "StatusConflict"

	InternalError            = "InternalError"
	StatusBadGateway         = "StatusBadGateway"
	StatusServiceUnavailable = "StatusServiceUnavailable"
	StatusGatewayTimeout     = "StatusGatewayTimeout"

)
