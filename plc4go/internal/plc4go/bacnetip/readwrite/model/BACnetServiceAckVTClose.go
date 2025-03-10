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

// BACnetServiceAckVTClose is the data-structure of this message
type BACnetServiceAckVTClose struct {
	*BACnetServiceAck

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetServiceAckVTClose is the corresponding interface of BACnetServiceAckVTClose
type IBACnetServiceAckVTClose interface {
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

func (m *BACnetServiceAckVTClose) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_CLOSE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckVTClose) InitializeParent(parent *BACnetServiceAck) {}

func (m *BACnetServiceAckVTClose) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

// NewBACnetServiceAckVTClose factory function for BACnetServiceAckVTClose
func NewBACnetServiceAckVTClose(serviceRequestLength uint16) *BACnetServiceAckVTClose {
	_result := &BACnetServiceAckVTClose{
		BACnetServiceAck: NewBACnetServiceAck(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckVTClose(structType interface{}) *BACnetServiceAckVTClose {
	if casted, ok := structType.(BACnetServiceAckVTClose); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckVTClose); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckVTClose(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckVTClose(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckVTClose) GetTypeName() string {
	return "BACnetServiceAckVTClose"
}

func (m *BACnetServiceAckVTClose) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckVTClose) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckVTClose) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckVTCloseParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetServiceAckVTClose, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckVTClose"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseValidationError{"TODO: implement me"}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckVTClose"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckVTClose{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckVTClose) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckVTClose"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckVTClose"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckVTClose) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
