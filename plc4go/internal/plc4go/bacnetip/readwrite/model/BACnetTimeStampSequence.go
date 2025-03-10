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

// BACnetTimeStampSequence is the data-structure of this message
type BACnetTimeStampSequence struct {
	*BACnetTimeStamp
	SequenceNumber *BACnetContextTagUnsignedInteger
}

// IBACnetTimeStampSequence is the corresponding interface of BACnetTimeStampSequence
type IBACnetTimeStampSequence interface {
	IBACnetTimeStamp
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() *BACnetContextTagUnsignedInteger
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetTimeStampSequence) InitializeParent(parent *BACnetTimeStamp, peekedTagHeader *BACnetTagHeader) {
	m.BACnetTimeStamp.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetTimeStampSequence) GetParent() *BACnetTimeStamp {
	return m.BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetTimeStampSequence) GetSequenceNumber() *BACnetContextTagUnsignedInteger {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampSequence factory function for BACnetTimeStampSequence
func NewBACnetTimeStampSequence(sequenceNumber *BACnetContextTagUnsignedInteger, peekedTagHeader *BACnetTagHeader) *BACnetTimeStampSequence {
	_result := &BACnetTimeStampSequence{
		SequenceNumber:  sequenceNumber,
		BACnetTimeStamp: NewBACnetTimeStamp(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetTimeStampSequence(structType interface{}) *BACnetTimeStampSequence {
	if casted, ok := structType.(BACnetTimeStampSequence); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTimeStampSequence); ok {
		return casted
	}
	if casted, ok := structType.(BACnetTimeStamp); ok {
		return CastBACnetTimeStampSequence(casted.Child)
	}
	if casted, ok := structType.(*BACnetTimeStamp); ok {
		return CastBACnetTimeStampSequence(casted.Child)
	}
	return nil
}

func (m *BACnetTimeStampSequence) GetTypeName() string {
	return "BACnetTimeStampSequence"
}

func (m *BACnetTimeStampSequence) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTimeStampSequence) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (sequenceNumber)
	lengthInBits += m.SequenceNumber.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetTimeStampSequence) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimeStampSequenceParse(readBuffer utils.ReadBuffer) (*BACnetTimeStampSequence, error) {
	if pullErr := readBuffer.PullContext("BACnetTimeStampSequence"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (sequenceNumber)
	if pullErr := readBuffer.PullContext("sequenceNumber"); pullErr != nil {
		return nil, pullErr
	}
	_sequenceNumber, _sequenceNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field")
	}
	sequenceNumber := CastBACnetContextTagUnsignedInteger(_sequenceNumber)
	if closeErr := readBuffer.CloseContext("sequenceNumber"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampSequence"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetTimeStampSequence{
		SequenceNumber:  CastBACnetContextTagUnsignedInteger(sequenceNumber),
		BACnetTimeStamp: &BACnetTimeStamp{},
	}
	_child.BACnetTimeStamp.Child = _child
	return _child, nil
}

func (m *BACnetTimeStampSequence) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampSequence"); pushErr != nil {
			return pushErr
		}

		// Simple Field (sequenceNumber)
		if pushErr := writeBuffer.PushContext("sequenceNumber"); pushErr != nil {
			return pushErr
		}
		_sequenceNumberErr := m.SequenceNumber.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("sequenceNumber"); popErr != nil {
			return popErr
		}
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampSequence"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetTimeStampSequence) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
