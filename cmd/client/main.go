package main

import (
	"encoding/json"
	"github.com/DiasOrazbaev/kazanexpress-test-task/internal/batch"
	"github.com/DiasOrazbaev/kazanexpress-test-task/internal/service/dto"
	"github.com/DiasOrazbaev/kazanexpress-test-task/pkg/slice"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	// here 17 items
	items := []batch.Item{
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
	}
	// get limits of our server
	resp, err := http.Get("http://localhost:4040/batch")
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var rp dto.GetLimitResponse
	err = json.Unmarshal(body, &rp)
	if err != nil {
		log.Fatal(err)
	}

	chunks := slice.ChunkSlice(items, int(rp.MaxItemCount))

	for _, itms := range chunks {

		bytes, err := json.Marshal(itms)
		if err != nil {
			log.Fatal(err)
		}

		_, err = http.Post("http://localhost:4040/batch", "application/json", strings.NewReader(string(bytes)))
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(rp.ProcessPeriod)
	}
}
