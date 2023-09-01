package handler

import (
	"math/big"
	"time"
)

// Ref: https://stackoverflow.com/questions/54362751/how-can-i-truncate-float64-number-to-a-particular-precision#:~:text=is%200.3...).-,Using%20big.Float,-To%20properly%20truncate
// Take a look at https://github.com/shopspring/decimal
func Truncate(f float64, unit float64) float64 {
	bf := big.NewFloat(0).SetPrec(1000).SetFloat64(f)
	bu := big.NewFloat(0).SetPrec(1000).SetFloat64(unit)

	bf.Quo(bf, bu)

	// Truncate:
	i := big.NewInt(0)
	bf.Int(i)
	bf.SetInt(i)

	f, _ = bf.Mul(bf, bu).Float64()
	return f
}

func ParseDBDate(dbDate []uint8) (time.Time, error) {
	dateString := string(dbDate)
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
