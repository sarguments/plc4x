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

// BACnetConstructedDataElement is the data-structure of this message
type BACnetConstructedDataElement struct {
	PeekedTagHeader *BACnetTagHeader
	ApplicationTag  *BACnetApplicationTag
	ContextTag      *BACnetContextTag
	ConstructedData *BACnetConstructedData

	// Arguments.
	ObjectType         BACnetObjectType
	PropertyIdentifier BACnetContextTagPropertyIdentifier
}

// IBACnetConstructedDataElement is the corresponding interface of BACnetConstructedDataElement
type IBACnetConstructedDataElement interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetApplicationTag returns ApplicationTag (property field)
	GetApplicationTag() *BACnetApplicationTag
	// GetContextTag returns ContextTag (property field)
	GetContextTag() *BACnetContextTag
	// GetConstructedData returns ConstructedData (property field)
	GetConstructedData() *BACnetConstructedData
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetIsApplicationTag returns IsApplicationTag (virtual field)
	GetIsApplicationTag() bool
	// GetIsConstructedData returns IsConstructedData (virtual field)
	GetIsConstructedData() bool
	// GetIsContextTag returns IsContextTag (virtual field)
	GetIsContextTag() bool
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

func (m *BACnetConstructedDataElement) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetConstructedDataElement) GetApplicationTag() *BACnetApplicationTag {
	return m.ApplicationTag
}

func (m *BACnetConstructedDataElement) GetContextTag() *BACnetContextTag {
	return m.ContextTag
}

func (m *BACnetConstructedDataElement) GetConstructedData() *BACnetConstructedData {
	return m.ConstructedData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataElement) GetPeekedTagNumber() uint8 {
	applicationTag := m.ApplicationTag
	_ = applicationTag
	contextTag := m.ContextTag
	_ = contextTag
	constructedData := m.ConstructedData
	_ = constructedData
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *BACnetConstructedDataElement) GetIsApplicationTag() bool {
	applicationTag := m.ApplicationTag
	_ = applicationTag
	contextTag := m.ContextTag
	_ = contextTag
	constructedData := m.ConstructedData
	_ = constructedData
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_APPLICATION_TAGS)))
}

func (m *BACnetConstructedDataElement) GetIsConstructedData() bool {
	applicationTag := m.ApplicationTag
	_ = applicationTag
	contextTag := m.ContextTag
	_ = contextTag
	constructedData := m.ConstructedData
	_ = constructedData
	return bool(bool(!(m.GetIsApplicationTag())) && bool(bool((m.GetPeekedTagHeader().GetLengthValueType()) == (0x6))))
}

func (m *BACnetConstructedDataElement) GetIsContextTag() bool {
	applicationTag := m.ApplicationTag
	_ = applicationTag
	contextTag := m.ContextTag
	_ = contextTag
	constructedData := m.ConstructedData
	_ = constructedData
	return bool(bool(!(m.GetIsConstructedData())) && bool(!(m.GetIsApplicationTag())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataElement factory function for BACnetConstructedDataElement
func NewBACnetConstructedDataElement(peekedTagHeader *BACnetTagHeader, applicationTag *BACnetApplicationTag, contextTag *BACnetContextTag, constructedData *BACnetConstructedData, objectType BACnetObjectType, propertyIdentifier BACnetContextTagPropertyIdentifier) *BACnetConstructedDataElement {
	return &BACnetConstructedDataElement{PeekedTagHeader: peekedTagHeader, ApplicationTag: applicationTag, ContextTag: contextTag, ConstructedData: constructedData, ObjectType: objectType, PropertyIdentifier: propertyIdentifier}
}

func CastBACnetConstructedDataElement(structType interface{}) *BACnetConstructedDataElement {
	if casted, ok := structType.(BACnetConstructedDataElement); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataElement); ok {
		return casted
	}
	return nil
}

func (m *BACnetConstructedDataElement) GetTypeName() string {
	return "BACnetConstructedDataElement"
}

func (m *BACnetConstructedDataElement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataElement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (applicationTag)
	if m.ApplicationTag != nil {
		lengthInBits += (*m.ApplicationTag).GetLengthInBits()
	}

	// Optional Field (contextTag)
	if m.ContextTag != nil {
		lengthInBits += (*m.ContextTag).GetLengthInBits()
	}

	// Optional Field (constructedData)
	if m.ConstructedData != nil {
		lengthInBits += (*m.ConstructedData).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConstructedDataElement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataElementParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType, propertyIdentifier *BACnetContextTagPropertyIdentifier) (*BACnetConstructedDataElement, error) {
	if pullErr := readBuffer.PullContext("BACnetConstructedDataElement"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = readBuffer.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Virtual field
	_isApplicationTag := bool((peekedTagHeader.GetTagClass()) == (TagClass_APPLICATION_TAGS))
	isApplicationTag := bool(_isApplicationTag)
	_ = isApplicationTag

	// Virtual field
	_isConstructedData := bool(!(isApplicationTag)) && bool(bool((peekedTagHeader.GetLengthValueType()) == (0x6)))
	isConstructedData := bool(_isConstructedData)
	_ = isConstructedData

	// Virtual field
	_isContextTag := bool(!(isConstructedData)) && bool(!(isApplicationTag))
	isContextTag := bool(_isContextTag)
	_ = isContextTag

	// Validation
	if !(bool(!(isContextTag)) || bool(bool(bool(isContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, utils.ParseValidationError{"unexpected closing tag"}
	}

	// Optional Field (applicationTag) (Can be skipped, if a given expression evaluates to false)
	var applicationTag *BACnetApplicationTag = nil
	if isApplicationTag {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("applicationTag"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'applicationTag' field")
		default:
			applicationTag = CastBACnetApplicationTag(_val)
			if closeErr := readBuffer.CloseContext("applicationTag"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (contextTag) (Can be skipped, if a given expression evaluates to false)
	var contextTag *BACnetContextTag = nil
	if isContextTag {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("contextTag"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, peekedTagNumber, GuessDataType(objectType, propertyIdentifier))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'contextTag' field")
		default:
			contextTag = CastBACnetContextTag(_val)
			if closeErr := readBuffer.CloseContext("contextTag"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (constructedData) (Can be skipped, if a given expression evaluates to false)
	var constructedData *BACnetConstructedData = nil
	if isConstructedData {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("constructedData"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, peekedTagNumber, objectType, propertyIdentifier)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'constructedData' field")
		default:
			constructedData = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("constructedData"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Validation
	if !(bool(bool(isApplicationTag) || bool(isContextTag)) || bool(isConstructedData)) {
		return nil, utils.ParseValidationError{"BACnetConstructedDataElement could not parse anything"}
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataElement"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetConstructedDataElement(peekedTagHeader, applicationTag, contextTag, constructedData, objectType, *propertyIdentifier), nil
}

func (m *BACnetConstructedDataElement) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetConstructedDataElement"); pushErr != nil {
		return pushErr
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	if _isApplicationTagErr := writeBuffer.WriteVirtual("isApplicationTag", m.GetIsApplicationTag()); _isApplicationTagErr != nil {
		return errors.Wrap(_isApplicationTagErr, "Error serializing 'isApplicationTag' field")
	}
	// Virtual field
	if _isConstructedDataErr := writeBuffer.WriteVirtual("isConstructedData", m.GetIsConstructedData()); _isConstructedDataErr != nil {
		return errors.Wrap(_isConstructedDataErr, "Error serializing 'isConstructedData' field")
	}
	// Virtual field
	if _isContextTagErr := writeBuffer.WriteVirtual("isContextTag", m.GetIsContextTag()); _isContextTagErr != nil {
		return errors.Wrap(_isContextTagErr, "Error serializing 'isContextTag' field")
	}

	// Optional Field (applicationTag) (Can be skipped, if the value is null)
	var applicationTag *BACnetApplicationTag = nil
	if m.ApplicationTag != nil {
		if pushErr := writeBuffer.PushContext("applicationTag"); pushErr != nil {
			return pushErr
		}
		applicationTag = m.ApplicationTag
		_applicationTagErr := applicationTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("applicationTag"); popErr != nil {
			return popErr
		}
		if _applicationTagErr != nil {
			return errors.Wrap(_applicationTagErr, "Error serializing 'applicationTag' field")
		}
	}

	// Optional Field (contextTag) (Can be skipped, if the value is null)
	var contextTag *BACnetContextTag = nil
	if m.ContextTag != nil {
		if pushErr := writeBuffer.PushContext("contextTag"); pushErr != nil {
			return pushErr
		}
		contextTag = m.ContextTag
		_contextTagErr := contextTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("contextTag"); popErr != nil {
			return popErr
		}
		if _contextTagErr != nil {
			return errors.Wrap(_contextTagErr, "Error serializing 'contextTag' field")
		}
	}

	// Optional Field (constructedData) (Can be skipped, if the value is null)
	var constructedData *BACnetConstructedData = nil
	if m.ConstructedData != nil {
		if pushErr := writeBuffer.PushContext("constructedData"); pushErr != nil {
			return pushErr
		}
		constructedData = m.ConstructedData
		_constructedDataErr := constructedData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("constructedData"); popErr != nil {
			return popErr
		}
		if _constructedDataErr != nil {
			return errors.Wrap(_constructedDataErr, "Error serializing 'constructedData' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetConstructedDataElement"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConstructedDataElement) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
