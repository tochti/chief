package chief

import (
	"log"
	"net/http"
)

func Example_Basic() {
	handle := func(url string) {
		resp, err := http.Get(url)

		if err != nil {
			log.Println(err)
		}

		log.Println(resp.Status)
	}

	decoder := func(j Job) {
		if url, ok := j.Order.(string); ok {
			handle(url)
		}
	}

	c := New(2, func(j Job) {
		decoder(j)
	})

	c.Start()
	urls := []string{
		"http://heise.de",
		"http://blog.fefe.de",
	}
	for _, url := range urls {
		c.Jobs <- Job{Order: url}
	}

}
