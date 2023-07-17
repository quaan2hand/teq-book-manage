package user

import (
	"fmt"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"

	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/presenter"
)

// GetList Users
// @Summary Get an User
// @Description Get User by id
// @Tags User
// @Accept json
// @Produce json
// @Security AuthToken
// @Success 200 {object} presenter.ListUserResponseWrapper
// @Router /Users [get] .
func (r *Route) GetAll(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.GetAllUserRequest{}
		resp *presenter.ListUserResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	fmt.Println("Unscoped: ", req.Unscoped)
	fmt.Println("Un: ", req.Un)
	resp, err := r.UseCase.User.GetAll(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
