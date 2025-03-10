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

// BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue is the data-structure of this message
type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue struct {
	PropertyIdentifier *BACnetContextTagPropertyIdentifier
	ArrayIndex         *BACnetContextTagUnsignedInteger
	PropertyValue      *BACnetConstructedData
	TimeOfChange       *BACnetContextTagTime

	// Arguments.
	ObjectType BACnetObjectType
}

// IBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue is the corresponding interface of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue
type IBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue interface {
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() *BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() *BACnetConstructedData
	// GetTimeOfChange returns TimeOfChange (property field)
	GetTimeOfChange() *BACnetContextTagTime
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

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier {
	return m.PropertyIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) GetArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) GetPropertyValue() *BACnetConstructedData {
	return m.PropertyValue
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) GetTimeOfChange() *BACnetContextTagTime {
	return m.TimeOfChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue factory function for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue
func NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue(propertyIdentifier *BACnetContextTagPropertyIdentifier, arrayIndex *BACnetContextTagUnsignedInteger, propertyValue *BACnetConstructedData, timeOfChange *BACnetContextTagTime, objectType BACnetObjectType) *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue {
	return &BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, TimeOfChange: timeOfChange, ObjectType: objectType}
}

func CastBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue(structType interface{}) *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue); ok {
		return casted
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue"
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).GetLengthInBits()
	}

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits()

	// Optional Field (timeOfChange)
	if m.TimeOfChange != nil {
		lengthInBits += (*m.TimeOfChange).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValueParse(readBuffer utils.ReadBuffer, objectType BACnetObjectType) (*BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_PROPERTY_IDENTIFIER))
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
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
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

	// Simple Field (propertyValue)
	if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
		return nil, pullErr
	}
	_propertyValue, _propertyValueErr := BACnetConstructedDataParse(readBuffer, uint8(uint8(2)), BACnetObjectType(objectType), propertyIdentifier)
	if _propertyValueErr != nil {
		return nil, errors.Wrap(_propertyValueErr, "Error parsing 'propertyValue' field")
	}
	propertyValue := CastBACnetConstructedData(_propertyValue)
	if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (timeOfChange) (Can be skipped, if a given expression evaluates to false)
	var timeOfChange *BACnetContextTagTime = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("timeOfChange"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_TIME)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'timeOfChange' field")
		default:
			timeOfChange = CastBACnetContextTagTime(_val)
			if closeErr := readBuffer.CloseContext("timeOfChange"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue(propertyIdentifier, arrayIndex, propertyValue, timeOfChange, objectType), nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue"); pushErr != nil {
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

	// Simple Field (propertyValue)
	if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
		return pushErr
	}
	_propertyValueErr := m.PropertyValue.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
		return popErr
	}
	if _propertyValueErr != nil {
		return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
	}

	// Optional Field (timeOfChange) (Can be skipped, if the value is null)
	var timeOfChange *BACnetContextTagTime = nil
	if m.TimeOfChange != nil {
		if pushErr := writeBuffer.PushContext("timeOfChange"); pushErr != nil {
			return pushErr
		}
		timeOfChange = m.TimeOfChange
		_timeOfChangeErr := timeOfChange.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeOfChange"); popErr != nil {
			return popErr
		}
		if _timeOfChangeErr != nil {
			return errors.Wrap(_timeOfChangeErr, "Error serializing 'timeOfChange' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
