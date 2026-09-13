package ffmpegArgs

type Video struct {
	Crf       string // -crf value is the Constant Rate Factor for controlling quality (e.g., "23" for good quality)
	Codec     string // -c:v codec is the video codec to use (e.g., "libx264", "libvpx", etc.)
	Aspect    string // -aspect value is the aspect ratio for the video stream (e.g., "16:9" for widescreen)
	Bitrate   string // -b:v value is the target bitrate for the video stream (e.g., "1000k" for 1000 kbps)
	Framerate string // -r value is the frame rate for the video stream (e.g., "30" for 30 fps)
	FrameSize string // -s value is the frame size for the video stream (e.g., "1920x1080" for 1080p)
}

func (st Video) Args() Args {
	args := make(Args, 0)

	args.Arg("-crf", st.Crf)
	args.Arg("-c:v", st.Codec)
	args.Arg("-aspect", st.Aspect)
	args.Arg("-b:v", st.Bitrate)
	args.Arg("-r", st.Framerate)
	args.Arg("-s", st.FrameSize)

	if len(args) == 0 {
		return nil
	}

	return args
}
