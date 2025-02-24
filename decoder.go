package trongrid

import (
	"math/big"
	"reflect"
	"time"

	"github.com/gorilla/schema"
)

func NewDecoder() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	decoder.RegisterConverter(time.Time{}, func(s string) reflect.Value {
		t, err := time.Parse(layout, s)
		if err != nil {
			return reflect.ValueOf("")
		}

		return reflect.ValueOf(t)
	})
	decoder.SetAliasTag("url")
	decoder.ZeroEmpty(true)

	return decoder
}

// ParseValue 金额转换示例（需导入math/big包）
func ParseValue(valueStr string, decimals int32) *big.Float {
	value := new(big.Float)
	value.SetString(valueStr)

	divisor := new(big.Float).SetInt(new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(int64(decimals)),
		nil,
	))

	result := new(big.Float).Quo(value, divisor)
	return result
}

// func ParseURI(s string, v any) (err error) {
//	var vs url.Values
//
//	if vs, err = url.ParseQuery(s); err != nil {
//		return err
//	}
//
//	if err = decoder.Decode(&v, vs); err != nil {
//		return err
//	}
//
//	return nil
// }.
