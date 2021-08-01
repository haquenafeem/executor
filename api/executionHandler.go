package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/executor/run"
)

type ExecutionRequest struct {
	Command string   `json:"command" binding:"required"`
	Args    []string `json:"args"`
}

type ExecutionResponse struct {
	ErrorMessage string `json:"error_message"`
	IsError      bool   `json:"is_error"`
	Output       string `json:"output"`
}

func executionHandler(c *gin.Context) {
	var req ExecutionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, ExecutionResponse{
			IsError:      true,
			ErrorMessage: "Error in data binding",
		})

		return
	}

	output, err := run.Execute(req.Command, req.Args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ExecutionResponse{
			IsError:      true,
			ErrorMessage: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, ExecutionResponse{
		IsError: false,
		Output:  output,
	})

	return
}
