package trongrid

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// SunToTRX converts a sun amount to TRX
func SunToTRX(sun int64) float64 {
	return float64(sun) / float64(SunPerTRX)
}

// TRXToSun converts a TRX amount to sun
func TRXToSun(trx float64) int64 {
	return int64(trx * float64(SunPerTRX))
}

// ParseValue parses a string value with given decimals to float64
func ParseValue(value string, decimals int32) float64 {
	if value == "" {
		return 0
	}

	// Convert string to big.Int
	val := new(big.Int)
	val.SetString(value, 10)

	// Calculate divisor (10^decimals)
	divisor := new(big.Float).SetFloat64(math.Pow10(int(decimals)))

	// Convert value to big.Float
	floatValue := new(big.Float).SetInt(val)

	// Perform division
	result := new(big.Float).Quo(floatValue, divisor)

	// Convert to float64
	f64, _ := result.Float64()
	return f64
}

// FormatValue formats a float64 value with given decimals to string
func FormatValue(value float64, decimals int32) string {
	// Multiply by 10^decimals
	multiplier := math.Pow10(int(decimals))
	intValue := int64(value * multiplier)

	return strconv.FormatInt(intValue, 10)
}

// AddressToHex converts a Tron address to hex format
func AddressToHex(address string) string {
	if strings.HasPrefix(address, "T") {
		address = address[1:]
	}
	return "41" + address
}

// HexToAddress converts a hex address to Tron format
func HexToAddress(hex string) string {
	if strings.HasPrefix(hex, "41") {
		hex = hex[2:]
	}
	return "T" + hex
}

// FormatAmount formats an amount with appropriate suffix (K, M, B, T)
func FormatAmount(amount float64) string {
	if amount == 0 {
		return "0"
	}

	abs := math.Abs(amount)
	if abs < 1000 {
		return fmt.Sprintf("%.2f", amount)
	}

	suffix := ""
	switch {
	case abs >= 1e12:
		amount /= 1e12
		suffix = "T"
	case abs >= 1e9:
		amount /= 1e9
		suffix = "B"
	case abs >= 1e6:
		amount /= 1e6
		suffix = "M"
	case abs >= 1e3:
		amount /= 1e3
		suffix = "K"
	}

	return fmt.Sprintf("%.2f%s", amount, suffix)
}

// FormatTimestamp formats a Unix timestamp to a human-readable string
func FormatTimestamp(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	// Convert milliseconds to seconds if necessary
	if timestamp > 1e11 {
		timestamp /= 1000
	}
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

// CalculateEnergyFee calculates the energy fee in TRX
func CalculateEnergyFee(energyUsed int64, energyFee int64) float64 {
	return SunToTRX(energyUsed * energyFee)
}

// CalculateBandwidthFee calculates the bandwidth fee in TRX
func CalculateBandwidthFee(bandwidthUsed int64, bandwidthFee int64) float64 {
	return SunToTRX(bandwidthUsed * bandwidthFee)
}

// IsValidTronAddress checks if the given address is a valid Tron address
func IsValidTronAddress(address string) bool {
	if !strings.HasPrefix(address, "T") {
		return false
	}

	// Tron addresses are base58 encoded and should be 34 characters long
	if len(address) != 34 {
		return false
	}

	// Additional validation can be added here if needed
	return true
}

// IsContract checks if the given address is a contract address
func IsContract(address string) bool {
	if !IsValidTronAddress(address) {
		return false
	}

	// Contract addresses in Tron start with "T" and their hex representation starts with "41"
	hex := AddressToHex(address)
	return strings.HasPrefix(hex, "41")
}

// ParseTRC20TransferData parses TRC20 transfer data
func ParseTRC20TransferData(data string) (to string, amount *big.Int, err error) {
	if len(data) < 138 {
		return "", nil, fmt.Errorf("invalid data length")
	}

	// Remove "0x" prefix if present
	data = strings.TrimPrefix(data, "0x")

	// Check if it's a transfer method (a9059cbb)
	if !strings.HasPrefix(data, "a9059cbb") {
		return "", nil, fmt.Errorf("not a transfer method")
	}

	// Extract to address (32 bytes, padded)
	addressHex := "41" + data[32:72]
	to = HexToAddress(addressHex)

	// Extract amount (32 bytes)
	amount = new(big.Int)
	amount.SetString(data[72:], 16)

	return to, amount, nil
}
