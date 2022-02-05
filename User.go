package HLF_ABAC

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type User struct {
	UserID     string            `json:"userID"`
	UserName   string            `json:"userName"`
	Attributes map[string]string `json:"attributes"`
}

func GetSubmittingClientIdentity(ctx contractapi.TransactionContextInterface) (string, error) {

	b64ID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", fmt.Errorf("Failed to read clientID: %v", err)
	}
	decodeID, err := base64.StdEncoding.DecodeString(b64ID)
	if err != nil {
		return "", fmt.Errorf("failed to base64 decode clientID: %v", err)
	}
	return string(decodeID), nil
}

func GetSubmittingClientPubKey(ctx contractapi.TransactionContextInterface) (string, error) {

	Cert, err := ctx.GetClientIdentity().GetX509Certificate()
	if err != nil {
		return "", fmt.Errorf("Failed to read clientID: %v", err)
	}
	pubKey := Cert["Certificate"]["Data"]["pub"]
	if pubKey != nil {
		return "", fmt.Errorf("failed to get client public key")
	}
	return pubKey, nil
}

func getUAPriv(ctx contractapi.TransactionContextInterface, attrName string) (string, error) {

	attrValue, found, err := ctx.GetClientIdentity().GetAttributeValue(ctx.GetStub(), attrName)

	if !found {
		return "", fmt.Errorf("No attribute with given name")
	}

	if err {
		return "", err
	}

	return attrValue, nil
}

func registerUAPub(ctx contractapi.TransactionContextInterface, userID string, userName string, attributes string) (string, error) {

	err := ctx.GetClientIdentity().AssertAttributeValue(ctx.GetStub(), "admin", "true")

	if err {
		return "", fmt.Errorf("Invoking client is not an Admin")
	}

	if len(userID) == 0 {
		return "", fmt.Errorf("Please enter valid Subject ID")
	}

	attrs := strings.Split(attributes, ",")

	var attrsmap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		attrsmap[attrPair[0]] = attrPair[1]
	}

	user := User{
		UserID:     userID,
		UserName:   userName,
		Attributes: attrsmap,
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("registerUAPub", userJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(userID, registerUAPub)
}

func updateUAPub(ctx contractapi.TransactionContextInterface, userID string, userName string, attributes string) (string, error) {

	err := ctx.GetClientIdentity().AssertAttributeValue(ctx.GetStub(), "admin", "true")

	if err {
		return "", fmt.Errorf("Invoking client is not an Admin")
	}

	if len(userID) == 0 {
		return "", fmt.Errorf("Please enter valid Subject ID")
	}

	userJSON, err := ctx.GetStub().GetState(userID)

	if err || userJSON == nil {
		return "", fmt.Errorf("Attributes for given user not present")
	}

	attrs := strings.Split(attributes, ",")

	var attrsmap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		attrsmap[attrPair[0]] = attrPair[1]
	}

	user := User{
		UserID:     userID,
		UserName:   userName,
		Attributes: attrsmap,
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("updateUAPub", userJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(userID, registerUAPub)
}

func getUAPub(ctx contractapi.TransactionContextInterface, userID string) *User {
	if len(userID) == 0 {
		return nil
	}

	userJSON, err := ctx.GetStub().GetState(userID)

	if err != nil || userJSON == nil {
		return nil
	}

	user := new(User)
	_ = json.Unmarshal(userJSON, user)

	return user
}
