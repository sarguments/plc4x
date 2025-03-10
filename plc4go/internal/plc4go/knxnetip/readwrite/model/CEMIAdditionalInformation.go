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

// CEMIAdditionalInformation is the data-structure of this message
type CEMIAdditionalInformation struct {
	Child ICEMIAdditionalInformationChild
}

// ICEMIAdditionalInformation is the corresponding interface of CEMIAdditionalInformation
type ICEMIAdditionalInformation interface {
	// GetAdditionalInformationType returns AdditionalInformationType (discriminator field)
	GetAdditionalInformationType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICEMIAdditionalInformationParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICEMIAdditionalInformation, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICEMIAdditionalInformationChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *CEMIAdditionalInformation)
	GetParent() *CEMIAdditionalInformation

	GetTypeName() string
	ICEMIAdditionalInformation
}

// NewCEMIAdditionalInformation factory function for CEMIAdditionalInformation
func NewCEMIAdditionalInformation() *CEMIAdditionalInformation {
	return &CEMIAdditionalInformation{}
}

func CastCEMIAdditionalInformation(structType interface{}) *CEMIAdditionalInformation {
	if casted, ok := structType.(CEMIAdditionalInformation); ok {
		return &casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformation); ok {
		return casted
	}
	if casted, ok := structType.(ICEMIAdditionalInformationChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *CEMIAdditionalInformation) GetTypeName() string {
	return "CEMIAdditionalInformation"
}

func (m *CEMIAdditionalInformation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CEMIAdditionalInformation) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *CEMIAdditionalInformation) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (additionalInformationType)
	lengthInBits += 8

	return lengthInBits
}

func (m *CEMIAdditionalInformation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CEMIAdditionalInformationParse(readBuffer utils.ReadBuffer) (*CEMIAdditionalInformation, error) {
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformation"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (additionalInformationType) (Used as input to a switch field)
	additionalInformationType, _additionalInformationTypeErr := readBuffer.ReadUint8("additionalInformationType", 8)
	if _additionalInformationTypeErr != nil {
		return nil, errors.Wrap(_additionalInformationTypeErr, "Error parsing 'additionalInformationType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CEMIAdditionalInformationChild interface {
		InitializeParent(*CEMIAdditionalInformation)
		GetParent() *CEMIAdditionalInformation
	}
	var _child CEMIAdditionalInformationChild
	var typeSwitchError error
	switch {
	case additionalInformationType == 0x03: // CEMIAdditionalInformationBusmonitorInfo
		_child, typeSwitchError = CEMIAdditionalInformationBusmonitorInfoParse(readBuffer)
	case additionalInformationType == 0x04: // CEMIAdditionalInformationRelativeTimestamp
		_child, typeSwitchError = CEMIAdditionalInformationRelativeTimestampParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformation"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *CEMIAdditionalInformation) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *CEMIAdditionalInformation) SerializeParent(writeBuffer utils.WriteBuffer, child ICEMIAdditionalInformation, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("CEMIAdditionalInformation"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (additionalInformationType) (Used as input to a switch field)
	additionalInformationType := uint8(child.GetAdditionalInformationType())
	_additionalInformationTypeErr := writeBuffer.WriteUint8("additionalInformationType", 8, (additionalInformationType))

	if _additionalInformationTypeErr != nil {
		return errors.Wrap(_additionalInformationTypeErr, "Error serializing 'additionalInformationType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CEMIAdditionalInformation"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *CEMIAdditionalInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
