package gmf
import "log"
import "unsafe"

type Multiplexer struct{
  Ds DataSink;
  tracks []Track
  ch chan Packet
  stream_count int
}

func(self * Multiplexer)AddTrack(enc * Encoder)*Track{
    if(self.ch==nil){
	self.ch=make(chan Packet)
        self.tracks=make([]Track, 5)
    }

    result:=Track{av_new_stream(self.Ds.ctx,self.stream_count),self.ch,0}
    enc.stream_index=self.stream_count
    enc.Track=&result
    result.codec=enc.Ctx.ctx
    result.time_base=enc.Ctx.ctx.time_base
    result.sample_aspect_ratio=enc.Ctx.ctx.sample_aspect_ratio
    //result.time_base.num=1//enc.Ctx.ctx.time_base.den
    //result.time_base.den=25//enc.Ctx.ctx.time_base.num
    log.Printf("TrackData TimeBase %d/%d",result.time_base.num,result.time_base.den)

    log.Printf("TrackData %s",result)
    self.tracks[self.stream_count]=result
    self.stream_count++
    return &result
    //stream.codec=track.codec
//    track.stream=self.ch
}

func(self * Multiplexer)Start(){
    av_write_header(self.Ds.ctx)
    dump_format(self.Ds.ctx)
    for i:=0;i<self.stream_count;i++ {
        log.Printf("Track %s",self.tracks[i].String())
    }
    for (true) {
	if(closed(self.ch)){
	    println("channel closed")
	    break
	}
	var p Packet=<-self.ch
	//println(self.Ds.ctx.ctx.preload)
        if(p.Size==0){
	    log.Printf("0 Size packet in multiplexer received for stream%d:",p.Stream)
	}
        if(p.avpacket==nil){
	    println("nil packet in multiplexer received")
	    continue
	}
	//if(p.Stream!=0&&p.Stream!=1){
	    //log.Printf("Stream not accepted")
	    //continue
	//}
        stream:=self.tracks[p.Stream]
        //log.Printf("Before:"+p.String())
        p.avpacket.size=(_Ctype_int)(p.Size)
	p.avpacket.data=(*_Ctypedef_uint8_t)(unsafe.Pointer(&p.Data[0]))
        p.avpacket.pts=(_Ctypedef_int64_t)(p.Pts.RescaleTo(Rational{int(stream.time_base.num),int(stream.time_base.den)}).Time)
        p.avpacket.duration=(_Ctype_int)(p.Duration.RescaleTo(Rational{int(stream.time_base.num),int(stream.time_base.den)}).Time)
        p.avpacket.flags=_Ctype_int(p.Flags)
        p.avpacket.stream_index=_Ctype_int(p.Stream)
        p.avpacket.dts=_Ctypedef_int64_t(AV_NOPTS_VALUE)
                //log.Printf("Multiplexer %s",p.String())
        //log.Printf("After:"+p.String())

	//println("try writing frame")
        if(p.avpacket.data==nil){
	    println("nil packet.data in multiplexer received")
	    continue
	}
	
	result:=av_interleaved_write_frame(self.Ds.ctx,&p)
	if(result!=0){
	    log.Printf("failed write packet to stream")
    	    log.Printf("Packet:"+p.String())
	}

	p.Free()
	//println("frame written")
    }
    log.Printf("Multiplexer End")
}

func(self * Multiplexer)Stop(){
    close(self.ch)
    log.Printf("Writing Trailer")
    av_write_trailer(self.Ds.ctx);
}

func NewMultiplexer(sink *DataSink)*Multiplexer{
    return &Multiplexer{Ds:*sink}
}