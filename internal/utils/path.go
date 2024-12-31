package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetStringPathParam(r *http.Request, paramName string) (string, error) {
	param := r.PathValue(paramName)

	if param == "" {
		return "", fmt.Errorf("parameter %s is empty", paramName)
	}

	return param, nil
}

func GetNumberPathParam(r *http.Request, paramName string) (uint64, error) {
	paramStr := r.PathValue(paramName)

	if paramStr != "" {
		return 0, fmt.Errorf("parameter %s is empty", paramName)
	}

	paramNum, err := strconv.ParseUint(paramStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number parameter for '%s'", paramName)
	}

	return paramNum, nil
}

func GetStringQueryParam(r *http.Request, paramName string) (string, error) {
	queryValues := r.URL.Query()
	
	param := queryValues.Get(paramName)
	
	if param == "" {
		return "", fmt.Errorf("query parameter %s is empty", paramName)
	}
	
	return param, nil
}

func GetNumberQueryParam(r *http.Request, paramName string, paramType string, errPointer *error) (uint64, error) {
	queryValues := r.URL.Query()

	paramStr := queryValues.Get(paramName)

	if paramStr == "" {
		return 0, fmt.Errorf("query parameter %s is empty", paramName)

	}

	paramNum, err := strconv.ParseUint(paramStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number parameter for '%s'", paramName)
	}
	return paramNum, nil
}