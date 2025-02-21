package format

import (
	"testing"

	"github.com/pion/rtp"
	"github.com/stretchr/testify/require"
)

func TestOpusAttributes(t *testing.T) {
	format := &Opus{
		PayloadTyp:   96,
		SampleRate:   48000,
		ChannelCount: 2,
	}
	require.Equal(t, "Opus", format.String())
	require.Equal(t, 48000, format.ClockRate())
	require.Equal(t, uint8(96), format.PayloadType())
	require.Equal(t, true, format.PTSEqualsDTS(&rtp.Packet{}))
}

func TestOpusMediaDescription(t *testing.T) {
	format := &Opus{
		PayloadTyp:   96,
		SampleRate:   48000,
		ChannelCount: 2,
	}

	rtpmap, fmtp := format.Marshal()
	require.Equal(t, "opus/48000/2", rtpmap)
	require.Equal(t, "sprop-stereo=1", fmtp)
}
