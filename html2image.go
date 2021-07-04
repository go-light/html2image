package html2image

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

//Params print page as image.
type Options struct {
	page.CaptureScreenshotParams
	CustomClip bool
}

type Html2Image interface {
	Convert() ([]byte, error)
}

type html2image struct {
	url                     string
	ctx                     context.Context
	captureScreenshotFormat CaptureScreenshotFormat

	// Sleep is an empty action that calls time.Sleep with the specified duration.
	//
	// Note: this is a temporary action definition for convenience, and will likely
	// be marked for deprecation in the future, after the remaining Actions have
	// been able to be written/tested.
	sleep time.Duration

	pageCaptureScreenshotParams page.CaptureScreenshotParams

	convertElapsed time.Duration
	buf            []byte
}

func NewHtml2Image(url string, opts ...Option) Html2Image {
	h2i := html2image{
		url:                     url,
		ctx:                     context.Background(),
		captureScreenshotFormat: CaptureScreenshotFormatPng,
	}

	for _, opt := range opts {
		opt.Apply(&h2i)
	}

	h2i.pageCaptureScreenshotParams = page.CaptureScreenshotParams{
		Format:  page.CaptureScreenshotFormat(h2i.captureScreenshotFormat),
		Quality: DefaultQuality,
		Clip: &page.Viewport{
			X:      0,
			Y:      0,
			Width:  0,
			Height: 0,
			Scale:  DefaultViewportScale,
		},
		FromSurface: DefaultFromSurface,
	}

	return &h2i
}

func (h2i *html2image) GetConvertElapsed() time.Duration {
	return h2i.convertElapsed
}

func (h2i *html2image) Convert() ([]byte, error) {
	start := time.Now()
	defer func() {
		h2i.convertElapsed = time.Since(start)
	}()

	ctx, cancel := chromedp.NewContext(h2i.ctx)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(h2i.url),
		chromedp.Sleep(h2i.sleep),
		chromedp.ActionFunc(func(ctx context.Context) error {

			// get layout metrics
			_, _, contentSize, _, _, _, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			params := h2i.pageCaptureScreenshotParams
			if params.Clip.X == 0 {
				params.Clip.X = contentSize.X
			}
			if params.Clip.Y == 0 {
				params.Clip.Y = contentSize.Y
			}
			if params.Clip.Width == 0 {
				params.Clip.Width = contentSize.Width
			}
			if params.Clip.Height == 0 {
				params.Clip.Height = contentSize.Height
			}

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(int64(params.Clip.Width), int64(params.Clip.Height), 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			h2i.buf, err = params.Do(ctx)
			return err
		}),
	); err != nil {
		return nil, err
	}

	return h2i.buf, nil
}
