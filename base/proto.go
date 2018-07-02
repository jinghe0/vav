package base

const (
	RTP_TYPE_VIDEOI uint8 = 0
	RTP_TYPE_VIDEOP uint8 = 1
	RTP_TYPE_VIDEOB uint8 = 2
	RTP_TYPE_AUDIO  uint8 = 3
	RTP_TYPE_RAW    uint8 = 4

	RTP_SEGMENT_COMPLETE uint8 = 0
	RTP_SEGMENT_FIRST    uint8 = 1
	RTP_SEGMENT_LAST     uint8 = 2
	RTP_SEGMENT_MID      uint8 = 3
)

type RTP struct {
	Serial             uint16
	SIM                string
	LogicalChannel     string
	Type               uint8
	Segment            uint8
	Timestamp          uint8
	LastIFrameInterval uint16
	LastFrameInterval  uint16
	Data               []byte
}
