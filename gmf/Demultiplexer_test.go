package gmf


import "testing"


func TestDemultiplex(t * testing.T){
println("start func TestDemultiplex(t * testing.T){")
    //loc:=MediaLocator{Filename:"/media/video/ChocolateFactory.ts"}
    loc:=MediaLocator{Filename:"../../../target/dependency/fixtures/testfile.flv"}
    source:=DataSource{Locator:loc}
    if(!source.Connect()){
    	t.Errorf("cold not open file : %s", loc.Filename)
    }
    _=Demultiplexer{Ds:source}
    //println(plexer.GetDuration().String())
    //println(plexer.GetTimestamp().String())
    source.Disconnect()
println("end func TestDemultiplex(t * testing.T){")
}



func TestDemultiplexTracks(t * testing.T){
    println("start func TestDemultiplexTracks(t * testing.T){")
    //loc:=MediaLocator{Filename:"/media/video/ChocolateFactory.ts"}
    loc:=MediaLocator{Filename:"../../../target/dependency/fixtures/testfile.flv"}
    source:=DataSource{Locator:loc}
    if(!source.Connect()){
    	t.Fatalf("cold not open file : %s", loc.Filename)
    }
    plex:=Demultiplexer{Ds:source}
    tracks:=plex.GetTracks()
    println(len(*tracks))
    //println(plexer.GetTimestamp().String())
    source.Disconnect()
    println("end func TestDemultiplexTracks(t * testing.T){")
}
