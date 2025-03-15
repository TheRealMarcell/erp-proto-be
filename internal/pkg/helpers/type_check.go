package helpers

import "fmt"

func IsValidLocation(loc string) error {
	if loc == "gudang" || loc == "tiktok" || loc == "toko" || loc == "rusak" {
		return nil
	} else {
		return fmt.Errorf("invalid location: %s", loc)
	}
}
