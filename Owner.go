package HLF_ABAC

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ResourceAttr struct {
	ResourceID string            `json:"resourceID"`
	OwnerID    string            `json:"ownerID"`
	OwnerKey   string            `json:ownerKey`
	Attributes map[string]string `json:"attributes"`
}

func registerRAPub(ctx contractapi.TransactionContextInterface, resourceID string, attributes string) (string, error) {

	if len(resourceID) == 0 {
		return "", fmt.Errorf("Please enter valid Resource ID")
	}

	ownerID, err := GetSubmittingClientIdentity(ctx)

	if err != nil {
		return "", err
	}

	ownerKey, err := GetSubmittingClientPubKey(ctx)

	if err != nil {
		return "", err
	}

	attrs := strings.Split(attributes, ",")

	var attrsmap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		attrsmap[attrPair[0]] = attrPair[1]
	}

	resourceAttr := ResourceAttr{
		ResourceID: resourceID,
		OwnerID:    ownerID,
		OwnerKey:   ownerKey,
		Attributes: attrsmap,
	}

	resourceAttrJSON, err := json.Marshal(resourceAttr)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("registerRAPub", resourceAttrJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(resourceID+"_RAPub", resourceAttrJSON)
}

func updateRAPub(ctx contractapi.TransactionContextInterface, resourceID string, attributes string) (string, error) {

	if len(resourceID) == 0 {
		return "", fmt.Errorf("Please enter valid Resource ID")
	}

	resourceAttrJSON, err := ctx.GetStub().GetState(resourceID + "_RAPub")

	if resourceAttrJSON == nil {
		return "", fmt.Errorf("No RAPub for Resource with given ID")
	}

	ownerID, err := GetSubmittingClientIdentity(ctx)

	if err != nil {
		return "", err
	}

	ownerKey, err := GetSubmittingClientPubKey(ctx)

	if err != nil {
		return "", err
	}

	attrs := strings.Split(attributes, ",")

	var attrsmap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		attrsmap[attrPair[0]] = attrPair[1]
	}

	resourceAttr := ResourceAttr{
		ResourceID: resourceID,
		OwnerID:    ownerID,
		OwnerKey:   ownerKey,
		Attributes: attrsmap,
	}

	resourceAttrJSON, err := json.Marshal(resourceAttr)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("updateRAPub", resourceAttrJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(resourceID+"_RAPub", resourceAttrJSON)
}

func getRAPub(ctx contractapi.TransactionContextInterface, resourceID string) *ResourceAttr {
	if len(resourceID) == 0 {
		return nil
	}

	resourceAttrJSON, err := ctx.GetStub().GetState(resourceID + "_RAPub")

	if err != nil || resourceAttrJSON == nil {
		return nil
	}

	resourceAttr := new(ResourceAttr)
	_ = json.Unmarshal(resourceAttrJSON, resourceAttr)

	return resourceAttr
}

func registerRAPriv(ctx contractapi.TransactionContextInterface, resourceID string, attributes string, collection string) (string, error) {

	if len(resourceID) == 0 {
		return "", fmt.Errorf("Please enter valid Resource ID")
	}

	ownerID, err := GetSubmittingClientIdentity(ctx)

	if err != nil {
		return "", err
	}

	ownerKey, err := GetSubmittingClientPubKey(ctx)

	if err != nil {
		return "", err
	}

	attrs := strings.Split(attributes, ",")

	var attrsmap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		attrsmap[attrPair[0]] = attrPair[1]
	}

	resourceAttr := ResourceAttr{
		ResourceID: resourceID,
		OwnerID:    ownerID,
		OwnerKey:   ownerKey,
		Attributes: attrsmap,
	}

	resourceAttrJSON, err := json.Marshal(resourceAttr)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("registerRAPriv", resourceAttrJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutPrivateData(resourceID+"_RAPriv", resourceAttrJSON)
}

func updateRAPriv(ctx contractapi.TransactionContextInterface, resourceID string, attributes string, collection string) (string, error) {

	if len(resourceID) == 0 {
		return "", fmt.Errorf("Please enter valid Resource ID")
	}

	resourceAttrJSON, err := ctx.GetStub().GetPrivateData(collection, resourceID+"_RAPriv")

	if resourceAttrJSON == nil {
		return "", fmt.Errorf("No RAPub for Resource with given ID")
	}

	ownerID, err := GetSubmittingClientIdentity(ctx)

	if err != nil {
		return "", err
	}

	ownerKey, err := GetSubmittingClientPubKey(ctx)

	if err != nil {
		return "", err
	}

	attrs := strings.Split(attributes, ",")

	var attrsmap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		attrsmap[attrPair[0]] = attrPair[1]
	}

	resourceAttr := ResourceAttr{
		ResourceID: resourceID,
		OwnerID:    ownerID,
		OwnerKey:   ownerKey,
		Attributes: attrsmap,
	}

	resourceAttrJSON, err := json.Marshal(resourceAttr)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("updateRAPriv", resourceAttrJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutPrivateData(resourceID+"_RAPriv", resourceAttrJSON)
}

func getRAPriv(ctx contractapi.TransactionContextInterface, resourceID string, collection string) *ResourceAttr {
	if len(resourceID) == 0 {
		return nil
	}

	resourceAttrJSON, err := ctx.GetStub().GetPrivateData(collection, resourceID+"_RAPriv")

	if err != nil || resourceAttrJSON == nil {
		return nil
	}

	resourceAttr := new(ResourceAttr)
	_ = json.Unmarshal(resourceAttrJSON, resourceAttr)

	return resourceAttr
}
