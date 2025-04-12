package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error"`
}

func (r Response) Error() string {
	return r.ErrorMessage
}

func JSON(c *fiber.Ctx, status int, data interface{}, errMsg string) error {
	return c.Status(status).JSON(Response{
		Data:         data,
		ErrorMessage: errMsg,
	})
}

func OK(c *fiber.Ctx, data interface{}) error {
	c.Status(fiber.StatusOK)
	return JSON(c, fiber.StatusOK, data, "")
}

func Fail(c *fiber.Ctx, msg string, status ...int) error {
	code := fiber.StatusInternalServerError
	if len(status) > 0 {
		code = status[0]
	}

	message := "An internal server error occurred (check logs)."
	if len(msg) > 0 {
		message = msg
	}

	return JSON(c, code, nil, message)
}
