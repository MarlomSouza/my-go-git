package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_when_endpoint_returns_internal_error(t *testing.T) {
	assert := assert.New(t)

	endpointWithError := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	sut := HandlerError(endpointWithError)

	sut.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func Test_HandlerError_when_endpoint_returns_ok(t *testing.T) {
	assert := assert.New(t)
	type bodyForTest struct {
		Id int
	}
	expectedObj := bodyForTest{Id: 2}
	endpointsWithDomainError := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return expectedObj, http.StatusOK, nil
	}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	sut := HandlerError(endpointsWithDomainError)

	sut.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	result := bodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &result)
	assert.Equal(expectedObj, result)
}
