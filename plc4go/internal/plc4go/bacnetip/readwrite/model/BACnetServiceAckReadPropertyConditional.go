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

// BACnetServiceAckReadPropertyConditional is the data-structure of this message
type BACnetServiceAckReadPropertyConditional struct {
	*BACnetServiceAck

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetServiceAckReadPropertyConditional is the corresponding interface of BACnetServiceAckReadPropertyConditional
type IBACnetServiceAckReadPropertyConditional interface {
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

func (m *BACnetServiceAckReadPropertyConditional) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckReadPropertyConditional) InitializeParent(parent *BACnetServiceAck) {}

func (m *BACnetServiceAckReadPropertyConditional) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

// NewBACnetServiceAckReadPropertyConditional factory function for BACnetServiceAckReadPropertyConditional
func NewBACnetServiceAckReadPropertyConditional(serviceRequestLength uint16) *BACnetServiceAckReadPropertyConditional {
	_result := &BACnetServiceAckReadPropertyConditional{
		BACnetServiceAck: NewBACnetServiceAck(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckReadPropertyConditional(structType interface{}) *BACnetServiceAckReadPropertyConditional {
	if casted, ok := structType.(BACnetServiceAckReadPropertyConditional); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadPropertyConditional); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckReadPropertyConditional(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckReadPropertyConditional(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckReadPropertyConditional) GetTypeName() string {
	return "BACnetServiceAckReadPropertyConditional"
}

func (m *BACnetServiceAckReadPropertyConditional) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckReadPropertyConditional) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckReadPropertyConditional) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckReadPropertyConditionalParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetServiceAckReadPropertyConditional, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadPropertyConditional"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseValidationError{"TODO: implement me"}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadPropertyConditional"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckReadPropertyConditional{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckReadPropertyConditional) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadPropertyConditional"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadPropertyConditional"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckReadPropertyConditional) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
