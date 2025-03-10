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

// BACnetTagPayloadObjectIdentifier is the data-structure of this message
type BACnetTagPayloadObjectIdentifier struct {
	ObjectType       BACnetObjectType
	ProprietaryValue uint16
	InstanceNumber   uint32
}

// IBACnetTagPayloadObjectIdentifier is the corresponding interface of BACnetTagPayloadObjectIdentifier
type IBACnetTagPayloadObjectIdentifier interface {
	// GetObjectType returns ObjectType (property field)
	GetObjectType() BACnetObjectType
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint16
	// GetInstanceNumber returns InstanceNumber (property field)
	GetInstanceNumber() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
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

func (m *BACnetTagPayloadObjectIdentifier) GetObjectType() BACnetObjectType {
	return m.ObjectType
}

func (m *BACnetTagPayloadObjectIdentifier) GetProprietaryValue() uint16 {
	return m.ProprietaryValue
}

func (m *BACnetTagPayloadObjectIdentifier) GetInstanceNumber() uint32 {
	return m.InstanceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetTagPayloadObjectIdentifier) GetIsProprietary() bool {
	return bool(bool((m.GetObjectType()) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadObjectIdentifier factory function for BACnetTagPayloadObjectIdentifier
func NewBACnetTagPayloadObjectIdentifier(objectType BACnetObjectType, proprietaryValue uint16, instanceNumber uint32) *BACnetTagPayloadObjectIdentifier {
	return &BACnetTagPayloadObjectIdentifier{ObjectType: objectType, ProprietaryValue: proprietaryValue, InstanceNumber: instanceNumber}
}

func CastBACnetTagPayloadObjectIdentifier(structType interface{}) *BACnetTagPayloadObjectIdentifier {
	if casted, ok := structType.(BACnetTagPayloadObjectIdentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTagPayloadObjectIdentifier); ok {
		return casted
	}
	return nil
}

func (m *BACnetTagPayloadObjectIdentifier) GetTypeName() string {
	return "BACnetTagPayloadObjectIdentifier"
}

func (m *BACnetTagPayloadObjectIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadObjectIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (objectType)
	lengthInBits += uint16(int32(10))

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(int32(0))

	// A virtual field doesn't have any in- or output.

	// Simple field (instanceNumber)
	lengthInBits += 22

	return lengthInBits
}

func (m *BACnetTagPayloadObjectIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadObjectIdentifierParse(readBuffer utils.ReadBuffer) (*BACnetTagPayloadObjectIdentifier, error) {
	if pullErr := readBuffer.PullContext("BACnetTagPayloadObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Manual Field (objectType)
	objectType, _objectTypeErr := ReadObjectType(readBuffer)
	if _objectTypeErr != nil {
		return nil, errors.Wrap(_objectTypeErr, "Error parsing 'objectType' field")
	}

	// Manual Field (proprietaryValue)
	proprietaryValue, _proprietaryValueErr := ReadProprietaryObjectType(readBuffer, objectType)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field")
	}

	// Virtual field
	_isProprietary := bool((objectType) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Simple Field (instanceNumber)
	_instanceNumber, _instanceNumberErr := readBuffer.ReadUint32("instanceNumber", 22)
	if _instanceNumberErr != nil {
		return nil, errors.Wrap(_instanceNumberErr, "Error parsing 'instanceNumber' field")
	}
	instanceNumber := _instanceNumber

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadObjectIdentifier(objectType, proprietaryValue, instanceNumber), nil
}

func (m *BACnetTagPayloadObjectIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadObjectIdentifier"); pushErr != nil {
		return pushErr
	}

	// Manual Field (objectType)
	_objectTypeErr := WriteObjectType(writeBuffer, m.GetObjectType())
	if _objectTypeErr != nil {
		return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryObjectType(writeBuffer, m.GetObjectType(), m.GetProprietaryValue())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}
	// Virtual field
	if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Simple Field (instanceNumber)
	instanceNumber := uint32(m.InstanceNumber)
	_instanceNumberErr := writeBuffer.WriteUint32("instanceNumber", 22, (instanceNumber))
	if _instanceNumberErr != nil {
		return errors.Wrap(_instanceNumberErr, "Error serializing 'instanceNumber' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadObjectIdentifier"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
