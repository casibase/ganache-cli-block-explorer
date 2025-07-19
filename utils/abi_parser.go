package utils

import (
	"fmt"
	"ganache-cli-block-explorer/conf"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// ContractInfo Store contract name and its ABI
type ContractInfo struct {
	Name string
	ABI  abi.ABI
}

// ParsedTxData Store parsed transaction data
type ParsedTxData struct {
	MethodName      string            `json:"method_name"`
	MethodSignature string            `json:"method_signature"`
	Parameters      []ParsedParameter `json:"parameters,omitempty"`
	Error           string            `json:"error,omitempty"`
}

// ParsedParameter Store parsed parameter data
type ParsedParameter struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Value   interface{} `json:"value"`
	Indexed bool        `json:"indexed,omitempty"` // 用于事件参数
}

// var contractABIs map[string]ContractInfo
type ContractABIs map[string]ContractInfo

// LoadContractABIs Load all contract ABIs
func LoadContractABIs(contracts []conf.ContractConfig) (*ContractABIs, error) {
	contractABIs := make(ContractABIs)

	for _, contract := range contracts {
		// Read ABI file
		abiBytes, err := os.ReadFile(contract.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to read ABI file %s: %v", contract.Path, err)
		}

		// Parse ABI
		contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
		if err != nil {
			return nil, fmt.Errorf("failed to parse ABI for contract %s: %v", contract.Name, err)
		}

		contractABIs[contract.Name] = ContractInfo{
			Name: contract.Name,
			ABI:  contractABI,
		}
	}

	return &contractABIs, nil
}

// ParseTransactionData Parse transaction data using loaded contract ABIs
// This function will try to match the method ID with the methods in the loaded ABIs
// If a match is found, it will unpack the input data and return the method name,
// method signature, and parameters.
func (contractABIs *ContractABIs) ParseTransactionData(data []byte) *ParsedTxData {
	if len(data) < 4 {
		return &ParsedTxData{
			Error: "transaction data too short",
		}
	}

	// Extract method selector (first 4 bytes)
	methodID := data[:4]

	// Try to match method in all contracts
	for _, contractInfo := range *contractABIs {
		method, err := contractInfo.ABI.MethodById(methodID)
		if err != nil {
			continue // Method not found in this contract, try next
		}

		// Found matching method, unpack parameters
		inputs, err := method.Inputs.Unpack(data[4:])
		if err != nil {
			return &ParsedTxData{
				MethodName:      method.Name,
				MethodSignature: method.Sig,
				Error:           fmt.Sprintf("failed to unpack inputs: %v", err),
			}
		}

		// Build parameter list
		var parameters []ParsedParameter
		for i, input := range method.Inputs {
			if i < len(inputs) {
				parameters = append(parameters, ParsedParameter{
					Name:    input.Name,
					Type:    input.Type.String(),
					Value:   formatValue(inputs[i]),
					Indexed: false, // Method parameters are not indexed
				})
			}
		}

		return &ParsedTxData{
			MethodName:      method.Name,
			MethodSignature: method.Sig,
			Parameters:      parameters,
		}
	}

	return &ParsedTxData{
		Error: "method not found in any loaded contract ABI",
	}
}

// formatValue Format value to a readable form
func formatValue(value interface{}) interface{} {
	switch v := value.(type) {
	case common.Address:
		return v.Hex()
	case []byte:
		return common.Bytes2Hex(v)
	case string:
		return v
	default:
		// For other types, just return the value as is
		return fmt.Sprintf("%v", v)
	}
}

// ParseEventLogs Parse event logs using loaded contract ABIs
// This function will try to match the event ID with the events in the loaded ABIs
// If a match is found, it will unpack the event data and return the event name,
// method signature, and parameters.
// It handles both indexed and non-indexed parameters.
// The first topic is the event ID, and the rest are indexed parameters.
// Non-indexed parameters are in the data field.
func (contractABIs *ContractABIs) ParseEventLogs(topics []common.Hash, data []byte) *ParsedTxData {
	if len(topics) == 0 {
		return &ParsedTxData{
			Error: "no topics in log",
		}
	}

	eventID := topics[0]

	// Try to match event in all contracts
	for _, contractInfo := range *contractABIs {
		event, err := contractInfo.ABI.EventByID(eventID)
		if err != nil {
			continue // Event not found in this contract, try next
		}

		// Parse event data
		eventData := make(map[string]interface{})

		// Parse non-indexed parameters (in data)
		if len(data) > 0 {
			nonIndexedArgs := make([]interface{}, 0)
			for _, input := range event.Inputs {
				if !input.Indexed {
					nonIndexedArgs = append(nonIndexedArgs, new(interface{}))
				}
			}

			if len(nonIndexedArgs) > 0 {
				err = event.Inputs.UnpackIntoMap(eventData, data)
				if err != nil {
					return &ParsedTxData{
						MethodName: event.Name,
						Error:      fmt.Sprintf("failed to unpack event data: %v", err),
					}
				}
			}
		}

		// Parse indexed parameters (in topics, skip first topic as it's event ID)
		topicIndex := 1
		for _, input := range event.Inputs {
			if input.Indexed && topicIndex < len(topics) {
				// Simple handling of indexed parameters
				eventData[input.Name] = topics[topicIndex].Hex()
				topicIndex++
			}
		}

		// Build parameter list
		var parameters []ParsedParameter
		for _, input := range event.Inputs {
			if value, exists := eventData[input.Name]; exists {
				parameters = append(parameters, ParsedParameter{
					Name:    input.Name,
					Type:    input.Type.String(),
					Value:   formatValue(value),
					Indexed: input.Indexed,
				})
			}
		}

		return &ParsedTxData{
			MethodName:      event.Name,
			MethodSignature: event.Sig,
			Parameters:      parameters,
		}
	}

	return &ParsedTxData{
		Error: "event not found in any loaded contract ABI",
	}
}
