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

// ApduDataExtLinkRead is the data-structure of this message
type ApduDataExtLinkRead struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtLinkRead is the corresponding interface of ApduDataExtLinkRead
type IApduDataExtLinkRead interface {
	IApduDataExt
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

func (m *ApduDataExtLinkRead) GetExtApciType() uint8 {
	return 0x25
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtLinkRead) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtLinkRead) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtLinkRead factory function for ApduDataExtLinkRead
func NewApduDataExtLinkRead(length uint8) *ApduDataExtLinkRead {
	_result := &ApduDataExtLinkRead{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtLinkRead(structType interface{}) *ApduDataExtLinkRead {
	if casted, ok := structType.(ApduDataExtLinkRead); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtLinkRead); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtLinkRead(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtLinkRead(casted.Child)
	}
	return nil
}

func (m *ApduDataExtLinkRead) GetTypeName() string {
	return "ApduDataExtLinkRead"
}

func (m *ApduDataExtLinkRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtLinkRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtLinkRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtLinkReadParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtLinkRead, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtLinkRead"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtLinkRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtLinkRead{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtLinkRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtLinkRead"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtLinkRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtLinkRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
