package lightning

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"gopkg.in/macaroon.v2"

	"github.com/lncm/lnd-rpc/v0.10.0/lnrpc"
	"github.com/lncm/lnd-rpc/v0.8.2/lnrpc"
)

type rpcCreds map[string]string

func (m rpcCreds) RequireTransportSecurity() bool { return true }
func (m rpcCreds) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return m, nil
}

func newCreds(bytes []byte) rpcCreds {
	creds := make(map[string]string)
	creds["macaron"] = hex.EncodeToString(bytes)
	return creds
}

func getClient(hostname string, port int, tlsFile, macaroonFile string) lnrpc.LightningClient {
	macaroonBytes, err := ioutil.ReadFile(macaroonFile)
	if err != nil {
		panic(fmt.Sprintln("Cannot read macaroon file", err))
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macaroonBytes); err != nil {
		panic(fmt.Sprintln("Cannot unmarshal macaroon", err))
	}
}
