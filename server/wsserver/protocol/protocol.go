package protocol

const CloseSignal = "CLOSE_CONNECTION"
const (
	CodeLogin          = 1
	CodeTaskRequest    = 2
	CodeTaskPush       = 3
	CodeTaskFinish     = 4
	CodeTaskFinishPush = 5
	CodeTaskStopPush   = 6
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
