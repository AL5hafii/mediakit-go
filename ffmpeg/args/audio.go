package ffmpegArgs

type Audio struct {
	Codec      string // -c:a codec is the audio codec to use (e.g., "aac", "mp3", etc.)
	Channels   string // -ac value is the number of audio channels (e.g., "2" for stereo)
	SampleRate string // -ar value is the audio sample rate (e.g., "44100" for 44.1 kHz)
}

func (st Audio) Args() Args {
	args := make(Args, 0)

	args.Arg("-c:a", st.Codec)
	args.Arg("-ac", st.Channels)
	args.Arg("-ar", st.SampleRate)

	if len(args) == 0 {
		return nil
	}

	return args
}
