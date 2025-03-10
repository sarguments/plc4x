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

// DeviceConfigurationAck is the data-structure of this message
type DeviceConfigurationAck struct {
	*KnxNetIpMessage
	DeviceConfigurationAckDataBlock *DeviceConfigurationAckDataBlock
}

// IDeviceConfigurationAck is the corresponding interface of DeviceConfigurationAck
type IDeviceConfigurationAck interface {
	IKnxNetIpMessage
	// GetDeviceConfigurationAckDataBlock returns DeviceConfigurationAckDataBlock (property field)
	GetDeviceConfigurationAckDataBlock() *DeviceConfigurationAckDataBlock
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

func (m *DeviceConfigurationAck) GetMsgType() uint16 {
	return 0x0311
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *DeviceConfigurationAck) InitializeParent(parent *KnxNetIpMessage) {}

func (m *DeviceConfigurationAck) GetParent() *KnxNetIpMessage {
	return m.KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *DeviceConfigurationAck) GetDeviceConfigurationAckDataBlock() *DeviceConfigurationAckDataBlock {
	return m.DeviceConfigurationAckDataBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceConfigurationAck factory function for DeviceConfigurationAck
func NewDeviceConfigurationAck(deviceConfigurationAckDataBlock *DeviceConfigurationAckDataBlock) *DeviceConfigurationAck {
	_result := &DeviceConfigurationAck{
		DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
		KnxNetIpMessage:                 NewKnxNetIpMessage(),
	}
	_result.Child = _result
	return _result
}

func CastDeviceConfigurationAck(structType interface{}) *DeviceConfigurationAck {
	if casted, ok := structType.(DeviceConfigurationAck); ok {
		return &casted
	}
	if casted, ok := structType.(*DeviceConfigurationAck); ok {
		return casted
	}
	if casted, ok := structType.(KnxNetIpMessage); ok {
		return CastDeviceConfigurationAck(casted.Child)
	}
	if casted, ok := structType.(*KnxNetIpMessage); ok {
		return CastDeviceConfigurationAck(casted.Child)
	}
	return nil
}

func (m *DeviceConfigurationAck) GetTypeName() string {
	return "DeviceConfigurationAck"
}

func (m *DeviceConfigurationAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DeviceConfigurationAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (deviceConfigurationAckDataBlock)
	lengthInBits += m.DeviceConfigurationAckDataBlock.GetLengthInBits()

	return lengthInBits
}

func (m *DeviceConfigurationAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DeviceConfigurationAckParse(readBuffer utils.ReadBuffer) (*DeviceConfigurationAck, error) {
	if pullErr := readBuffer.PullContext("DeviceConfigurationAck"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (deviceConfigurationAckDataBlock)
	if pullErr := readBuffer.PullContext("deviceConfigurationAckDataBlock"); pullErr != nil {
		return nil, pullErr
	}
	_deviceConfigurationAckDataBlock, _deviceConfigurationAckDataBlockErr := DeviceConfigurationAckDataBlockParse(readBuffer)
	if _deviceConfigurationAckDataBlockErr != nil {
		return nil, errors.Wrap(_deviceConfigurationAckDataBlockErr, "Error parsing 'deviceConfigurationAckDataBlock' field")
	}
	deviceConfigurationAckDataBlock := CastDeviceConfigurationAckDataBlock(_deviceConfigurationAckDataBlock)
	if closeErr := readBuffer.CloseContext("deviceConfigurationAckDataBlock"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DeviceConfigurationAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DeviceConfigurationAck{
		DeviceConfigurationAckDataBlock: CastDeviceConfigurationAckDataBlock(deviceConfigurationAckDataBlock),
		KnxNetIpMessage:                 &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child, nil
}

func (m *DeviceConfigurationAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationAck"); pushErr != nil {
			return pushErr
		}

		// Simple Field (deviceConfigurationAckDataBlock)
		if pushErr := writeBuffer.PushContext("deviceConfigurationAckDataBlock"); pushErr != nil {
			return pushErr
		}
		_deviceConfigurationAckDataBlockErr := m.DeviceConfigurationAckDataBlock.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceConfigurationAckDataBlock"); popErr != nil {
			return popErr
		}
		if _deviceConfigurationAckDataBlockErr != nil {
			return errors.Wrap(_deviceConfigurationAckDataBlockErr, "Error serializing 'deviceConfigurationAckDataBlock' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *DeviceConfigurationAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
