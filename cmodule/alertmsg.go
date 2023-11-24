package cmodule

type AlertMsg struct {
	Rid          string `json:"rid"`
	Fid          string `json:"fid"`
	Sys          string `json:"sys"`
	Job          string `json:"job"`
	Ctx          string `json:"ctx"`
	Status       string `json:"status"`
	AlertRule    string `json:"alert_rule"`
	AlertContact string `json:"alert_contact"`
	AlertMsg     string `json:"alert_msg"`
	CreateAt     string `json:"create_at"`
}

type RetBean struct {
	Code   string
	Status string
	Err    string
	Data   interface{}
}
