package repo_test

import (
	"testing"

	"github.com/kopia/kopia/internal/testutil"
	"github.com/kopia/kopia/repo/content"
)

func TestMain(m *testing.M) { testutil.MyTestMain(m) }

type formatSpecificTestSuite struct {
	formatVersion content.FormatVersion
}

func TestFormatV1(t *testing.T) {
	testutil.RunAllTestsWithParam(t, &formatSpecificTestSuite{content.FormatVersion1})
}

func TestFormatV2(t *testing.T) {
	testutil.RunAllTestsWithParam(t, &formatSpecificTestSuite{content.FormatVersion2})
}
