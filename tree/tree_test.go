package tree_test

import (
	"testing"

	. "github.com/imrenagi/pp-challenge/tree"
	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	nd := NewNode(0, nil, nil)
	assert.NotNil(t, nd)
}
