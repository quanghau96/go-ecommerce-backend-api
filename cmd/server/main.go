package main

import (
	"github.com/quanghau96/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8002")
}
