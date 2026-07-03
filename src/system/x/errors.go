package x

const (
	CodeServiceBootFailed ErrorCode = "SERVICE_BOOT_FAILED"
	CodeTCPStartFailed    ErrorCode = "TCP_START_FAILED"
	CodeConfigLoadFailed  ErrorCode = "CONFIG_LOAD_FAILED"
	CodeConfigNotFound    ErrorCode = "CONFIG_NOT_FOUND"
)

var (
	ErrServiceBootFailed = NewFolionError(CodeServiceBootFailed, "Failed to boot Folion service", true)
	ErrServerStartFailed = NewFolionError(CodeTCPStartFailed, "Failed to start server", true)
	ErrConfigLoadFailed  = NewFolionError(CodeConfigLoadFailed, "Failed to load configuration file", true)
	ErrConfigNotFound    = NewFolionError(CodeConfigNotFound, "Config file not found", false)
)
