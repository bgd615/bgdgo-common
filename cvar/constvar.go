package cvar

const (
	//status
	JOB_STATUS_SUBMIT  = "submit"
	JOB_STATUS_RUNNING = "running"
	JOB_STATUS_SUCC    = "succ"
	JOB_STATUS_FAIL    = "fail"
	JOB_STATUS_READY   = "ready"
	JOB_STATUS_PENDING = "pending"
	JOB_STATUS_CURRENT = "current"
	JOB_STATUS_GO      = "go"
	///
	API_EXIT_CODE_ERR    = 700
	API_EXIT_CODE_SUCC   = 200
	API_EXIT_CODE_NOFREE = 999
	//
	HEART_EXPIRED_TIME = 300

	// ENV_VAR_CTX_DATE      = "CTX_DATE"
	// ENV_VAR_CTX_TIME      = "CTX_TIME"
	// ENV_VAR_CTX_TIMESTAMP = "CTX_TIMESTAMP"
	ENV_VAR_CTX_ID  = "CTX_ID"
	EVN_VAR_CTX_STR = "CTX_STR"
	ENV_VAR_CTX_SYS = "CTX_SYS"
	ENV_VAR_CTX_JOB = "CTX_JOB"
	ENV_VAR_DATE    = "yyyymmdd"
	//
	ENV_VAR_FORMAT_CTX_ID        = "ENV_CTX_ID"
	EVN_VAR_FORMAT_CTX_STR       = "ENV_CTX_STR"
	ENV_VAR_FORMAT_CTX_SYS       = "ENV_CTX_SYS"
	ENV_VAR_FORMAT_CTX_JOB       = "ENV_CTX_JOB"
	ENV_VAR_FORMAT_CTX_DATE      = "ENV_CTX_DATE"
	ENV_VAR_FORMAT_CTX_TIME      = "ENV_CTX_TIME"
	ENV_VAR_FORMAT_CTX_TIMESTAMP = "ENV_CTX_TIMESTAMP"
	ENV_VAR_FORMAT_CTX_YEAR      = "ENV_CTX_YEAR"
	ENV_VAR_FORMAT_CTX_MONTH     = "ENV_CTX_MONTH"
	ENV_VAR_FORMAT_CTX_DAY       = "ENV_CTX_DAY"
	ENV_VAR_FORMAT_CTX_HOUR      = "ENV_CTX_HOUR"
	ENV_VAR_FORMAT_CTX_MINUTE    = "ENV_CTX_MINUTE"
	ENV_VAR_FORMAT_CTX_SECOND    = "ENV_CTX_SECOND"
	//
	ENV_VAR_FORMAT_DATE    = "yyyymmdd"
	ENV_VAR_FORMAT_TIME    = "HHMMSS"
	ENV_VAR_FORMAT_YEAR    = "yyyy"
	ENV_VAR_FORMAT_MONTH   = "mm"
	ENV_VAR_FORMAT_DAY     = "dd"
	ENV_VAR_FORMAT_HOUR    = "HH"
	ENV_VAR_FORMAT_MINUTE  = "MM"
	ENV_VAR_FORMAT_SECOND  = "SS"
	ENV_VAR_FORMAT_DATE_1D = "yyyymmdd_1d"
	ENV_VAR_FORMAT_HOUR_1H = "HH_1h"
	ENV_VAR_FORMAT_DATE_2D = "yyyymmdd_2d"
	ENV_VAR_FORMAT_HOUR_2H = "HH_2h"

	//
	VAR_ALERT_MSG_RID           = "VAR_ALERT_MSG_RID"
	VAR_ALERT_MSG_FID           = "VAR_ALERT_MSG_FID"
	VAR_ALERT_MSG_SYS           = "VAR_ALERT_MSG_SYS"
	VAR_ALERT_MSG_JOB           = "VAR_ALERT_MSG_JOB"
	VAR_ALERT_MSG_CTX           = "VAR_ALERT_MSG_CTX"
	VAR_ALERT_MSG_STATUS        = "VAR_ALERT_MSG_STATUS"
	VAR_ALERT_MSG_STARTAT       = "VAR_ALERT_MSG_STARTAT"
	VAR_ALERT_MSG_ENDAT         = "VAR_ALERT_MSG_ENDAT"
	VAR_ALERT_MSG_WORKER_SERVER = "VAR_ALERT_MSG_WORKER_SERVER"
	VAR_ALERT_MSG_WORKER_IP     = "VAR_ALERT_MSG_WORKER_IP"
	VAR_ALERT_MSG_WORKER_PORT   = "VAR_ALERT_MSG_WORKER_PORT"
)
