package html2image

// Option represents the Html2image options
type Option interface {
	Apply(*html2image)
}

// OptionFunc is a function that configures a Html2image.
type OptionFunc func(*html2image)

// Apply calls f(client)
func (f OptionFunc) Apply(h2i *html2image) {
	f(h2i)
}

func WithCaptureScreenshotFormat(f CaptureScreenshotFormat) Option {
	return OptionFunc(func(h2i *html2image) {
		h2i.captureScreenshotFormat = f
	})
}
