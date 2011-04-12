package gmf

import "fmt"

//A Track abstracts the information specific to an individual track in a media stream. 
//A media stream might contain multiple media tracks, such as separate tracks for audio, video, and midi data. 
//A Track is the output of a Demultiplexer. 
type Track struct {
	*Stream
	stream   chan Packet
	next_pts int64
}

func (self *Track) String() string {
	return fmt.Sprintf("Idx:%d;CTB:%d/%d;STB:%d/%d", self.index, self.codec.time_base.num, self.codec.time_base.den, self.time_base.num, self.time_base.den)
}

func (self *Track) GetFormat() Format {
	if self.codec.codec_type == CODEC_TYPE_VIDEO {
		return VideoFormat{
			Rational{int(self.r_frame_rate.num), int(self.r_frame_rate.den)},
			Rational{int(self.codec.time_base.num), int(self.codec.time_base.den)},
			int(self.codec.width),
			int(self.codec.height)}
	}
	return AudioFormat{
		int(self.codec.channels),
		int(self.codec.frame_size),
		int(self.codec.sample_rate),
		av_get_bits_per_sample_fmt(self.codec.sample_fmt) / 8}
}

func (self *Track) GetStreamIndex() int {
	return int(self.index)
}

func (self *Track) GetStartTime() Timestamp {
	return Timestamp{}
}

func (self *Track) ReadPacket(p *Packet) bool {
	*p = <-self.stream
	if p==nil {
		return false
	}
	return true
}

func (self *Track) WritePacket(p *Packet) bool {
	if self.stream == nil  {
		return false
	}
	p.Stream = int(self.index)

	if self.next_pts > 0 && p.Pts.Time != self.next_pts {
		//log.Printf("Fail: next_pts=%d incoming pts=%d", self.next_pts, p.Pts.Time)
	} else {
		self.next_pts = p.Pts.Time
	}
	self.next_pts += p.Duration.Time
	self.stream <- *p
	return true
}

func (self *Track) GetDecoder() *Decoder {
	coder := Decoder{}         //NewCoder()
	coder.Ctx.ctx = self.codec //dpx.Ds.Ctx.streams[streamid].codec
	coder.frame_rate = Rational{int(self.r_frame_rate.num), int(self.r_frame_rate.den)}
	coder.time_base = Rational{int(self.codec.time_base.num), int(self.codec.time_base.den)}
	return &coder
}
