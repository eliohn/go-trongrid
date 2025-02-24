package trongrid

import (
	"time"
)

// Network constants
const (
	// MainnetURI is the URI for the Tron mainnet
	MainnetURI = "https://api.trongrid.io"
	// TestnetURI is the URI for the Tron testnet (Shasta)
	TestnetURI = "https://api.shasta.trongrid.io"
	// NileURI is the URI for the Tron testnet (Nile)
	NileURI = "https://nile.trongrid.io"
)

const (
	Mainnet = "mainnet"
	Testnet = "testnet"
	Nile    = "nile"
)

// Contract types
const (
	// ContractTypeTRX represents TRX transfer contract
	ContractTypeTRX = "TransferContract"
	// ContractTypeTRC10 represents TRC10 token transfer contract
	ContractTypeTRC10 = "TransferAssetContract"
	// ContractTypeTRC20 represents TRC20 token transfer contract
	ContractTypeTRC20 = "TriggerSmartContract"
	// ContractTypeAccountCreate represents account creation contract
	ContractTypeAccountCreate = "AccountCreateContract"
	// ContractTypeAccountUpdate represents account update contract
	ContractTypeAccountUpdate = "AccountUpdateContract"
	// ContractTypeFreeze represents TRX freeze contract
	ContractTypeFreeze = "FreezeBalanceContract"
	// ContractTypeUnfreeze represents TRX unfreeze contract
	ContractTypeUnfreeze = "UnfreezeBalanceContract"
	// ContractTypeVote represents vote contract
	ContractTypeVote = "VoteWitnessContract"
)

// Transaction status
const (
	// TxStatusSuccess represents a successful transaction
	TxStatusSuccess = "SUCCESS"
	// TxStatusFailed represents a failed transaction
	TxStatusFailed = "FAILED"
	// TxStatusPending represents a pending transaction
	TxStatusPending = "PENDING"
)

// API endpoints
const (
	// EndpointAccounts is the endpoint for account related operations
	EndpointAccounts = "/v1/accounts"
	// EndpointTransactions is the endpoint for transaction related operations
	EndpointTransactions = "/v1/transactions"
	// EndpointContracts is the endpoint for contract related operations
	EndpointContracts = "/v1/contracts"
	// EndpointEvents is the endpoint for event related operations
	EndpointEvents = "/v1/events"
)

// Common constants
const (
	// SunPerTRX represents the number of sun per TRX (1 TRX = 1,000,000 sun)
	SunPerTRX = 1_000_000
	// DefaultEnergyLimit represents the default energy limit for smart contract calls
	DefaultEnergyLimit = 10_000_000
	// MaxTransactionLifetime represents the maximum lifetime of a transaction in hours
	MaxTransactionLifetime = 24
	// DefaultTransactionTimeout represents the default timeout for transaction confirmation in seconds
	DefaultTransactionTimeout = 60
	// MaxBatchSize represents the maximum number of items in a batch request
	MaxBatchSize = 200
)

// Order constants
const (
	// OrderByTimestampDesc orders by timestamp in descending order
	OrderByTimestampDesc = "block_timestamp,desc"
	// OrderByTimestampAsc orders by timestamp in ascending order
	OrderByTimestampAsc = "block_timestamp,asc"
)

// Resource types
const (
	// ResourceBandwidth represents bandwidth resource
	ResourceBandwidth = "BANDWIDTH"
	// ResourceEnergy represents energy resource
	ResourceEnergy = "ENERGY"
	// ResourceTron represents TRON resource
	ResourceTron = "TRON"
)

const layout = "2006-01-02T15:04:05"

const timeout = time.Second * 10

const TransactionTypeTransfer TransactionType = "Transfer"

const URI = "https://api.trongrid.io/"
