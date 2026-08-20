package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask05(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	require.NoError(t, s.CheckOperatorTraining(context.Background(), compliantStore(now), "quality-v2"))
}
