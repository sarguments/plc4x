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

// SearchRequest is the data-structure of this message
type SearchRequest struct {
	*KnxNetIpMessage
	HpaiIDiscoveryEndpoint *HPAIDiscoveryEndpoint
}

// ISearchRequest is the corresponding interface of SearchRequest
type ISearchRequest interface {
	IKnxNetIpMessage
	// GetHpaiIDiscoveryEndpoint returns HpaiIDiscoveryEndpoint (property field)
	GetHpaiIDiscoveryEndpoint() *HPAIDiscoveryEndpoint
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

func (m *SearchRequest) GetMsgType() uint16 {
	return 0x0201
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SearchRequest) InitializeParent(parent *KnxNetIpMessage) {}

func (m *SearchRequest) GetParent() *KnxNetIpMessage {
	return m.KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *SearchRequest) GetHpaiIDiscoveryEndpoint() *HPAIDiscoveryEndpoint {
	return m.HpaiIDiscoveryEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSearchRequest factory function for SearchRequest
func NewSearchRequest(hpaiIDiscoveryEndpoint *HPAIDiscoveryEndpoint) *SearchRequest {
	_result := &SearchRequest{
		HpaiIDiscoveryEndpoint: hpaiIDiscoveryEndpoint,
		KnxNetIpMessage:        NewKnxNetIpMessage(),
	}
	_result.Child = _result
	return _result
}

func CastSearchRequest(structType interface{}) *SearchRequest {
	if casted, ok := structType.(SearchRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*SearchRequest); ok {
		return casted
	}
	if casted, ok := structType.(KnxNetIpMessage); ok {
		return CastSearchRequest(casted.Child)
	}
	if casted, ok := structType.(*KnxNetIpMessage); ok {
		return CastSearchRequest(casted.Child)
	}
	return nil
}

func (m *SearchRequest) GetTypeName() string {
	return "SearchRequest"
}

func (m *SearchRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SearchRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (hpaiIDiscoveryEndpoint)
	lengthInBits += m.HpaiIDiscoveryEndpoint.GetLengthInBits()

	return lengthInBits
}

func (m *SearchRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SearchRequestParse(readBuffer utils.ReadBuffer) (*SearchRequest, error) {
	if pullErr := readBuffer.PullContext("SearchRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (hpaiIDiscoveryEndpoint)
	if pullErr := readBuffer.PullContext("hpaiIDiscoveryEndpoint"); pullErr != nil {
		return nil, pullErr
	}
	_hpaiIDiscoveryEndpoint, _hpaiIDiscoveryEndpointErr := HPAIDiscoveryEndpointParse(readBuffer)
	if _hpaiIDiscoveryEndpointErr != nil {
		return nil, errors.Wrap(_hpaiIDiscoveryEndpointErr, "Error parsing 'hpaiIDiscoveryEndpoint' field")
	}
	hpaiIDiscoveryEndpoint := CastHPAIDiscoveryEndpoint(_hpaiIDiscoveryEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiIDiscoveryEndpoint"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("SearchRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SearchRequest{
		HpaiIDiscoveryEndpoint: CastHPAIDiscoveryEndpoint(hpaiIDiscoveryEndpoint),
		KnxNetIpMessage:        &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child, nil
}

func (m *SearchRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SearchRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (hpaiIDiscoveryEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiIDiscoveryEndpoint"); pushErr != nil {
			return pushErr
		}
		_hpaiIDiscoveryEndpointErr := m.HpaiIDiscoveryEndpoint.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("hpaiIDiscoveryEndpoint"); popErr != nil {
			return popErr
		}
		if _hpaiIDiscoveryEndpointErr != nil {
			return errors.Wrap(_hpaiIDiscoveryEndpointErr, "Error serializing 'hpaiIDiscoveryEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("SearchRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SearchRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
