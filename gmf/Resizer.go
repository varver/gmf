package gmf

import "log"
//import "unsafe"
type Resizer struct{
    ctx * SwsContext
    width int
    height int
    fmt int
//    frame Frame
//    buffer []byte
}


func (self * Resizer)Init(dec * Decoder, enc * Encoder){
    self.ctx=new(SwsContext)
    sws_scale_getcontext(self.ctx,int(dec.Ctx.ctx.width), int(dec.Ctx.ctx.height), int(dec.Ctx.ctx.pix_fmt), int(enc.Ctx.ctx.width), int(enc.Ctx.ctx.height), int(enc.Ctx.ctx.pix_fmt), 1)
    self.width=int(enc.Ctx.ctx.width)
    self.height=int(enc.Ctx.ctx.height)
    self.fmt=int(enc.Ctx.ctx.pix_fmt)
    log.Printf("setting swscale from %d/%d:%d to %d/%d:%d",int(dec.Ctx.ctx.width),int(dec.Ctx.ctx.height),int(dec.Ctx.ctx.pix_fmt),int(enc.Ctx.ctx.width),int(enc.Ctx.ctx.height),int(enc.Ctx.ctx.pix_fmt))
    /*
    numBytes:= avpicture_get_size(uint32(self.fmt), self.width, self.height)
    if(numBytes>0){
	self.buffer=av_malloc(numBytes)
	avpicture_fill(&self.frame, self.buffer, 0, self.width, self.height);
    }*/
//func sws_scale_getcontext(ctx * SwsContext, srcwidth, srcheight, srcfmt, trgwidth,trgheight,trgfmt,flags int){

}

func(self*Resizer)Resize(in* Frame)*Frame{
    frame := NewFrame(self.fmt, self.width, self.height)
    //return frame
    
    if result:=sws_scale(self.ctx, in, frame);result<=0 {
	log.Printf("failed to resize the image")
    }else{
	//log.Printf("frame result=%d", result)
    }
    //log.Printf("in frame size %d", in.size)
    //log.Printf("number bytes %d", numBytes)
    //frame.size=numBytes
    frame.Pts=in.Pts
    frame.Duration=in.Duration
    //frame.avframe.pts=in.avframe.pts
    //frame.width=self.width
    //frame.height=self.height
    //println("data pointer")
    //println(&frame.avframe.data)
    return frame
    //return in
}

func (self*Resizer)Close(){
    sws_free_context(self.ctx)
}