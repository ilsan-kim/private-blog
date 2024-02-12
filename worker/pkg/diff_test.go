package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostDiffItem_EqualTo(t *testing.T) {
	testSet := []struct {
		name      string
		fileAPath string
		fileBPath string
		expect    bool
	}{
		{
			name:      "not equal",
			fileAPath: "../../md/posts/1.md",
			fileBPath: "../../md/posts/2.md",
			expect:    false,
		},
		{
			name:      "equal",
			fileAPath: "../../md/posts/1.md",
			fileBPath: "../../md/posts/1.md",
			expect:    true,
		},
		{
			name:      "same content, but different meta info -> not equal",
			fileAPath: "../../md/posts/1.md",
			fileBPath: "../../md/posts/1-2.md",
			expect:    false,
		},
	}

	for _, tc := range testSet {
		t.Run(tc.name, func(t *testing.T) {
			fileA, errA := NewDiffItem("CONTENT", tc.fileAPath)
			fileB, errB := NewDiffItem("CONTENT", tc.fileBPath)

			assert.NoError(t, errA)
			assert.NoError(t, errB)

			res := fileA.EqualTo(fileB)
			assert.Equal(t, tc.expect, res)
		})
	}
}
