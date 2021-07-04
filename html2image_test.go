package html2image

import (
	"io/ioutil"
	"testing"
)

func TestHtml2image_Convert(t *testing.T) {
	h2i := NewHtml2Image("https://www.baidu.com/s?ie=UTF-8&wd=github")

	buf, err := h2i.Convert()
	if err != nil {
		t.Error(err)
		return
	}

	err = ioutil.WriteFile("test.png", buf, 0644)
	if err != nil {
		t.Error(err)
		return
	}
}
