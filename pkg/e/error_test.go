package e

import (
	"testing"
)

func TestCustomError_Error(t *testing.T) {
	t.Log(NewCustomError(WsParseMsg), WsOfflineReply.String())
}
