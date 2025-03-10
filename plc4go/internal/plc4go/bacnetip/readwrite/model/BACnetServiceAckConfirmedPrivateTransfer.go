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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckConfirmedPrivateTransfer is the data-structure of this message
type BACnetServiceAckConfirmedPrivateTransfer struct {
	*BACnetServiceAck
	VendorId      *BACnetContextTagUnsignedInteger
	ServiceNumber *BACnetContextTagUnsignedInteger
	ResultBlock   *BACnetConstructedData

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetServiceAckConfirmedPrivateTransfer is the corresponding interface of BACnetServiceAckConfirmedPrivateTransfer
type IBACnetServiceAckConfirmedPrivateTransfer interface {
	IBACnetServiceAck
	// GetVendorId returns VendorId (property field)
	GetVendorId() *BACnetContextTagUnsignedInteger
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() *BACnetContextTagUnsignedInteger
	// GetResultBlock returns ResultBlock (property field)
	GetResultBlock() *BACnetConstructedData
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

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckConfirmedPrivateTransfer) InitializeParent(parent *BACnetServiceAck) {}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetVendorId() *BACnetContextTagUnsignedInteger {
	return m.VendorId
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetServiceNumber() *BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetResultBlock() *BACnetConstructedData {
	return m.ResultBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckConfirmedPrivateTransfer factory function for BACnetServiceAckConfirmedPrivateTransfer
func NewBACnetServiceAckConfirmedPrivateTransfer(vendorId *BACnetContextTagUnsignedInteger, serviceNumber *BACnetContextTagUnsignedInteger, resultBlock *BACnetConstructedData, serviceRequestLength uint16) *BACnetServiceAckConfirmedPrivateTransfer {
	_result := &BACnetServiceAckConfirmedPrivateTransfer{
		VendorId:         vendorId,
		ServiceNumber:    serviceNumber,
		ResultBlock:      resultBlock,
		BACnetServiceAck: NewBACnetServiceAck(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckConfirmedPrivateTransfer(structType interface{}) *BACnetServiceAckConfirmedPrivateTransfer {
	if casted, ok := structType.(BACnetServiceAckConfirmedPrivateTransfer); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckConfirmedPrivateTransfer); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckConfirmedPrivateTransfer(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckConfirmedPrivateTransfer(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetServiceAckConfirmedPrivateTransfer"
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits()

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits()

	// Optional Field (resultBlock)
	if m.ResultBlock != nil {
		lengthInBits += (*m.ResultBlock).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckConfirmedPrivateTransferParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetServiceAckConfirmedPrivateTransfer, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckConfirmedPrivateTransfer"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, pullErr
	}
	_vendorId, _vendorIdErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}
	vendorId := CastBACnetContextTagUnsignedInteger(_vendorId)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (serviceNumber)
	if pullErr := readBuffer.PullContext("serviceNumber"); pullErr != nil {
		return nil, pullErr
	}
	_serviceNumber, _serviceNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _serviceNumberErr != nil {
		return nil, errors.Wrap(_serviceNumberErr, "Error parsing 'serviceNumber' field")
	}
	serviceNumber := CastBACnetContextTagUnsignedInteger(_serviceNumber)
	if closeErr := readBuffer.CloseContext("serviceNumber"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (resultBlock) (Can be skipped, if a given expression evaluates to false)
	var resultBlock *BACnetConstructedData = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("resultBlock"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(2), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, DummyPropertyIdentifier())
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'resultBlock' field")
		default:
			resultBlock = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("resultBlock"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckConfirmedPrivateTransfer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckConfirmedPrivateTransfer{
		VendorId:         CastBACnetContextTagUnsignedInteger(vendorId),
		ServiceNumber:    CastBACnetContextTagUnsignedInteger(serviceNumber),
		ResultBlock:      CastBACnetConstructedData(resultBlock),
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckConfirmedPrivateTransfer"); pushErr != nil {
			return pushErr
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return pushErr
		}
		_vendorIdErr := m.VendorId.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return popErr
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (serviceNumber)
		if pushErr := writeBuffer.PushContext("serviceNumber"); pushErr != nil {
			return pushErr
		}
		_serviceNumberErr := m.ServiceNumber.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("serviceNumber"); popErr != nil {
			return popErr
		}
		if _serviceNumberErr != nil {
			return errors.Wrap(_serviceNumberErr, "Error serializing 'serviceNumber' field")
		}

		// Optional Field (resultBlock) (Can be skipped, if the value is null)
		var resultBlock *BACnetConstructedData = nil
		if m.ResultBlock != nil {
			if pushErr := writeBuffer.PushContext("resultBlock"); pushErr != nil {
				return pushErr
			}
			resultBlock = m.ResultBlock
			_resultBlockErr := resultBlock.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("resultBlock"); popErr != nil {
				return popErr
			}
			if _resultBlockErr != nil {
				return errors.Wrap(_resultBlockErr, "Error serializing 'resultBlock' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckConfirmedPrivateTransfer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
