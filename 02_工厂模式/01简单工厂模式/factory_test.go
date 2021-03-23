package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRuleConfigParse(t *testing.T) {
	_, err := Load1("json")
	assert.Nil(t, err)
	_, err = Load2("json")
	assert.Nil(t, err)
}

