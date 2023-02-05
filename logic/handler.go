package logic

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func QueryAPI(logger *zap.Logger, storage Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		tid := c.QueryParam("tid")
		if len(tid) == 0 {
			return c.JSON(http.StatusBadRequest, Resp[*string]{
				Code: http.StatusBadRequest,
				Msg:  "require parameter tid",
			})
		}
		value, err := storage.Get(c.Request().Context(), tid)
		if err != nil {
			logger.Error("get failed", zap.Error(err))
			return c.JSON(http.StatusNotFound, Resp[*string]{
				Code: http.StatusNotFound,
				Msg:  "not found",
			})
		}
		return c.JSON(http.StatusOK, Resp[string]{
			Code: 0,
			Msg:  "success",
			Data: value,
		})
	}
}

func ShareAPI(logger *zap.Logger, storage Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		content := c.FormValue("content")
		if len(content) == 0 {
			return c.JSON(http.StatusBadRequest, Resp[*string]{
				Code: http.StatusBadRequest,
				Msg:  "require parameter content",
			})
		}
		key, err := storage.Put(c.Request().Context(), content, time.Minute)
		if err != nil {
			logger.Error("put failed", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, Resp[*string]{
				Code: http.StatusInternalServerError,
				Msg:  "fail",
			})
		}
		return c.JSON(http.StatusOK, Resp[string]{
			Code: 0,
			Msg:  "success",
			Data: key,
		})
	}
}
