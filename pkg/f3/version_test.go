package f3_test

import (
	"github.com/xeus2001/interview-accountapi/pkg/f3"
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	const versionRegExp string = `^([0-9]+)\.([0-9]+)\.([0-9]+)$`
	versionRegex := regexp.MustCompile(versionRegExp)
	match := versionRegex.FindStringSubmatch(f3.Version)
	if match == nil {
		t.Errorf("Version string %v does not conform to X.X.X", f3.Version)
	}

	invalidVersion := "wrong.5.1"
	match = versionRegex.FindStringSubmatch(invalidVersion)
	if match != nil {
		t.Errorf("Failed to detect invalid version in test string '%s'", invalidVersion)
	}
}
