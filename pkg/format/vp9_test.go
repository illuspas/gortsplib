package format

import (
	"testing"

	"github.com/pion/rtp"
	"github.com/stretchr/testify/require"
)

func TestVP9Attributes(t *testing.T) {
	format := &VP9{
		PayloadTyp: 100,
	}
	require.Equal(t, "VP9", format.String())
	require.Equal(t, 90000, format.ClockRate())
	require.Equal(t, uint8(100), format.PayloadType())
	require.Equal(t, true, format.PTSEqualsDTS(&rtp.Packet{}))
}

func TestVP9MediaDescription(t *testing.T) {
	maxFR := 123
	maxFS := 456
	profileID := 789
	format := &VP9{
		PayloadTyp: 96,
		MaxFR:      &maxFR,
		MaxFS:      &maxFS,
		ProfileID:  &profileID,
	}

	rtpmap, fmtp := format.Marshal()
	require.Equal(t, "VP9/90000", rtpmap)
	require.Equal(t, "max-fr=123;max-fs=456;profile-id=789", fmtp)
}
