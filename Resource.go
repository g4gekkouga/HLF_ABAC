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

func AddResource(ctx contractapi.TransactionContextInterface, resourceID string, data string) (string, error) {

	if len(resourceID) == 0 {
		return "", fmt.Errorf("Please enter valid Resource ID")
	}

	resource := Resource{
		RID:  resourceID,
		Data: data,
	}

	resourceJSON, err := json.Marshal(resource)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("AddResource", resourceJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(resourceID, resourceJSON)
}
