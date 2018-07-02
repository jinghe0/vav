package protocol

import (
	"bytes"
	"github.com/giskook/vav/base"
)

func Parse(p []byte) *base.RTP {
	r := bytes.NewReader(p)
	r.Seek(6, 0)
	serial := base.ReadWord(r)
	sim := base.ReadBcdString(r, 6)
	logical_channel := base.ReadByte(r)
	type_seg := base.ReadByte(r)
	t := (type_seg & 0xf0) >> 4
	seg := (type_seg & 0x0f)
	timestamp := uint8(0)
	if t != base.RTP_TYPE_RAW {
		timestamp = base.ReadByte(r)
	}
	var last_i_frame_interval, last_frame_interval uint16 = 0, 0
	if t == base.RTP_TYPE_VIDEOI ||
		t == base.RTP_TYPE_VIDEOP ||
		t == base.RTP_TYPE_VIDEOB {
		last_i_frame_interval = base.ReadWord(r)
		last_frame_interval = base.ReadWord(r)
	}

	length := base.ReadWord(r)
	data := base.ReadBytes(r, length)

	return &base.RTP{
		Serial:             serial,
		SIM:                sim,
		LogicalChannel:     logical_channel,
		Type:               t,
		Segment:            seg,
		Timestamp:          timestamp,
		LastIFrameInterval: last_i_frame_interval,
		LastFrameInterval:  last_frame_interval,
		Data:               data,
	}
}
