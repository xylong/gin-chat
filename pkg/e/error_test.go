package e

import "testing"

func TestErrCode_String(t *testing.T) {
	t.Log(OK.String(), WsUpgrade.String())
}
