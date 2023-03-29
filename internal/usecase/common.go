package usecase

import (
	"context"
	"fmt"

	"github.com/krobus00/auth-service/internal/constant"
)

func getUserIDFromCtx(ctx context.Context) string {
	ctxUserID := ctx.Value(constant.KeyUserIDCtx)

	userID := fmt.Sprintf("%v", ctxUserID)
	if userID == "" {
		return constant.GuestID
	}
	return userID
}
