package randomGen

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRandomString_Length(t *testing.T) {
	size := 10
	result := NewRandomString(size)

	// проверяем, что длина строки соответствует ожидаемой
	assert.Equal(t, size, len(result), "Generated string length should be %d", size)
}

func TestNewRandomString_ValidCharacters(t *testing.T) {
	size := 100
	result := NewRandomString(size)

	// Допустимые символы
	allowedChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

	// проверяем, что все символы в строке допустимы
	for _, char := range result {
		assert.Contains(t, allowedChars, string(char), "Character %v is not allowed", string(char))
	}
}

func TestNewRandomString_Uniqueness(t *testing.T) {
	// Тестируем уникальность строк
	size := 10
	result1 := NewRandomString(size)

	time.Sleep(1 * time.Millisecond)

	result2 := NewRandomString(size)

	// Проверяем, что строки разные
	assert.NotEqual(t, result1, result2, "Generated strings should be different")
}

func TestNewRandomString_SpecialCases(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"Size 1", 1},
		{"Size 100", 100},
		{"Size 1000", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewRandomString(tt.size)
			assert.Equal(t, tt.size, len(result), "Generated string length should be %d", tt.size)
		})
	}
}
