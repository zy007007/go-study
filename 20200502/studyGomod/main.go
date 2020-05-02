package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/toolkits/smtp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://golang.org/pkg/time/`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Expand All" link
		chromedp.Click(`#pkg-examples > div`, chromedp.NodeVisible),
		// retrieve the value of the textarea
		chromedp.Value(`#example_After .play .input textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)

	// 个人新增第二个包
	s := smtp.NewSMTP("smtp.exmail.qq.com:25", "notify@a.com", "password", false, false, false)
	log.Println(s.SendMail("notify@a.com", "ulric@b.com;rain@c.com", "这是subject", "这是body,<font color=red>red</font>"))
}
