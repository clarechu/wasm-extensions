package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrite(t *testing.T) {
	err := Write()
	assert.Equal(t, nil, err)
}
