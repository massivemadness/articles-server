package v1

type ArticlesResponse struct {
	Articles []string `json:"articles"`
}

type ArticleResponse struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
