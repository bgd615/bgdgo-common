package cmodule

type TopicsWorkerBody struct {
	WorkerId     string `json:"worker_id"`
	WId          string `json:"wid"`
	Ip           string `json:"ip"`
	Port         string `json:"port"`
	StartTime    string `json:"start_time"`
	UpdateTime   string `json:"update_time"`
	CurCpu       string `json:"cur_cpu"`
	TotalCpu     string `json:"total_cpu"`
	CurMem       string `json:"cur_mem"`
	TotalMem     string `json:"total_mem"`
	MaxCount     string `json:"max_count"`
	RunningCount string `json:"running_count"`
}

type TopicsWorkerBucketBody struct {
	Id          string          `json:"id"`
	WorkerId    string          `json:"worker_id"`
	CreateTime  string          `json:"create_time"`
	ExpiredTime string          `json:"expired_time"`
	Expired     int             `json:"expired"`
	WId         string          `json:"wid"`
	Flag        int             `json:"flag"`
	Task        *TopicsJobSpool `json:"task"`
	Ip          string          `json:"ip"`
	Port        string          `json:"port"`
	NoExpired   string          `json:"no_expired"`
}

type TopicsWorkerCmdLogBody struct {
	Cmd        string `json:"cmd"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Seq        int    `json:"seq"`
	RetryTime  int    `json:"retry_time"`
	Log        string `json:"log"`
	ReturnCode string `json:"return_code"`
	LogMd5     string `json:"log_md5"`
}

type TopicsWorkerServer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Server []string `json:"server"`
	Instance string `json:"instance"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
	Ct   string `json:"ct"`
	Das  string `json:"das"`
}


type TopicsActiveWorker struct {
	ActiveRoutines int `json:"active_routines"`
}
