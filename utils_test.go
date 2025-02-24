package trongrid

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSunToTRX(t *testing.T) {
	tests := []struct {
		name     string
		sun      int64
		expected float64
	}{
		{"zero", 0, 0},
		{"one TRX", 1_000_000, 1},
		{"half TRX", 500_000, 0.5},
		{"large amount", 1_234_567_890_000, 1234567.89},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SunToTRX(tt.sun)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTRXToSun(t *testing.T) {
	tests := []struct {
		name     string
		trx      float64
		expected int64
	}{
		{"zero", 0, 0},
		{"one TRX", 1, 1_000_000},
		{"half TRX", 0.5, 500_000},
		{"large amount", 1234567.89, 1_234_567_890_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TRXToSun(tt.trx)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseValue(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		decimals int32
		expected float64
	}{
		{"zero", "0", 6, 0},
		{"one token", "1000000", 6, 1},
		{"half token", "500000", 6, 0.5},
		{"large amount", "1234567890000", 6, 1234567.89},
		{"different decimals", "1000", 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseValue(tt.value, tt.decimals)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatValue(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		decimals int32
		expected string
	}{
		{"zero", 0, 6, "0"},
		{"one token", 1, 6, "1000000"},
		{"half token", 0.5, 6, "500000"},
		{"large amount", 1234567.89, 6, "1234567890000"},
		{"different decimals", 1, 3, "1000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatValue(tt.value, tt.decimals)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAddressConversion(t *testing.T) {
	tests := []struct {
		name    string
		address string
		hex     string
	}{
		{
			"valid address",
			"TJRabPrwbZy45sbavfcjinPJC18kjpRTv8",
			"41a614f803b6fd780986a42c78ec9c7f77e6ded13c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hex := AddressToHex(tt.address)
			assert.Equal(t, tt.hex, hex)

			address := HexToAddress(tt.hex)
			assert.Equal(t, tt.address, address)
		})
	}
}

func TestFormatAmount(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		expected string
	}{
		{"zero", 0, "0"},
		{"small", 123.45, "123.45"},
		{"thousands", 1234.56, "1.23K"},
		{"millions", 1234567.89, "1.23M"},
		{"billions", 1234567890.12, "1.23B"},
		{"trillions", 1234567890123.45, "1.23T"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatAmount(tt.amount)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatTimestamp(t *testing.T) {
	timestamp := time.Date(2024, 2, 24, 12, 0, 0, 0, time.UTC).Unix()
	expected := "2024-02-24 12:00:00"

	result := FormatTimestamp(timestamp)
	assert.Equal(t, expected, result)

	// Test millisecond timestamp
	result = FormatTimestamp(timestamp * 1000)
	assert.Equal(t, expected, result)
}

func TestCalculateEnergyFee(t *testing.T) {
	tests := []struct {
		name       string
		energyUsed int64
		energyFee  int64
		expected   float64
	}{
		{"zero", 0, 0, 0},
		{"basic", 1000, 420, 0.42},
		{"large", 10000, 420, 4.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateEnergyFee(tt.energyUsed, tt.energyFee)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsValidTronAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		valid   bool
	}{
		{"valid address", "TJRabPrwbZy45sbavfcjinPJC18kjpRTv8", true},
		{"invalid prefix", "XJRabPrwbZy45sbavfcjinPJC18kjpRTv8", false},
		{"invalid length", "TJRabPrwbZy45sbavfcjinPJC18kjpRTv", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidTronAddress(tt.address)
			assert.Equal(t, tt.valid, result)
		})
	}
}

func TestParseTRC20TransferData(t *testing.T) {
	tests := []struct {
		name          string
		data          string
		expectedTo    string
		expectedValue *big.Int
		expectError   bool
	}{
		{
			name:       "valid transfer",
			data:       "a9059cbb000000000000000000000041a614f803b6fd780986a42c78ec9c7f77e6ded13c0000000000000000000000000000000000000000000000000de0b6b3a7640000",
			expectedTo: "TJRabPrwbZy45sbavfcjinPJC18kjpRTv8",
			expectedValue: func() *big.Int {
				val := new(big.Int)
				val.SetString("1000000000000000000", 10)
				return val
			}(),
			expectError: false,
		},
		{
			name:        "invalid data",
			data:        "invalid",
			expectError: true,
		},
		{
			name:        "not transfer method",
			data:        "0123456789000000000000000000000041a614f803b6fd780986a42c78ec9c7f77e6ded13c0000000000000000000000000000000000000000000000000de0b6b3a7640000",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			to, value, err := ParseTRC20TransferData(tt.data)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTo, to)
				assert.Equal(t, tt.expectedValue, value)
			}
		})
	}
}
