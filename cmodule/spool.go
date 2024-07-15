package cmodule

import (
	"os/exec"
)

type TopicsJobSpool struct {
	Rid               string        `json:"rid"`
	Fid               string        `json:"fid"`
	Pid               string        `json:"pid"`
	Sid               string        `json:"sid"`
	Sys               string        `json:"sys"`
	Job               string        `json:"job"`
	Ctx               string        `json:"ctx"`
	Status            string        `json:"status"`
	Server            string        `json:"server"`
	CreateTime        string        `json:"create_time"`
	StartTime         string        `json:"start_time"`
	EndTime           string        `json:"end_time"`
	UpdateTime        string        `json:"update_time"`
	WorkerServer      string        `json:"worker_server"`
	WorkerIp          string        `json:"worker_ip"`
	WorkerPort        string        `json:"worker_port"`
	Priority          string        `json:"priority"`
	Alert             string        `json:"alert"`
	TimeWindow        string        `json:"timewindow"`
	TimeOut           string        `json:"timeout"`
	JobType           string        `json:"job_type"`
	Retry             string        `json:"retry"`
	FailFuse          string        `json:"fail_fuse"`
	QueueCreateTime   string        `json:"queue_create_time"`
	QueueRcvTime      string        `json:"queue_rcv_time"`
	QueueStartTime    string        `json:"queue_start_time"`
	QueueEndTime      string        `json:"queue_end_time"`
	QueueServerName   string        `json:"queue_server_name"`
	PendingRcvTime    string        `json:"pending_rcv_time"`
	PendingStartTime  string        `json:"pending_start_time"`
	PendingEndTime    string        `json:"pending_end_time"`
	PendingServerName string        `json:"pending_server_name"`
	SubmitRcvTime     string        `json:"submit_rcv_time"`
	SubmitStartTime   string        `json:"submit_start_time"`
	SubmitEndTime     string        `json:"submit_end_time"`
	SubmitServerName  string        `json:"submit_server_name"`
	GoRcvTime         string        `json:"go_rcv_time"`
	GoStartTime       string        `json:"go_start_time"`
	GoEndTime         string        `json:"go_end_time"`
	GoServerName      string        `json:"go_server_name"`
	WorkerRcvTime     string        `json:"worker_rcv_time"`
	WorkerStartTime   string        `json:"worker_start_time"`
	WorkerEndTime     string        `json:"worker_end_time"`
	WorkerServerName  string        `json:"worker_server_name"`
	EtcdEndpoints     string        `json:"etcd_endpoints"`
	EtcdHome          string        `json:"etcd_home"`
	Cmds              []string      `json:"cmds"`
	Logs              []interface{} `json:"logs"`
	Parameters        []string      `json:"parameters"`
	ReturnCode        string        `json:"return_code"`
	ExecCmds          []string      `json:"exec_cmds"`
	Command           *exec.Cmd     `json:"command"`
	CmdRunning        string        `json:"cmdrunning"`
	Msg               string        `json:"msg"`
	Enable            string        `json:"enable"`
	WorkerReturnCode  string        `json:"worker_return_code"`
	WorkerReturnTxt   string        `json:"worker_return_txt"`
	QueueRetryCount   string        `json:"queue_retry_count"`
}
