package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {

	listenAddress := ":4000"

	tr := NEWTCPTransport(listenAddress)

	assert.Equal(t, tr.listenAddress, listenAddress)

}
