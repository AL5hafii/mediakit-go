package ffmpegArgs

// AUDIO struct represents the audio stream parameters for ffmpeg command-line arguments
type AUDIO struct {
	Map    string   // -map value is the input stream mapping for the output file (e.g., "0:a" for the first audio stream of the first input file)
	Filter []string // -af value is the audio filter to apply (e.g., "volume=1.5" to increase volume by 50%)

	Codec, // -c:a codec is the audio codec to use (e.g., "aac", "mp3", etc.)
	Profile, // -profile:a value is the profile for the audio stream (e.g., "aac_low" for AAC Low Complexity)
	Channels, // -ac value is the number of audio channels (e.g., "2" for stereo)
	SampleRate string // -ar value is the audio sample rate (e.g., "44100" for 44.1 kHz)
}

// Audio appends the audio parameters to the Args
// filter must be applied before map and map before other parameters
func (a *ARGS) Audio(aud AUDIO) {
	// spicial handling for filter
	a.Filter("af", aud.Filter...)
	// spicial handling for map
	a.Maps(aud.Map)

	// Append the audio parameters to the Args
	a.Append("c:a", aud.Codec)
	a.Append("ac", aud.Channels)
	a.Append("ar", aud.SampleRate)
	a.Append("profile:a", aud.Profile)
}
