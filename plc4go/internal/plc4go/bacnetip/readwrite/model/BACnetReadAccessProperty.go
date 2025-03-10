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

// BACnetReadAccessProperty is the data-structure of this message
type BACnetReadAccessProperty struct {
	PropertyIdentifier  *BACnetContextTagPropertyIdentifier
	ArrayIndex          *BACnetContextTagUnsignedInteger
	PropertyValue       *BACnetConstructedData
	PropertyAccessError *BACnetReadAccessPropertyError

	// Arguments.
	ObjectType BACnetObjectType
}

// IBACnetReadAccessProperty is the corresponding interface of BACnetReadAccessProperty
type IBACnetReadAccessProperty interface {
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() *BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() *BACnetConstructedData
	// GetPropertyAccessError returns PropertyAccessError (property field)
	GetPropertyAccessError() *BACnetReadAccessPropertyError
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

func (m *BACnetReadAccessProperty) GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier {
	return m.PropertyIdentifier
}

func (m *BACnetReadAccessProperty) GetArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *BACnetReadAccessProperty) GetPropertyValue() *BACnetConstructedData {
	return m.PropertyValue
}

func (m *BACnetReadAccessProperty) GetPropertyAccessError() *BACnetReadAccessPropertyError {
	return m.PropertyAccessError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetReadAccessProperty factory function for BACnetReadAccessProperty
func NewBACnetReadAccessProperty(propertyIdentifier *BACnetContextTagPropertyIdentifier, arrayIndex *BACnetContextTagUnsignedInteger, propertyValue *BACnetConstructedData, propertyAccessError *BACnetReadAccessPropertyError, objectType BACnetObjectType) *BACnetReadAccessProperty {
	return &BACnetReadAccessProperty{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, PropertyAccessError: propertyAccessError, ObjectType: objectType}
}

func CastBACnetReadAccessProperty(structType interface{}) *BACnetReadAccessProperty {
	if casted, ok := structType.(BACnetReadAccessProperty); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetReadAccessProperty); ok {
		return casted
	}
	return nil
}

func (m *BACnetReadAccessProperty) GetTypeName() string {
	return "BACnetReadAccessProperty"
}

func (m *BACnetReadAccessProperty) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetReadAccessProperty) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).GetLengthInBits()
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += (*m.PropertyValue).GetLengthInBits()
	}

	// Optional Field (propertyAccessError)
	if m.PropertyAccessError != nil {
		lengthInBits += (*m.PropertyAccessError).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetReadAccessProperty) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetReadAccessPropertyParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType) (*BACnetReadAccessProperty, error) {
	if pullErr := readBuffer.PullContext("BACnetReadAccessProperty"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BACNET_PROPERTY_IDENTIFIER))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetContextTagPropertyIdentifier(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		default:
			arrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue *BACnetConstructedData = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(4), objectType, propertyIdentifier)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field")
		default:
			propertyValue = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (propertyAccessError) (Can be skipped, if a given expression evaluates to false)
	var propertyAccessError *BACnetReadAccessPropertyError = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("propertyAccessError"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetReadAccessPropertyErrorParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyAccessError' field")
		default:
			propertyAccessError = CastBACnetReadAccessPropertyError(_val)
			if closeErr := readBuffer.CloseContext("propertyAccessError"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessProperty"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetReadAccessProperty(propertyIdentifier, arrayIndex, propertyValue, propertyAccessError, objectType), nil
}

func (m *BACnetReadAccessProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetReadAccessProperty"); pushErr != nil {
		return pushErr
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return pushErr
	}
	_propertyIdentifierErr := m.PropertyIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return popErr
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	if m.ArrayIndex != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return pushErr
		}
		arrayIndex = m.ArrayIndex
		_arrayIndexErr := arrayIndex.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return popErr
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue *BACnetConstructedData = nil
	if m.PropertyValue != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return pushErr
		}
		propertyValue = m.PropertyValue
		_propertyValueErr := propertyValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return popErr
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (propertyAccessError) (Can be skipped, if the value is null)
	var propertyAccessError *BACnetReadAccessPropertyError = nil
	if m.PropertyAccessError != nil {
		if pushErr := writeBuffer.PushContext("propertyAccessError"); pushErr != nil {
			return pushErr
		}
		propertyAccessError = m.PropertyAccessError
		_propertyAccessErrorErr := propertyAccessError.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyAccessError"); popErr != nil {
			return popErr
		}
		if _propertyAccessErrorErr != nil {
			return errors.Wrap(_propertyAccessErrorErr, "Error serializing 'propertyAccessError' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetReadAccessProperty"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetReadAccessProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
