package CommonStruct

type Article struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ID struct {
	Id uint64 `json:"id"`
}

type GetArticlesReply struct {
	Status  uint      `json:"status"`
	Message string    `json:"message"`
	Data    []Article `json:"data"`
}

type CreateArticlesReply struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Data    ID     `json:"data"`
}

type CreateArticlesRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

const (
	SuccessStatus = "Success"
)
