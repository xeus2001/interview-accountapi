package f3_test

import (
	"github.com/xeus2001/interview-accountapi/src/f3"
	"testing"
)

func TestAccountAttr_SetStatus(t *testing.T) {
	attr := f3.AccountAttr{}
	err := attr.SetStatus(f3.StatusClosed, "must be ignored")
	if err != nil {
		t.Errorf("Setting StatusClosed failed, but must not, reason: %s", err.Error())
	} else {
		if attr.Status == nil || *attr.Status != f3.StatusClosed {
			t.Error("Settings StatusClosed failed")
		}
		if attr.StatusReason != nil {
			t.Error("The status reason is set, but must not be")
		}
	}

	err = attr.SetStatus(f3.StatusFailed, "Required")
	if err != nil {
		t.Errorf("Setting StatusFailed failed, but must not, reason: %s", err.Error())
	} else {
		if attr.Status == nil || *attr.Status != f3.StatusFailed {
			t.Error("Settings StatusFailed failed")
		}
		if attr.StatusReason == nil || *attr.StatusReason != "Required" {
			t.Error("The status reason is not correctly set for StatusFailed")
		}
	}

	err = attr.SetStatus(f3.StatusFailed, "")
	if err == nil {
		t.Errorf("Setting StatusFailed with empty reason, this operation must fail")
	}
}
