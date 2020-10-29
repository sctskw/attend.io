package models

import "github.com/go-openapi/swag"

func NewSystemStatus(code int64, msg string) *SystemStatus {
	return &SystemStatus{
		Code:    code,
		Message: swag.String(msg),
	}
}
