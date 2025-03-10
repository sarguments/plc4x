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

// BACnetDeviceObjectReference is the data-structure of this message
type BACnetDeviceObjectReference struct {
	DeviceIdentifier *BACnetContextTagObjectIdentifier
	ObjectIdentifier *BACnetContextTagObjectIdentifier
}

// IBACnetDeviceObjectReference is the corresponding interface of BACnetDeviceObjectReference
type IBACnetDeviceObjectReference interface {
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() *BACnetContextTagObjectIdentifier
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
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

func (m *BACnetDeviceObjectReference) GetDeviceIdentifier() *BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *BACnetDeviceObjectReference) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDeviceObjectReference factory function for BACnetDeviceObjectReference
func NewBACnetDeviceObjectReference(deviceIdentifier *BACnetContextTagObjectIdentifier, objectIdentifier *BACnetContextTagObjectIdentifier) *BACnetDeviceObjectReference {
	return &BACnetDeviceObjectReference{DeviceIdentifier: deviceIdentifier, ObjectIdentifier: objectIdentifier}
}

func CastBACnetDeviceObjectReference(structType interface{}) *BACnetDeviceObjectReference {
	if casted, ok := structType.(BACnetDeviceObjectReference); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetDeviceObjectReference); ok {
		return casted
	}
	return nil
}

func (m *BACnetDeviceObjectReference) GetTypeName() string {
	return "BACnetDeviceObjectReference"
}

func (m *BACnetDeviceObjectReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetDeviceObjectReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += (*m.DeviceIdentifier).GetLengthInBits()
	}

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetDeviceObjectReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDeviceObjectReferenceParse(readBuffer utils.ReadBuffer) (*BACnetDeviceObjectReference, error) {
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectReference"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Optional Field (deviceIdentifier) (Can be skipped, if a given expression evaluates to false)
	var deviceIdentifier *BACnetContextTagObjectIdentifier = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceIdentifier' field")
		default:
			deviceIdentifier = CastBACnetContextTagObjectIdentifier(_val)
			if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectReference"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetDeviceObjectReference(deviceIdentifier, objectIdentifier), nil
}

func (m *BACnetDeviceObjectReference) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectReference"); pushErr != nil {
		return pushErr
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
	var deviceIdentifier *BACnetContextTagObjectIdentifier = nil
	if m.DeviceIdentifier != nil {
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return pushErr
		}
		deviceIdentifier = m.DeviceIdentifier
		_deviceIdentifierErr := deviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return popErr
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}
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

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectReference"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetDeviceObjectReference) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
