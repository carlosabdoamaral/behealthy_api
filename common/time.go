package common

import (
	"time"

	googleProtobuf "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimestampToTime(ts *googleProtobuf.Timestamp) time.Time {
	return ts.AsTime()
}

func TimeToTimestamp(t time.Time) *googleProtobuf.Timestamp {
	return timestamppb.New(t)
}