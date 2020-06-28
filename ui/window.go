package ui

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
	"upm/app"

	"github.com/zserge/lorca"
)

func Start(filename string) {
	data, _ := ioutil.ReadFile("public/loading.html")

	ui, _ := lorca.New("data:text/html,"+url.PathEscape(string(data)), "", 800, 600)

	uiBind(ui)

	srv := web()

	u := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("localhost:%v", port),
		Path:   "/",
		RawQuery: func() string {
			if filename == "" {
				return ""
			}
			values := url.Values{}
			values.Add("file", filename)
			return values.Encode()
		}(),
	}
	go func() {
		resp, err := http.Post(fmt.Sprintf("http://localhost:%v/hello", port), "application/json", nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
		ui.Load(u.String())
	}()

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	if srv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("server shutdown:", err)
		}
	}
	fmt.Println("exiting...")
}

func uiBind(ui lorca.UI) {
	ui.Bind("_openpackage", func(filename string) string {
		info := app.OpenPackage(filename)
		if info == nil {
			return "error: fail"
		}
		data, err := json.Marshal(info)
		if err != nil {
			return "error: " + err.Error()
		}
		return string(data)
	})
}
