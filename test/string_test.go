package test

import (
	"hello-go/util/stringUtil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert.True(t, true, stringUtil.IsEmpty(""))
	assert.True(t, true, stringUtil.IsNotEmpty("a"))
}

func TestBlank(t *testing.T) {
	assert.True(t, true, stringUtil.IsBlank(""))
	assert.True(t, true, stringUtil.IsBlank("   "))
	assert.True(t, true, stringUtil.IsNotBlank("a"))
	assert.True(t, true, stringUtil.IsNotBlank(" a "))
}

func TestGetInt(t *testing.T) {
	assert.Equal(t, 1, stringUtil.GetInt("1", -1))
	assert.Equal(t, -1, stringUtil.GetInt("a", -1))
	assert.Equal(t, -1, stringUtil.GetInt("", -1))
}

func TestGetInt32(t *testing.T) {
	assert.Equal(t, int32(1), stringUtil.GetInt32("1", -1))
	assert.Equal(t, int32(-1), stringUtil.GetInt32("a", -1))
	assert.Equal(t, int32(-1), stringUtil.GetInt32("", -1))
}

func TestGetInt64(t *testing.T) {
	assert.Equal(t, int64(1), stringUtil.GetInt64("1", -1))
	assert.Equal(t, int64(-1), stringUtil.GetInt64("a", -1))
	assert.Equal(t, int64(-1), stringUtil.GetInt64("", -1))
}

func TestGetFloat32(t *testing.T) {
	assert.Equal(t, float32(1.0), stringUtil.GetFloat32("1.0", -1))
	assert.Equal(t, float32(1.0), stringUtil.GetFloat32("1", -1))
	assert.Equal(t, float32(-1.0), stringUtil.GetFloat32("a", -1))
	assert.Equal(t, float32(-1.0), stringUtil.GetFloat32("", -1))
}

func TestGetFloat64(t *testing.T) {
	assert.Equal(t, float64(1.0), stringUtil.GetFloat64("1.0", -1))
	assert.Equal(t, float64(1.0), stringUtil.GetFloat64("1", -1))
	assert.Equal(t, float64(-1.0), stringUtil.GetFloat64("a", -1))
	assert.Equal(t, float64(-1.0), stringUtil.GetFloat64("", -1))
}

func TestIntToString(t *testing.T) {
	assert.Equal(t, "1", stringUtil.FormatInt(1))
	assert.Equal(t, "0", stringUtil.FormatInt(0))
	assert.Equal(t, "-1", stringUtil.FormatInt(-1))
}

func TestInt32ToString(t *testing.T) {
	assert.Equal(t, "1", stringUtil.FormatInt32(1))
	assert.Equal(t, "0", stringUtil.FormatInt32(0))
	assert.Equal(t, "-1", stringUtil.FormatInt32(-1))
}

func TestInt64ToString(t *testing.T) {
	assert.Equal(t, "1", stringUtil.FormatInt64(1))
	assert.Equal(t, "0", stringUtil.FormatInt64(0))
	assert.Equal(t, "-1", stringUtil.FormatInt64(-1))
}

func TestFloat32ToString(t *testing.T) {
	assert.Equal(t, "10.1002", stringUtil.FormatFloat32(10.1002))
	assert.Equal(t, "10.01", stringUtil.FormatFloat32(10.01))
	assert.Equal(t, "-10.1", stringUtil.FormatFloat32(-10.100))
}

func TestFloat64ToString(t *testing.T) {
	assert.Equal(t, "10.1002", stringUtil.FormatFloat64(10.1002))
	assert.Equal(t, "10.01", stringUtil.FormatFloat64(10.01))
	assert.Equal(t, "-10.1", stringUtil.FormatFloat64(-10.100))
}

func TestBoolToString(t *testing.T) {
	assert.Equal(t, "true", stringUtil.FormatBool(true))
	assert.Equal(t, "false", stringUtil.FormatBool(false))
}
