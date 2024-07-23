package protocol

const CloseSignal = "CLOSE_CONNECTION"
const (
	CodeLogin       = 1
	CodeTaskRequest = 2
	CodeTaskPush    = 3
	CodeTaskStart   = 4
	CodeTaskFinish  = 5
	CodeTaskStop    = 6
)

type Message struct {
	Code uint        `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

type Login struct {
	Token string `json:"token"`
}
type LoginPush struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

type TaskRequest struct {
	Token string `json:"token"`
}

type TaskFinish struct {
	TaskId string `json:"taskId"`
	Error  string `json:"error"`
}

type TaskStart struct {
	TaskId string `json:"taskId"`
	Token  string `json:"token"`
}

type TaskPush struct {
	TaskId string `json:"taskId"`
	Script string `json:"script"`
	Error  string `json:"error"`
}

type TaskStopPush struct {
	TaskId string `json:"taskId"`
	Error  string `json:"error"`
}
