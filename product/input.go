package product

type FindProductByNameInput struct {
	Keyword string `json:"keyword" binding:"required"`
}

type FindProductByIdInput struct {
	Id int `json:"id" binding:"required"`
}
