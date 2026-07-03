package x

const (
	CodeServiceBootFailed    ErrorCode = "SERVICE_BOOT_FAILED"
	CodeTCPStartFailed       ErrorCode = "TCP_START_FAILED"
	CodeConfigLoadFailed     ErrorCode = "CONFIG_LOAD_FAILED"
	CodeConfigNotFound       ErrorCode = "CONFIG_NOT_FOUND"
	CodeSocketWriteFailed    ErrorCode = "SOCKET_WRITE_FAILED"
	CodeSessionRespondFailed ErrorCode = "SESSION_RESPOND_FAILED"
	CodeRequestDecodeFailed  ErrorCode = "REQUEST_DECODE_FAILED"
)

var (
	ErrServiceBootFailed    = NewFolionError(CodeServiceBootFailed, "Failed to boot Folion service", true)
	ErrServerStartFailed    = NewFolionError(CodeTCPStartFailed, "Failed to start server", true)
	ErrConfigLoadFailed     = NewFolionError(CodeConfigLoadFailed, "Failed to load configuration file", true)
	ErrConfigNotFound       = NewFolionError(CodeConfigNotFound, "Config file not found", false)
	ErrSocketWriteFailed    = NewFolionError(CodeSocketWriteFailed, "Failed to write data to socket", false)
	ErrSessionRespondFailed = NewFolionError(CodeSessionRespondFailed, "Failed to respond to a session", false)
	ErrRequestDecodeFailed  = NewFolionError(CodeRequestDecodeFailed, "Failed to decode request", false)
)
