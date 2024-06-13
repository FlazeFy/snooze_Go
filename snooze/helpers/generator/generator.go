package generator

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"snooze/helpers/typography"
	"time"
)

func GenerateQueryMsg(ctx string, total int) string {
	ctx = typography.UcFirst(ctx)

	if total > 0 {
		return ctx + " found"
	} else {
		return ctx + " not found"
	}
}

func GenerateCommandMsg(ctx, cmd string, total int) string {
	ctx = typography.UcFirst(ctx)

	if total > 0 {
		return ctx + " " + cmd
	} else {
		return ctx + " failed to " + cmd
	}
}

func GenerateValidatorMsg(ctx string, min, max int) string {
	if ctx != "Valid until" {
		if min != 0 && max != 0 {
			return fmt.Sprintf("%s must be between %d and %d characters", ctx, min, max)
		} else if min == 0 {
			return fmt.Sprintf("%s must be below than %d characters", ctx, max)
		} else {
			return fmt.Sprintf("%s must be more than %d characters", ctx, min)
		}
	} else {
		return fmt.Sprintf("%s must be between year %d and %d", ctx, min, max)
	}
}

func GenerateTimeNow(name string) string {
	if name == "timestamp" {
		now := time.Now()
		res := now.Format("2006-01-02 15:04:05")
		return res
	}
	return ""
}

func GenerateUUID(len int) (string, error) {
	var id string
	uuid := make([]byte, 16)

	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	if len == 16 {
		id = fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	} else if len == 32 {
		id = hex.EncodeToString(uuid)
	} else {
		id = ""
	}

	return id, nil
}
