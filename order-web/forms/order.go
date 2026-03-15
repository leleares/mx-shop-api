package forms

type CreateOrderForm struct {
	Address string `json:"address" form:"address" binding:"required,max=50"`
	Name    string `json:"name" form:"name" binding:"required"`
	Mobile  string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Post    string `json:"post" form:"post" binding:"max=50"`
}
