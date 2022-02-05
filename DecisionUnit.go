package HLF_ABAC

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func getContext() *map[string]string {
	return nil
}

func isValid(rule *Policy, UAPub *map[string]string, RAPub *map[string]string, RAPriv *map[string]string, EA *map[string]string, OP string) bool {
	return true
}

func validateAccess(ctx contractapi.TransactionContextInterface, userIDHash string, resourceID string, operation string, collection string) (*Resource, error) {

	if len(userIDHash) == 0 {
		return nil, fmt.Errorf("Please enter valid Subject ID")
	}

	if len(resourceID) == 0 {
		return nil, fmt.Errorf("Please enter valid Object ID")
	}

	userID, err := GetSubmittingClientIdentity(ctx)
	userPubKey, err := GetSubmittingClientPubKey(ctx)

	// User Verification

	UAPub := getUAPub(ctx, userID).Attributes
	RAPub := getRAPub(ctx, resourceID).Attributes
	RAPriv := getRAPriv(ctx, resourceID, collection).Attributes
	EA := getContext()

	POL := getPolicySet(ctx)

	for _, rule := range POL {
		if !isValid(rule, UAPub, RAPub, RAPriv, EA, operation) {
			continue
		}

		return getResource(ctx, resourceID, collection), nil
	}

	return nil, fmt.Errorf("Access Denied")
}
