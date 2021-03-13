package helo

type HelloResponse struct {
	Name    string `json:"name"`    // 调用者的名字
	Content string `json:"content"` // 给调用者的回话
}
