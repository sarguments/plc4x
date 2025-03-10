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

// S7Parameter is the data-structure of this message
type S7Parameter struct {
	Child IS7ParameterChild
}

// IS7Parameter is the corresponding interface of S7Parameter
type IS7Parameter interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7ParameterParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7Parameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7ParameterChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7Parameter)
	GetParent() *S7Parameter

	GetTypeName() string
	IS7Parameter
}

// NewS7Parameter factory function for S7Parameter
func NewS7Parameter() *S7Parameter {
	return &S7Parameter{}
}

func CastS7Parameter(structType interface{}) *S7Parameter {
	if casted, ok := structType.(S7Parameter); ok {
		return &casted
	}
	if casted, ok := structType.(*S7Parameter); ok {
		return casted
	}
	if casted, ok := structType.(IS7ParameterChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *S7Parameter) GetTypeName() string {
	return "S7Parameter"
}

func (m *S7Parameter) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7Parameter) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *S7Parameter) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (parameterType)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7Parameter) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterParse(readBuffer utils.ReadBuffer, messageType uint8) (*S7Parameter, error) {
	if pullErr := readBuffer.PullContext("S7Parameter"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType, _parameterTypeErr := readBuffer.ReadUint8("parameterType", 8)
	if _parameterTypeErr != nil {
		return nil, errors.Wrap(_parameterTypeErr, "Error parsing 'parameterType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7ParameterChild interface {
		InitializeParent(*S7Parameter)
		GetParent() *S7Parameter
	}
	var _child S7ParameterChild
	var typeSwitchError error
	switch {
	case parameterType == 0xF0: // S7ParameterSetupCommunication
		_child, typeSwitchError = S7ParameterSetupCommunicationParse(readBuffer, messageType)
	case parameterType == 0x04 && messageType == 0x01: // S7ParameterReadVarRequest
		_child, typeSwitchError = S7ParameterReadVarRequestParse(readBuffer, messageType)
	case parameterType == 0x04 && messageType == 0x03: // S7ParameterReadVarResponse
		_child, typeSwitchError = S7ParameterReadVarResponseParse(readBuffer, messageType)
	case parameterType == 0x05 && messageType == 0x01: // S7ParameterWriteVarRequest
		_child, typeSwitchError = S7ParameterWriteVarRequestParse(readBuffer, messageType)
	case parameterType == 0x05 && messageType == 0x03: // S7ParameterWriteVarResponse
		_child, typeSwitchError = S7ParameterWriteVarResponseParse(readBuffer, messageType)
	case parameterType == 0x00 && messageType == 0x07: // S7ParameterUserData
		_child, typeSwitchError = S7ParameterUserDataParse(readBuffer, messageType)
	case parameterType == 0x01 && messageType == 0x07: // S7ParameterModeTransition
		_child, typeSwitchError = S7ParameterModeTransitionParse(readBuffer, messageType)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7Parameter"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *S7Parameter) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7Parameter) SerializeParent(writeBuffer utils.WriteBuffer, child IS7Parameter, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7Parameter"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType := uint8(child.GetParameterType())
	_parameterTypeErr := writeBuffer.WriteUint8("parameterType", 8, (parameterType))

	if _parameterTypeErr != nil {
		return errors.Wrap(_parameterTypeErr, "Error serializing 'parameterType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Parameter"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7Parameter) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
