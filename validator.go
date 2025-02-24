package trongrid

import (
	"regexp"
	"time"
)

var (
	// tronAddressRegex represents the regex pattern for a valid TRON address
	tronAddressRegex = regexp.MustCompile("^T[A-Za-z1-9]{33}$")
)

// validateListTransactionsRequest validates the ListTransactionsRequest
func validateListTransactionsRequest(req *ListTransactionsRequest) error {
	if req == nil {
		return ErrInvalidRequest
	}

	// Validate address
	if req.Address == "" {
		return ErrMissingAddress
	}
	if !tronAddressRegex.MatchString(req.Address) {
		return ErrInvalidAddress
	}

	// Validate time range if both are set
	if !req.MinTimestamp.IsZero() && !req.MaxTimestamp.IsZero() {
		if req.MinTimestamp.After(req.MaxTimestamp) {
			return ErrInvalidTimeRange
		}
		// Check if time range is not too large (e.g., max 90 days)
		if req.MaxTimestamp.Sub(req.MinTimestamp) > 90*24*time.Hour {
			return ErrInvalidTimeRange
		}
	}

	// Validate limit
	if req.Limit < 0 || req.Limit > 200 {
		return ErrInvalidLimit
	}

	// Set default limit if not specified
	if req.Limit == 0 {
		req.Limit = 20
	}

	return nil
}
