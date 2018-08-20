package db

// users table
type (
	userSignupForm struct {
		Username    string `form:"username" json:"username" binding:"required"`
		Password    string `form:"password" json:"password" binding:"required,min=6"`
		ConfirmPass string `form:"cofirmPass" json:"cofirmPass" binding:"eqfield=Password"`
		Email       string `form:"email" json:"email" binding:"omitempty,email"`
	}
	userQueryForm struct {
		Username string `form:"username" json:"username" binding:"required"`
	}
	userDeleteForm struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required,min=6"`
	}
	userUpdateForm struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required,min=6"`
		NewPass  string `form:"newPass" json:"newPass" binding:"required,min=6,nefield=Password"`
		Email    string `form:"email" json:"email" binding:"omitempty,email"`
	}
	userLoginForm struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required,min=6"`
	}
)

type (
	contentCreateForm struct {
		Title  string `form:"title" json:"title" binding:"required"`
		Text   string `form:"text" json:"text" binding:"required"`
		Author string `"form:author" json:"author" binding:"required"`
		Status string `"form:"status" json:"status"`
	}
)
