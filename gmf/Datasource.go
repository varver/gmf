package gmf



type DataSource struct{
  loc MediaLocator;
  ctx * FormatContext
  valid bool
}


func (src * DataSource) Connect() bool{
  src.valid=false
  src.ctx=avformat_alloc_context();
  result:=av_open_input_file(src.ctx, src.loc.Filename, nil,0,nil);
  if(result!=0){
    return src.valid
  }
  result=av_find_stream_info(src.ctx)
  if(result<0){
    return src.valid
  }
  src.valid=true
  return src.valid;
}

func (src * DataSource) Disconnect() bool{
  if(src.valid){
    av_close_input_file(src.ctx);
  }
  return true;
}
/*
func (src * DataSource) GetContentType() string{
  return "video/ffmpeg";
}*/


func NewDataSource(loc MediaLocator)*DataSource{
    return &DataSource{loc:loc,ctx:nil,valid:false}
}
