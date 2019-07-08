package mocktime

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	os.Exit(m.Run())
}

func TestMockTimeNow(t *testing.T) {
	a := assert.New(t)

	defer func() {
		now := timeNowFunc()
		a.Equal(uint(time.Now().Unix()), now)

		clientTime := defaultCtx.Value(clientTimeKey)
		a.Nil(clientTime)
	}()

	mockTime := uint(25)
	defer mockTimeNow(mockTime)()

	now := timeNowFunc()
	a.Equal(mockTime, now)

	clientTime := defaultCtx.Value(clientTimeKey).(int64)
	a.Equal(mockTime, uint(clientTime))

}
