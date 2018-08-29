/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
*
* ─────────────────────────────────────────────────────────────────────────────────────────────────┐
* Code Review: IBM Apex App Connect Platform Event
* Class Name: NerosidioCoin
* Description: Executes upsert of disaster data driven from current API
* ──────────────────────────────────────────────────────────────────────────────────────────────────
* @secondaryDeveloper         TEST     test@test@ibm.com
* @Tester     Tester   email
* @ReviewedBy   Brandon Kravitz    brandon.kravitz@bluewolfgroup.com
* @userStory        US-0001
* @created         2018-11-08
* @modified        2018-11-08
* @systemLayer    Test
* @thisCodeReviewTestClass            appConnect_DisasterEvent
* @see            AppConnect Documentation
* ──────────────────────────────────────────────────────────────────────────────────────────────────
* @changes
* v1.0             brandon.kravitz@ibm.com
* 2018.01.08     Initial Commit
* ─────────────────────────────────────────────────────────────────────────────────────────────────┘
*/		
package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the NerosidioCoin structure, with 4 properties.  Structure tags are used by encoding/json library
type NerosidioCoin struct {
	DisasterType   string `json:"disastertype"`
	Location  string `json:"location"`
	Donation string `json:"donation"`
	Owner  string `json:"owner"`
}

/*
 * The Init method is called when the Smart Contract "NerosidioCoin" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "NerosidioCoin"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryNerosidioCoin" {
		return s.queryNerosidioCoin(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createNerosidioCoin" {
		return s.createNerosidioCoin(APIstub, args)
	} else if function == "queryAllNerosidioCoins" {
		return s.queryAllNerosidioCoins(APIstub)
	} else if function == "changeNerosidioCoinOwner" {
		return s.changeNerosidioCoinOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryNerosidioCoin(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	NerosidioCoinAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(NerosidioCoinAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	NerosidioCoins := []NerosidioCoin{
		NerosidioCoin{DisasterType: "Fire", Location: "[-33.009, 110.070]", donation: "5gb", Owner: "Brandon"},
		NerosidioCoin{DisasterType: "Fire", Location: "[-33.009, 110.070]", donation: "1gb", Owner: "Anita"},
		NerosidioCoin{DisasterType: "Fire", Location: "[-33.009, 110.070]", donation: "2gb", Owner: "Michel"},
		NerosidioCoin{DisasterType: "Fire", Location: "[-33.009, 110.070]", donation: "5gb", Owner: "Jade"},
		NerosidioCoin{DisasterType: "Fire", Location: "[-33.009, 110.070]", donation: "1gb", Owner: "Liz"}
	}

	i := 0
	for i < len(NerosidioCoins) {
		fmt.Println("i is ", i)
		NerosidioCoinAsBytes, _ := json.Marshal(NerosidioCoins[i])
		APIstub.PutState("NerosidioCoin"+strconv.Itoa(i), NerosidioCoinAsBytes)
		fmt.Println("Added", NerosidioCoins[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createNerosidioCoin(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of params. Disastertype, Location and Donation Amounts")
	}

	var NerosidioCoin = NerosidioCoin{DisasterType: args[1], Location: args[2], donation: args[3], Owner: args[4]}

	NerosidioCoinAsBytes, _ := json.Marshal(NerosidioCoin)
	APIstub.PutState(args[0], NerosidioCoinAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllNerosidioCoins(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "NerosidioCoin0"
	endKey := "NerosidioCoin999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllNerosidioCoins:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeNerosidioCoinOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Error changing owner")
	}

	NerosidioCoinAsBytes, _ := APIstub.GetState(args[0])
	NerosidioCoin := NerosidioCoin{}

	json.Unmarshal(NerosidioCoinAsBytes, &NerosidioCoin)
	NerosidioCoin.Owner = args[1]

	NerosidioCoinAsBytes, _ = json.Marshal(NerosidioCoin)
	APIstub.PutState(args[0], NerosidioCoinAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error submitting donation: %s", err)
	}
}