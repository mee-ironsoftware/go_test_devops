package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("Meee")
	assert.Equal(t, "Hello Meee. Welcome!", greeting)

	another_greeting := starter.SayHello("Fern")
	assert.Equal(t, "Hello Fern. Welcome!", another_greeting)
}

func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Even Numbers", func(t *testing.T) {
		assert.Equal(t, "50 is an even number", starter.OddOrEven(50))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("Check Negative Even Numbers", func(t *testing.T) {
		assert.Equal(t, "-50 is an even number", starter.OddOrEven(-50))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
		assert.Equal(t, "-100 is an even number", starter.OddOrEven(-100))
	})
	t.Run("Check Non Negative Odd Numbers", func(t *testing.T) {
		assert.Equal(t, "3 is an odd number", starter.OddOrEven((3)))
		assert.Equal(t, "55 is an odd number", starter.OddOrEven(55))
		assert.Equal(t, "101 is an odd number", starter.OddOrEven(101))
	})
	t.Run("Check Negative Odd Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-1 is an odd number", starter.OddOrEven(-1))
		assert.Equal(t, "-101 is an odd number", starter.OddOrEven(-101))
	})
}

func TestCheckHealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		starter.CheckHealth(writer, req)
		response := writer.Result()
		body, _ := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
	})
}
