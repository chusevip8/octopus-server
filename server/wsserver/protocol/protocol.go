package protocol

const CloseSignal = "CLOSE_CONNECTION"

type Message struct {
	Code uint        `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

type HeartbeatReq struct {
	Ping string `json:"ping"`
}
type HeartbeatRes struct {
	Pong string `json:"pong"`
}

type LoginReq struct {
	Token string `json:"token"`
}
type LoginRes struct {
	Code  uint   `json:"code"`
	Error string `json:"error"`
}

type TaskRequest struct {
	Device string `json:"device"`
}

type TaskFinishReq struct {
	TaskId string `json:"taskId"`
	Error  string `json:"error"`
}
type TaskFinishRes struct {
	TaskId string `json:"taskId"`
}

type TaskStartReq struct {
	TaskId string `json:"taskId"`
}
type TaskStartRes struct {
	TaskId string `json:"taskId"`
}

type TaskPush struct {
	TaskId string `json:"taskId"`
	Script string `json:"script"`
}

type TaskStop struct {
	TaskId string `json:"taskId"`
}
