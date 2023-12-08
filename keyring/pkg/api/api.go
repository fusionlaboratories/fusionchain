package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/qredo/fusionchain/keyring/pkg/common"
	"github.com/qredo/fusionchain/keyring/pkg/database"

	"github.com/qredo/fusionchain/keyring/pkg/rpc"
	"github.com/sirupsen/logrus"
)

const (

	// API - TODO create single package for multiple services
	StatusEndPnt   = "/status"
	HealthEndPnt   = "/healthcheck"
	KeyringEndPnt  = "/keyring"  // Password protected
	PubKeysEndPnt  = "/pubkeys"  // Password protected
	MnemonicEndPnt = "/mnemonic" // Password protected

	pwdHeaderKey = "password"
	pkPrefix     = "pk"
)

var (
	errInvalidPswd = errors.New("invalid password")
)

// Response represents the superset of Status and PubKey API responses.
type Response struct {
	Message       string    `json:"message,omitempty"`
	Version       string    `json:"version,omitempty"`
	Service       string    `json:"service,omitempty"`
	KeyRing       string    `json:"keyring,omitempty"`
	KeyringSigner string    `json:"keyring_signer,omitempty"`
	PubKeys       []*PubKey `json:"pubkeys,omitempty"`
	Mnemonic      string    `json:"mnemonic,omitempty"`
	PasswordUsed  bool      `json:"password_protected,omitempty"`
}

// HealthResponse represents the healthcheck API with no omitted fields.
type HealthResponse struct {
	Version  string   `json:"version"`
	Service  string   `json:"service"`
	Failures []string `json:"failures"`
}

type PubKey struct {
	KeyID      string `json:"key_id"`
	PubKeyData PkData `json:"pubkey_data"`
}

type PkData struct {
	PublicKey string `json:"pubkey"`
	Created   string `json:"created"`
	LastUsed  string `json:"last_used"`
}

type KeyringService interface { // Keyring service APIs
	Status(w http.ResponseWriter, r *http.Request)
	HealthCheck(w http.ResponseWriter, r *http.Request)
	Keyring(w http.ResponseWriter, r *http.Request)
	PubKeys(w http.ResponseWriter, r *http.Request)
	Mnemonic(w http.ResponseWriter, r *http.Request)
}

// HandleStatusRequest handles the /status query and will always respond OK
func HandleStatusRequest(w http.ResponseWriter, logger *logrus.Entry, serviceName string) {
	resp := Response{Message: "OK", Version: common.FullVersion, Service: serviceName}
	if err := rpc.RespondWithJSON(w, http.StatusOK, resp); err != nil {
		logger.Error(err)
	}
}

// HandleHealthcheckRequest handles the the /healthcheck query.
func HandleHealthcheckRequest(w http.ResponseWriter, modules []Module, logger *logrus.Entry, serviceName string) {
	health := &HealthResponse{
		Service: serviceName,
		Version: common.FullVersion,
	}
	var failures = []string{}

	for _, sub := range modules {
		// verify all subprocesses are healthy
		r := sub.Healthcheck()
		failures = append(failures, r.Failures...)
	}

	health.Failures = failures
	if len(failures) > 0 {
		if err := rpc.RespondWithJSON(w, http.StatusServiceUnavailable, health); err != nil {
			logger.Error(err)
		}
		return
	}
	if err := rpc.RespondWithJSON(w, http.StatusOK, health); err != nil {
		logger.Error(err)
	}
}

// HandleKeyringRequest implements the /keyring endpoint, keyring address registered for the service.
// PASSWORD PROTECTION is used, the http header must contain the correct password for the service.
func HandleKeyringRequest(w http.ResponseWriter, r *http.Request, logger *logrus.Entry, password, keyringAddr, keyRingSigner, serviceName string) {
	pwd := r.Header.Get(pwdHeaderKey)
	if password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	if err := rpc.RespondWithJSON(w, http.StatusOK, &Response{
		Service:       serviceName,
		Version:       common.FullVersion,
		Message:       "OK",
		KeyRing:       keyringAddr,
		KeyringSigner: keyRingSigner,
	}); err != nil {
		logger.Error(err)
	}
}

// HandlePubKeyRequest implements the /pubkeys endpoint, returning a list of registered keyID and public keys
// stored in the local database. PASSWORD PROTECTION is used, the http header must contain the correct password for
// the service.
func HandlePubKeyRequest(w http.ResponseWriter, r *http.Request, logger *logrus.Entry, db database.Database, password, serviceName string) {
	pwd := r.Header.Get(pwdHeaderKey)
	if password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	pKeyResponse := &Response{
		Service: serviceName,
		Version: common.FullVersion,
		Message: "OK",
	}

	keyMap, err := db.Read(pkPrefix)
	if err != nil {
		pKeyResponse.Message = err.Error()
		if err := rpc.RespondWithJSON(w, http.StatusInternalServerError, pKeyResponse); err != nil {
			logger.Error(err)
		}
		return
	}

	var pubKeyList []*PubKey
	for keyID, pKDat := range keyMap {
		dt := PkData{}
		if err := json.Unmarshal(pKDat, &dt); err != nil {
			rpc.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("could not unmarshal data '%s': %v", pKDat, err))
			return
		}
		pubKeyList = append(pubKeyList, &PubKey{KeyID: keyID, PubKeyData: dt})
	}
	pKeyResponse.PubKeys = pubKeyList

	if err := rpc.RespondWithJSON(w, http.StatusOK, pKeyResponse); err != nil {
		logger.Error(err)
	}
}

// HandleMnemonicRequest implements the /mnemonic endpoint, returning the BIP39 seed phrase used to derive the keyring's master seed.
// PASSWORD PROTECTION is used, the http header must contain the correct password for the service.
func HandleMnemonicRequest(w http.ResponseWriter, r *http.Request, logger *logrus.Entry, password, mnemonic, serviceName string) {
	pwd := r.Header.Get(pwdHeaderKey)
	if password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	if err := rpc.RespondWithJSON(w, http.StatusOK, &Response{
		Service:      serviceName,
		Version:      common.FullVersion,
		Message:      "OK",
		Mnemonic:     mnemonic,
		PasswordUsed: (password != ""),
	}); err != nil {
		logger.Error(err)
	}
}
