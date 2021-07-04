package html2image

import "github.com/chromedp/cdproto/page"

type CaptureScreenshotFormat string

// String returns the CaptureScreenshotFormat as string value.
func (t CaptureScreenshotFormat) String() string {
	return string(t)
}

// CaptureScreenshotFormat values.
const (
	CaptureScreenshotFormatJpeg = CaptureScreenshotFormat(page.CaptureScreenshotFormatJpeg)
	CaptureScreenshotFormatPng  = CaptureScreenshotFormat(page.CaptureScreenshotFormatPng)
)

//DefaultQuality Compression quality from range [0..100] (jpeg only).
const DefaultQuality = 100

//DefaultFromSurface Capture the screenshot from the surface, rather than the view. Defaults to true.
const DefaultFromSurface = true

//DefaultViewportX Capture the screenshot of a given region only.
// X offset in device independent pixels (dip).
const DefaultViewportX = 0

//DefaultViewportY Y offset in device independent pixels (dip).
const DefaultViewportY = 0

//DefaultViewportWidth Rectangle width in device independent pixels (dip).
const DefaultViewportWidth = 0

//DefaultViewportHeight Rectangle height in device independent pixels (dip).
const DefaultViewportHeight = 0

//DefaultViewportScale Page scale factor.
const DefaultViewportScale = 1
