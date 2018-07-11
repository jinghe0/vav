package socket_server

import (
	"github.com/gansidui/gotcp"
	"github.com/giskook/vav/protocol"
	"log"
	//"runtime/debug"
	"sync/atomic"
)

func (ss *SocketServer) OnConnect(c *gotcp.Conn) bool {
	connection := NewConnection(c, &ConnConf{
		read_limit: ss.conf.TCP.ReadLimit,
		uuid:       atomic.AddUint32(&ss.conn_uuid, 1),
	})

	c.PutExtraData(connection)
	log.Printf("<CNT> %x \n", c.GetRawConn())

	return true
}

func (ss *SocketServer) OnClose(c *gotcp.Conn) {
	connection := c.GetExtraData().(*Connection)
	ss.cm.Del(connection)
	connection.Close()
	log.Printf("<DIS> %x\n", c.GetRawConn())
	//debug.PrintStack()
}

func (ss *SocketServer) OnMessage(c *gotcp.Conn, p gotcp.Packet) bool {
	connection := c.GetExtraData().(*Connection)
	connection.SetReadDeadline()
	connection.RecvBuffer.Write(p.Serialize())
	for {
		protocol_id, protocol_length := protocol.CheckProtocol(connection.RecvBuffer)
		buf := make([]byte, protocol_length)
		connection.RecvBuffer.Read(buf)
		switch protocol_id {
		case protocol.PROTOCOL_HALF_PACK:
			return true
		case protocol.PROTOCOL_ILLEGAL:
			return true
		case protocol.PROTOCOL_RTP:
			rtp := protocol.Parse(buf)
			log.Printf("<INF> sim %s chan %s type %x, seg %d  len %d timestamp %d lastI %d lastF %d \n", rtp.SIM, rtp.LogicalChannel, rtp.Type, rtp.Segment, len(rtp.Data), rtp.Timestamp, rtp.LastIFrameInterval, rtp.LastFrameInterval)
			if rtp.Segment > base.RTP_SEGMENT_COMPLETE {
				connection.Buffer = append(connection.Buffer, rtp.Data...)
			}
			if rtp.Segment == base.RTP_SEGMENT_LAST {
				ss.ChanFrame <- &base.Frame{
					SIM:                rtp.SIM,
					LogicalChannel:     rtp.LogicalChannel,
					Type:               rtp.Type,
					Timestamp:          rtp.Timestamp,
					LastIFrameInterval: rtp.LastIFrameInterval,
					LastFrameInterval:  rtp.LastFrameInterval,
					Data:               connection.Buffer,
				}
				connection.Buffer = nil
			} else if rtp.Segment == base.RTP_SEGMENT_COMPLETE {
				ss.ChanFrame <- &base.Frame{
					SIM:                rtp.SIM,
					LogicalChannel:     rtp.LogicalChannel,
					Type:               rtp.Type,
					Timestamp:          rtp.Timestamp,
					LastIFrameInterval: rtp.LastIFrameInterval,
					LastFrameInterval:  rtp.LastFrameInterval,
					Data:               rtp.Data,
				}
			}
		}
	}
}
