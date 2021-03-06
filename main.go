package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {

	// Starting Chrome
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	// Want to turn these into command line args using flag
	url := "https://www.nytimes.com/"
	filename := "nyt.png"

	// Run Tasks
	// List of actions to run in sequence (which also fills our image buffer)
	var imageBuf []byte
	if err := chromedp.Run(ctx, TakeScreenshot(url, &imageBuf)); err != nil {
		log.Fatal(err)
	}

	// This writes the image to file
	if err := ioutil.WriteFile(filename, imageBuf, 0644); err != nil {
		log.Fatal(err)
	}
}

func TakeScreenshot(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) (err error) {

			// Cool little addition, if I wanted to make a PDF instead of a PNG I'd change this line to
			// *imageBuf, _, err = page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			*imageBuf, err = page.CaptureScreenshot().WithQuality(90).Do(ctx)
			return err
		}),
	}
}
