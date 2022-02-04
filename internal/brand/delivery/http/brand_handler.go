package http

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ipan97/simple-pos/internal/core"
)

type brandHTTPHandler struct {
	core.BaseHandler
}

type Product struct {
	ID        uint
	Name      string
	Price     float32
	PhotoUrl  string //nolint
	CreatedAt time.Time
}

func Test() {
	fmt.Println("Hello, World!")
	product := &Product{
		ID: 1,
	}

	// Unmarshal JSON
	json.Unmarshal([]byte(`{id: "1"}`), &product)
}
