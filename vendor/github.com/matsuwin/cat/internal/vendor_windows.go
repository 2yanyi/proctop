package internal

import "github.com/jaypipes/ghw"

func vendor() (_ string) {
	if product, _ := ghw.Product(); product != nil {
		return product.Vendor
	}
	return
}
