package blockchain

type (
	Service interface {
		// Start tries start connection to blockchain client.
		Start() error

		// Stop tries stop connection to blockchain client.
		Stop() error

		// Reconnect tries to reconnect to blockchain client.
		Reconnect() error
	}
)
