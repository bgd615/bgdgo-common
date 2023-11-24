package cmodule

type TopicsLookupBody struct {
	MgrId     string `json:"mgr_id"`
	Ip        string `json:"ip"`
	Port      string `json:"port"`
	StartTime string `json:"start_time"`
	CurCpu    string `json:"cur_cpu"`
	TotalCpu  string `json:"total_cpu"`
	CurMem    string `json:"cur_mem"`
	TotalMem  string `json:"total_mem"`
}
