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

// ModbusPDUReadFifoQueueRequest is the data-structure of this message
type ModbusPDUReadFifoQueueRequest struct {
	*ModbusPDU
	FifoPointerAddress uint16
}

// IModbusPDUReadFifoQueueRequest is the corresponding interface of ModbusPDUReadFifoQueueRequest
type IModbusPDUReadFifoQueueRequest interface {
	IModbusPDU
	// GetFifoPointerAddress returns FifoPointerAddress (property field)
	GetFifoPointerAddress() uint16
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

func (m *ModbusPDUReadFifoQueueRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadFifoQueueRequest) GetFunctionFlag() uint8 {
	return 0x18
}

func (m *ModbusPDUReadFifoQueueRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUReadFifoQueueRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUReadFifoQueueRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusPDUReadFifoQueueRequest) GetFifoPointerAddress() uint16 {
	return m.FifoPointerAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadFifoQueueRequest factory function for ModbusPDUReadFifoQueueRequest
func NewModbusPDUReadFifoQueueRequest(fifoPointerAddress uint16) *ModbusPDUReadFifoQueueRequest {
	_result := &ModbusPDUReadFifoQueueRequest{
		FifoPointerAddress: fifoPointerAddress,
		ModbusPDU:          NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUReadFifoQueueRequest(structType interface{}) *ModbusPDUReadFifoQueueRequest {
	if casted, ok := structType.(ModbusPDUReadFifoQueueRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadFifoQueueRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReadFifoQueueRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReadFifoQueueRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReadFifoQueueRequest) GetTypeName() string {
	return "ModbusPDUReadFifoQueueRequest"
}

func (m *ModbusPDUReadFifoQueueRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadFifoQueueRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fifoPointerAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUReadFifoQueueRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadFifoQueueRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUReadFifoQueueRequest, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadFifoQueueRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (fifoPointerAddress)
	_fifoPointerAddress, _fifoPointerAddressErr := readBuffer.ReadUint16("fifoPointerAddress", 16)
	if _fifoPointerAddressErr != nil {
		return nil, errors.Wrap(_fifoPointerAddressErr, "Error parsing 'fifoPointerAddress' field")
	}
	fifoPointerAddress := _fifoPointerAddress

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFifoQueueRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadFifoQueueRequest{
		FifoPointerAddress: fifoPointerAddress,
		ModbusPDU:          &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUReadFifoQueueRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFifoQueueRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (fifoPointerAddress)
		fifoPointerAddress := uint16(m.FifoPointerAddress)
		_fifoPointerAddressErr := writeBuffer.WriteUint16("fifoPointerAddress", 16, (fifoPointerAddress))
		if _fifoPointerAddressErr != nil {
			return errors.Wrap(_fifoPointerAddressErr, "Error serializing 'fifoPointerAddress' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFifoQueueRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadFifoQueueRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
