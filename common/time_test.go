package common

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTimestampToTimeSuccess(t *testing.T) {
	ts := timestamppb.Now()
	tm := TimestampToTime(ts).UTC()
	tmNow := time.Now().UTC()

	assert.Equal(t, tm.Day(), tmNow.Day())
	assert.Equal(t, tm.Month(), tmNow.Month())
	assert.Equal(t, tm.Year(), tmNow.Year())

	assert.Equal(t, tm.Hour(), tmNow.Hour())
	assert.Equal(t, tm.Minute(), tmNow.Minute())
	assert.Equal(t, tm.Second(), tmNow.Second())
}

func TestTimestampToTimeDifferentDateError(t *testing.T) {
	ts := timestamppb.Now()
	tm := TimestampToTime(ts).UTC()
	tmNow := time.Now().AddDate(2, 2, 2).UTC()

	assert.NotEqual(t, tm.Day(), tmNow.Day())
	assert.NotEqual(t, tm.Month(), tmNow.Month())
	assert.NotEqual(t, tm.Year(), tmNow.Year())

	assert.Equal(t, tm.Hour(), tmNow.Hour())
	assert.Equal(t, tm.Minute(), tmNow.Minute())
	assert.Equal(t, tm.Second(), tmNow.Second())
}

func TestTimestampToTimeDifferentTimezoneError(t *testing.T) {
	ts := timestamppb.Now()
	tm := TimestampToTime(ts).UTC()
	tmNow := time.Now()

	assert.NotEqual(t, tm.Hour(), tmNow.Hour())
	assert.Equal(t, tm.Minute(), tmNow.Minute())
	assert.Equal(t, tm.Second(), tmNow.Second())
}

func TestTimeToTimestampSuccess(t *testing.T) {
	tm := time.Now().UTC()
	ts := TimeToTimestamp(tm)
	tsAsTime := TimestampToTime(ts).UTC()

	assert.Equal(t, tm.Day(), tsAsTime.Day())
	assert.Equal(t, tm.Month(), tsAsTime.Month())
	assert.Equal(t, tm.Year(), tsAsTime.Year())
	assert.Equal(t, tm.Hour(), tsAsTime.Hour())
	assert.Equal(t, tm.Minute(), tsAsTime.Minute())
	assert.Equal(t, tm.Second(), tsAsTime.Second())
	assert.Equal(t, tm.Nanosecond(), tsAsTime.Nanosecond())
}
