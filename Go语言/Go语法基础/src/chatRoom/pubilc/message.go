package pubilc


const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
)






type Message struct {

	Type string `json:"type"`// 消息类型
	Data string  `json:"data"`// 消息的类型
}

// 定义2个消息
type LoginMes struct {
	UserId int  `json:"userid"`
	Password string `json:"password"`
	Name string `json:"name"`
}

type LoginResMes struct {
	Code int `json:"code"`   // 状态码
	Error string `json:"error"` // c错误信息
}