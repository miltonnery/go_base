package log

import (
	"context"
	"github.com/google/uuid"
	"os"
)

//Fill context for logs
func FillContextForLogs(ctx context.Context) context.Context {
	hostname, _ := os.Hostname()
	requestID := uuid.New().String()
	cwv := context.WithValue(ctx, "pod", hostname)
	cwv = context.WithValue(cwv, "request-id", requestID)
	return cwv
}
