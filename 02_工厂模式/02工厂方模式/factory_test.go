package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRuleConfigParse(t *testing.T) {
	_, err := rcpFactory.createParser("json")
	assert.Nil(t, err)
}
