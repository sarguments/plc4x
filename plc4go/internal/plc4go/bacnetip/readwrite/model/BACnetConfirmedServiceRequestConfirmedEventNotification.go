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

// BACnetConfirmedServiceRequestConfirmedEventNotification is the data-structure of this message
type BACnetConfirmedServiceRequestConfirmedEventNotification struct {
	*BACnetConfirmedServiceRequest
	ProcessIdentifier          *BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier *BACnetContextTagObjectIdentifier
	EventObjectIdentifier      *BACnetContextTagObjectIdentifier
	Timestamp                  *BACnetTimeStampEnclosed
	NotificationClass          *BACnetContextTagUnsignedInteger
	Priority                   *BACnetContextTagUnsignedInteger
	EventType                  *BACnetContextTagEventType
	MessageText                *BACnetContextTagCharacterString
	NotifyType                 *BACnetContextTagNotifyType
	AckRequired                *BACnetContextTagBoolean
	FromState                  *BACnetContextTagEventState
	ToState                    *BACnetContextTagEventState
	EventValues                *BACnetNotificationParameters

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestConfirmedEventNotification is the corresponding interface of BACnetConfirmedServiceRequestConfirmedEventNotification
type IBACnetConfirmedServiceRequestConfirmedEventNotification interface {
	IBACnetConfirmedServiceRequest
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() *BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() *BACnetContextTagObjectIdentifier
	// GetEventObjectIdentifier returns EventObjectIdentifier (property field)
	GetEventObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() *BACnetTimeStampEnclosed
	// GetNotificationClass returns NotificationClass (property field)
	GetNotificationClass() *BACnetContextTagUnsignedInteger
	// GetPriority returns Priority (property field)
	GetPriority() *BACnetContextTagUnsignedInteger
	// GetEventType returns EventType (property field)
	GetEventType() *BACnetContextTagEventType
	// GetMessageText returns MessageText (property field)
	GetMessageText() *BACnetContextTagCharacterString
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() *BACnetContextTagNotifyType
	// GetAckRequired returns AckRequired (property field)
	GetAckRequired() *BACnetContextTagBoolean
	// GetFromState returns FromState (property field)
	GetFromState() *BACnetContextTagEventState
	// GetToState returns ToState (property field)
	GetToState() *BACnetContextTagEventState
	// GetEventValues returns EventValues (property field)
	GetEventValues() *BACnetNotificationParameters
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

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetProcessIdentifier() *BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetInitiatingDeviceIdentifier() *BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetEventObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.EventObjectIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetTimestamp() *BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetNotificationClass() *BACnetContextTagUnsignedInteger {
	return m.NotificationClass
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetPriority() *BACnetContextTagUnsignedInteger {
	return m.Priority
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetEventType() *BACnetContextTagEventType {
	return m.EventType
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetMessageText() *BACnetContextTagCharacterString {
	return m.MessageText
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetNotifyType() *BACnetContextTagNotifyType {
	return m.NotifyType
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetAckRequired() *BACnetContextTagBoolean {
	return m.AckRequired
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetFromState() *BACnetContextTagEventState {
	return m.FromState
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetToState() *BACnetContextTagEventState {
	return m.ToState
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetEventValues() *BACnetNotificationParameters {
	return m.EventValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedEventNotification factory function for BACnetConfirmedServiceRequestConfirmedEventNotification
func NewBACnetConfirmedServiceRequestConfirmedEventNotification(processIdentifier *BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier *BACnetContextTagObjectIdentifier, eventObjectIdentifier *BACnetContextTagObjectIdentifier, timestamp *BACnetTimeStampEnclosed, notificationClass *BACnetContextTagUnsignedInteger, priority *BACnetContextTagUnsignedInteger, eventType *BACnetContextTagEventType, messageText *BACnetContextTagCharacterString, notifyType *BACnetContextTagNotifyType, ackRequired *BACnetContextTagBoolean, fromState *BACnetContextTagEventState, toState *BACnetContextTagEventState, eventValues *BACnetNotificationParameters, serviceRequestLength uint16) *BACnetConfirmedServiceRequestConfirmedEventNotification {
	_result := &BACnetConfirmedServiceRequestConfirmedEventNotification{
		ProcessIdentifier:             processIdentifier,
		InitiatingDeviceIdentifier:    initiatingDeviceIdentifier,
		EventObjectIdentifier:         eventObjectIdentifier,
		Timestamp:                     timestamp,
		NotificationClass:             notificationClass,
		Priority:                      priority,
		EventType:                     eventType,
		MessageText:                   messageText,
		NotifyType:                    notifyType,
		AckRequired:                   ackRequired,
		FromState:                     fromState,
		ToState:                       toState,
		EventValues:                   eventValues,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestConfirmedEventNotification(structType interface{}) *BACnetConfirmedServiceRequestConfirmedEventNotification {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedEventNotification); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedEventNotification); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestConfirmedEventNotification(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestConfirmedEventNotification(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedEventNotification"
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits()

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits()

	// Simple field (eventObjectIdentifier)
	lengthInBits += m.EventObjectIdentifier.GetLengthInBits()

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits()

	// Simple field (notificationClass)
	lengthInBits += m.NotificationClass.GetLengthInBits()

	// Simple field (priority)
	lengthInBits += m.Priority.GetLengthInBits()

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits()

	// Optional Field (messageText)
	if m.MessageText != nil {
		lengthInBits += (*m.MessageText).GetLengthInBits()
	}

	// Simple field (notifyType)
	lengthInBits += m.NotifyType.GetLengthInBits()

	// Optional Field (ackRequired)
	if m.AckRequired != nil {
		lengthInBits += (*m.AckRequired).GetLengthInBits()
	}

	// Optional Field (fromState)
	if m.FromState != nil {
		lengthInBits += (*m.FromState).GetLengthInBits()
	}

	// Simple field (toState)
	lengthInBits += m.ToState.GetLengthInBits()

	// Optional Field (eventValues)
	if m.EventValues != nil {
		lengthInBits += (*m.EventValues).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedEventNotificationParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestConfirmedEventNotification, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedEventNotification"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (processIdentifier)
	if pullErr := readBuffer.PullContext("processIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_processIdentifier, _processIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _processIdentifierErr != nil {
		return nil, errors.Wrap(_processIdentifierErr, "Error parsing 'processIdentifier' field")
	}
	processIdentifier := CastBACnetContextTagUnsignedInteger(_processIdentifier)
	if closeErr := readBuffer.CloseContext("processIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pullErr := readBuffer.PullContext("initiatingDeviceIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_initiatingDeviceIdentifier, _initiatingDeviceIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _initiatingDeviceIdentifierErr != nil {
		return nil, errors.Wrap(_initiatingDeviceIdentifierErr, "Error parsing 'initiatingDeviceIdentifier' field")
	}
	initiatingDeviceIdentifier := CastBACnetContextTagObjectIdentifier(_initiatingDeviceIdentifier)
	if closeErr := readBuffer.CloseContext("initiatingDeviceIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (eventObjectIdentifier)
	if pullErr := readBuffer.PullContext("eventObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_eventObjectIdentifier, _eventObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _eventObjectIdentifierErr != nil {
		return nil, errors.Wrap(_eventObjectIdentifierErr, "Error parsing 'eventObjectIdentifier' field")
	}
	eventObjectIdentifier := CastBACnetContextTagObjectIdentifier(_eventObjectIdentifier)
	if closeErr := readBuffer.CloseContext("eventObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, pullErr
	}
	_timestamp, _timestampErr := BACnetTimeStampEnclosedParse(readBuffer, uint8(uint8(3)))
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}
	timestamp := CastBACnetTimeStampEnclosed(_timestamp)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (notificationClass)
	if pullErr := readBuffer.PullContext("notificationClass"); pullErr != nil {
		return nil, pullErr
	}
	_notificationClass, _notificationClassErr := BACnetContextTagParse(readBuffer, uint8(uint8(4)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _notificationClassErr != nil {
		return nil, errors.Wrap(_notificationClassErr, "Error parsing 'notificationClass' field")
	}
	notificationClass := CastBACnetContextTagUnsignedInteger(_notificationClass)
	if closeErr := readBuffer.CloseContext("notificationClass"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, pullErr
	}
	_priority, _priorityErr := BACnetContextTagParse(readBuffer, uint8(uint8(5)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field")
	}
	priority := CastBACnetContextTagUnsignedInteger(_priority)
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (eventType)
	if pullErr := readBuffer.PullContext("eventType"); pullErr != nil {
		return nil, pullErr
	}
	_eventType, _eventTypeErr := BACnetContextTagParse(readBuffer, uint8(uint8(6)), BACnetDataType(BACnetDataType_EVENT_TYPE))
	if _eventTypeErr != nil {
		return nil, errors.Wrap(_eventTypeErr, "Error parsing 'eventType' field")
	}
	eventType := CastBACnetContextTagEventType(_eventType)
	if closeErr := readBuffer.CloseContext("eventType"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (messageText) (Can be skipped, if a given expression evaluates to false)
	var messageText *BACnetContextTagCharacterString = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("messageText"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(7), BACnetDataType_CHARACTER_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'messageText' field")
		default:
			messageText = CastBACnetContextTagCharacterString(_val)
			if closeErr := readBuffer.CloseContext("messageText"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (notifyType)
	if pullErr := readBuffer.PullContext("notifyType"); pullErr != nil {
		return nil, pullErr
	}
	_notifyType, _notifyTypeErr := BACnetContextTagParse(readBuffer, uint8(uint8(8)), BACnetDataType(BACnetDataType_NOTIFY_TYPE))
	if _notifyTypeErr != nil {
		return nil, errors.Wrap(_notifyTypeErr, "Error parsing 'notifyType' field")
	}
	notifyType := CastBACnetContextTagNotifyType(_notifyType)
	if closeErr := readBuffer.CloseContext("notifyType"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (ackRequired) (Can be skipped, if a given expression evaluates to false)
	var ackRequired *BACnetContextTagBoolean = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("ackRequired"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(9), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'ackRequired' field")
		default:
			ackRequired = CastBACnetContextTagBoolean(_val)
			if closeErr := readBuffer.CloseContext("ackRequired"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (fromState) (Can be skipped, if a given expression evaluates to false)
	var fromState *BACnetContextTagEventState = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("fromState"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(10), BACnetDataType_EVENT_STATE)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'fromState' field")
		default:
			fromState = CastBACnetContextTagEventState(_val)
			if closeErr := readBuffer.CloseContext("fromState"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (toState)
	if pullErr := readBuffer.PullContext("toState"); pullErr != nil {
		return nil, pullErr
	}
	_toState, _toStateErr := BACnetContextTagParse(readBuffer, uint8(uint8(11)), BACnetDataType(BACnetDataType_EVENT_STATE))
	if _toStateErr != nil {
		return nil, errors.Wrap(_toStateErr, "Error parsing 'toState' field")
	}
	toState := CastBACnetContextTagEventState(_toState)
	if closeErr := readBuffer.CloseContext("toState"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (eventValues) (Can be skipped, if a given expression evaluates to false)
	var eventValues *BACnetNotificationParameters = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("eventValues"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetNotificationParametersParse(readBuffer, uint8(12), eventObjectIdentifier.GetObjectType())
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'eventValues' field")
		default:
			eventValues = CastBACnetNotificationParameters(_val)
			if closeErr := readBuffer.CloseContext("eventValues"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedEventNotification"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestConfirmedEventNotification{
		ProcessIdentifier:             CastBACnetContextTagUnsignedInteger(processIdentifier),
		InitiatingDeviceIdentifier:    CastBACnetContextTagObjectIdentifier(initiatingDeviceIdentifier),
		EventObjectIdentifier:         CastBACnetContextTagObjectIdentifier(eventObjectIdentifier),
		Timestamp:                     CastBACnetTimeStampEnclosed(timestamp),
		NotificationClass:             CastBACnetContextTagUnsignedInteger(notificationClass),
		Priority:                      CastBACnetContextTagUnsignedInteger(priority),
		EventType:                     CastBACnetContextTagEventType(eventType),
		MessageText:                   CastBACnetContextTagCharacterString(messageText),
		NotifyType:                    CastBACnetContextTagNotifyType(notifyType),
		AckRequired:                   CastBACnetContextTagBoolean(ackRequired),
		FromState:                     CastBACnetContextTagEventState(fromState),
		ToState:                       CastBACnetContextTagEventState(toState),
		EventValues:                   CastBACnetNotificationParameters(eventValues),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedEventNotification"); pushErr != nil {
			return pushErr
		}

		// Simple Field (processIdentifier)
		if pushErr := writeBuffer.PushContext("processIdentifier"); pushErr != nil {
			return pushErr
		}
		_processIdentifierErr := m.ProcessIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
			return popErr
		}
		if _processIdentifierErr != nil {
			return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
		}

		// Simple Field (initiatingDeviceIdentifier)
		if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
			return pushErr
		}
		_initiatingDeviceIdentifierErr := m.InitiatingDeviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
			return popErr
		}
		if _initiatingDeviceIdentifierErr != nil {
			return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		// Simple Field (eventObjectIdentifier)
		if pushErr := writeBuffer.PushContext("eventObjectIdentifier"); pushErr != nil {
			return pushErr
		}
		_eventObjectIdentifierErr := m.EventObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("eventObjectIdentifier"); popErr != nil {
			return popErr
		}
		if _eventObjectIdentifierErr != nil {
			return errors.Wrap(_eventObjectIdentifierErr, "Error serializing 'eventObjectIdentifier' field")
		}

		// Simple Field (timestamp)
		if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
			return pushErr
		}
		_timestampErr := m.Timestamp.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
			return popErr
		}
		if _timestampErr != nil {
			return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
		}

		// Simple Field (notificationClass)
		if pushErr := writeBuffer.PushContext("notificationClass"); pushErr != nil {
			return pushErr
		}
		_notificationClassErr := m.NotificationClass.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("notificationClass"); popErr != nil {
			return popErr
		}
		if _notificationClassErr != nil {
			return errors.Wrap(_notificationClassErr, "Error serializing 'notificationClass' field")
		}

		// Simple Field (priority)
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return pushErr
		}
		_priorityErr := m.Priority.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return popErr
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}

		// Simple Field (eventType)
		if pushErr := writeBuffer.PushContext("eventType"); pushErr != nil {
			return pushErr
		}
		_eventTypeErr := m.EventType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("eventType"); popErr != nil {
			return popErr
		}
		if _eventTypeErr != nil {
			return errors.Wrap(_eventTypeErr, "Error serializing 'eventType' field")
		}

		// Optional Field (messageText) (Can be skipped, if the value is null)
		var messageText *BACnetContextTagCharacterString = nil
		if m.MessageText != nil {
			if pushErr := writeBuffer.PushContext("messageText"); pushErr != nil {
				return pushErr
			}
			messageText = m.MessageText
			_messageTextErr := messageText.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("messageText"); popErr != nil {
				return popErr
			}
			if _messageTextErr != nil {
				return errors.Wrap(_messageTextErr, "Error serializing 'messageText' field")
			}
		}

		// Simple Field (notifyType)
		if pushErr := writeBuffer.PushContext("notifyType"); pushErr != nil {
			return pushErr
		}
		_notifyTypeErr := m.NotifyType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("notifyType"); popErr != nil {
			return popErr
		}
		if _notifyTypeErr != nil {
			return errors.Wrap(_notifyTypeErr, "Error serializing 'notifyType' field")
		}

		// Optional Field (ackRequired) (Can be skipped, if the value is null)
		var ackRequired *BACnetContextTagBoolean = nil
		if m.AckRequired != nil {
			if pushErr := writeBuffer.PushContext("ackRequired"); pushErr != nil {
				return pushErr
			}
			ackRequired = m.AckRequired
			_ackRequiredErr := ackRequired.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("ackRequired"); popErr != nil {
				return popErr
			}
			if _ackRequiredErr != nil {
				return errors.Wrap(_ackRequiredErr, "Error serializing 'ackRequired' field")
			}
		}

		// Optional Field (fromState) (Can be skipped, if the value is null)
		var fromState *BACnetContextTagEventState = nil
		if m.FromState != nil {
			if pushErr := writeBuffer.PushContext("fromState"); pushErr != nil {
				return pushErr
			}
			fromState = m.FromState
			_fromStateErr := fromState.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("fromState"); popErr != nil {
				return popErr
			}
			if _fromStateErr != nil {
				return errors.Wrap(_fromStateErr, "Error serializing 'fromState' field")
			}
		}

		// Simple Field (toState)
		if pushErr := writeBuffer.PushContext("toState"); pushErr != nil {
			return pushErr
		}
		_toStateErr := m.ToState.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("toState"); popErr != nil {
			return popErr
		}
		if _toStateErr != nil {
			return errors.Wrap(_toStateErr, "Error serializing 'toState' field")
		}

		// Optional Field (eventValues) (Can be skipped, if the value is null)
		var eventValues *BACnetNotificationParameters = nil
		if m.EventValues != nil {
			if pushErr := writeBuffer.PushContext("eventValues"); pushErr != nil {
				return pushErr
			}
			eventValues = m.EventValues
			_eventValuesErr := eventValues.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("eventValues"); popErr != nil {
				return popErr
			}
			if _eventValuesErr != nil {
				return errors.Wrap(_eventValuesErr, "Error serializing 'eventValues' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedEventNotification"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestConfirmedEventNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
