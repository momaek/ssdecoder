package ssrdecoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	result, err := ssParser("Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTpUc2VobVBjcXA2VTlPemFkR3J5cQ@venus.cn.v4.hdsky.org:28000#%F0%9F%8F%B7%E9%82%80%E8%AF%B7%E6%9C%8B%E5%8F%8B%E8%8E%B7%E5%BE%9720%25%E4%BD%A3%E9%87%91%E5%A5%96%E5%8A%B1")
	assert.Nil(t, err)
	assert.Equal(t, "venus.cn.v4.hdsky.org", result.Server)
	assert.Equal(t, 28000, result.ServerPort)
	assert.Equal(t, "TsehmPcqp6U9OzadGryq", result.Password)
	assert.Equal(t, "chacha20-ietf-poly1305", result.Method)

	result, err = ssParser("YWVzLTEyOC1jdHI6dmllbmNvZGluZy5jb21AMTUyLjg5LjIwOC4xNDY6MjMzMw")
	assert.Nil(t, err)
	assert.Equal(t, "152.89.208.146", result.Server)
	assert.Equal(t, 2333, result.ServerPort)
	assert.Equal(t, "viencoding.com", result.Password)

}
