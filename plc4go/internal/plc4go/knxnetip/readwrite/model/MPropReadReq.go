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

// MPropReadReq is the data-structure of this message
type MPropReadReq struct {
	*CEMI
	InterfaceObjectType uint16
	ObjectInstance      uint8
	PropertyId          uint8
	NumberOfElements    uint8
	StartIndex          uint16

	// Arguments.
	Size uint16
}

// IMPropReadReq is the corresponding interface of MPropReadReq
type IMPropReadReq interface {
	ICEMI
	// GetInterfaceObjectType returns InterfaceObjectType (property field)
	GetInterfaceObjectType() uint16
	// GetObjectInstance returns ObjectInstance (property field)
	GetObjectInstance() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetNumberOfElements returns NumberOfElements (property field)
	GetNumberOfElements() uint8
	// GetStartIndex returns StartIndex (property field)
	GetStartIndex() uint16
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

func (m *MPropReadReq) GetMessageCode() uint8 {
	return 0xFC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *MPropReadReq) InitializeParent(parent *CEMI) {}

func (m *MPropReadReq) GetParent() *CEMI {
	return m.CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *MPropReadReq) GetInterfaceObjectType() uint16 {
	return m.InterfaceObjectType
}

func (m *MPropReadReq) GetObjectInstance() uint8 {
	return m.ObjectInstance
}

func (m *MPropReadReq) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *MPropReadReq) GetNumberOfElements() uint8 {
	return m.NumberOfElements
}

func (m *MPropReadReq) GetStartIndex() uint16 {
	return m.StartIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMPropReadReq factory function for MPropReadReq
func NewMPropReadReq(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, size uint16) *MPropReadReq {
	_result := &MPropReadReq{
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		CEMI:                NewCEMI(size),
	}
	_result.Child = _result
	return _result
}

func CastMPropReadReq(structType interface{}) *MPropReadReq {
	if casted, ok := structType.(MPropReadReq); ok {
		return &casted
	}
	if casted, ok := structType.(*MPropReadReq); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastMPropReadReq(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastMPropReadReq(casted.Child)
	}
	return nil
}

func (m *MPropReadReq) GetTypeName() string {
	return "MPropReadReq"
}

func (m *MPropReadReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *MPropReadReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (interfaceObjectType)
	lengthInBits += 16

	// Simple field (objectInstance)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 4

	// Simple field (startIndex)
	lengthInBits += 12

	return lengthInBits
}

func (m *MPropReadReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MPropReadReqParse(readBuffer utils.ReadBuffer, size uint16) (*MPropReadReq, error) {
	if pullErr := readBuffer.PullContext("MPropReadReq"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (interfaceObjectType)
	_interfaceObjectType, _interfaceObjectTypeErr := readBuffer.ReadUint16("interfaceObjectType", 16)
	if _interfaceObjectTypeErr != nil {
		return nil, errors.Wrap(_interfaceObjectTypeErr, "Error parsing 'interfaceObjectType' field")
	}
	interfaceObjectType := _interfaceObjectType

	// Simple Field (objectInstance)
	_objectInstance, _objectInstanceErr := readBuffer.ReadUint8("objectInstance", 8)
	if _objectInstanceErr != nil {
		return nil, errors.Wrap(_objectInstanceErr, "Error parsing 'objectInstance' field")
	}
	objectInstance := _objectInstance

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}
	propertyId := _propertyId

	// Simple Field (numberOfElements)
	_numberOfElements, _numberOfElementsErr := readBuffer.ReadUint8("numberOfElements", 4)
	if _numberOfElementsErr != nil {
		return nil, errors.Wrap(_numberOfElementsErr, "Error parsing 'numberOfElements' field")
	}
	numberOfElements := _numberOfElements

	// Simple Field (startIndex)
	_startIndex, _startIndexErr := readBuffer.ReadUint16("startIndex", 12)
	if _startIndexErr != nil {
		return nil, errors.Wrap(_startIndexErr, "Error parsing 'startIndex' field")
	}
	startIndex := _startIndex

	if closeErr := readBuffer.CloseContext("MPropReadReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MPropReadReq{
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		CEMI:                &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child, nil
}

func (m *MPropReadReq) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropReadReq"); pushErr != nil {
			return pushErr
		}

		// Simple Field (interfaceObjectType)
		interfaceObjectType := uint16(m.InterfaceObjectType)
		_interfaceObjectTypeErr := writeBuffer.WriteUint16("interfaceObjectType", 16, (interfaceObjectType))
		if _interfaceObjectTypeErr != nil {
			return errors.Wrap(_interfaceObjectTypeErr, "Error serializing 'interfaceObjectType' field")
		}

		// Simple Field (objectInstance)
		objectInstance := uint8(m.ObjectInstance)
		_objectInstanceErr := writeBuffer.WriteUint8("objectInstance", 8, (objectInstance))
		if _objectInstanceErr != nil {
			return errors.Wrap(_objectInstanceErr, "Error serializing 'objectInstance' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.PropertyId)
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (numberOfElements)
		numberOfElements := uint8(m.NumberOfElements)
		_numberOfElementsErr := writeBuffer.WriteUint8("numberOfElements", 4, (numberOfElements))
		if _numberOfElementsErr != nil {
			return errors.Wrap(_numberOfElementsErr, "Error serializing 'numberOfElements' field")
		}

		// Simple Field (startIndex)
		startIndex := uint16(m.StartIndex)
		_startIndexErr := writeBuffer.WriteUint16("startIndex", 12, (startIndex))
		if _startIndexErr != nil {
			return errors.Wrap(_startIndexErr, "Error serializing 'startIndex' field")
		}

		if popErr := writeBuffer.PopContext("MPropReadReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MPropReadReq) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
