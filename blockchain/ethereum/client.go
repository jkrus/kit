package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jkrus/kit/blockchain"
)

type (
	// EthereumClientI describes ethereum client interface.
	EthereumClientI interface {
		// Service use embedded common interface.
		blockchain.Service
	}
)

type (
	ethereumClient struct {
		pathIPC string
		Client  *ethclient.Client
	}
)

// NewEthClient return EthereumClientI interface.
func NewEthClient(pathIPC string) EthereumClientI {
	c := ethereumClient{Client: nil, pathIPC: pathIPC}
	return &c
}

// Reconnect implement Service interface.
func (c *ethereumClient) Reconnect() error {
	return c.Start()
}

// Start implement Service interface.
func (c *ethereumClient) Start() error {
	rpcClient, err := rpc.Dial(c.pathIPC)
	ethClient := ethclient.NewClient(rpcClient)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	c.Client = ethClient
	log.Println("Successfully connected to Ethereum client.")

	return nil
}

// Stop implement Service interface.
func (c *ethereumClient) Stop() error {
	log.Println("Close connection to Ethereum client...")
	c.Client.Close()
	log.Println("Ethereum client connection closed.")

	return nil
}
