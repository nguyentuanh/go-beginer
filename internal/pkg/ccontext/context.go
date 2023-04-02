package ccontext

import (
	"context"
	"strconv"

	constant "go-template/internal/constants"
)

// GetWorkspaceID get workspace id from context
func GetWorkspaceID(ctx context.Context) string {
	value := ctx.Value(constant.XGapoWorkspaceId)
	if id, ok := value.(string); ok {
		return id
	}
	return ""
}

// GetUserID get workspace id from context
func GetUserID(ctx context.Context) string {
	value := ctx.Value(constant.XGapoUserId)
	if id, ok := value.(string); ok {
		return id
	}
	return ""
}

// GetRole get workspace id from context
func GetRole(ctx context.Context) string {
	value := ctx.Value(constant.XGapoRole)
	if id, ok := value.(string); ok {
		return id
	}
	return ""
}

// GetApiKey get workspace id from context
func GetApiKey(ctx context.Context) string {
	value := ctx.Value(constant.XGapoApiKey)
	if id, ok := value.(string); ok {
		return id
	}
	return ""
}

func GetLang(ctx context.Context) string {
	value := ctx.Value(constant.XGapoLang)
	id, ok := value.(string)

	if !ok {
		return constant.DefaultLang
	}

	if _, ok := constant.SupportLang[id]; ok {
		return id
	}
	return constant.DefaultLang
}

func GetTimezone(ctx context.Context) int {
	value := ctx.Value(constant.HeaderTimezoneOffset)
	if value == nil {
		return 0
	}
	timezoneOffset, err := strconv.Atoi(value.(string))
	if err != nil {
		return 0
	}
	return timezoneOffset
}
