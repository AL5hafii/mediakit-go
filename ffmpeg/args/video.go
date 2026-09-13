package ffmpegArgs

// VIDEO struct represents the video stream parameters for ffmpeg command-line arguments
type VIDEO struct {
	Map    string   // -map value is the input stream mapping for the output file (e.g., "0:v" for the first video stream of the first input file)
	Filter []string // -vf value is the video filter to apply (e.g., "scale=1280:720" to resize to 720p)

	Crf, // -crf value is the Constant Rate Factor for controlling quality (e.g., "23" for good quality)
	Codec, // -c:v codec is the video codec to use (e.g., "libx264", "libvpx", etc.)
	Aspect, // -aspect value is the aspect ratio for the video stream (e.g., "16:9" for widescreen)
	BitRate, // -b:v value is the target bitrate for the video stream (e.g., "1000k" for 1000 kbps)
	Profile, // -profile:v value is the profile for the video stream (e.g., "high" for H.264)
	FrameRate, // -r value is the frame rate for the video stream (e.g., "30" for 30 fps)
	FrameSize, // -s value is the frame size for the video stream (e.g., "1920x1080" for 1080p)
	Pixfmt string // -pix_fmt value is the pixel format for the video stream (e.g., "yuv420p" for standard color format)
}

// Video appends the video parameters to the Args
// filter must be applied before map and map before other parameters
func (a *ARGS) Video(vid VIDEO) {
	// spicial handling for filter
	a.Filter("vf", vid.Filter...)
	// spicial handling for map
	a.Maps(vid.Map)

	// Append the video parameters to the Args
	a.Append("crf", vid.Crf)
	a.Append("c:v", vid.Codec)
	a.Append("b:v", vid.BitRate)
	a.Append("r", vid.FrameRate)
	a.Append("s", vid.FrameSize)
	a.Append("aspect", vid.Aspect)
	a.Append("pix_fmt", vid.Pixfmt)
	a.Append("profile:v", vid.Profile)
}
