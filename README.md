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

Input: 
1. Context - contractapi.TransactionContextInterface  
2. Resource ID - string
3. Resource Data - string
4. Private Collection - string

Output: 
1. Transaction ID - string
2. Error if any - error

### getResource()

Input: 
1. Context - contractapi.TransactionContextInterface  
2. Resource ID - string
4. Private Collection - string

Output: 
1. Resource - Type: Resource Struct
2. Error if any - error


## Methods for Resource Attributes  

### registerRAPub() and updateRAPub()

Input:  
1. Context - contractapi.TransactionContextInterface  
2. Resource ID - string
3. Attributes - string : Key-value pairs in 'key:value' format, seperated by ','

Output:  
1. Transaction ID - string
2. Error if any - error

### getRAPub()

Input:  
1. Context - contractapi.TransactionContextInterface  
2. Resource ID - string

Output:  
1. Resource Attributes - Type: ResourceAttr Struct
2. Error if any - error


### registerRAPriv() and updateRAPriv()

Input:  
1. Context - contractapi.TransactionContextInterface  
2. Resource ID - string
3. Attributes - string : Key-value pairs in 'key:value' format, seperated by ','
4. Private Collection - string

Output:  
1. Transaction ID - string
2. Error if any - error

### getRAPriv()

Input:  
1. Context - contractapi.TransactionContextInterface  
2. Resource ID - string
3. Private Collection - string

Output:  
1. Resource Attributes - Type: ResourceAttr Struct
2. Error if any - error


## Methods for User Attributes

### registerUAPub() and updateUAPub()

Input:  
1. Context - contractapi.TransactionContextInterface  
2. User ID - string
3. User Name - string
4. Attributes - string : Key-value pairs in 'key:value' format, seperated by ','

Output:  
1. Transaction ID - string
2. Error if any - error

### getUAPub()

Input:  
1. Context - contractapi.TransactionContextInterface  
2. User ID - string

Output:  
1. User Attributes - Type: UserAttr Struct
2. Error if any - error

### getUAPriv()

Input:  
1. Context - contractapi.TransactionContextInterface
2. Attribute Name - string

Output:  
1. Attribute Value - string
2. Error if any - error

## Methods for Policy

### registerPolicy() and updatePolicy()

Input:
1. Context - contractapi.TransactionContextInterface
2. Policy ID - string
3. User Attributes - string : Required User Attribute names separated by ','
4. Resource Attributes - string : Required Resource Attribute names separated by ','
5. Environment Attributes - string : Required Environment Attribute names separated by ','
6. Operation - string
7. Rules - string : Key-value pairs in 'key:value' format, seperated by ','

Output:
1. Transaction ID - string
2. Error if any - error

## getPolicy()

Input:
1. Context - contractapi.TransactionContextInterface
2. Policy ID - string

Output:
1. Policy - Type: Policy Struct
2. Error if any - error

## getPolicySet()

Input:
1. Context - contractapi.TransactionContextInterface

Output:
1. All Policies - []Policy
2. Error if any - error

## Methods for Decision Unit

### validateAccess() 

Input:  
1. Context - contractapi.TransactionContextInterface
2. Hash(User ID, User Priv Key) - string
3. Resource ID - string
4. Operation - string
5. Private Collection - string

Output:
1. Resource - Type: Resource Struct
2. Error if any - error

## Note
The current version of the package supports only minimal use case required for latency testing purpose of proposed approach. We are currently in development more robust version to support generalized use cases and error support for boundary situation.

