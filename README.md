# HLF_ABAC

This is a Golang Package to provide necessary methods for supporting fully functional Attribute-Based Access Control in Hyperledger Fabric. Fabric chaincodes can use these methods directly by importing this package.

## Methods for Resource

### Resource Struct

type Resource struct {  
	ResourceID string `json:"resourceID"`  
	OwnerID    string `json:"ownerID"`  
	OwnerKey   string `json:ownerKey`  
	Data       string `json:"data"`  
}  

### addResource() and updateResource()

Input : 
1. Context - contractapi.TransactionContextInterface  
2. Resource ID - string
3. Resource Data - string
4. Private Collection - string

Output : 
1. Transaction ID - string
2. Error if any - error

### getResource()

Input : 
1. Context - contractapi.TransactionContextInterface  
2. Resource ID - string
4. Private Collection - string

Output : 
1. Resource - Type: Resource Struct
2. Error if any - error


## Methods for Resource Attributes

## Methods for User Attributes

## Methods for Policy

## Methods for Decision Unit

## Defined Structures

