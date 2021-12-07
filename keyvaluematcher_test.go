package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testNamespace = "test_name"
	testTaskType  = "test_task"
	testService   = "test_service"
)

var q = `
# test_name: demo
# test_task: demo2
# test_service: demo3`

func TestNewKeyValueMatcher(t *testing.T) {
	namespace, _ := NewKeyValueMatcher(testNamespace)
	taskType, _ := NewKeyValueMatcher(testTaskType)
	service, _ := NewKeyValueMatcher(testService)

	assert.Equal(t, "demo", namespace.Match(q))
	assert.Equal(t, "demo2", taskType.Match(q))
	assert.Equal(t, "demo3", service.Match(q))

}
