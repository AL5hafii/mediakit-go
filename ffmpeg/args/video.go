package ffmpegArgs

type Video struct {
	Codec string
}

func (st Video) Args() Args {
	args := make(Args, 0)

	args.Arg("-c:v", st.Codec)

	return args
}
