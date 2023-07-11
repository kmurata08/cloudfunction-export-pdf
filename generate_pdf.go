package function

import (
	"context"
	"net/http"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func GeneratePDF(w http.ResponseWriter, r *http.Request) {
	html := `
		<html>
		<head>
			<style>
				.border-box {
					border: 2px solid black;
					padding: 20px;
					margin: 20px;
				}
			</style>
		</head>
		<body>
			<h1>Hello, PDF!</h1>
			<div class="border-box">
				<p>This is a paragraph with a border box.</p>
				<img src="https://placekitten.com/200/300" alt="A placeholder kitten">
			</div>
		</body>
		</html>`

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(html, &buf)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="simple.pdf"`)

	if _, err := w.Write(buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func printToPDF(html string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`data:text/html,` + html),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// capture screenshot
			var err error
			*res, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			return err
		}),
	}
}
