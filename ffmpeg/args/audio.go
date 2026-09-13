package ffmpegArgs

type Audio struct {
	Codec,
	Channels,
	SampleRate string
}

func (st Audio) Args() Args {
	args := make(Args, 0)

	args.Arg("-c:a", st.Codec)
	args.Arg("-ac", st.Channels)
	args.Arg("-ar", st.SampleRate)

	return args
}
