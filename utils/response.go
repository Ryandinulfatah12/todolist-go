package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
    Status bool        `json:"status"`
    Data   interface{} `json:"data"`
    Error  string      `json:"error,omitempty"`
}

func SuccessResponse(data interface{}) Response {
    return Response{
        Status: true,
        Data:   data,
    }
}

func ErrorResponse(err error) Response {
    return Response{
        Status: false,
        Data:   nil,
        Error:  err.Error(),
    }
}

func SendSuccess(c *fiber.Ctx, data interface{}) error {
    return c.JSON(SuccessResponse(data))
}

func SendError(c *fiber.Ctx, err error) error {
    return c.JSON(ErrorResponse(err))
}
