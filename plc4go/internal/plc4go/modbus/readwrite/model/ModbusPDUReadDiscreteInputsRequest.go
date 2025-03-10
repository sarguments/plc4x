/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUReadDiscreteInputsRequest is the data-structure of this message
type ModbusPDUReadDiscreteInputsRequest struct {
	*ModbusPDU
	StartingAddress uint16
	Quantity        uint16
}

// IModbusPDUReadDiscreteInputsRequest is the corresponding interface of ModbusPDUReadDiscreteInputsRequest
type IModbusPDUReadDiscreteInputsRequest interface {
	IModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *ModbusPDUReadDiscreteInputsRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadDiscreteInputsRequest) GetFunctionFlag() uint8 {
	return 0x02
}

func (m *ModbusPDUReadDiscreteInputsRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUReadDiscreteInputsRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUReadDiscreteInputsRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusPDUReadDiscreteInputsRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *ModbusPDUReadDiscreteInputsRequest) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadDiscreteInputsRequest factory function for ModbusPDUReadDiscreteInputsRequest
func NewModbusPDUReadDiscreteInputsRequest(startingAddress uint16, quantity uint16) *ModbusPDUReadDiscreteInputsRequest {
	_result := &ModbusPDUReadDiscreteInputsRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		ModbusPDU:       NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUReadDiscreteInputsRequest(structType interface{}) *ModbusPDUReadDiscreteInputsRequest {
	if casted, ok := structType.(ModbusPDUReadDiscreteInputsRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadDiscreteInputsRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReadDiscreteInputsRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReadDiscreteInputsRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReadDiscreteInputsRequest) GetTypeName() string {
	return "ModbusPDUReadDiscreteInputsRequest"
}

func (m *ModbusPDUReadDiscreteInputsRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadDiscreteInputsRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUReadDiscreteInputsRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadDiscreteInputsRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUReadDiscreteInputsRequest, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadDiscreteInputsRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field")
	}
	quantity := _quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDiscreteInputsRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadDiscreteInputsRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		ModbusPDU:       &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUReadDiscreteInputsRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDiscreteInputsRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.StartingAddress)
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.Quantity)
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDiscreteInputsRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadDiscreteInputsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
