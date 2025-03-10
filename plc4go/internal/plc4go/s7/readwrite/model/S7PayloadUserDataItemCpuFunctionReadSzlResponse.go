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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH uint16 = uint16(28)

// S7PayloadUserDataItemCpuFunctionReadSzlResponse is the data-structure of this message
type S7PayloadUserDataItemCpuFunctionReadSzlResponse struct {
	*S7PayloadUserDataItem
	SzlId    *SzlId
	SzlIndex uint16
	Items    []*SzlDataTreeItem
}

// IS7PayloadUserDataItemCpuFunctionReadSzlResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionReadSzlResponse
type IS7PayloadUserDataItemCpuFunctionReadSzlResponse interface {
	IS7PayloadUserDataItem
	// GetSzlId returns SzlId (property field)
	GetSzlId() *SzlId
	// GetSzlIndex returns SzlIndex (property field)
	GetSzlIndex() uint16
	// GetItems returns Items (property field)
	GetItems() []*SzlDataTreeItem
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

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuSubfunction() uint8 {
	return 0x01
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetParent() *S7PayloadUserDataItem {
	return m.S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetSzlId() *SzlId {
	return m.SzlId
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetSzlIndex() uint16 {
	return m.SzlIndex
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetItems() []*SzlDataTreeItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetSzlItemLength() uint16 {
	return S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionReadSzlResponse factory function for S7PayloadUserDataItemCpuFunctionReadSzlResponse
func NewS7PayloadUserDataItemCpuFunctionReadSzlResponse(szlId *SzlId, szlIndex uint16, items []*SzlDataTreeItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	_result := &S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		SzlId:                 szlId,
		SzlIndex:              szlIndex,
		Items:                 items,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result.Child = _result
	return _result
}

func CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(structType interface{}) *S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
		return casted
	}
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(casted.Child)
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(casted.Child)
	}
	return nil
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionReadSzlResponse"
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (szlId)
	lengthInBits += m.SzlId.GetLengthInBits()

	// Simple field (szlIndex)
	lengthInBits += 16

	// Const Field (szlItemLength)
	lengthInBits += 16

	// Implicit Field (szlItemCount)
	lengthInBits += 16

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionReadSzlResponseParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItemCpuFunctionReadSzlResponse, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (szlId)
	if pullErr := readBuffer.PullContext("szlId"); pullErr != nil {
		return nil, pullErr
	}
	_szlId, _szlIdErr := SzlIdParse(readBuffer)
	if _szlIdErr != nil {
		return nil, errors.Wrap(_szlIdErr, "Error parsing 'szlId' field")
	}
	szlId := CastSzlId(_szlId)
	if closeErr := readBuffer.CloseContext("szlId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (szlIndex)
	_szlIndex, _szlIndexErr := readBuffer.ReadUint16("szlIndex", 16)
	if _szlIndexErr != nil {
		return nil, errors.Wrap(_szlIndexErr, "Error parsing 'szlIndex' field")
	}
	szlIndex := _szlIndex

	// Const Field (szlItemLength)
	szlItemLength, _szlItemLengthErr := readBuffer.ReadUint16("szlItemLength", 16)
	if _szlItemLengthErr != nil {
		return nil, errors.Wrap(_szlItemLengthErr, "Error parsing 'szlItemLength' field")
	}
	if szlItemLength != S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH) + " but got " + fmt.Sprintf("%d", szlItemLength))
	}

	// Implicit Field (szlItemCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	szlItemCount, _szlItemCountErr := readBuffer.ReadUint16("szlItemCount", 16)
	_ = szlItemCount
	if _szlItemCountErr != nil {
		return nil, errors.Wrap(_szlItemCountErr, "Error parsing 'szlItemCount' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*SzlDataTreeItem, szlItemCount)
	{
		for curItem := uint16(0); curItem < uint16(szlItemCount); curItem++ {
			_item, _err := SzlDataTreeItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items[curItem] = CastSzlDataTreeItem(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		SzlId:                 CastSzlId(szlId),
		SzlIndex:              szlIndex,
		Items:                 items,
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child, nil
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (szlId)
		if pushErr := writeBuffer.PushContext("szlId"); pushErr != nil {
			return pushErr
		}
		_szlIdErr := m.SzlId.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("szlId"); popErr != nil {
			return popErr
		}
		if _szlIdErr != nil {
			return errors.Wrap(_szlIdErr, "Error serializing 'szlId' field")
		}

		// Simple Field (szlIndex)
		szlIndex := uint16(m.SzlIndex)
		_szlIndexErr := writeBuffer.WriteUint16("szlIndex", 16, (szlIndex))
		if _szlIndexErr != nil {
			return errors.Wrap(_szlIndexErr, "Error serializing 'szlIndex' field")
		}

		// Const Field (szlItemLength)
		_szlItemLengthErr := writeBuffer.WriteUint16("szlItemLength", 16, 28)
		if _szlItemLengthErr != nil {
			return errors.Wrap(_szlItemLengthErr, "Error serializing 'szlItemLength' field")
		}

		// Implicit Field (szlItemCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		szlItemCount := uint16(uint16(len(m.GetItems())))
		_szlItemCountErr := writeBuffer.WriteUint16("szlItemCount", 16, (szlItemCount))
		if _szlItemCountErr != nil {
			return errors.Wrap(_szlItemCountErr, "Error serializing 'szlItemCount' field")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
