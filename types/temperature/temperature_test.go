package temperature

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTemperature(t *testing.T) {
	temp := New(15.0)
	assert.Equal(t, "Temperature", reflect.TypeOf(temp).Name())
}
