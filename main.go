package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func printToPDF(urlstr string, res *[]byte, injectData interface{}) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Evaluate(fmt.Sprintf(`
			injectValue(%v);
		`, injectData), &injectData),
		chromedp.WaitReady("div#root div"),
		chromedp.Sleep(1 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	ts := httptest.NewServer(http.FileServer(http.Dir("./template/dist")))
	defer ts.Close()

	obj := struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"createdAt"`
	}{
		Name:      "lilonghe",
		CreatedAt: time.Now(),
	}

	injectData, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
		return
	}

	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(ts.URL+"/index.html", &buf, string(injectData))); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("sample.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("sample.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	
	ts.Close()

	output := http.FileServer(http.Dir("/app"))
	http.Handle("/", output)

	log.Printf("Starting server on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
