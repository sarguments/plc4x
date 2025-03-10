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

// BACnetTimeStampEnclosed is the data-structure of this message
type BACnetTimeStampEnclosed struct {
	OpeningTag *BACnetOpeningTag
	Timestamp  *BACnetTimeStamp
	ClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetTimeStampEnclosed is the corresponding interface of BACnetTimeStampEnclosed
type IBACnetTimeStampEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() *BACnetTimeStamp
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetTimeStampEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetTimeStampEnclosed) GetTimestamp() *BACnetTimeStamp {
	return m.Timestamp
}

func (m *BACnetTimeStampEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampEnclosed factory function for BACnetTimeStampEnclosed
func NewBACnetTimeStampEnclosed(openingTag *BACnetOpeningTag, timestamp *BACnetTimeStamp, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetTimeStampEnclosed {
	return &BACnetTimeStampEnclosed{OpeningTag: openingTag, Timestamp: timestamp, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetTimeStampEnclosed(structType interface{}) *BACnetTimeStampEnclosed {
	if casted, ok := structType.(BACnetTimeStampEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTimeStampEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetTimeStampEnclosed) GetTypeName() string {
	return "BACnetTimeStampEnclosed"
}

func (m *BACnetTimeStampEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTimeStampEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetTimeStampEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimeStampEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetTimeStampEnclosed, error) {
	if pullErr := readBuffer.PullContext("BACnetTimeStampEnclosed"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_OPENING_TAG))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, pullErr
	}
	_timestamp, _timestampErr := BACnetTimeStampParse(readBuffer)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}
	timestamp := CastBACnetTimeStamp(_timestamp)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_CLOSING_TAG))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampEnclosed"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTimeStampEnclosed(openingTag, timestamp, closingTag, tagNumber), nil
}

func (m *BACnetTimeStampEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTimeStampEnclosed"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (timestamp)
	if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
		return pushErr
	}
	_timestampErr := m.Timestamp.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
		return popErr
	}
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeStampEnclosed"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTimeStampEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
