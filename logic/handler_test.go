package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

type MockStorage struct {
	mapping    map[string]string
	raiseError bool
	key        uint64
}

func (m *MockStorage) Put(_ context.Context, value string, _ time.Duration) (key string, err error) {
	if m.raiseError {
		return "", errors.New("error")
	}
	id := strconv.FormatUint(m.key, 10)
	m.key++
	m.mapping[id] = value
	return id, nil
}

func (m *MockStorage) Get(_ context.Context, key string) (value string, err error) {
	if m.raiseError {
		return "", errors.New("error")
	}
	v, ok := m.mapping[key]
	if !ok {
		return "", errors.New("not exist")
	}
	return v, nil
}

func TestShareAPIOk(t *testing.T) {
	e, req, w := echo.New(), httptest.NewRequest(http.MethodPost, "/share", strings.NewReader("content=test")), httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	c := e.NewContext(req, w)
	storage := &MockStorage{
		mapping:    map[string]string{},
		raiseError: false,
		key:        0,
	}
	assert.NoError(t, ShareAPI(zap.L(), storage)(c))
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"code":0, "data":"0", "msg":"success"}`, w.Body.String())
}
func TestShareAPIParamError(t *testing.T) {
	e, req, w := echo.New(), httptest.NewRequest(http.MethodPost, "/share", strings.NewReader("")), httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	c := e.NewContext(req, w)
	// following request missing param content
	storage := &MockStorage{
		mapping:    map[string]string{},
		raiseError: false,
		key:        0,
	}
	assert.NoError(t, ShareAPI(zap.L(), storage)(c))
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"code":400, "msg":"require parameter content"}`, w.Body.String())
}
func TestShareAPIFail(t *testing.T) {
	e, req, w := echo.New(), httptest.NewRequest(http.MethodPost, "/share", strings.NewReader("content=test")), httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	storage := &MockStorage{
		mapping:    map[string]string{},
		raiseError: false,
		key:        0,
	}
	storage.raiseError = true
	c := e.NewContext(req, w)
	assert.NoError(t, ShareAPI(zap.L(), storage)(c))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"code":500, "msg":"fail"}`, w.Body.String())
}

func TestQueryAPIOk(t *testing.T) {
	storage := &MockStorage{
		mapping:    map[string]string{},
		raiseError: false,
		key:        0,
	}
	testVal := "a quick fox jumps over a lazy dog"
	key, _ := storage.Put(context.Background(), testVal, time.Second)
	e, req, w := echo.New(), httptest.NewRequest(http.MethodGet, fmt.Sprintf("/query?tid=%s", key), nil), httptest.NewRecorder()
	c := e.NewContext(req, w)
	assert.NoError(t, QueryAPI(zap.L(), storage)(c))
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, fmt.Sprintf(`{"code":0, "data":"%s", "msg":"success"}`, testVal), w.Body.String())
}

func TestQueryAPIFail(t *testing.T) {
	storage := &MockStorage{
		mapping:    map[string]string{},
		raiseError: false,
		key:        0,
	}
	testVal := "a quick fox jumps over a lazy dog"
	key, _ := storage.Put(context.Background(), testVal, time.Second)
	e, req, w := echo.New(), httptest.NewRequest(http.MethodGet, fmt.Sprintf("/query?tid=%s", key), nil), httptest.NewRecorder()
	c := e.NewContext(req, w)
	storage.raiseError = true
	assert.NoError(t, QueryAPI(zap.L(), storage)(c))
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, `{"code":404, "msg":"not found"}`, w.Body.String())
}

func TestQueryAPIParamError(t *testing.T) {
	storage := &MockStorage{
		mapping:    map[string]string{},
		raiseError: false,
		key:        0,
	}
	// following request do not contain the 'tid' parameter
	testVal := "a quick fox jumps over a lazy dog"
	key, _ := storage.Put(context.Background(), testVal, time.Second)
	e, req, w := echo.New(), httptest.NewRequest(http.MethodGet, fmt.Sprintf("/query?x=%s", key), nil), httptest.NewRecorder()
	c := e.NewContext(req, w)
	assert.NoError(t, QueryAPI(zap.L(), storage)(c))
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"code":400, "msg":"require parameter tid"}`, w.Body.String())
}
