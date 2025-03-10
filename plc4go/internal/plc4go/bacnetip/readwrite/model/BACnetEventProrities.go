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

// BACnetEventProrities is the data-structure of this message
type BACnetEventProrities struct {
	OpeningTag  *BACnetOpeningTag
	ToOffnormal *BACnetApplicationTagUnsignedInteger
	ToFault     *BACnetApplicationTagUnsignedInteger
	ToNormal    *BACnetApplicationTagUnsignedInteger
	ClosingTag  *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetEventProrities is the corresponding interface of BACnetEventProrities
type IBACnetEventProrities interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetToOffnormal returns ToOffnormal (property field)
	GetToOffnormal() *BACnetApplicationTagUnsignedInteger
	// GetToFault returns ToFault (property field)
	GetToFault() *BACnetApplicationTagUnsignedInteger
	// GetToNormal returns ToNormal (property field)
	GetToNormal() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetEventProrities) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetEventProrities) GetToOffnormal() *BACnetApplicationTagUnsignedInteger {
	return m.ToOffnormal
}

func (m *BACnetEventProrities) GetToFault() *BACnetApplicationTagUnsignedInteger {
	return m.ToFault
}

func (m *BACnetEventProrities) GetToNormal() *BACnetApplicationTagUnsignedInteger {
	return m.ToNormal
}

func (m *BACnetEventProrities) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventProrities factory function for BACnetEventProrities
func NewBACnetEventProrities(openingTag *BACnetOpeningTag, toOffnormal *BACnetApplicationTagUnsignedInteger, toFault *BACnetApplicationTagUnsignedInteger, toNormal *BACnetApplicationTagUnsignedInteger, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetEventProrities {
	return &BACnetEventProrities{OpeningTag: openingTag, ToOffnormal: toOffnormal, ToFault: toFault, ToNormal: toNormal, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetEventProrities(structType interface{}) *BACnetEventProrities {
	if casted, ok := structType.(BACnetEventProrities); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventProrities); ok {
		return casted
	}
	return nil
}

func (m *BACnetEventProrities) GetTypeName() string {
	return "BACnetEventProrities"
}

func (m *BACnetEventProrities) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventProrities) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (toOffnormal)
	lengthInBits += m.ToOffnormal.GetLengthInBits()

	// Simple field (toFault)
	lengthInBits += m.ToFault.GetLengthInBits()

	// Simple field (toNormal)
	lengthInBits += m.ToNormal.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventProrities) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventProritiesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetEventProrities, error) {
	if pullErr := readBuffer.PullContext("BACnetEventProrities"); pullErr != nil {
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

	// Simple Field (toOffnormal)
	if pullErr := readBuffer.PullContext("toOffnormal"); pullErr != nil {
		return nil, pullErr
	}
	_toOffnormal, _toOffnormalErr := BACnetApplicationTagParse(readBuffer)
	if _toOffnormalErr != nil {
		return nil, errors.Wrap(_toOffnormalErr, "Error parsing 'toOffnormal' field")
	}
	toOffnormal := CastBACnetApplicationTagUnsignedInteger(_toOffnormal)
	if closeErr := readBuffer.CloseContext("toOffnormal"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toFault)
	if pullErr := readBuffer.PullContext("toFault"); pullErr != nil {
		return nil, pullErr
	}
	_toFault, _toFaultErr := BACnetApplicationTagParse(readBuffer)
	if _toFaultErr != nil {
		return nil, errors.Wrap(_toFaultErr, "Error parsing 'toFault' field")
	}
	toFault := CastBACnetApplicationTagUnsignedInteger(_toFault)
	if closeErr := readBuffer.CloseContext("toFault"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toNormal)
	if pullErr := readBuffer.PullContext("toNormal"); pullErr != nil {
		return nil, pullErr
	}
	_toNormal, _toNormalErr := BACnetApplicationTagParse(readBuffer)
	if _toNormalErr != nil {
		return nil, errors.Wrap(_toNormalErr, "Error parsing 'toNormal' field")
	}
	toNormal := CastBACnetApplicationTagUnsignedInteger(_toNormal)
	if closeErr := readBuffer.CloseContext("toNormal"); closeErr != nil {
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

	if closeErr := readBuffer.CloseContext("BACnetEventProrities"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetEventProrities(openingTag, toOffnormal, toFault, toNormal, closingTag, tagNumber), nil
}

func (m *BACnetEventProrities) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetEventProrities"); pushErr != nil {
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

	// Simple Field (toOffnormal)
	if pushErr := writeBuffer.PushContext("toOffnormal"); pushErr != nil {
		return pushErr
	}
	_toOffnormalErr := m.ToOffnormal.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("toOffnormal"); popErr != nil {
		return popErr
	}
	if _toOffnormalErr != nil {
		return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
	}

	// Simple Field (toFault)
	if pushErr := writeBuffer.PushContext("toFault"); pushErr != nil {
		return pushErr
	}
	_toFaultErr := m.ToFault.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("toFault"); popErr != nil {
		return popErr
	}
	if _toFaultErr != nil {
		return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
	}

	// Simple Field (toNormal)
	if pushErr := writeBuffer.PushContext("toNormal"); pushErr != nil {
		return pushErr
	}
	_toNormalErr := m.ToNormal.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("toNormal"); popErr != nil {
		return popErr
	}
	if _toNormalErr != nil {
		return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
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

	if popErr := writeBuffer.PopContext("BACnetEventProrities"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetEventProrities) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
