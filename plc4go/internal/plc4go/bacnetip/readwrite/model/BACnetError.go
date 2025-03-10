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

// BACnetError is the data-structure of this message
type BACnetError struct {
	ErrorClass *BACnetApplicationTagEnumerated
	ErrorCode  *BACnetApplicationTagEnumerated
	Child      IBACnetErrorChild
}

// IBACnetError is the corresponding interface of BACnetError
type IBACnetError interface {
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() uint8
	// GetErrorClass returns ErrorClass (property field)
	GetErrorClass() *BACnetApplicationTagEnumerated
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() *BACnetApplicationTagEnumerated
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetErrorParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetError, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetErrorChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetError, errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated)
	GetParent() *BACnetError

	GetTypeName() string
	IBACnetError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetError) GetErrorClass() *BACnetApplicationTagEnumerated {
	return m.ErrorClass
}

func (m *BACnetError) GetErrorCode() *BACnetApplicationTagEnumerated {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetError factory function for BACnetError
func NewBACnetError(errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) *BACnetError {
	return &BACnetError{ErrorClass: errorClass, ErrorCode: errorCode}
}

func CastBACnetError(structType interface{}) *BACnetError {
	if casted, ok := structType.(BACnetError); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetError); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetErrorChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetError) GetTypeName() string {
	return "BACnetError"
}

func (m *BACnetError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetError) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetError) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	// Simple field (errorClass)
	lengthInBits += m.ErrorClass.GetLengthInBits()

	// Simple field (errorCode)
	lengthInBits += m.ErrorCode.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetErrorParse(readBuffer utils.ReadBuffer) (*BACnetError, error) {
	if pullErr := readBuffer.PullContext("BACnetError"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetErrorChild interface {
		InitializeParent(*BACnetError, *BACnetApplicationTagEnumerated, *BACnetApplicationTagEnumerated)
		GetParent() *BACnetError
	}
	var _child BACnetErrorChild
	var typeSwitchError error
	switch {
	case serviceChoice == 0x00: // BACnetErrorAcknowledgeAlarm
		_child, typeSwitchError = BACnetErrorAcknowledgeAlarmParse(readBuffer)
	case serviceChoice == 0x03: // BACnetErrorGetAlarmSummary
		_child, typeSwitchError = BACnetErrorGetAlarmSummaryParse(readBuffer)
	case serviceChoice == 0x02: // BACnetErrorConfirmedEventNotification
		_child, typeSwitchError = BACnetErrorConfirmedEventNotificationParse(readBuffer)
	case serviceChoice == 0x04: // BACnetErrorGetEnrollmentSummary
		_child, typeSwitchError = BACnetErrorGetEnrollmentSummaryParse(readBuffer)
	case serviceChoice == 0x05: // BACnetErrorDeviceCommunicationProtocol
		_child, typeSwitchError = BACnetErrorDeviceCommunicationProtocolParse(readBuffer)
	case serviceChoice == 0x1D: // BACnetErrorGetEventInformation
		_child, typeSwitchError = BACnetErrorGetEventInformationParse(readBuffer)
	case serviceChoice == 0x06: // BACnetErrorAtomicReadFile
		_child, typeSwitchError = BACnetErrorAtomicReadFileParse(readBuffer)
	case serviceChoice == 0x07: // BACnetErrorAtomicWriteFile
		_child, typeSwitchError = BACnetErrorAtomicWriteFileParse(readBuffer)
	case serviceChoice == 0x0A: // BACnetErrorCreateObject
		_child, typeSwitchError = BACnetErrorCreateObjectParse(readBuffer)
	case serviceChoice == 0x0C: // BACnetErrorReadProperty
		_child, typeSwitchError = BACnetErrorReadPropertyParse(readBuffer)
	case serviceChoice == 0x0E: // BACnetErrorReadPropertyMultiple
		_child, typeSwitchError = BACnetErrorReadPropertyMultipleParse(readBuffer)
	case serviceChoice == 0x0F: // BACnetErrorWriteProperty
		_child, typeSwitchError = BACnetErrorWritePropertyParse(readBuffer)
	case serviceChoice == 0x1A: // BACnetErrorReadRange
		_child, typeSwitchError = BACnetErrorReadRangeParse(readBuffer)
	case serviceChoice == 0x11: // BACnetErrorDeviceCommunicationProtocol
		_child, typeSwitchError = BACnetErrorDeviceCommunicationProtocolParse(readBuffer)
	case serviceChoice == 0x12: // BACnetErrorConfirmedPrivateTransfer
		_child, typeSwitchError = BACnetErrorConfirmedPrivateTransferParse(readBuffer)
	case serviceChoice == 0x14: // BACnetErrorPasswordFailure
		_child, typeSwitchError = BACnetErrorPasswordFailureParse(readBuffer)
	case serviceChoice == 0x15: // BACnetErrorVTOpen
		_child, typeSwitchError = BACnetErrorVTOpenParse(readBuffer)
	case serviceChoice == 0x17: // BACnetErrorVTData
		_child, typeSwitchError = BACnetErrorVTDataParse(readBuffer)
	case serviceChoice == 0x18: // BACnetErrorRemovedAuthenticate
		_child, typeSwitchError = BACnetErrorRemovedAuthenticateParse(readBuffer)
	case serviceChoice == 0x0D: // BACnetErrorRemovedReadPropertyConditional
		_child, typeSwitchError = BACnetErrorRemovedReadPropertyConditionalParse(readBuffer)
	case true: // BACnetErrorUnknown
		_child, typeSwitchError = BACnetErrorUnknownParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (errorClass)
	if pullErr := readBuffer.PullContext("errorClass"); pullErr != nil {
		return nil, pullErr
	}
	_errorClass, _errorClassErr := BACnetApplicationTagParse(readBuffer)
	if _errorClassErr != nil {
		return nil, errors.Wrap(_errorClassErr, "Error parsing 'errorClass' field")
	}
	errorClass := CastBACnetApplicationTagEnumerated(_errorClass)
	if closeErr := readBuffer.CloseContext("errorClass"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (errorCode)
	if pullErr := readBuffer.PullContext("errorCode"); pullErr != nil {
		return nil, pullErr
	}
	_errorCode, _errorCodeErr := BACnetApplicationTagParse(readBuffer)
	if _errorCodeErr != nil {
		return nil, errors.Wrap(_errorCodeErr, "Error parsing 'errorCode' field")
	}
	errorCode := CastBACnetApplicationTagEnumerated(_errorCode)
	if closeErr := readBuffer.CloseContext("errorCode"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetError"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), errorClass, errorCode)
	return _child.GetParent(), nil
}

func (m *BACnetError) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetError) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetError, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetError"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(child.GetServiceChoice())
	_serviceChoiceErr := writeBuffer.WriteUint8("serviceChoice", 8, (serviceChoice))

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (errorClass)
	if pushErr := writeBuffer.PushContext("errorClass"); pushErr != nil {
		return pushErr
	}
	_errorClassErr := m.ErrorClass.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("errorClass"); popErr != nil {
		return popErr
	}
	if _errorClassErr != nil {
		return errors.Wrap(_errorClassErr, "Error serializing 'errorClass' field")
	}

	// Simple Field (errorCode)
	if pushErr := writeBuffer.PushContext("errorCode"); pushErr != nil {
		return pushErr
	}
	_errorCodeErr := m.ErrorCode.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("errorCode"); popErr != nil {
		return popErr
	}
	if _errorCodeErr != nil {
		return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
	}

	if popErr := writeBuffer.PopContext("BACnetError"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetError) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
