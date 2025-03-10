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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesBoolean is the data-structure of this message
type BACnetPropertyStatesBoolean struct {
	*BACnetPropertyStates
	BooleanValue *BACnetContextTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetPropertyStatesBoolean is the corresponding interface of BACnetPropertyStatesBoolean
type IBACnetPropertyStatesBoolean interface {
	IBACnetPropertyStates
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() *BACnetContextTagBoolean
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

func (m *BACnetPropertyStatesBoolean) InitializeParent(parent *BACnetPropertyStates, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetPropertyStates.OpeningTag = openingTag
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
	m.BACnetPropertyStates.ClosingTag = closingTag
}

func (m *BACnetPropertyStatesBoolean) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesBoolean) GetBooleanValue() *BACnetContextTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBoolean factory function for BACnetPropertyStatesBoolean
func NewBACnetPropertyStatesBoolean(booleanValue *BACnetContextTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetPropertyStatesBoolean {
	_result := &BACnetPropertyStatesBoolean{
		BooleanValue:         booleanValue,
		BACnetPropertyStates: NewBACnetPropertyStates(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesBoolean(structType interface{}) *BACnetPropertyStatesBoolean {
	if casted, ok := structType.(BACnetPropertyStatesBoolean); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBoolean); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBoolean(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBoolean(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesBoolean) GetTypeName() string {
	return "BACnetPropertyStatesBoolean"
}

func (m *BACnetPropertyStatesBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (booleanValue)
	if m.BooleanValue != nil {
		lengthInBits += (*m.BooleanValue).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetPropertyStatesBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesBooleanParse(readBuffer utils.ReadBuffer, tagNumber uint8, peekedTagNumber uint8) (*BACnetPropertyStatesBoolean, error) {
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBoolean"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Optional Field (booleanValue) (Can be skipped, if a given expression evaluates to false)
	var booleanValue *BACnetContextTagBoolean = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, peekedTagNumber, BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'booleanValue' field")
		default:
			booleanValue = CastBACnetContextTagBoolean(_val)
			if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBoolean"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesBoolean{
		BooleanValue:         CastBACnetContextTagBoolean(booleanValue),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBoolean"); pushErr != nil {
			return pushErr
		}

		// Optional Field (booleanValue) (Can be skipped, if the value is null)
		var booleanValue *BACnetContextTagBoolean = nil
		if m.BooleanValue != nil {
			if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
				return pushErr
			}
			booleanValue = m.BooleanValue
			_booleanValueErr := booleanValue.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
				return popErr
			}
			if _booleanValueErr != nil {
				return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBoolean"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
