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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckConfirmedTextMessage is the data-structure of this message
type BACnetServiceAckConfirmedTextMessage struct {
	*BACnetServiceAck

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetServiceAckConfirmedTextMessage is the corresponding interface of BACnetServiceAckConfirmedTextMessage
type IBACnetServiceAckConfirmedTextMessage interface {
	IBACnetServiceAck
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

func (m *BACnetServiceAckConfirmedTextMessage) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckConfirmedTextMessage) InitializeParent(parent *BACnetServiceAck) {}

func (m *BACnetServiceAckConfirmedTextMessage) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

// NewBACnetServiceAckConfirmedTextMessage factory function for BACnetServiceAckConfirmedTextMessage
func NewBACnetServiceAckConfirmedTextMessage(serviceRequestLength uint16) *BACnetServiceAckConfirmedTextMessage {
	_result := &BACnetServiceAckConfirmedTextMessage{
		BACnetServiceAck: NewBACnetServiceAck(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckConfirmedTextMessage(structType interface{}) *BACnetServiceAckConfirmedTextMessage {
	if casted, ok := structType.(BACnetServiceAckConfirmedTextMessage); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckConfirmedTextMessage); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckConfirmedTextMessage(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckConfirmedTextMessage(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckConfirmedTextMessage) GetTypeName() string {
	return "BACnetServiceAckConfirmedTextMessage"
}

func (m *BACnetServiceAckConfirmedTextMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckConfirmedTextMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckConfirmedTextMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckConfirmedTextMessageParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetServiceAckConfirmedTextMessage, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckConfirmedTextMessage"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseValidationError{"TODO: implement me"}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckConfirmedTextMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckConfirmedTextMessage{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckConfirmedTextMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckConfirmedTextMessage"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckConfirmedTextMessage"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckConfirmedTextMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
