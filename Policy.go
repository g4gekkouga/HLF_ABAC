package HLF_ABAC

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Policy struct {
	PolicyID     string            `json:"policyID"`
	UserAttr     map[string]string `json:"userAttr"`
	ResourceAttr map[string]string `json:"resourceAttr"`
	EnvAttr      map[string]string `json:"envAttr"`
	Operation    string            `json:"op"`
	Rules        map[string]string `json:"rules"`
}

func registerPolicy(ctx contractapi.TransactionContextInterface, policyID string, userAttr string, resourceAttr string, envAttr string, operation string, rules string) (string, error) {

	if len(policyID) == 0 {
		return "", fmt.Errorf("Please enter valid Policy ID")
	}

	attrs := strings.Split(userAttr, ",")

	var userAttrsMap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		userAttrsMap[attrPair[0]] = attrPair[1]
	}

	attrs = strings.Split(resourceAttr, ",")

	var resourceAttrsMap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		resourceAttrsMap[attrPair[0]] = attrPair[1]
	}

	attrs = strings.Split(envAttr, ",")

	var envAttrsMap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		envAttrsMap[attrPair[0]] = attrPair[1]
	}

	attrs = strings.Split(rules, ",")

	var rulesMap = make(map[string]string)

	for i := 0; i < len(attrs); i++ {
		attrPair := strings.Split(attrs[i], ":")
		rulesMap[attrPair[0]] = attrPair[1]
	}

	policy := Policy{
		PolicyID:     policyID,
		UserAttr:     userAttrsMap,
		ResourceAttr: resourceAttrsMap,
		EnvAttr:      envAttrsMap,
		Operation:    operation,
		Rules:        rulesMap,
	}

	policyJSON, err := json.Marshal(policy)
	if err != nil {
		return "", err
	}

	ctx.GetStub().SetEvent("RegisterPolicy", policyJSON)

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(policyID, policyJSON)
}

func getPolicy(ctx contractapi.TransactionContextInterface, policyID string) (*Policy, error) {
	if len(policyID) == 0 {
		return nil, fmt.Errorf("Please enter valid Policy ID")
	}

	policyJSON, err := ctx.GetStub().GetState(policyID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if policyJSON == nil {
		return nil, fmt.Errorf("%s does not exist", policyID)
	}

	policy := new(Policy)
	_ = json.Unmarshal(policyJSON, policy)

	return policy, nil

}

func getPolicySet(ctx contractapi.TransactionContextInterface) []*Policy {

	resultsIterator, err := ctx.GetStub().GetStateByRange("Policy", "RA")
	if err != nil {
		return nil
	}
	defer resultsIterator.Close()

	var policies []*Policy
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil
		}

		var policy Policy
		err = json.Unmarshal(queryResponse.Value, &policy)
		if err != nil {
			return nil
		}
		policies = append(policies, &policy)
	}

	return policies
}
