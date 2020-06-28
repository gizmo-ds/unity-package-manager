package ui

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/zserge/lorca"
)

func Start(filename string) {
	data, _ := ioutil.ReadFile("public/loading.html")

	ui, _ := lorca.New("data:text/html,"+url.PathEscape(string(data)), "", 800, 600)
	srv := web()

	values := url.Values{}
	values.Add("file", filename)
	u := url.URL{
		Scheme:     "http",
		Host:       fmt.Sprintf("localhost:%v", port),
		Path:       "/",
		ForceQuery: true,
		RawQuery:   values.Encode(),
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
