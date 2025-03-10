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

// BACnetReadAccessResult is the data-structure of this message
type BACnetReadAccessResult struct {
	ObjectIdentifier         *BACnetContextTagObjectIdentifier
	OpeningTag               *BACnetOpeningTag
	ListOfReadAccessProperty []*BACnetReadAccessProperty
	ClosingTag               *BACnetClosingTag
}

// IBACnetReadAccessResult is the corresponding interface of BACnetReadAccessResult
type IBACnetReadAccessResult interface {
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetListOfReadAccessProperty returns ListOfReadAccessProperty (property field)
	GetListOfReadAccessProperty() []*BACnetReadAccessProperty
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

func (m *BACnetReadAccessResult) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetReadAccessResult) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetReadAccessResult) GetListOfReadAccessProperty() []*BACnetReadAccessProperty {
	return m.ListOfReadAccessProperty
}

func (m *BACnetReadAccessResult) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetReadAccessResult factory function for BACnetReadAccessResult
func NewBACnetReadAccessResult(objectIdentifier *BACnetContextTagObjectIdentifier, openingTag *BACnetOpeningTag, listOfReadAccessProperty []*BACnetReadAccessProperty, closingTag *BACnetClosingTag) *BACnetReadAccessResult {
	return &BACnetReadAccessResult{ObjectIdentifier: objectIdentifier, OpeningTag: openingTag, ListOfReadAccessProperty: listOfReadAccessProperty, ClosingTag: closingTag}
}

func CastBACnetReadAccessResult(structType interface{}) *BACnetReadAccessResult {
	if casted, ok := structType.(BACnetReadAccessResult); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetReadAccessResult); ok {
		return casted
	}
	return nil
}

func (m *BACnetReadAccessResult) GetTypeName() string {
	return "BACnetReadAccessResult"
}

func (m *BACnetReadAccessResult) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetReadAccessResult) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfReadAccessProperty) > 0 {
		for _, element := range m.ListOfReadAccessProperty {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetReadAccessResult) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetReadAccessResultParse(readBuffer utils.ReadBuffer) (*BACnetReadAccessResult, error) {
	if pullErr := readBuffer.PullContext("BACnetReadAccessResult"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_OPENING_TAG))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (listOfReadAccessProperty)
	if pullErr := readBuffer.PullContext("listOfReadAccessProperty", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	listOfReadAccessProperty := make([]*BACnetReadAccessProperty, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, 1)) {
			_item, _err := BACnetReadAccessPropertyParse(readBuffer, objectIdentifier.GetObjectType())
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfReadAccessProperty' field")
			}
			listOfReadAccessProperty = append(listOfReadAccessProperty, CastBACnetReadAccessProperty(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfReadAccessProperty", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_CLOSING_TAG))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessResult"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetReadAccessResult(objectIdentifier, openingTag, listOfReadAccessProperty, closingTag), nil
}

func (m *BACnetReadAccessResult) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetReadAccessResult"); pushErr != nil {
		return pushErr
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return pushErr
	}
	_objectIdentifierErr := m.ObjectIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return popErr
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
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

	// Array Field (listOfReadAccessProperty)
	if m.ListOfReadAccessProperty != nil {
		if pushErr := writeBuffer.PushContext("listOfReadAccessProperty", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.ListOfReadAccessProperty {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'listOfReadAccessProperty' field")
			}
		}
		if popErr := writeBuffer.PopContext("listOfReadAccessProperty", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
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

	if popErr := writeBuffer.PopContext("BACnetReadAccessResult"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetReadAccessResult) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
