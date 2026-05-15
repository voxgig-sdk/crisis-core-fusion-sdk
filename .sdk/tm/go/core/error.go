package core

type CrisisCoreFusionError struct {
	IsCrisisCoreFusionError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewCrisisCoreFusionError(code string, msg string, ctx *Context) *CrisisCoreFusionError {
	return &CrisisCoreFusionError{
		IsCrisisCoreFusionError: true,
		Sdk:              "CrisisCoreFusion",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *CrisisCoreFusionError) Error() string {
	return e.Msg
}
