package protocol

const CloseSignal = "CLOSE_CONNECTION"
const (
	CodeLogin          = 1
	CodeLoginPush      = 2
	CodeTaskRequest    = 3
	CodeTaskPush       = 4
	CodeTaskFinish     = 5
	CodeTaskFinishPush = 6
	CodeTaskStopPush   = 7
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
	Token  string `json:"token"`
	TaskId string `json:"taskId"`
	Error  string `json:"error"`
}

type TaskFinishPush struct {
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
