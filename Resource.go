package HLF_ABAC

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Resource struct {
	ResourceID string `json:"resourceID"`
	OwnerID    string `json:"ownerID"`
	OwnerKey   string `json:ownerKey`
	Data       string `json:"data"`
}

func addResource(ctx contractapi.TransactionContextInterface, resourceID string, data string, collection string) (string, error) {

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

	resource := Resource{
		ResourceID: resourceID,
		OwnerID:    ownerID,
		OwnerKey:   ownerKey,
		Data:       data,
	}

	resourceJSON, err := json.Marshal(resource)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("AddResource", resourceJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutPrivateData(collection, resourceID, resourceJSON)
}

func updateResource(ctx contractapi.TransactionContextInterface, resourceID string, data string, collection string) (string, error) {

	if len(resourceID) == 0 {
		return "", fmt.Errorf("Please enter valid Resource ID")
	}

	resourceJSON, err := ctx.GetStub().GetPrivateData(collection, resourceID)

	if resourceJSON == nil {
		return "", fmt.Errorf("No Resource with given ID")
	}

	ownerID, err := GetSubmittingClientIdentity(ctx)

	if err != nil {
		return "", err
	}

	ownerKey, err := GetSubmittingClientPubKey(ctx)

	if err != nil {
		return "", err
	}

	resource := Resource{
		ResourceID: resourceID,
		OwnerID:    ownerID,
		OwnerKey:   ownerKey,
		Data:       data,
	}

	resourceJSON, err = json.Marshal(resource)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("UpdateResource", resourceJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutPrivateData(collection, resourceID, resourceJSON)
}

func getResource(ctx contractapi.TransactionContextInterface, resourceID string, collection string) *Resource {
	if len(resourceID) == 0 {
		return nil
	}

	resourceJSON, err := ctx.GetStub().GetPrivateData(collection, resourceID)

	if err != nil || resourceJSON == nil {
		return nil
	}

	resource := new(Resource)
	_ = json.Unmarshal(resourceJSON, resource)

	return resource
}
