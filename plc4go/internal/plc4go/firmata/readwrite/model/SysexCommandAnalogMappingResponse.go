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
)

// Code generated by code-generation. DO NOT EDIT.

// SysexCommandAnalogMappingResponse is the data-structure of this message
type SysexCommandAnalogMappingResponse struct {
	*SysexCommand
}

// ISysexCommandAnalogMappingResponse is the corresponding interface of SysexCommandAnalogMappingResponse
type ISysexCommandAnalogMappingResponse interface {
	ISysexCommand
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

func (m *SysexCommandAnalogMappingResponse) GetCommandType() uint8 {
	return 0x6A
}

func (m *SysexCommandAnalogMappingResponse) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SysexCommandAnalogMappingResponse) InitializeParent(parent *SysexCommand) {}

func (m *SysexCommandAnalogMappingResponse) GetParent() *SysexCommand {
	return m.SysexCommand
}

// NewSysexCommandAnalogMappingResponse factory function for SysexCommandAnalogMappingResponse
func NewSysexCommandAnalogMappingResponse() *SysexCommandAnalogMappingResponse {
	_result := &SysexCommandAnalogMappingResponse{
		SysexCommand: NewSysexCommand(),
	}
	_result.Child = _result
	return _result
}

func CastSysexCommandAnalogMappingResponse(structType interface{}) *SysexCommandAnalogMappingResponse {
	if casted, ok := structType.(SysexCommandAnalogMappingResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommandAnalogMappingResponse); ok {
		return casted
	}
	if casted, ok := structType.(SysexCommand); ok {
		return CastSysexCommandAnalogMappingResponse(casted.Child)
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return CastSysexCommandAnalogMappingResponse(casted.Child)
	}
	return nil
}

func (m *SysexCommandAnalogMappingResponse) GetTypeName() string {
	return "SysexCommandAnalogMappingResponse"
}

func (m *SysexCommandAnalogMappingResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommandAnalogMappingResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandAnalogMappingResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandAnalogMappingResponseParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommandAnalogMappingResponse, error) {
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandAnalogMappingResponse{
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child, nil
}

func (m *SysexCommandAnalogMappingResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandAnalogMappingResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
