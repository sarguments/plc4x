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

// NLMInitalizeRoutingTableAck is the data-structure of this message
type NLMInitalizeRoutingTableAck struct {
	*NLM
	NumberOfPorts uint8
	PortMappings  []*NLMInitalizeRoutingTablePortMapping

	// Arguments.
	ApduLength uint16
}

// INLMInitalizeRoutingTableAck is the corresponding interface of NLMInitalizeRoutingTableAck
type INLMInitalizeRoutingTableAck interface {
	INLM
	// GetNumberOfPorts returns NumberOfPorts (property field)
	GetNumberOfPorts() uint8
	// GetPortMappings returns PortMappings (property field)
	GetPortMappings() []*NLMInitalizeRoutingTablePortMapping
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

func (m *NLMInitalizeRoutingTableAck) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *NLMInitalizeRoutingTableAck) InitializeParent(parent *NLM, vendorId *uint16) {
	m.NLM.VendorId = vendorId
}

func (m *NLMInitalizeRoutingTableAck) GetParent() *NLM {
	return m.NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *NLMInitalizeRoutingTableAck) GetNumberOfPorts() uint8 {
	return m.NumberOfPorts
}

func (m *NLMInitalizeRoutingTableAck) GetPortMappings() []*NLMInitalizeRoutingTablePortMapping {
	return m.PortMappings
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMInitalizeRoutingTableAck factory function for NLMInitalizeRoutingTableAck
func NewNLMInitalizeRoutingTableAck(numberOfPorts uint8, portMappings []*NLMInitalizeRoutingTablePortMapping, vendorId *uint16, apduLength uint16) *NLMInitalizeRoutingTableAck {
	_result := &NLMInitalizeRoutingTableAck{
		NumberOfPorts: numberOfPorts,
		PortMappings:  portMappings,
		NLM:           NewNLM(vendorId, apduLength),
	}
	_result.Child = _result
	return _result
}

func CastNLMInitalizeRoutingTableAck(structType interface{}) *NLMInitalizeRoutingTableAck {
	if casted, ok := structType.(NLMInitalizeRoutingTableAck); ok {
		return &casted
	}
	if casted, ok := structType.(*NLMInitalizeRoutingTableAck); ok {
		return casted
	}
	if casted, ok := structType.(NLM); ok {
		return CastNLMInitalizeRoutingTableAck(casted.Child)
	}
	if casted, ok := structType.(*NLM); ok {
		return CastNLMInitalizeRoutingTableAck(casted.Child)
	}
	return nil
}

func (m *NLMInitalizeRoutingTableAck) GetTypeName() string {
	return "NLMInitalizeRoutingTableAck"
}

func (m *NLMInitalizeRoutingTableAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NLMInitalizeRoutingTableAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numberOfPorts)
	lengthInBits += 8

	// Array field
	if len(m.PortMappings) > 0 {
		for i, element := range m.PortMappings {
			last := i == len(m.PortMappings)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *NLMInitalizeRoutingTableAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMInitalizeRoutingTableAckParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (*NLMInitalizeRoutingTableAck, error) {
	if pullErr := readBuffer.PullContext("NLMInitalizeRoutingTableAck"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (numberOfPorts)
	_numberOfPorts, _numberOfPortsErr := readBuffer.ReadUint8("numberOfPorts", 8)
	if _numberOfPortsErr != nil {
		return nil, errors.Wrap(_numberOfPortsErr, "Error parsing 'numberOfPorts' field")
	}
	numberOfPorts := _numberOfPorts

	// Array field (portMappings)
	if pullErr := readBuffer.PullContext("portMappings", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	portMappings := make([]*NLMInitalizeRoutingTablePortMapping, numberOfPorts)
	{
		for curItem := uint16(0); curItem < uint16(numberOfPorts); curItem++ {
			_item, _err := NLMInitalizeRoutingTablePortMappingParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'portMappings' field")
			}
			portMappings[curItem] = CastNLMInitalizeRoutingTablePortMapping(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("portMappings", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("NLMInitalizeRoutingTableAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NLMInitalizeRoutingTableAck{
		NumberOfPorts: numberOfPorts,
		PortMappings:  portMappings,
		NLM:           &NLM{},
	}
	_child.NLM.Child = _child
	return _child, nil
}

func (m *NLMInitalizeRoutingTableAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMInitalizeRoutingTableAck"); pushErr != nil {
			return pushErr
		}

		// Simple Field (numberOfPorts)
		numberOfPorts := uint8(m.NumberOfPorts)
		_numberOfPortsErr := writeBuffer.WriteUint8("numberOfPorts", 8, (numberOfPorts))
		if _numberOfPortsErr != nil {
			return errors.Wrap(_numberOfPortsErr, "Error serializing 'numberOfPorts' field")
		}

		// Array Field (portMappings)
		if m.PortMappings != nil {
			if pushErr := writeBuffer.PushContext("portMappings", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.PortMappings {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'portMappings' field")
				}
			}
			if popErr := writeBuffer.PopContext("portMappings", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("NLMInitalizeRoutingTableAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NLMInitalizeRoutingTableAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
