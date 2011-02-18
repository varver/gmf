package gmf


type Format interface{}


//Encapsulates format information for video data. 
//The attributes of a VideoFormat include the encoding type, frame size, frame rate, and the data type. 
type VideoFormat struct {
	//Real base framerate of the stream. .
	FrameRate Rational
	//This is the fundamental unit of time (in seconds) in terms of which frame timestamps are represented. 
	TimeBase Rational
	//A Dimension that specifies the frame width.
	Width int
	//A Dimension that specifies the frame height.
	Height int
}
//Encapsulates format information for audio data. 
//The attributes of an AudioFormat include the sample rate, bits per sample, and number of channels. 
type AudioFormat struct {
	//The number of channels as an integer.
	Channels int
	//The frame size of this AudioFormat in bits.
	FrameSize int
	//The sample rate.
	SampleRate int
	//The sample size in bits.
	SampleSize int
}
