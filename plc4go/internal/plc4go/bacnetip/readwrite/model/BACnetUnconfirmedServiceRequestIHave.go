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

// BACnetUnconfirmedServiceRequestIHave is the data-structure of this message
type BACnetUnconfirmedServiceRequestIHave struct {
	*BACnetUnconfirmedServiceRequest
	DeviceIdentifier *BACnetApplicationTagObjectIdentifier
	ObjectIdentifier *BACnetApplicationTagObjectIdentifier
	ObjectName       *BACnetApplicationTagCharacterString

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetUnconfirmedServiceRequestIHave is the corresponding interface of BACnetUnconfirmedServiceRequestIHave
type IBACnetUnconfirmedServiceRequestIHave interface {
	IBACnetUnconfirmedServiceRequest
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() *BACnetApplicationTagObjectIdentifier
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetApplicationTagObjectIdentifier
	// GetObjectName returns ObjectName (property field)
	GetObjectName() *BACnetApplicationTagCharacterString
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

func (m *BACnetUnconfirmedServiceRequestIHave) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_I_HAVE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestIHave) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestIHave) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetUnconfirmedServiceRequestIHave) GetDeviceIdentifier() *BACnetApplicationTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *BACnetUnconfirmedServiceRequestIHave) GetObjectIdentifier() *BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetUnconfirmedServiceRequestIHave) GetObjectName() *BACnetApplicationTagCharacterString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestIHave factory function for BACnetUnconfirmedServiceRequestIHave
func NewBACnetUnconfirmedServiceRequestIHave(deviceIdentifier *BACnetApplicationTagObjectIdentifier, objectIdentifier *BACnetApplicationTagObjectIdentifier, objectName *BACnetApplicationTagCharacterString, serviceRequestLength uint16) *BACnetUnconfirmedServiceRequestIHave {
	_result := &BACnetUnconfirmedServiceRequestIHave{
		DeviceIdentifier:                deviceIdentifier,
		ObjectIdentifier:                objectIdentifier,
		ObjectName:                      objectName,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetUnconfirmedServiceRequestIHave(structType interface{}) *BACnetUnconfirmedServiceRequestIHave {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestIHave); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestIHave); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestIHave(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestIHave(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestIHave) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestIHave"
}

func (m *BACnetUnconfirmedServiceRequestIHave) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestIHave) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits()

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestIHave) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestIHaveParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetUnconfirmedServiceRequestIHave, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestIHave"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (deviceIdentifier)
	if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_deviceIdentifier, _deviceIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _deviceIdentifierErr != nil {
		return nil, errors.Wrap(_deviceIdentifierErr, "Error parsing 'deviceIdentifier' field")
	}
	deviceIdentifier := CastBACnetApplicationTagObjectIdentifier(_deviceIdentifier)
	if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetApplicationTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (objectName)
	if pullErr := readBuffer.PullContext("objectName"); pullErr != nil {
		return nil, pullErr
	}
	_objectName, _objectNameErr := BACnetApplicationTagParse(readBuffer)
	if _objectNameErr != nil {
		return nil, errors.Wrap(_objectNameErr, "Error parsing 'objectName' field")
	}
	objectName := CastBACnetApplicationTagCharacterString(_objectName)
	if closeErr := readBuffer.CloseContext("objectName"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestIHave"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestIHave{
		DeviceIdentifier:                CastBACnetApplicationTagObjectIdentifier(deviceIdentifier),
		ObjectIdentifier:                CastBACnetApplicationTagObjectIdentifier(objectIdentifier),
		ObjectName:                      CastBACnetApplicationTagCharacterString(objectName),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetUnconfirmedServiceRequestIHave) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestIHave"); pushErr != nil {
			return pushErr
		}

		// Simple Field (deviceIdentifier)
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return pushErr
		}
		_deviceIdentifierErr := m.DeviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return popErr
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
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

		// Simple Field (objectName)
		if pushErr := writeBuffer.PushContext("objectName"); pushErr != nil {
			return pushErr
		}
		_objectNameErr := m.ObjectName.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("objectName"); popErr != nil {
			return popErr
		}
		if _objectNameErr != nil {
			return errors.Wrap(_objectNameErr, "Error serializing 'objectName' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestIHave"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestIHave) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
