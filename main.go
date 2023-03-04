package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
  const url = "https://re-library.com/original/the-sword-saints-second-life-as-a-fox-girl/volume-1/1-1-second-wind/";

  req, err := http.NewRequest("GET", url, nil);
  if err != nil {
    log.Fatal("Error making http request");
  }

  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36");
  resp, err := http.DefaultClient.Do(req);
  if err != nil {
    log.Fatal("Error setting http header");
  }

  body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
    log.Fatal("Error reading body");
  }

  fmt.Println(string(body));
}
