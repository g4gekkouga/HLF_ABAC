# HLF_ABAC

This is a Golang Package to provide necessary methods for supporting fully functional Attribute-Based Access Control in Hyperledger Fabric. Fabric chaincodes can use these methods directly by importing this package.

## Defined Structures

### Resource Struct Fields

1. ResourceID - string
2. OwnerID - string
3. OwnerKey - string
4. Data - string  

### Resource Attributes Struct Fields

1. ResourceID - string
2. OwnerID - string
3. OwnerKey - string
4. Attribures - map[string]string 

### User Attribues Struct Fields

1. UserID - string
2. UserName - string
3. Attributes - map[string]string

### Policy Struct Fields

1. PolicyID - string
2. UserAttr - map[string]string 
3. ResourceAttr - map[string]string 
4. EnvAttr - map[string]string 
5. Operation - string
6. Rules - map[string]string

## Methods for Resource

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

