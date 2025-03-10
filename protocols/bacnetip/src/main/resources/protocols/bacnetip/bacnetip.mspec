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

[discriminatedType BVLC byteOrder='BIG_ENDIAN'
    [const         uint 8   bacnetType   0x81                   ]
    [discriminator uint 8   bvlcFunction                        ]
    [implicit      uint 16  bvlcLength          'lengthInBytes' ]
    [virtual       uint 16  bvlcPayloadLength   'bvlcLength-4'  ]
    [typeSwitch bvlcFunction
        ['0x00' BVLCResult
            [simple BVLCResultCode
                            code                                ]
        ]
        ['0x01' BVLCWriteBroadcastDistributionTable(uint 16 bvlcPayloadLength)
            [array BVLCBroadcastDistributionTableEntry
                            table
                            length 'bvlcPayloadLength'          ]
        ]
        ['0x02' BVLCReadBroadcastDistributionTable
        ]
        ['0x03' BVLCReadBroadcastDistributionTableAck(uint 16 bvlcPayloadLength)
            [array BVLCBroadcastDistributionTableEntry
                            table
                            length 'bvlcPayloadLength'          ]
        ]
        ['0x04' BVLCForwardedNPDU(uint 16 bvlcPayloadLength)
            [array  uint 8  ip    count '4'                     ]
            [simple uint 16 port                                ]
            [simple NPDU('bvlcPayloadLength - 6')
                            npdu                                ]
        ]
        ['0x05' BVLCRegisterForeignDevice
            [simple uint 16 ttl]
        ]
        ['0x06' BVLCReadForeignDeviceTable
        ]
        ['0x07' BVLCReadForeignDeviceTableAck(uint 16 bvlcPayloadLength)
            [array BVLCForeignDeviceTableEntry
                            table
                            length 'bvlcPayloadLength'          ]
        ]
        ['0x08' BVLCDeleteForeignDeviceTableEntry
            [array  uint 8      ip                          count '4'       ]
            [simple uint 16     port                                        ]
        ]
        ['0x09' BVLCDistributeBroadcastToNetwork(uint 16 bvlcPayloadLength)
            [simple NPDU('bvlcPayloadLength')
                            npdu                                ]
        ]
        ['0x0A' BVLCOriginalUnicastNPDU(uint 16 bvlcPayloadLength)
            [simple NPDU('bvlcPayloadLength')
                            npdu                                ]
        ]
        ['0x0B' BVLCOriginalBroadcastNPDU(uint 16 bvlcPayloadLength)
            [simple NPDU('bvlcPayloadLength')
                            npdu                                ]
        ]
        ['0x0C' BVLCSecureBVLL(uint 16 bvlcPayloadLength)
            [array byte     securityWrapper
                            length 'bvlcPayloadLength'          ]
        ]
    ]
]

[type BVLCBroadcastDistributionTableEntry
    [array  uint 8      ip                          count '4'       ]
    [simple uint 16     port                                        ]
    [array  uint 8      broadcastDistributionMap    count '4'       ]
]

[type BVLCForeignDeviceTableEntry
    [array  uint 8      ip                          count '4'       ]
    [simple uint 16     port                                        ]
    [simple uint 16     ttl                                         ]
    [simple uint 16     secondRemainingBeforePurge                  ]
]

[type NPDU(uint 16 npduLength)
    [simple     uint 8      protocolVersionNumber                                                                   ]
    [simple     NPDUControl control                                                                                 ]
    [optional   uint 16     destinationNetworkAddress   'control.destinationSpecified'                              ]
    [optional   uint 8      destinationLength           'control.destinationSpecified'                              ]
    [array      uint 8      destinationAddress count    'control.destinationSpecified ? destinationLength : 0'      ]
    [optional   uint 16     sourceNetworkAddress        'control.sourceSpecified'                                   ]
    [optional   uint 8      sourceLength                'control.sourceSpecified'                                   ]
    [array      uint 8      sourceAddress count         'control.sourceSpecified ? sourceLength : 0'                ]
    [optional   uint 8      hopCount                    'control.destinationSpecified'                              ]
    [virtual    uint 16     sourceLengthAddon           'control.sourceSpecified ? 3 + sourceLength : 0'            ]
    [virtual    uint 16     destinationLengthAddon      'control.destinationSpecified ? 3 + destinationLength : 0'  ]
    [virtual    uint 16     payloadSubtraction         '2 + (sourceLengthAddon + destinationLengthAddon + ((control.destinationSpecified || control.sourceSpecified) ? 1 : 0))'     ]
    [optional   NLM('npduLength - payloadSubtraction')
                            nlm
                                                        'control.messageTypeFieldPresent'                           ]
    [optional   APDU('npduLength - payloadSubtraction')
                            apdu
                                                        '!control.messageTypeFieldPresent'                          ]
    [validation    'nlm != null || apdu != null'        "something is wrong here... apdu and nlm not set"           ]
]

[type NPDUControl
    [simple     bit         messageTypeFieldPresent]
    [reserved   uint 1      '0']
    [simple     bit         destinationSpecified]
    [reserved   uint 1      '0']
    [simple     bit         sourceSpecified]
    [simple     bit         expectingReply]
    [simple     NPDUNetworkPriority
                            networkPriority
    ]
]

[enum uint 2 NPDUNetworkPriority
    ['3' LIFE_SAVETY_MESSAGE        ]
    ['2' CRITICAL_EQUIPMENT_MESSAGE ]
    ['1' URGENT_MESSAGE             ]
    ['0' NORMAL_MESSAGE             ]
]

[discriminatedType NLM(uint 16 apduLength)
    [discriminator uint 8  messageType]
    [optional      uint 16 vendorId '(messageType >= 128) && (messageType <= 255)']
    [typeSwitch messageType
        ['0x00' NLMWhoIsRouterToNetwork(uint 8 messageType)
            [array      uint 16     destinationNetworkAddress length 'apduLength - (((messageType >= 128) && (messageType <= 255)) ? 3 : 1)']
        ]
        ['0x01' NLMIAmRouterToNetwork(uint 8 messageType)
            [array      uint 16     destinationNetworkAddress length 'apduLength - (((messageType >= 128) && (messageType <= 255)) ? 3 : 1)']
        ]
        ['0x02' NLMICouldBeRouterToNetwork(uint 8 messageType)
            [simple     uint 16     destinationNetworkAddress   ]
            [simple     uint 8      performanceIndex            ]
        ]
        ['0x03' NLMRejectRouterToNetwork(uint 8 messageType)
            [simple     NLMRejectRouterToNetworkRejectReason
                                    rejectReason                ]
            [simple     uint 16     destinationNetworkAddress   ]
        ]
        ['0x04' NLMRouterBusyToNetwork(uint 8 messageType)
            [array      uint 16     destinationNetworkAddress length 'apduLength - (((messageType >= 128) && (messageType <= 255)) ? 3 : 1)']
        ]
        ['0x05' NLMRouterAvailableToNetwork(uint 8 messageType)
            [array      uint 16     destinationNetworkAddress length 'apduLength - (((messageType >= 128) && (messageType <= 255)) ? 3 : 1)']
        ]
        ['0x06' NLMInitalizeRoutingTable(uint 8 messageType)
            [simple     uint 8      numberOfPorts                   ]
            [array      NLMInitalizeRoutingTablePortMapping
                                    portMappings
                        count 'numberOfPorts'                       ]
        ]
        ['0x07' NLMInitalizeRoutingTableAck(uint 8 messageType)
            [simple     uint 8      numberOfPorts                   ]
            [array      NLMInitalizeRoutingTablePortMapping
                                    portMappings
                        count 'numberOfPorts'                       ]
        ]
        ['0x08' NLMEstablishConnectionToNetwork(uint 8 messageType)
            [simple     uint 16     destinationNetworkAddress   ]
            [simple     uint 8      terminationTime             ]
        ]
        ['0x09' NLMDisconnectConnectionToNetwork(uint 8 messageType)
            [simple     uint 16     destinationNetworkAddress   ]
        ]
    ]
]

[enum uint 8  NLMRejectRouterToNetworkRejectReason
    ['0'    OTHER]
    ['1'    NOT_DIRECTLY_CONNECTED]
    ['2'    BUSY]
    ['3'    UNKNOWN_NLMT]
    ['4'    TOO_LONG]
    ['5'    SECURITY_ERROR]
    ['6'    ADDRESSING_ERROR]
]

[type NLMInitalizeRoutingTablePortMapping
    [simple     uint 16     destinationNetworkAddress       ]
    [simple     uint 8      portId                          ]
    [simple     uint 8      portInfoLength                  ]
    [array      byte        portInfo count 'portInfoLength' ]
]

[discriminatedType APDU(uint 16 apduLength)
    [discriminator uint 4 apduType]
    [typeSwitch apduType
        ['0x0' APDUConfirmedRequest
            [simple   bit       segmentedMessage                         ]
            [simple   bit       moreFollows                              ]
            [simple   bit       segmentedResponseAccepted                ]
            [reserved uint 2    '0'                                      ]
            [simple   MaxSegmentsAccepted
                                maxSegmentsAccepted                      ]
            [simple   MaxApduLengthAccepted
                                maxApduLengthAccepted                    ]
            [simple   uint 8    invokeId                                 ]
            [optional uint 8    sequenceNumber       'segmentedMessage'  ]
            [optional uint 8    proposedWindowSize   'segmentedMessage'  ]
            [optional BACnetConfirmedServiceRequest('apduLength - (4 + (segmentedMessage ? 2 : 0))')
                                serviceRequest       '!segmentedMessage' ]
            // TODO: maybe we should put this in the discriminated types below
            [optional uint 8    segmentServiceChoice 'segmentedMessage && sequenceNumber != 0']
            [array    byte      segment
                                length
                                'segmentedMessage?((apduLength>0)?(apduLength - ((sequenceNumber != 0)?6:5)):0):0'
        ]
        ['0x1' APDUUnconfirmedRequest
            [reserved uint 4                          '0'               ]
            [simple   BACnetUnconfirmedServiceRequest('apduLength - 1')
                                serviceRequest                          ]
        ]
        ['0x2' APDUSimpleAck
            [reserved uint 4    '0'                                     ]
            [simple   uint 8    originalInvokeId                        ]
            [simple   uint 8    serviceChoice                           ]
        ]
        ['0x3' APDUComplexAck
            [simple   bit       segmentedMessage                        ]
            [simple   bit       moreFollows                             ]
            [reserved uint 2    '0'                                     ]
            [simple   uint 8    originalInvokeId                        ]
            [optional uint 8    sequenceNumber     'segmentedMessage'   ]
            [optional uint 8    proposedWindowSize 'segmentedMessage'   ]
            [optional BACnetServiceAck('apduLength - (3 + (segmentedMessage ? 2 : 0))')
                                serviceAck         '!segmentedMessage'  ]
            // TODO: maybe we should put this in the discriminated types below
            [optional uint 8    segmentServiceChoice 'segmentedMessage && sequenceNumber != 0']
            [array    byte      segment
                                length
                                'segmentedMessage?((apduLength>0)?(apduLength - ((sequenceNumber != 0)?5:4)):0):0'
                                                                        ]
        ]
        ['0x4' APDUSegmentAck
            [reserved uint 2    '0x00'                                  ]
            [simple   bit       negativeAck                             ]
            [simple   bit       server                                  ]
            [simple   uint 8    originalInvokeId                        ]
            [simple   uint 8    sequenceNumber                          ]
            [simple   uint 8    proposedWindowSize                      ]
        ]
        ['0x5' APDUError
            [reserved uint 4    '0x00'                                  ]
            [simple   uint 8    originalInvokeId]
            [simple   BACnetError
                                error                                   ]
        ]
        ['0x6' APDUReject
            [reserved uint 4    '0x00'                                  ]
            [simple   uint 8    originalInvokeId                        ]
            [simple   uint 8    rejectReason                            ]
        ]
        ['0x7' APDUAbort
            [reserved uint 3    '0x00'                                  ]
            [simple   bit       server                                  ]
            [simple   uint 8    originalInvokeId                        ]
            [simple   uint 8    abortReason                             ]
        ]
        [APDUUnknown
            [array    byte      unknownBytes length '(apduLength>0)?(apduLength - 1):0'    ]
        ]
    ]
]

[enum uint 3 MaxSegmentsAccepted
    ['0x0' UNSPECIFIED             ]
    ['0x1' NUM_SEGMENTS_02         ]
    ['0x2' NUM_SEGMENTS_04         ]
    ['0x3' NUM_SEGMENTS_08         ]
    ['0x4' NUM_SEGMENTS_16         ]
    ['0x5' NUM_SEGMENTS_32         ]
    ['0x6' NUM_SEGMENTS_64         ]
    ['0x7' MORE_THAN_64_SEGMENTS   ]
]

[enum uint 4 MaxApduLengthAccepted
    ['0x0' MINIMUM_MESSAGE_SIZE    ] // 50 octets
    ['0x1' NUM_OCTETS_128          ]
    ['0x2' NUM_OCTETS_206          ] // fits in a LonTalk frame
    ['0x3' NUM_OCTETS_480          ] // fits in an ARCNET frame
    ['0x4' NUM_OCTETS_1024         ]
    ['0x5' NUM_OCTETS_1476         ] // fits in an Ethernet frame
    ['0x6' RESERVED_BY_ASHRAE_01   ]
    ['0x7' RESERVED_BY_ASHRAE_02   ]
    ['0x8' RESERVED_BY_ASHRAE_03   ]
    ['0x9' RESERVED_BY_ASHRAE_04   ]
    ['0xA' RESERVED_BY_ASHRAE_05   ]
    ['0xB' RESERVED_BY_ASHRAE_06   ]
    ['0xC' RESERVED_BY_ASHRAE_07   ]
    ['0xD' RESERVED_BY_ASHRAE_08   ]
    ['0xE' RESERVED_BY_ASHRAE_09   ]
    ['0xF' RESERVED_BY_ASHRAE_10   ]
]

[discriminatedType BACnetConfirmedServiceRequest(uint 16 serviceRequestLength)
    [discriminator BACnetConfirmedServiceChoice serviceChoice]
    [typeSwitch serviceChoice
        ////
        // Alarm and Event Services

        ['ACKNOWLEDGE_ALARM' BACnetConfirmedServiceRequestAcknowledgeAlarm
            [simple   BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          acknowledgingProcessIdentifier ]
            [simple   BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') eventObjectIdentifier          ]
            [simple   BACnetContextTagEventState('2', 'BACnetDataType.EVENT_STATE')                    eventStateAcknowledged         ]
            [simple   BACnetTimeStampEnclosed('3')                                                     timestamp                      ]
            [simple   BACnetContextTagCharacterString('4', 'BACnetDataType.CHARACTER_STRING')          acknowledgmentSource           ]
            [simple   BACnetTimeStampEnclosed('5')                                                     timeOfAcknowledgment           ]
        ]
        ['CONFIRMED_COV_NOTIFICATION' BACnetConfirmedServiceRequestConfirmedCOVNotification
            [simple   BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          subscriberProcessIdentifier ]
            [simple   BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') initiatingDeviceIdentifier  ]
            [simple   BACnetContextTagObjectIdentifier('2', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') monitoredObjectIdentifier   ]
            [simple   BACnetContextTagUnsignedInteger('3', 'BACnetDataType.UNSIGNED_INTEGER')          lifetimeInSeconds           ]
            [simple   BACnetPropertyValues('4', 'monitoredObjectIdentifier.objectType')                listOfValues                ]
        ]
        ['CONFIRMED_COV_NOTIFICATION_MULTIPLE' BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
            [simple   BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          subscriberProcessIdentifier ]
            [simple   BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') initiatingDeviceIdentifier  ]
            [simple   BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')          timeRemaining               ]
            [optional BACnetTimeStampEnclosed('3')                                                     timestamp                   ]
            [simple   BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotifications('4')
                                                                                                       listOfCovNotifications      ]
        ]
        ['CONFIRMED_EVENT_NOTIFICATION' BACnetConfirmedServiceRequestConfirmedEventNotification // Spec complete
            [simple   BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          processIdentifier            ]
            [simple   BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') initiatingDeviceIdentifier   ]
            [simple   BACnetContextTagObjectIdentifier('2', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') eventObjectIdentifier        ]
            [simple   BACnetTimeStampEnclosed('3')                                                     timestamp                    ]
            [simple   BACnetContextTagUnsignedInteger('4', 'BACnetDataType.UNSIGNED_INTEGER')          notificationClass            ]
            [simple   BACnetContextTagUnsignedInteger('5', 'BACnetDataType.UNSIGNED_INTEGER')          priority                     ]
            [simple   BACnetContextTagEventType('6', 'BACnetDataType.EVENT_TYPE')                      eventType                    ]
            [optional BACnetContextTagCharacterString('7', 'BACnetDataType.CHARACTER_STRING')          messageText                  ]
            [simple   BACnetContextTagNotifyType('8', 'BACnetDataType.NOTIFY_TYPE')                    notifyType                   ]
            [optional BACnetContextTagBoolean('9', 'BACnetDataType.BOOLEAN')                           ackRequired                  ]
            [optional BACnetContextTagEventState('10', 'BACnetDataType.EVENT_STATE')                   fromState                    ]
            [simple   BACnetContextTagEventState('11', 'BACnetDataType.EVENT_STATE')                   toState                      ]
            [optional BACnetNotificationParameters('12', 'eventObjectIdentifier.objectType')           eventValues                  ]
        ]
        ['GET_ENROLLMENT_SUMMARY' BACnetConfirmedServiceRequestGetEnrollmentSummary
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['GET_EVENT_INFORMATION' BACnetConfirmedServiceRequestGetEventInformation
            [optional BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') lastReceivedObjectIdentifier   ]
        ]
        ['LIFE_SAFETY_OPERATION' BACnetConfirmedServiceRequestLifeSafetyOperation
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['SUBSCRIBE_COV' BACnetConfirmedServiceRequestSubscribeCOV
            [simple BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          subscriberProcessIdentifier ]
            [simple BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') monitoredObjectIdentifier   ]
            [simple BACnetContextTagBoolean('2', 'BACnetDataType.BOOLEAN')                           issueConfirmed              ]
            [simple BACnetContextTagUnsignedInteger('3', 'BACnetDataType.UNSIGNED_INTEGER')          lifetimeInSeconds           ]
        ]
        ['SUBSCRIBE_COV_PROPERTY' BACnetConfirmedServiceRequestSubscribeCOVProperty
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['SUBSCRIBE_COV_PROPERTY_MULTIPLE' BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        //
        ////

        ////
        // File Access Services

        ['ATOMIC_READ_FILE' BACnetConfirmedServiceRequestAtomicReadFile
            [simple BACnetApplicationTagObjectIdentifier                        fileIdentifier      ]
            [simple BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord   accessMethod        ]
        ]
        ['ATOMIC_WRITE_FILE' BACnetConfirmedServiceRequestAtomicWriteFile
            [simple BACnetApplicationTagObjectIdentifier                  deviceIdentifier    ]
            [optional BACnetOpeningTag('0', 'BACnetDataType.OPENING_TAG') openingTag          ]
            [simple BACnetApplicationTagSignedInteger                     fileStartPosition   ]
            [simple BACnetApplicationTagOctetString                       fileData            ]
            [optional BACnetClosingTag('0', 'BACnetDataType.CLOSING_TAG') closingTag          ]
        ]
        //
        ////

        ////
        // Object Access Services
        ['ADD_LIST_ELEMENT' BACnetConfirmedServiceRequestAddListElement
            [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')     objectIdentifier    ]
            [simple   BACnetContextTagPropertyIdentifier('1', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER') propertyIdentifier  ]
            [optional BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')              arrayIndex          ]
            [optional BACnetConstructedData('3', 'objectIdentifier.objectType', 'propertyIdentifier')      listOfElements      ]
        ]
        ['REMOVE_LIST_ELEMENT' BACnetConfirmedServiceRequestRemoveListElement
            [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')     objectIdentifier    ]
            [simple   BACnetContextTagPropertyIdentifier('1', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER') propertyIdentifier  ]
            [optional BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')              arrayIndex          ]
            [optional BACnetConstructedData('3', 'objectIdentifier.objectType', 'propertyIdentifier')      listOfElements      ]
        ]
        ['CREATE_OBJECT' BACnetConfirmedServiceRequestCreateObject
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['DELETE_OBJECT' BACnetConfirmedServiceRequestDeleteObject
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['READ_PROPERTY' BACnetConfirmedServiceRequestReadProperty
            [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                            objectIdentifier        ]
            [simple   BACnetContextTagPropertyIdentifier('1', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER')
                            propertyIdentifier      ]
            [optional BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')
                            arrayIndex              ]
        ]
        ['READ_PROPERTY_MULTIPLE' BACnetConfirmedServiceRequestReadPropertyMultiple
            [array    BACnetReadAccessSpecification
                            data
                            length
                            'serviceRequestLength'                   ]
        ]
        ['READ_RANGE' BACnetConfirmedServiceRequestReadRange
            [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')     objectIdentifier    ]
            [simple   BACnetContextTagPropertyIdentifier('1', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER') propertyIdentifier  ]
            [optional BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')              propertyArrayIndex  ]
            // TODO: this attribute should be named range but this is a keyword in golang (so at this point we should build a language translator which makes keywords safe)
            [optional BACnetConfirmedServiceRequestReadRangeRange                                          readRange           ]
        ]
        ['WRITE_PROPERTY' BACnetConfirmedServiceRequestWriteProperty
            [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')     objectIdentifier    ]
            [simple   BACnetContextTagPropertyIdentifier('1', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER') propertyIdentifier  ]
            [optional BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')              arrayIndex          ]
            [simple   BACnetConstructedData('3', 'objectIdentifier.objectType', 'propertyIdentifier')      propertyValue       ]
            [optional BACnetContextTagUnsignedInteger('4', 'BACnetDataType.UNSIGNED_INTEGER')              priority            ]
        ]
        ['WRITE_PROPERTY_MULTIPLE' BACnetConfirmedServiceRequestWritePropertyMultiple
            [array    BACnetWriteAccessSpecification
                            data
                            length
                            'serviceRequestLength'                   ]
        ]
        //
        ////

        ////
        // Remote Device Management Services

        ['DEVICE_COMMUNICATION_CONTROL' BACnetConfirmedServiceRequestDeviceCommunicationControl
            [optional BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')              timeDuration        ]
            [simple   BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable('1')                    enableDisable       ]
            [optional BACnetContextTagCharacterString('2', 'BACnetDataType.CHARACTER_STRING')              password            ]

        ]
        ['CONFIRMED_PRIVATE_TRANSFER' BACnetConfirmedServiceRequestConfirmedPrivateTransfer
            [simple     BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          vendorId                    ]// TODO: vendor list?
            [simple     BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')          serviceNumber               ]
            // TODO: temporary dummy property identifier... get rid of that
            [optional BACnetConstructedData('2', 'BACnetObjectType.VENDOR_PROPRIETARY_VALUE', 'STATIC_CALL("dummyPropertyIdentifier")') serviceParameters           ]
        ]
        ['CONFIRMED_TEXT_MESSAGE' BACnetConfirmedServiceRequestConfirmedTextMessage
             // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['REINITIALIZE_DEVICE' BACnetConfirmedServiceRequestReinitializeDevice
          [simple BACnetContextTagDeviceState('0', 'BACnetDataType.BACNET_DEVICE_STATE')     reinitializedStateOfDevice  ]
          [optional BACnetContextTagCharacterString('1', 'BACnetDataType.CHARACTER_STRING')  password                    ]
        ]

        ////
        //  Virtual Terminal Services

        ['VT_OPEN' BACnetConfirmedServiceRequestVTOpen
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['VT_CLOSE' BACnetConfirmedServiceRequestVTClose
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['VT_DATA' BACnetConfirmedServiceRequestVTData
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        //
        ////

        ////
        //  Removed Services

        ['AUTHENTICATE' BACnetConfirmedServiceRequestAuthenticate
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['REQUEST_KEY' BACnetConfirmedServiceRequestRequestKey
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['READ_PROPERTY_CONDITIONAL' BACnetConfirmedServiceRequestReadPropertyConditional
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        //
        ////

        [BACnetConfirmedServiceRequestConfirmedUnknown
            [array  byte    unknownBytes length '(serviceRequestLength>0)?(serviceRequestLength - 1):0']
        ]
    ]
]

[enum uint 8 BACnetConfirmedServiceChoice
    ////
    // Alarm and Event Services

    ['0x00' ACKNOWLEDGE_ALARM                       ]
    ['0x01' CONFIRMED_COV_NOTIFICATION              ]
    ['0x1F' CONFIRMED_COV_NOTIFICATION_MULTIPLE     ]
    ['0x02' CONFIRMED_EVENT_NOTIFICATION            ]
    ['0x03' GET_ALARM_SUMMARY                       ]
    ['0x04' GET_ENROLLMENT_SUMMARY                  ]
    ['0x1D' GET_EVENT_INFORMATION                   ]
    ['0x1B' LIFE_SAFETY_OPERATION                   ]
    ['0x05' SUBSCRIBE_COV                           ]
    ['0x1C' SUBSCRIBE_COV_PROPERTY                  ]
    ['0x1E' SUBSCRIBE_COV_PROPERTY_MULTIPLE         ]
    //
    ////

    ////
    // File Access Services

    ['0x06' ATOMIC_READ_FILE                        ]
    ['0x07' ATOMIC_WRITE_FILE                       ]
    //
    ////

    ////
    // Object Access Services

    ['0x08' ADD_LIST_ELEMENT                        ]
    ['0x09' REMOVE_LIST_ELEMENT                     ]
    ['0x0A' CREATE_OBJECT                           ]
    ['0x0B' DELETE_OBJECT                           ]
    ['0x0C' READ_PROPERTY                           ]
    ['0x0E' READ_PROPERTY_MULTIPLE                  ]
    ['0x1A' READ_RANGE                              ]
    ['0x0F' WRITE_PROPERTY                          ]
    ['0x10' WRITE_PROPERTY_MULTIPLE                 ]
    //
    ////

    ////
    // Remote Device Management Services

    ['0x11' DEVICE_COMMUNICATION_CONTROL            ]
    ['0x12' CONFIRMED_PRIVATE_TRANSFER              ]
    ['0x13' CONFIRMED_TEXT_MESSAGE                  ]
    ['0x14' REINITIALIZE_DEVICE                     ]
    //
    ////

    ////
    //  Virtual Terminal Services

    ['0x15' VT_OPEN                                 ]
    ['0x16' VT_CLOSE                                ]
    ['0x17' VT_DATA                                 ]
    //
    ////

    ////
    //  Removed Services

    ['0x18' AUTHENTICATE                            ]
    ['0x19' REQUEST_KEY                             ]
    ['0x0D' READ_PROPERTY_CONDITIONAL               ]
    //
    ////
]

[type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotifications(uint 8 tagNumber)
    [simple     BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                     openingTag                     ]
    [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') monitoredObjectIdentifier   ]
    [simple     BACnetOpeningTag('1', 'BACnetDataType.OPENING_TAG')
                         innerOpeningTag            ]
    [array    BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue('monitoredObjectIdentifier.objectType')
                        listOfValues
                        terminated
                        'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, 1)'
    ]
    [simple     BACnetClosingTag('1', 'BACnetDataType.CLOSING_TAG')
                     innerClosingTag                ]
    [simple     BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                     closingTag                     ]
]

[type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleListOfCovNotificationsValue(BACnetObjectType objectType)
    [simple   BACnetContextTagPropertyIdentifier('0', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER')
                                propertyIdentifier      ]
    [optional BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')
                                arrayIndex              ]
    [simple   BACnetConstructedData('2', 'objectType', 'propertyIdentifier')
                                propertyValue           ]
    [optional BACnetContextTagTime('3', 'BACnetDataType.TIME')
                                timeOfChange            ]
]

[type BACnetReadAccessSpecification
    [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                    objectIdentifier                ]
    [simple     BACnetOpeningTag('1', 'BACnetDataType.OPENING_TAG')
                     openingTag                     ]
    [array    BACnetPropertyReference
                    listOfPropertyReferences
                    terminated
                    'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, 1)'
    ]
    [simple     BACnetClosingTag('1', 'BACnetDataType.CLOSING_TAG')
                     closingTag                     ]
]

[type BACnetPropertyReference
    [simple     BACnetContextTagPropertyIdentifier('0', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER')
                    propertyIdentifier              ]
    [optional   BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')
                    arrayIndex                      ]
]

[type BACnetWriteAccessSpecification
    [simple     BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                    objectIdentifier                ]
    [simple     BACnetOpeningTag('1', 'BACnetDataType.OPENING_TAG')
                     openingTag                     ]
    [array      BACnetPropertyWriteDefinition('objectIdentifier.objectType')
                    listOfPropertyWriteDefinition
                    terminated
                    'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, 1)'
    ]
    [simple     BACnetClosingTag('1', 'BACnetDataType.CLOSING_TAG')
                     closingTag                     ]
]

[type BACnetConfirmedServiceRequestReadRangeRange
    [peek       BACnetTagHeader
                    peekedTagHeader                 ]
    [simple     BACnetOpeningTag('peekedTagHeader.actualTagNumber', 'BACnetDataType.OPENING_TAG')
                     openingTag                     ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber']
    [typeSwitch peekedTagNumber
        ['0x3' BACnetConfirmedServiceRequestReadRangeRangeByPosition
            [simple BACnetApplicationTagUnsignedInteger                   referenceIndex            ]
            [simple BACnetApplicationTagSignedInteger                     count                     ]
        ]
        ['0x6' BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
            [simple BACnetApplicationTagUnsignedInteger                   referenceSequenceNumber   ]
            [simple BACnetApplicationTagSignedInteger                     count                     ]
        ]
        ['0x7' BACnetConfirmedServiceRequestReadRangeRangeByTime
            [simple BACnetDateTime                                        referenceTime             ]
            [simple BACnetApplicationTagSignedInteger                     count                     ]
        ]
    ]
    [simple     BACnetClosingTag('peekedTagHeader.actualTagNumber', 'BACnetDataType.CLOSING_TAG')
                     closingTag
    ]
]

[type BACnetPropertyWriteDefinition(BACnetObjectType objectType)
    [simple     BACnetContextTagPropertyIdentifier('0', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER')
                    propertyIdentifier              ]
    [optional   BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')
                    arrayIndex                      ]
    [optional   BACnetConstructedData('2', 'objectType', 'propertyIdentifier')
                    propertyValue                   ]
    [optional   BACnetContextTagUnsignedInteger('3', 'BACnetDataType.UNSIGNED_INTEGER')
                    priority                        ]
]

// TODO: this is a enum so we should build a static call which maps a enum (could be solved by using only the tag header with a length validation and the enum itself)
[type BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable(uint 8 tagNumber)
    [optional   BACnetContextTagEnumerated('tagNumber', 'BACnetDataType.ENUMERATED')
                    rawData                         ]
    [virtual    bit isEnable            'rawData != null && rawData.payload.actualValue == 0']
    [virtual    bit isDisable           'rawData != null && rawData.payload.actualValue == 1']
    [virtual    bit isDisableInitiation 'rawData != null && rawData.payload.actualValue == 2']
]

[type BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
    [peek       BACnetTagHeader
                    peekedTagHeader                 ]
    [simple     BACnetOpeningTag('peekedTagHeader.actualTagNumber', 'BACnetDataType.OPENING_TAG')
                     openingTag                     ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber']
    [typeSwitch peekedTagNumber
        ['0x0' BACnetConfirmedServiceRequestAtomicReadFileStream
            [simple BACnetApplicationTagSignedInteger                     fileStartPosition   ]
            [simple BACnetApplicationTagUnsignedInteger                   requestOctetCount   ]
        ]
        ['0x1' BACnetConfirmedServiceRequestAtomicReadFileRecord
            [simple BACnetApplicationTagSignedInteger                     fileStartRecord     ]
            [simple BACnetApplicationTagUnsignedInteger                   requestRecordCount  ]
        ]
    ]
    [simple     BACnetClosingTag('peekedTagHeader.actualTagNumber', 'BACnetDataType.CLOSING_TAG')
                     closingTag
    ]
]

[discriminatedType BACnetUnconfirmedServiceRequest(uint 16 serviceRequestLength)
    [discriminator BACnetUnconfirmedServiceChoice serviceChoice]
    [typeSwitch serviceChoice
        ['I_AM' BACnetUnconfirmedServiceRequestIAm
            [simple     BACnetApplicationTagObjectIdentifier    deviceIdentifier                ]
            [simple     BACnetApplicationTagUnsignedInteger     maximumApduLengthAcceptedLength ]
            [simple     BACnetSegmentation                      segmentationSupported           ]
            [simple     BACnetApplicationTagUnsignedInteger     vendorId                        ] // TODO: vendor list?
        ]
        ['I_HAVE' BACnetUnconfirmedServiceRequestIHave
            [simple     BACnetApplicationTagObjectIdentifier    deviceIdentifier    ]
            [simple     BACnetApplicationTagObjectIdentifier    objectIdentifier    ]
            [simple     BACnetApplicationTagCharacterString     objectName          ]
        ]
        ['UNCONFIRMED_COV_NOTIFICATION' BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
            [simple     BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          subscriberProcessIdentifier ]
            [simple     BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') initiatingDeviceIdentifier  ]
            [simple     BACnetContextTagObjectIdentifier('2', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') monitoredObjectIdentifier   ]
            [simple     BACnetContextTagUnsignedInteger('3', 'BACnetDataType.UNSIGNED_INTEGER')          lifetimeInSeconds           ]
            [simple     BACnetPropertyValues('4', 'monitoredObjectIdentifier.objectType')                listOfValues                ]
        ]
        ['UNCONFIRMED_EVENT_NOTIFICATION' BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
            [simple   BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          processIdentifier            ]
            [simple   BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') initiatingDeviceIdentifier   ]
            [simple   BACnetContextTagObjectIdentifier('2', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER') eventObjectIdentifier        ]
            [simple   BACnetTimeStampEnclosed('3')                                                     timestamp                    ]
            [simple   BACnetContextTagUnsignedInteger('4', 'BACnetDataType.UNSIGNED_INTEGER')          notificationClass            ]
            [simple   BACnetContextTagUnsignedInteger('5', 'BACnetDataType.UNSIGNED_INTEGER')          priority                     ]
            [simple   BACnetContextTagEventType('6', 'BACnetDataType.EVENT_TYPE')                      eventType                    ]
            [optional BACnetContextTagCharacterString('7', 'BACnetDataType.CHARACTER_STRING')          messageText                  ]
            [simple   BACnetContextTagNotifyType('8', 'BACnetDataType.NOTIFY_TYPE')                    notifyType                   ]
            [optional BACnetContextTagBoolean('9', 'BACnetDataType.BOOLEAN')                           ackRequired                  ]
            [optional BACnetContextTagEventState('10', 'BACnetDataType.EVENT_STATE')                   fromState                    ]
            [simple   BACnetContextTagEventState('11', 'BACnetDataType.EVENT_STATE')                   toState                      ]
            [optional BACnetNotificationParameters('12', 'eventObjectIdentifier.objectType')           eventValues                  ]
        ]
        ['UNCONFIRMED_PRIVATE_TRANSFER' BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
            [simple     BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          vendorId                    ]// TODO: vendor list?
            [simple     BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')          serviceNumber               ]
            // TODO: temporary dummy property identifier... get rid of that
            [optional BACnetConstructedData('2', 'BACnetObjectType.VENDOR_PROPRIETARY_VALUE', 'STATIC_CALL("dummyPropertyIdentifier")') serviceParameters           ]
        ]
        ['UNCONFIRMED_TEXT_MESSAGE' BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['TIME_SYNCHRONIZATION' BACnetUnconfirmedServiceRequestTimeSynchronization
            [simple BACnetApplicationTagDate synchronizedDate]
            [simple BACnetApplicationTagTime synchronizedTime]
        ]
        ['WHO_HAS' BACnetUnconfirmedServiceRequestWhoHas
            [optional BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')           deviceInstanceRangeLowLimit                                         ]
            [optional BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')           deviceInstanceRangeHighLimit  'deviceInstanceRangeLowLimit != null' ]
            [optional BACnetContextTagObjectIdentifier('2', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')  objectIdentifier                                                    ]
            [optional BACnetContextTagOctetString('3', 'BACnetDataType.OCTET_STRING')                   objectName                    'objectIdentifier == null'            ]
        ]
        ['WHO_IS' BACnetUnconfirmedServiceRequestWhoIs
            [optional BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')   deviceInstanceRangeLowLimit                                                 ]
            [optional BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')   deviceInstanceRangeHighLimit  'deviceInstanceRangeLowLimit != null'         ]
        ]
        ['UTC_TIME_SYNCHRONIZATION' BACnetUnconfirmedServiceRequestUTCTimeSynchronization
            [simple BACnetApplicationTagDate synchronizedDate]
            [simple BACnetApplicationTagTime synchronizedTime]
        ]
        ['WRITE_GROUP' BACnetUnconfirmedServiceRequestWriteGroup
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['UNCONFIRMED_COV_NOTIFICATION_MULTIPLE' BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        [BACnetUnconfirmedServiceRequestUnconfirmedUnknown
            [array  byte    unknownBytes length '(serviceRequestLength>0)?(serviceRequestLength - 1):0']
        ]
    ]
]

[enum uint 8 BACnetUnconfirmedServiceChoice
    ['0x00' I_AM                                        ]
    ['0x01' I_HAVE                                      ]
    ['0x02' UNCONFIRMED_COV_NOTIFICATION                ]
    ['0x03' UNCONFIRMED_EVENT_NOTIFICATION              ]
    ['0x04' UNCONFIRMED_PRIVATE_TRANSFER                ]
    ['0x05' UNCONFIRMED_TEXT_MESSAGE                    ]
    ['0x06' TIME_SYNCHRONIZATION                        ]
    ['0x07' WHO_HAS                                     ]
    ['0x08' WHO_IS                                      ]
    ['0x09' UTC_TIME_SYNCHRONIZATION                    ]
    ['0x0A' WRITE_GROUP                                 ]
    ['0x0B' UNCONFIRMED_COV_NOTIFICATION_MULTIPLE       ]
]

// TODO: this is a enum so we should build a static call which maps a enum (could be solved by using only the tag header with a length validation and the enum itself)
[type BACnetSegmentation
    [simple BACnetApplicationTagEnumerated          rawData ]
    [virtual    bit isSegmentedBoth           'rawData != null && rawData.payload.actualValue == 0']
    [virtual    bit isSegmentedTransmit       'rawData != null && rawData.payload.actualValue == 1']
    [virtual    bit isSegmentedReceive        'rawData != null && rawData.payload.actualValue == 3']
    [virtual    bit isNoSegmentation          'rawData != null && rawData.payload.actualValue == 4']
]

[discriminatedType BACnetServiceAck(uint 16 serviceRequestLength)
    [discriminator   BACnetConfirmedServiceChoice
                        serviceChoice                   ]
    [typeSwitch serviceChoice
        ////
        // Alarm and Event Services

        ['ACKNOWLEDGE_ALARM' BACnetServiceAckAcknowledgeAlarm
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['CONFIRMED_COV_NOTIFICATION' BACnetServiceAckConfirmedCovNotification
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['CONFIRMED_COV_NOTIFICATION_MULTIPLE' BACnetServiceAckConfirmedCovNotificationMultiple
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['CONFIRMED_EVENT_NOTIFICATION' BACnetServiceAckConfirmedEventNotification
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['GET_ALARM_SUMMARY' BACnetServiceAckGetAlarmSummary
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['GET_ENROLLMENT_SUMMARY' BACnetServiceAckGetEnrollmentSummary
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['GET_EVENT_INFORMATION' BACnetServiceAckGetEventInformation
            [simple BACnetEventSummaries
                        listOfEventSummaries        ]
            [simple BACnetContextTagBoolean('0', 'BACnetDataType.BOOLEAN')
                        moreEvents                  ]
        ]
        ['LIFE_SAFETY_OPERATION' BACnetServiceAckLifeSafetyOperation
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['SUBSCRIBE_COV' BACnetServiceAckSubscribeCov
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['SUBSCRIBE_COV_PROPERTY' BACnetServiceAckSubscribeCovProperty
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['SUBSCRIBE_COV_PROPERTY_MULTIPLE' BACnetServiceAckSubscribeCovPropertyMultiple
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        //
        ////

        ////
        // File Access Services

        ['ATOMIC_READ_FILE' BACnetServiceAckAtomicReadFile
            [simple BACnetApplicationTagBoolean
                            endOfFile               ]
            [simple BACnetServiceAckAtomicReadFileStreamOrRecord
                            accessMethod            ]
        ]
        ['ATOMIC_WRITE_FILE' BACnetServiceAckAtomicWriteFile
            [simple BACnetContextTagSignedInteger('0', 'BACnetDataType.SIGNED_INTEGER')
                            fileStartPosition       ]
        ]
        //
        ////

        ////
        // Object Access Services
        ['ADD_LIST_ELEMENT' BACnetServiceAckAddListElement
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['REMOVE_LIST_ELEMENT' BACnetServiceAckRemoveListElement
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['CREATE_OBJECT' BACnetServiceAckCreateObject
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['READ_PROPERTY' BACnetServiceAckReadProperty
            [simple     BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                            objectIdentifier        ]
            [simple     BACnetContextTagPropertyIdentifier('1', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER')
                            propertyIdentifier      ]
            [optional   BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')
                            arrayIndex              ]
            [optional   BACnetConstructedData('3', 'objectIdentifier.objectType', 'propertyIdentifier')
                            values                  ]
        ]
        ['READ_PROPERTY_MULTIPLE' BACnetServiceAckReadPropertyMultiple
            [array    BACnetReadAccessResult
                            data
                            length
                            'serviceRequestLength'                   ]
        ]
        ['READ_RANGE' BACnetServiceAckReadRange
            [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')     objectIdentifier    ]
            [simple   BACnetContextTagPropertyIdentifier('1', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER') propertyIdentifier  ]
            [optional BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')              propertyArrayIndex  ]
            [simple   BACnetResultFlags('3')                                                               resultFlags         ]
            [simple   BACnetContextTagUnsignedInteger('4', 'BACnetDataType.UNSIGNED_INTEGER')              itemCount           ]
            [optional BACnetConstructedData('5', 'objectIdentifier.objectType', 'propertyIdentifier')      itemData            ]
            [optional BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')              firstSequenceNumber ]
        ]
        ['WRITE_PROPERTY' BACnetServiceAckWriteProperty
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['WRITE_PROPERTY_MULTIPLE' BACnetServiceAckWritePropertyMultiple
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        //
        ////


        ////
        // Remote Device Management Services

        ['DEVICE_COMMUNICATION_CONTROL' BACnetServiceAckDeviceCommunicationControl
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['CONFIRMED_PRIVATE_TRANSFER' BACnetServiceAckConfirmedPrivateTransfer
            [simple     BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')          vendorId                    ]// TODO: vendor list?
            [simple     BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')          serviceNumber               ]
            // TODO: temporary dummy property identifier... get rid of that
            [optional BACnetConstructedData('2', 'BACnetObjectType.VENDOR_PROPRIETARY_VALUE', 'STATIC_CALL("dummyPropertyIdentifier")') resultBlock                 ]
        ]
        ['CONFIRMED_TEXT_MESSAGE' BACnetServiceAckConfirmedTextMessage
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['REINITIALIZE_DEVICE' BACnetServiceAckReinitializeDevice
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        //
        ////

        ////
        //  Virtual Terminal Services

        ['VT_OPEN' BACnetServiceAckVTOpen
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['VT_CLOSE' BACnetServiceAckVTClose
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['VT_DATA' BACnetServiceAckVTData
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        //
        ////


        ////
        //  Removed Services

        ['AUTHENTICATE' BACnetServiceAckAuthenticate
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['REQUEST_KEY' BACnetServiceAckRequestKey
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        ['READ_PROPERTY_CONDITIONAL' BACnetServiceAckReadPropertyConditional
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
        //
        ////
    ]
]

// TODO: check if we should embed this type into BACnetServiceAckGetEventInformation as we double map
[type BACnetEventSummaries
    [simple     BACnetOpeningTag('0', 'BACnetDataType.OPENING_TAG')
                     openingTag                     ]
    [array    BACnetEventSummary
                         listOfEventSummaries
                         terminated
                         'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, 1)'
    ]
    [simple     BACnetClosingTag('0', 'BACnetDataType.CLOSING_TAG')
                     closingTag                     ]
]

[type BACnetEventSummary
    [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                    objectIdentifier                ]
    [simple   BACnetContextTagEventState('1', 'BACnetDataType.EVENT_STATE')
                    eventState                      ]
    [simple   BACnetEventTransitionBits('2')
                    acknowledgedTransitions         ]
    [simple   BACnetEventTimestamps('3')
                    eventTimestamps                 ]
    [simple   BACnetContextTagNotifyType('4', 'BACnetDataType.NOTIFY_TYPE')
                    notifyType                      ]
    [simple   BACnetEventTransitionBits('5')
                    eventEnable                     ]
    [simple   BACnetEventProrities('6')
                    eventPriorities                 ]
]

[type BACnetEventTimestamps(uint 8 tagNumber)
    [simple  BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                    openingTag
    ]
    [simple  BACnetTimeStamp
                    toOffnormal                     ]
    [simple  BACnetTimeStamp
                    toFault                         ]
    [simple  BACnetTimeStamp
                    toNormal                        ]
    [simple  BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                    closingTag
    ]
]

[type BACnetEventProrities(uint 8 tagNumber)
    [simple  BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                    openingTag
    ]
    [simple  BACnetApplicationTagUnsignedInteger
                    toOffnormal                     ]
    [simple  BACnetApplicationTagUnsignedInteger
                    toFault                         ]
    [simple  BACnetApplicationTagUnsignedInteger
                    toNormal                        ]
    [simple  BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                    closingTag
    ]
]

// TODO: this is a enum so we should build a static call which maps a enum (could be solved by using only the tag header with a length validation and the enum itself)
[type BACnetEventTransitionBits(uint 8 tagNumber)
    [simple BACnetContextTagBitString('tagNumber', 'BACnetDataType.BIT_STRING')
        rawBits
    ]
    [virtual    bit toOffnormal         'rawBits.payload.data[0]']
    [virtual    bit toFault             'rawBits.payload.data[1]']
    [virtual    bit toNormal            'rawBits.payload.data[2]']
]

[type BACnetReadAccessResult
    [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                    objectIdentifier                ]
    [simple     BACnetOpeningTag('1', 'BACnetDataType.OPENING_TAG')
                     openingTag                     ]
    [array    BACnetReadAccessProperty('objectIdentifier.objectType')
                    listOfReadAccessProperty
                    terminated
                    'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, 1)'
    ]
    [simple     BACnetClosingTag('1', 'BACnetDataType.CLOSING_TAG')
                     closingTag                     ]
]

[type BACnetReadAccessProperty(BACnetObjectType objectType)
    [simple     BACnetContextTagPropertyIdentifier('2', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER')
                    propertyIdentifier              ]
    [optional   BACnetContextTagUnsignedInteger('3', 'BACnetDataType.UNSIGNED_INTEGER')
                    arrayIndex                      ]
    [optional   BACnetConstructedData('4', 'objectType', 'propertyIdentifier')
                    propertyValue                   ]
    [optional   BACnetReadAccessPropertyError
                    propertyAccessError             ]
]

// TODO: this need to be completely refactored with BACnet error below
[type BACnetReadAccessPropertyError
    [simple     BACnetOpeningTag('5', 'BACnetDataType.OPENING_TAG')
                     openingTag                     ]
    [simple BACnetApplicationTagEnumerated errorClass]
    [simple BACnetApplicationTagEnumerated errorCode]
    [simple     BACnetClosingTag('5', 'BACnetDataType.CLOSING_TAG')
                     closingTag                     ]
]

// TODO: this is a enum so we should build a static call which maps a enum (could be solved by using only the tag header with a length validation and the enum itself)
[type BACnetResultFlags(uint 8 tagNumber)
    [simple BACnetContextTagBitString('tagNumber', 'BACnetDataType.BIT_STRING')
        rawBits
    ]
    [virtual    bit firstItem           'rawBits.payload.data[0]']
    [virtual    bit lastItem            'rawBits.payload.data[1]']
    [virtual    bit moreItems           'rawBits.payload.data[2]']
]

[type BACnetServiceAckAtomicReadFileStreamOrRecord
    [peek       BACnetTagHeader
                            peekedTagHeader
    ]
    [simple     BACnetOpeningTag('peekedTagHeader.actualTagNumber', 'BACnetDataType.OPENING_TAG')
                     openingTag
    ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber']
    [typeSwitch peekedTagNumber
        ['0x0' BACnetServiceAckAtomicReadFileStream
            [simple BACnetApplicationTagSignedInteger
                            fileStartPosition           ]
            [simple BACnetApplicationTagOctetString
                            fileData                    ]
        ]
        ['0x1' BACnetServiceAckAtomicReadFileRecord
            [simple BACnetApplicationTagSignedInteger
                            fileStartRecord             ]
            [simple BACnetApplicationTagUnsignedInteger
                            returnedRecordCount         ]
            [array  BACnetApplicationTagOctetString
                            fileRecordData
                            count
                            'returnedRecordCount.payload.actualValue'   ]
        ]
    ]
    [simple     BACnetClosingTag('peekedTagHeader.actualTagNumber', 'BACnetDataType.CLOSING_TAG')
                     closingTag
    ]
]

// TODO: this need to be completly refactored
[discriminatedType BACnetError
    [discriminator uint 8 serviceChoice]
    [typeSwitch serviceChoice
        ['0x00' BACnetErrorAcknowledgeAlarm
        ]
        ['0x03' BACnetErrorGetAlarmSummary
        ]
        ['0x02' BACnetErrorConfirmedEventNotification
        ]
        ['0x04' BACnetErrorGetEnrollmentSummary
        ]
        ['0x05' BACnetErrorDeviceCommunicationProtocol
        ]
        ['0x1D' BACnetErrorGetEventInformation
        ]
        ['0x06' BACnetErrorAtomicReadFile
        ]
        ['0x07' BACnetErrorAtomicWriteFile
        ]
        ['0x0A' BACnetErrorCreateObject
        ]
        ['0x0C' BACnetErrorReadProperty
        ]
        ['0x0E' BACnetErrorReadPropertyMultiple
        ]
        ['0x0F' BACnetErrorWriteProperty
        ]
        ['0x1A' BACnetErrorReadRange
        ]
        ['0x11' BACnetErrorDeviceCommunicationProtocol
        ]
        ['0x12' BACnetErrorConfirmedPrivateTransfer
        ]
        ['0x14' BACnetErrorPasswordFailure
        ]
        ['0x15' BACnetErrorVTOpen
        ]
        ['0x17' BACnetErrorVTData
        ]
        ['0x18' BACnetErrorRemovedAuthenticate
        ]
        ['0x0D' BACnetErrorRemovedReadPropertyConditional
        ]
        [BACnetErrorUnknown
        ]
    ]
    [simple BACnetApplicationTagEnumerated errorClass]
    [simple BACnetApplicationTagEnumerated errorCode]
]

[type BACnetNotificationParameters(uint 8 tagNumber, BACnetObjectType objectType)
    [simple     BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                            openingTag
    ]
    [peek       BACnetTagHeader
                            peekedTagHeader
    ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber']
    [typeSwitch peekedTagNumber
        ['0' BACnetNotificationParametersChangeOfBitString(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                    innerOpeningTag
            ]
            [simple BACnetContextTagBitString('0', 'BACnetDataType.BIT_STRING')
                    changeOfBitString
            ]
            [simple BACnetStatusFlags('1')
                    statusFlags
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        ['1' BACnetNotificationParametersChangeOfState(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                    innerOpeningTag
            ]
            [simple BACnetPropertyStates('0')
                    changeOfState
            ]
            [simple BACnetStatusFlags('1')
                    statusFlags
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        ['2' BACnetNotificationParametersChangeOfValue(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                    innerOpeningTag
            ]
            [simple BACnetNotificationParametersChangeOfValueNewValue('0')
                    newValue
            ]
            [simple BACnetStatusFlags('1')
                    statusFlags
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        ['3' BACnetNotificationParametersCommandFailure(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                    innerOpeningTag
            ]
            // TODO: temporary dummy property identifier... get rid of that
            [simple BACnetConstructedData('0', 'objectType', 'STATIC_CALL("dummyPropertyIdentifier")')
                    commandValue
            ]
            [simple BACnetStatusFlags('1')
                    statusFlags
            ]
            // TODO: temporary dummy property identifier... get rid of that
            [simple BACnetConstructedData('2', 'objectType', 'STATIC_CALL("dummyPropertyIdentifier")')
                    feedbackValue
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        ['4' BACnetNotificationParametersFloatingLimit(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                            innerOpeningTag
            ]
            [simple BACnetContextTagReal('0', 'BACnetDataType.REAL')
                    referenceValue
            ]
            [simple BACnetStatusFlags('1')
                    statusFlags
            ]
            [simple BACnetContextTagReal('2', 'BACnetDataType.REAL')
                    setPointValue
            ]
            [simple BACnetContextTagReal('3', 'BACnetDataType.REAL')
                    errorLimit
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        ['5' BACnetNotificationParametersOutOfRange(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                            innerOpeningTag
            ]
            [simple BACnetContextTagReal('0', 'BACnetDataType.REAL')
                    exceedingValue
            ]
            [simple BACnetStatusFlags('1')
                    statusFlags
            ]
            [simple BACnetContextTagReal('2', 'BACnetDataType.REAL')
                    deadband
            ]
            [simple BACnetContextTagReal('3', 'BACnetDataType.REAL')
                    exceededLimit
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        ['6' BACnetNotificationParametersComplexEventType(uint 8 peekedTagNumber)
            [simple     BACnetPropertyValues('peekedTagNumber', 'objectType')
                        listOfValues
            ]
        ]
        // TODO: implement other cases
        ['9' BACnetNotificationParametersExtended(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                    innerOpeningTag
            ]
            [simple BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')
                    vendorId
            ]
            [simple BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')
                    extendedEventType
            ]
            [simple BACnetNotificationParametersExtendedParameters('2')
                    parameters
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        ['10' BACnetNotificationParametersBufferReady(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                    innerOpeningTag
            ]
            [simple BACnetDeviceObjectPropertyReferenceEnclosed('0')
                    bufferProperty
            ]
            [simple BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')
                    previousNotification
            ]
            [simple BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')
                    currentNotification
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        ['11' BACnetNotificationParametersUnsignedRange(uint 8 peekedTagNumber)
            [simple BACnetOpeningTag('peekedTagNumber', 'BACnetDataType.OPENING_TAG')
                    innerOpeningTag
            ]
            [simple BACnetContextTagUnsignedInteger('0', 'BACnetDataType.UNSIGNED_INTEGER')
                    sequenceNumber
            ]
            [simple BACnetStatusFlags('1')
                    statusFlags
            ]
            [simple BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')
                    exceededLimit
            ]
            [simple BACnetClosingTag('peekedTagNumber', 'BACnetDataType.CLOSING_TAG')
                    innerClosingTag
            ]
        ]
        // TODO: implement other cases
        [BACnetNotificationParametersUnmapped
            // TODO: implement me
            [validation    '1 == 2'    "TODO: implement me"]
        ]
    ]
    [simple     BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                closingTag
    ]
]

// TODO: this could be inlined once we can support nested types
[type BACnetNotificationParametersChangeOfValueNewValue(uint 8 tagNumber)
    [simple     BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                openingTag
    ]
    [peek       BACnetTagHeader
                            peekedTagHeader
    ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber']
    [typeSwitch peekedTagNumber
        ['0' BACnetNotificationParametersChangeOfValueNewValueChangedBits(uint 8 peekedTagNumber)
            [simple BACnetContextTagBitString('0', 'BACnetDataType.BIT_STRING')
                    changedBits
            ]
        ]
        ['1' BACnetNotificationParametersChangeOfValueNewValueChangedValue(uint 8 peekedTagNumber)
            [simple BACnetContextTagReal('0', 'BACnetDataType.REAL')
                    changedValue
            ]
        ]
    ]
    [simple     BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                closingTag
    ]
]

// TODO: this could be inlined once we can support nested types
[type BACnetNotificationParametersExtendedParameters(uint 8 tagNumber)
    [simple     BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                openingTag
    ]
    [peek       BACnetTagHeader
                            peekedTagHeader
    ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber'       ]
    [virtual    bit         isOpeningTag        'peekedTagHeader.lengthValueType == 0x6']
    [virtual    bit         isClosingTag        'peekedTagHeader.lengthValueType == 0x7']
    [optional   BACnetApplicationTagNull
                nullValue
                    'peekedTagNumber == 0x0 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagReal
                realValue
                    'peekedTagNumber == 0x4 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagUnsignedInteger
                unsignedValue
                    'peekedTagNumber == 0x2 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagBoolean
                booleanValue
                    'peekedTagNumber == 0x1 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagSignedInteger
                integerValue
                    'peekedTagNumber == 0x3 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagDouble
                doubleValue
                    'peekedTagNumber == 0x5 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagOctetString
                octetStringValue
                    'peekedTagNumber == 0x6 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagCharacterString
                characterStringValue
                    'peekedTagNumber == 0x7 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagBitString
                bitStringValue
                    'peekedTagNumber == 0x8 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagEnumerated
                enumeratedValue
                    'peekedTagNumber == 0x9 && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagDate
                dateValue
                    'peekedTagNumber == 0xA && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagTime
                timeValue
                    'peekedTagNumber == 0xB && !isOpeningTag && !isClosingTag'          ]
    [optional   BACnetApplicationTagObjectIdentifier
                objectIdentifier
                    'peekedTagNumber == 0xC && !isOpeningTag'                           ]
    [optional   BACnetDeviceObjectPropertyReferenceEnclosed('0')
                reference
                    'isOpeningTag && !isClosingTag'                                     ]
    [simple     BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                closingTag
    ]
]

[type BACnetPropertyValues(uint 8 tagNumber, BACnetObjectType objectType)
    [simple BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
            innerOpeningTag
    ]
    [array  BACnetPropertyValue('objectType')
            data
            terminated
            'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, tagNumber)'
    ]
    [simple BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
            innerClosingTag
    ]
]

[type BACnetPropertyValue(BACnetObjectType objectType)
    [simple   BACnetContextTagPropertyIdentifier('0', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER') propertyIdentifier  ]
    [optional BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')              propertyArrayIndex  ]
    [optional BACnetConstructedDataElement('objectType', 'propertyIdentifier')                     propertyValue       ]
    [optional BACnetContextTagUnsignedInteger('3', 'BACnetDataType.UNSIGNED_INTEGER')              priority            ]
]

[type BACnetDeviceObjectPropertyReferenceEnclosed(uint 8 tagNumber)
    [simple   BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
              openingTag
    ]
    [simple   BACnetDeviceObjectPropertyReference     value                                                            ]
    [simple   BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
              closingTag
    ]
]

// TODO: this is a enum so we should build a static call which maps a enum (could be solved by using only the tag header with a length validation and the enum itself)
[type BACnetStatusFlags(uint 8 tagNumber)
    [simple BACnetContextTagBitString('tagNumber', 'BACnetDataType.BIT_STRING')
        rawBits
    ]
    [virtual    bit inAlarm         'rawBits.payload.data[0]']
    [virtual    bit fault           'rawBits.payload.data[1]']
    [virtual    bit overriden       'rawBits.payload.data[2]']
    [virtual    bit outOfService    'rawBits.payload.data[3]']
]

// TODO: this is a enum so we should build a static call which maps a enum (could be solved by using only the tag header with a length validation and the enum itself)
[type BACnetAction(uint 8 tagNumber)
    [optional   BACnetContextTagEnumerated('tagNumber', 'BACnetDataType.ENUMERATED')
                rawData
    ]
    [virtual    bit isDirect         'rawData != null && rawData.payload.actualValue == 0']
    [virtual    bit isReverse        'rawData != null && rawData.payload.actualValue == 1']
]

[type BACnetActionCommand
    [optional   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                    deviceIdentifier
    ]
    [simple     BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                    objectIdentifier
    ]
    [simple     BACnetContextTagPropertyIdentifier('2', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER')
                    propertyIdentifier
    ]
    [optional   BACnetContextTagUnsignedInteger('3', 'BACnetDataType.UNSIGNED_INTEGER')
                    arrayIndex
    ]
    [optional   BACnetConstructedData('4', 'objectIdentifier.objectType', 'propertyIdentifier')
                    propertyValue
    ]
    [optional     BACnetContextTagUnsignedInteger('5', 'BACnetDataType.UNSIGNED_INTEGER')
                    priority
    ]
    [optional   BACnetContextTagBoolean('6', 'BACnetDataType.BOOLEAN')
                    postDelay
    ]
    [simple   BACnetContextTagBoolean('7', 'BACnetDataType.BOOLEAN')
                    quitOnFailure
    ]
    [simple   BACnetContextTagBoolean('8', 'BACnetDataType.BOOLEAN')
                    writeSuccessful
    ]
]

// TODO: this is a enum so we should build a static call which maps a enum (could be solved by using only the tag header with a length validation and the enum itself)
[type BACnetBinaryPV(uint 8 tagNumber)
    [optional   BACnetContextTagEnumerated('tagNumber', 'BACnetDataType.ENUMERATED')
                rawData
    ]
    [virtual    bit isInactive         'rawData != null && rawData.payload.actualValue == 0']
    [virtual    bit isActive           'rawData != null && rawData.payload.actualValue == 1']
]

[type BACnetPropertyStates(uint 8 tagNumber)
    [simple     BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                openingTag
    ]
    [peek       BACnetTagHeader
                            peekedTagHeader
    ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber']
    [typeSwitch peekedTagNumber
        ['0' BACnetPropertyStatesBoolean(uint 8 peekedTagNumber)
            [optional   BACnetContextTagBoolean('peekedTagNumber', 'BACnetDataType.BOOLEAN')
                        booleanValue
            ]
        ]
        ['1' BACnetPropertyStatesBinaryValue(uint 8 peekedTagNumber)
            [optional   BACnetBinaryPV('peekedTagNumber')
                        binaryValue
            ]
        ]
        // TODO: add missing type
        ['16' BACnetPropertyStatesAction(uint 8 peekedTagNumber)
            [optional   BACnetAction('peekedTagNumber')
                        action
            ]
        ]
        // TODO: add missing type
        [BACnetPropertyStateActionUnmapped
                // TODO: implement me
                [validation    '1 == 2'    "TODO: implement me"]
        ]
    ]
    [simple     BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                closingTag
    ]
]

[type BACnetTimeStamp
    [peek       BACnetTagHeader
                            peekedTagHeader
    ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber']
    [typeSwitch peekedTagNumber
        ['0' BACnetTimeStampTime
            [simple BACnetContextTagTime('0', 'BACnetDataType.TIME')
                    timeValue
            ]
        ]
        ['1' BACnetTimeStampSequence
            [simple BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')
                    sequenceNumber
            ]
        ]
        ['2' BACnetTimeStampDateTime
            [simple BACnetDateTimeEnclosed('2')
                    dateTimeValue
            ]
        ]
    ]
]

[type BACnetTimeStampEnclosed(uint 8 tagNumber)
    [simple     BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                    openingTag          ]
    [simple     BACnetTimeStamp
                    timestamp           ]
    [simple     BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                    closingTag          ]
]

[type BACnetTimeStampsEnclosed(uint 8 tagNumber)
    [simple         BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                    openingTag
    ]
    [array  BACnetTimeStamp
                timestamps
                terminated
                'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, tagNumber)'
    ]
    [simple         BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                    closingTag
    ]
]

[type BACnetDateTime
    [simple     BACnetApplicationTagDate
                dateValue
    ]
    [simple     BACnetApplicationTagTime
                timeValue
    ]
]

[type BACnetDateTimeEnclosed(uint 8 tagNumber)
    [simple     BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                openingTag
    ]
    [simple     BACnetDateTime
                dateTimeValue
    ]
    [simple     BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                closingTag
    ]
]

[type BACnetAddress
    [array  uint 8 address count '4']
    [simple uint 16 port]
]

[discriminatedType BACnetTagHeader
    [simple        uint 4   tagNumber                                                                                   ]
    [simple        TagClass tagClass                                                                                    ]
    [simple        uint 3   lengthValueType                                                                             ]
    [optional      uint 8   extTagNumber    'tagNumber == 15'                                                           ]
    [virtual       uint 8   actualTagNumber 'tagNumber < 15 ? tagNumber : extTagNumber'                                 ]
    [virtual       bit      isBoolean       'tagNumber == 1 && tagClass == TagClass.APPLICATION_TAGS'                   ]
    [virtual       bit      isConstructed   'tagClass == TagClass.CONTEXT_SPECIFIC_TAGS && lengthValueType == 6'        ]
    [virtual       bit      isPrimitiveAndNotBoolean '!isConstructed && !isBoolean'                                     ]
    [optional      uint 8   extLength       'isPrimitiveAndNotBoolean && lengthValueType == 5'                          ]
    [optional      uint 16  extExtLength    'isPrimitiveAndNotBoolean && lengthValueType == 5 && extLength == 254'      ]
    [optional      uint 32  extExtExtLength 'isPrimitiveAndNotBoolean && lengthValueType == 5 && extLength == 255'      ]
    [virtual       uint 32  actualLength    'lengthValueType == 5 && extLength == 255 ? extExtExtLength : (lengthValueType == 5 && extLength == 254 ? extExtLength : (lengthValueType == 5 ? extLength : lengthValueType))']
]

[discriminatedType BACnetApplicationTag
    [simple        BACnetTagHeader
                            header
    ]
    [validation    'header.tagClass == TagClass.APPLICATION_TAGS'    "should be a application tag"                      ]
    [virtual       uint 8   actualTagNumber 'header.actualTagNumber'                                                    ]
    [virtual       uint 32  actualLength    'header.actualLength'                                                       ]
    [typeSwitch actualTagNumber
        ['0x0' BACnetApplicationTagNull
        ]
        ['0x1' BACnetApplicationTagBoolean(BACnetTagHeader header)
            [simple BACnetTagPayloadBoolean('header.actualLength')
                                payload                                                                                 ]
            [virtual    bit     actualValue 'payload.value'                                                             ]
        ]
        ['0x2' BACnetApplicationTagUnsignedInteger(BACnetTagHeader header)
            [simple BACnetTagPayloadUnsignedInteger('header.actualLength')
                                payload                                                                                 ]
            [virtual    uint 64 actualValue   'payload.actualValue'                                                     ]
        ]
        ['0x3' BACnetApplicationTagSignedInteger(BACnetTagHeader header)
            [simple BACnetTagPayloadSignedInteger('header.actualLength')
                                payload                                                                                 ]
            [virtual    uint 64    actualValue   'payload.actualValue'                                                  ]
        ]
        ['0x4' BACnetApplicationTagReal
            [simple BACnetTagPayloadReal
                                payload                                                                                 ]

            [virtual    float 32     actualValue 'payload.value'                                                        ]
        ]
        ['0x5' BACnetApplicationTagDouble
            [simple BACnetTagPayloadDouble
                                payload                                                                                 ]
            [virtual    float 64     actualValue 'payload.value'                                                        ]
        ]
        ['0x6' BACnetApplicationTagOctetString(BACnetTagHeader header)
            [simple BACnetTagPayloadOctetString('header.actualLength')
                                payload                                                                                 ]
            [virtual vstring     value             'payload.value'                                                      ]
        ]
        ['0x7' BACnetApplicationTagCharacterString(BACnetTagHeader header)
            [simple BACnetTagPayloadCharacterString('header.actualLength')
                                payload                                                                                 ]
            [virtual vstring     value             'payload.value'                                                      ]
        ]
        ['0x8' BACnetApplicationTagBitString(BACnetTagHeader header)
            [simple BACnetTagPayloadBitString('header.actualLength')
                                payload                                                                                 ]
        ]
        ['0x9' BACnetApplicationTagEnumerated(BACnetTagHeader header)
            [simple BACnetTagPayloadEnumerated('header.actualLength')
                                payload                                                                                 ]
            [virtual  uint 32   actualValue 'payload.actualValue'                                                       ]
        ]
        ['0xA' BACnetApplicationTagDate
            [simple BACnetTagPayloadDate
                                payload                                                                                 ]
        ]
        ['0xB' BACnetApplicationTagTime
            [simple BACnetTagPayloadTime
                                payload                                                                                 ]
        ]
        ['0xC' BACnetApplicationTagObjectIdentifier
            [simple BACnetTagPayloadObjectIdentifier
                                payload                                                                                 ]
            [virtual    BACnetObjectType
                                objectType
                                               'payload.objectType'                                                     ]
            [virtual  uint 22   instanceNumber
                                               'payload.instanceNumber'                                                 ]
        ]
    ]
]

[discriminatedType BACnetContextTag(uint 8 tagNumberArgument, BACnetDataType dataType)
    [simple        BACnetTagHeader
                            header
    ]
    [validation    'header.actualTagNumber == tagNumberArgument' "tagnumber doesn't match" shouldFail=false             ]
    [validation    'header.tagClass == TagClass.CONTEXT_SPECIFIC_TAGS' "should be a context tag"                        ]
    [virtual       uint 4   tagNumber     'header.tagNumber'                                                            ]
    [virtual       uint 32  actualLength  'header.actualLength'                                                         ]
    [virtual       bit      isNotOpeningOrClosingTag    'header.lengthValueType != 6 && header.lengthValueType != 7'    ]
    [typeSwitch dataType
        ['NULL' BACnetContextTagNull(bit isNotOpeningOrClosingTag, BACnetTagHeader header)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [validation 'header.actualLength == 0' "length field should be 0"                                           ]
        ]
        ['BOOLEAN' BACnetContextTagBoolean(bit isNotOpeningOrClosingTag, BACnetTagHeader header)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [validation 'header.actualLength == 1' "length field should be 1"                                           ]
            [simple  uint 8 value                                                                                       ]
            [simple BACnetTagPayloadBoolean('value')
                            payload                                                                                     ]
            [virtual bit    actualValue 'payload.value'                                                                 ]
        ]
        ['UNSIGNED_INTEGER' BACnetContextTagUnsignedInteger(bit isNotOpeningOrClosingTag, BACnetTagHeader header)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [simple BACnetTagPayloadUnsignedInteger('header.actualLength')
                                payload                                                                                 ]
            [virtual    uint 64 actualValue 'payload.actualValue'                                                     ]
        ]
        ['SIGNED_INTEGER' BACnetContextTagSignedInteger(bit isNotOpeningOrClosingTag, BACnetTagHeader header)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [simple BACnetTagPayloadSignedInteger('header.actualLength')
                                payload                                                                                 ]
            [virtual    uint 64     actualValue 'payload.actualValue'                                                 ]
        ]
        ['REAL' BACnetContextTagReal(bit isNotOpeningOrClosingTag)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag"]
            [simple BACnetTagPayloadReal
                                    payload                                                                             ]
            [virtual    float 32     actualValue 'payload.value'                                                        ]
        ]
        ['DOUBLE' BACnetContextTagDouble(bit isNotOpeningOrClosingTag)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag"]
            [simple BACnetTagPayloadDouble
                                payload                                                                                 ]

            [virtual    float 64     actualValue 'payload.value'                                                        ]
        ]
        ['OCTET_STRING' BACnetContextTagOctetString(bit isNotOpeningOrClosingTag, BACnetTagHeader header)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag"]
            [simple BACnetTagPayloadOctetString('header.actualLength')
                                payload                                                                                 ]
            [virtual vstring     value             'payload.value'                                                      ]
        ]
        ['CHARACTER_STRING' BACnetContextTagCharacterString(bit isNotOpeningOrClosingTag, BACnetTagHeader header)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [simple BACnetTagPayloadCharacterString('header.actualLength')
                                payload                                                                                 ]
            [virtual vstring     value             'payload.value'                                                      ]
        ]
        ['BIT_STRING' BACnetContextTagBitString(bit isNotOpeningOrClosingTag, BACnetTagHeader header)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [simple BACnetTagPayloadBitString('header.actualLength')
                                payload                                                                                 ]
        ]
        ['ENUMERATED' BACnetContextTagEnumerated(bit isNotOpeningOrClosingTag, BACnetTagHeader header)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [simple BACnetTagPayloadEnumerated('header.actualLength')
                                payload                                                                                 ]
            [virtual  uint 32   actualValue 'payload.actualValue'                                                       ]
        ]
        ['DATE' BACnetContextTagDate(bit isNotOpeningOrClosingTag)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [simple BACnetTagPayloadDate
                                payload                                                                                 ]
        ]
        ['TIME' BACnetContextTagTime(bit isNotOpeningOrClosingTag)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [simple     BACnetTagPayloadTime
                                payload                                                                                 ]
        ]
        ['BACNET_OBJECT_IDENTIFIER' BACnetContextTagObjectIdentifier(bit isNotOpeningOrClosingTag)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [simple  BACnetTagPayloadObjectIdentifier
                                payload                                                                                 ]
            [virtual BACnetObjectType
                                objectType 'payload.objectType'                                                         ]
            [virtual uint 22    instanceNumber
                                               'payload.instanceNumber'                                                 ]
        ]
        ['BACNET_PROPERTY_IDENTIFIER' BACnetContextTagPropertyIdentifier(bit isNotOpeningOrClosingTag, uint 32 actualLength)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag" shouldFail=false]
            [manual     BACnetPropertyIdentifier
                                propertyIdentifier
                                'STATIC_CALL("readPropertyIdentifier", readBuffer, actualLength)'
                                'STATIC_CALL("writePropertyIdentifier", writeBuffer, propertyIdentifier)'
                                '_value.actualLength*8'                                                                 ]
            [manual     uint 32
                                proprietaryValue
                                'STATIC_CALL("readProprietaryPropertyIdentifier", readBuffer, propertyIdentifier, actualLength)'
                                'STATIC_CALL("writeProprietaryPropertyIdentifier", writeBuffer, propertyIdentifier, proprietaryValue)'
                                '0'                                                                                     ]
            [virtual    bit
                                isProprietary
                                'propertyIdentifier == BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE'               ]
        ]
        ['EVENT_TYPE' BACnetContextTagEventType(bit isNotOpeningOrClosingTag, uint 32 actualLength)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag"]
            [manual     BACnetEventType
                                eventType
                                'STATIC_CALL("readEventType", readBuffer, actualLength)'
                                'STATIC_CALL("writeEventType", writeBuffer, eventType)'
                                '_value.actualLength*8'                                                                 ]
            [manual     uint 32
                                proprietaryValue
                                'STATIC_CALL("readProprietaryEventType", readBuffer, eventType, actualLength)'
                                'STATIC_CALL("writeProprietaryEventType", writeBuffer, eventType, proprietaryValue)'
                                '0'                                                                                     ]
            [virtual    bit
                                isProprietary
                                'eventType == BACnetEventType.VENDOR_PROPRIETARY_VALUE'                                 ]
        ]
        ['EVENT_STATE' BACnetContextTagEventState(bit isNotOpeningOrClosingTag, uint 32 actualLength)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag"]
            [manual     BACnetEventState
                                eventState
                                'STATIC_CALL("readEventState", readBuffer, actualLength)'
                                'STATIC_CALL("writeEventState", writeBuffer, eventState)'
                                '_value.actualLength*8'                                                                 ]
            [manual     uint 32
                                proprietaryValue
                                'STATIC_CALL("readProprietaryEventState", readBuffer, eventState, actualLength)'
                                'STATIC_CALL("writeProprietaryEventState", writeBuffer, eventState, proprietaryValue)'
                                '0'                                                                                     ]
            [virtual    bit
                                isProprietary
                                'eventState == BACnetEventState.VENDOR_PROPRIETARY_VALUE'                               ]
        ]
        ['NOTIFY_TYPE' BACnetContextTagNotifyType(bit isNotOpeningOrClosingTag, uint 32 actualLength)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag"]
            [simple  BACnetNotifyType
                                value                                                                                   ]
        ]
        ['BACNET_DEVICE_STATE' BACnetContextTagDeviceState(bit isNotOpeningOrClosingTag)
            [validation 'isNotOpeningOrClosingTag' "length 6 and 7 reserved for opening and closing tag"]
            [simple BACnetDeviceState
                                state                                                                                   ]
        ]
        ['OPENING_TAG' BACnetOpeningTag(uint 32 actualLength)
            [validation 'actualLength == 6' "opening tag should have a value of 6"]
        ]
        ['CLOSING_TAG' BACnetClosingTag(uint 32 actualLength)
            [validation 'actualLength == 7' "closing tag should have a value of 7"]
        ]
        [BACnetContextTagEmpty
        ]
    ]
]

[type BACnetTagPayloadBoolean(uint 32 actualLength)
    [virtual bit value   'actualLength == 1'    ]
    [virtual bit isTrue  'value'                ]
    [virtual bit isFalse '!value'               ]
]

[type BACnetTagPayloadUnsignedInteger(uint 32 actualLength)
    [virtual    bit         isUint8         'actualLength == 1'  ]
    [optional   uint  8     valueUint8      'isUint8'            ]
    [virtual    bit         isUint16        'actualLength == 2'  ]
    [optional   uint 16     valueUint16     'isUint16'           ]
    [virtual    bit         isUint24        'actualLength == 3'  ]
    [optional   uint 24     valueUint24     'isUint24'           ]
    [virtual    bit         isUint32        'actualLength == 4'  ]
    [optional   uint 32     valueUint32     'isUint32'           ]
    [virtual    bit         isUint40        'actualLength == 5'  ]
    [optional   uint 40     valueUint40     'isUint40'           ]
    [virtual    bit         isUint48        'actualLength == 6'  ]
    [optional   uint 48     valueUint48     'isUint48'           ]
    [virtual    bit         isUint56        'actualLength == 7'  ]
    [optional   uint 56     valueUint56     'isUint56'           ]
    [virtual    bit         isUint64        'actualLength == 8'  ]
    [optional   uint 64     valueUint64     'isUint64'           ]
    [validation 'isUint8 || isUint16 || isUint24 || isUint32 || isUint40 || isUint48 || isUint56 || isUint64' "unmapped integer length"]
    [virtual    uint 64     actualValue     'isUint8?valueUint8:(isUint16?valueUint16:(isUint24?valueUint24:(isUint32?valueUint32:(isUint40?valueUint40:(isUint48?valueUint48:(isUint56?valueUint56:valueUint64))))))']
]

[type BACnetTagPayloadSignedInteger(uint 32 actualLength)
    [virtual    bit         isInt8          'actualLength == 1'  ]
    [optional   int 8       valueInt8       'isInt8'             ]
    [virtual    bit         isInt16         'actualLength == 2'  ]
    [optional   int 16      valueInt16      'isInt16'            ]
    [virtual    bit         isInt24         'actualLength == 3'  ]
    [optional   int 24      valueInt24      'isInt24'            ]
    [virtual    bit         isInt32         'actualLength == 4'  ]
    [optional   int 32      valueInt32      'isInt32'            ]
    [virtual    bit         isInt40         'actualLength == 5'  ]
    [optional   int 40      valueInt40      'isInt40'            ]
    [virtual    bit         isInt48         'actualLength == 6'  ]
    [optional   int 48      valueInt48      'isInt48'            ]
    [virtual    bit         isInt56         'actualLength == 7'  ]
    [optional   int 56      valueInt56      'isInt56'            ]
    [virtual    bit         isInt64         'actualLength == 8'  ]
    [optional   int 64      valueInt64      'isInt64'            ]
    [validation 'isInt8 || isInt16 || isInt24 || isInt32 || isInt40 || isInt48 || isInt56 || isInt64' "unmapped integer length"]
    [virtual    uint 64     actualValue     'isInt8?valueInt8:(isInt16?valueInt16:(isInt24?valueInt24:(isInt32?valueInt32:(isInt40?valueInt40:(isInt48?valueInt48:(isInt56?valueInt56:valueInt64))))))']
]

[type BACnetTagPayloadReal
    [simple float 32 value]
]

[type BACnetTagPayloadDouble
    [simple float 64 value]
]

[type BACnetTagPayloadOctetString(uint 32 actualLength)
    // TODO: The reader expects int but uint32 gets mapped to long so even uint32 would easily overflow...
    [virtual    uint     16                   actualLengthInBit 'actualLength * 8']
    [simple     vstring 'actualLengthInBit'  value encoding='"ASCII"']
]

[type BACnetTagPayloadCharacterString(uint 32 actualLength)
    [simple     BACnetCharacterEncoding      encoding]
    // TODO: The reader expects int but uint32 gets mapped to long so even uint32 would easily overflow...
    [virtual    uint     16                  actualLengthInBit 'actualLength * 8 - 8']
    // TODO: call to string on encoding or add type conversion so we can use the enum above
    [simple     vstring 'actualLengthInBit'  value encoding='"UTF-8"']
]

[type BACnetTagPayloadBitString(uint 32 actualLength)
    [simple     uint 8      unusedBits                                           ]
    [array      bit         data count '((actualLength - 1) * 8) - unusedBits'   ]
    [array      bit         unused count 'unusedBits'                            ]
]

[type BACnetTagPayloadEnumerated(uint 32 actualLength)
    [array   byte       data length 'actualLength']
    [virtual uint 32    actualValue 'STATIC_CALL("parseVarUint", data)'  ]
]

[type BACnetTagPayloadDate
    [virtual uint  8 wildcard '0xFF'                                 ]
    [simple  uint  8 yearMinus1900                                   ]
    [virtual bit    yearIsWildcard 'yearMinus1900 == wildcard'      ]
    [virtual uint 16 year 'yearMinus1900 + 1900'                     ]
    [simple  uint  8 month                                           ]
    [virtual bit    monthIsWildcard 'month == wildcard'             ]
    [virtual bit    oddMonthWildcard 'month == 13'                  ]
    [virtual bit    evenMonthWildcard 'month == 14'                 ]
    [simple  uint  8 dayOfMonth                                      ]
    [virtual bit    dayOfMonthIsWildcard 'dayOfMonth == wildcard'   ]
    [virtual bit    lastDayOfMonthWildcard 'dayOfMonth == 32'       ]
    [virtual bit    oddDayOfMonthWildcard 'dayOfMonth == 33'        ]
    [virtual bit    evenDayOfMonthWildcard 'dayOfMonth == 34'       ]
    [simple  uint  8 dayOfWeek                                       ]
    [virtual bit    dayOfWeekIsWildcard 'dayOfWeek == wildcard'     ]
]

[type BACnetTagPayloadTime
    [virtual uint  8 wildcard '0xFF'                                 ]
    [simple  uint  8 hour                                            ]
    [virtual bit    hourIsWildcard 'hour == wildcard'               ]
    [simple  uint  8 minute                                          ]
    [virtual bit    minuteIsWildcard 'minute == wildcard'           ]
    [simple  uint  8 second                                          ]
    [virtual bit    secondIsWildcard 'second == wildcard'           ]
    [simple  uint  8 fractional                                      ]
    [virtual bit    fractionalIsWildcard 'fractional == wildcard'   ]
]

[type BACnetTagPayloadObjectIdentifier
    [manual     BACnetObjectType    objectType         'STATIC_CALL("readObjectType", readBuffer)' 'STATIC_CALL("writeObjectType", writeBuffer, objectType)' '10']
    [manual     uint 10             proprietaryValue   'STATIC_CALL("readProprietaryObjectType", readBuffer, objectType)' 'STATIC_CALL("writeProprietaryObjectType", writeBuffer, objectType, proprietaryValue)' '0']
    [virtual    bit                 isProprietary      'objectType == BACnetObjectType.VENDOR_PROPRIETARY_VALUE']
    [simple     uint 22             instanceNumber  ]
]

[type BACnetConstructedData(uint 8 tagNumber, BACnetObjectType objectType, BACnetContextTagPropertyIdentifier propertyIdentifierArgument)
    [simple     BACnetOpeningTag('tagNumber', 'BACnetDataType.OPENING_TAG')
                        openingTag                                                                              ]
    [virtual    BACnetPropertyIdentifier
                        propertyIdentifierEnum  'propertyIdentifierArgument.propertyIdentifier']
    [typeSwitch objectType, propertyIdentifierEnum
        ['COMMAND'                                      BACnetConstructedDataCommand
            [simple       BACnetOpeningTag('0', 'BACnetDataType.OPENING_TAG')
                                innerOpeningTag                                                                 ]
            [array  BACnetActionCommand
                        action
                    terminated
                    'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, 0)'                    ]
            [simple       BACnetClosingTag('0', 'BACnetDataType.CLOSING_TAG')
                                innerClosingTag                                                                 ]
        ]
        /////
        // LifeSafetyZone

        ['LIFE_SAFETY_ZONE', 'ZONE_MEMBERS'             BACnetConstructedDataLifeSafetyZoneMembers
            [array  BACnetDeviceObjectReference
                        members
                    terminated
                    'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, tagNumber)'            ]
        ]
        ['LIFE_SAFETY_ZONE', 'MEMBER_OF'                BACnetConstructedDataLifeSafetyZoneMemberOf
            [array  BACnetDeviceObjectReference
                        zones
                    terminated
                    'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, tagNumber)'            ]
        ]
        //
        /////

        /////
        // Generic Mapping
        //[*, 'ABSENTEE_LIMIT'                          BACnetConstructedDataAbsenteeLimit]
        //[*, 'ACCEPTED_MODES'                          BACnetConstructedDataAcceptedModes]
        //[*, 'ACCESS_ALARM_EVENTS'                     BACnetConstructedDataAccessAlarmEvents]
        //[*, 'ACCESS_DOORS'                            BACnetConstructedDataAccessDoors]
        //[*, 'ACCESS_EVENT'                            BACnetConstructedDataAccessEvent]
        //[*, 'ACCESS_EVENT_AUTHENTICATION_FACTOR'      BACnetConstructedDataAccessEventAuthenticationFactor]
        //[*, 'ACCESS_EVENT_CREDENTIAL'                 BACnetConstructedDataAccessEventCredential]
        //[*, 'ACCESS_EVENT_TAG'                        BACnetConstructedDataAccessEventTag]
        //[*, 'ACCESS_EVENT_TIME'                       BACnetConstructedDataAccessEventTime]
        //[*, 'ACCESS_TRANSACTION_EVENTS'               BACnetConstructedDataAccessTransactionEvents]
        //[*, 'ACCOMPANIMENT'                           BACnetConstructedDataAccompaniment]
        //[*, 'ACCOMPANIMENT_TIME'                      BACnetConstructedDataAccompanimentTime]
        //[*, 'ACK_REQUIRED'                            BACnetConstructedDataAckRequired]
        //[*, 'ACKED_TRANSITIONS'                       BACnetConstructedDataAckedTransitions]
        //[*, 'ACTION'                                  BACnetConstructedDataAction]
        //[*, 'ACTION_TEXT'                             BACnetConstructedDataActionText]
        //[*, 'ACTIVATION_TIME'                         BACnetConstructedDataActivationTime]
        //[*, 'ACTIVE_AUTHENTICATION_POLICY'            BACnetConstructedDataActiveAuthenticationPolicy]
        //[*, 'ACTIVE_COV_MULTIPLE_SUBSCRIPTIONS'       BACnetConstructedDataActiveCOVMultipleSubscriptions]
        //[*, 'ACTIVE_COV_SUBSCRIPTIONS'                BACnetConstructedDataActiveCOVSubscriptions]
        //[*, 'ACTIVE_TEXT'                             BACnetConstructedDataActiveText]
        //[*, 'ACTIVE_VT_SESSIONS'                      BACnetConstructedDataActiveVTSessions]
        //[*, 'ACTUAL_SHED_LEVEL'                       BACnetConstructedDataActualShedLevel]
        //[*, 'ADJUST_VALUE'                            BACnetConstructedDataAdjustValue]
        //[*, 'ALARM_VALUE'                             BACnetConstructedDataAlarmValue]
        //[*, 'ALARM_VALUES'                            BACnetConstructedDataAlarmValues]
        //[*, 'ALIGN_INTERVALS'                         BACnetConstructedDataAlignIntervals]
        //[*, 'ALL'                                     BACnetConstructedDataAll]
        //[*, 'ALL_WRITES_SUCCESSFUL'                   BACnetConstructedDataAllWritesSuccessful]
        //[*, 'ALLOW_GROUP_DELAY_INHIBIT'               BACnetConstructedDataAllowGroupDelayInhibit]
        //[*, 'APDU_LENGTH'                             BACnetConstructedDataAPDULength]
        //[*, 'APDU_SEGMENT_TIMEOUT'                    BACnetConstructedDataApduSegmentTimeout]
        //[*, 'APDU_TIMEOUT'                            BACnetConstructedDataAPDUTimeout]
        //[*, 'APPLICATION_SOFTWARE_VERSION'            BACnetConstructedDataApplicationSoftwareVersion]
        //[*, 'ARCHIVE'                                 BACnetConstructedDataArchive]
        //[*, 'ASSIGNED_ACCESS_RIGHTS'                  BACnetConstructedDataAssignedAccessRights]
        //[*, 'ASSIGNED_LANDING_CALLS'                  BACnetConstructedDataAssignedLandingCalls]
        //[*, 'ATTEMPTED_SAMPLES'                       BACnetConstructedDataAttemptedSamples]
        //[*, 'AUTHENTICATION_FACTORS'                  BACnetConstructedDataAuthenticationFactors]
        //[*, 'AUTHENTICATION_POLICY_LIST'              BACnetConstructedDataAuthenticationPolicyList]
        //[*, 'AUTHENTICATION_POLICY_NAMES'             BACnetConstructedDataAuthenticationPolicyNames]
        //[*, 'AUTHENTICATION_STATUS'                   BACnetConstructedDataAuthenticationStatus]
        //[*, 'AUTHORIZATION_EXEMPTIONS'                BACnetConstructedDataAuthorizationExemptions]
        //[*, 'AUTHORIZATION_MODE'                      BACnetConstructedDataAuthorizationMode]
        //[*, 'AUTO_SLAVE_DISCOVERY'                    BACnetConstructedDataAutoSlaveDiscovery]
        //[*, 'AVERAGE_VALUE'                           BACnetConstructedDataAverageValue]
        //[*, 'BACKUP_AND_RESTORE_STATE'                BACnetConstructedDataBackupAndRestoreState]
        //[*, 'BACKUP_FAILURE_TIMEOUT'                  BACnetConstructedDataBackupFailureTimeout]
        //[*, 'BACKUP_PREPARATION_TIME'                 BACnetConstructedDataBackupPreparationTime]
        //[*, 'BACNET_IP_GLOBAL_ADDRESS'                BACnetConstructedDataBACnetIpGlobalAddress]
        //[*, 'BACNET_IP_MODE'                          BACnetConstructedDataBACnetIpMode]
        //[*, 'BACNET_IP_MULTICAST_ADDRESS'             BACnetConstructedDataBACnetIpMulticastAddress]
        //[*, 'BACNET_IP_NAT_TRAVERSAL'                 BACnetConstructedDataBACnetIpNatTraversal]
        //[*, 'BACNET_IP_UDP_PORT'                      BACnetConstructedDataBACnetIpUdpPort]
        //[*, 'BACNET_IPV6_MODE'                        BACnetConstructedDataBACnetIpV6Mode]
        //[*, 'BACNET_IPV6_UDP_PORT'                    BACnetConstructedDataBACnetIpV6UdpPort]
        //[*, 'BACNET_IPV6_MULTICAST_ADDRESS'           BACnetConstructedDataBACnetIpV6MulticastAddress]
        //[*, 'BASE_DEVICE_SECURITY_POLICY'             BACnetConstructedDataBaseDeviceSecurityPolicy]
        //[*, 'BBMD_ACCEPT_FD_REGISTRATIONS'            BACnetConstructedDataBBMDAcceptFDRegistrations]
        //[*, 'BBMD_BROADCAST_DISTRIBUTION_TABLE'       BACnetConstructedDataBBMDBroadcastDistributionTable]
        //[*, 'BBMD_FOREIGN_DEVICE_TABLE'               BACnetConstructedDataBBMDForeignDeviceTable]
        //[*, 'BELONGS_TO'                              BACnetConstructedDataBelongsTo]
        //[*, 'BIAS'                                    BACnetConstructedDataBias]
        //[*, 'BIT_MASK'                                BACnetConstructedDataBitMask]
        //[*, 'BIT_TEXT'                                BACnetConstructedDataBitText]
        //[*, 'BLINK_WARN_ENABLE'                       BACnetConstructedDataBlinkWarnEnable]
        //[*, 'BUFFER_SIZE'                             BACnetConstructedDataBufferSize]
        //[*, 'CAR_ASSIGNED_DIRECTION'                  BACnetConstructedDataCarAssignedDirection]
        //[*, 'CAR_DOOR_COMMAND'                        BACnetConstructedDataCarDoorCommand]
        //[*, 'CAR_DOOR_STATUS'                         BACnetConstructedDataCarDoorStatus]
        //[*, 'CAR_DOOR_TEXT'                           BACnetConstructedDataCarDoorText]
        //[*, 'CAR_DOOR_ZONE'                           BACnetConstructedDataCarDoorZone]
        //[*, 'CAR_DRIVE_STATUS'                        BACnetConstructedDataCarDriveStatus]
        //[*, 'CAR_LOAD'                                BACnetConstructedDataCarLoad]
        //[*, 'CAR_LOAD_UNITS'                          BACnetConstructedDataCarLoadUnits]
        //[*, 'CAR_MODE'                                BACnetConstructedDataCarMode]
        //[*, 'CAR_MOVING_DIRECTION'                    BACnetConstructedDataCarMovingDirection]
        //[*, 'CAR_POSITION'                            BACnetConstructedDataCarPosition]
        //[*, 'CHANGE_OF_STATE_COUNT'                   BACnetConstructedDataChangeOfStateCount]
        //[*, 'CHANGE_OF_STATE_TIME'                    BACnetConstructedDataChangeOfStateTime]
        //[*, 'CHANGES_PENDING'                         BACnetConstructedDataChangesPending]
        //[*, 'CHANNEL_NUMBER'                          BACnetConstructedDataChannelNumber]
        //[*, 'CLIENT_COV_INCREMENT'                    BACnetConstructedDataClientCovIncrement]
        //[*, 'COMMAND'                                 BACnetConstructedDataCommand]
        //[*, 'COMMAND_TIME_ARRAY'                      BACnetConstructedDataCommandTimeArray]
        //[*, 'CONFIGURATION_FILES'                     BACnetConstructedDataConfigurationFiles]
        //[*, 'CONTROL_GROUPS'                          BACnetConstructedDataControlGroups]
        //[*, 'CONTROLLED_VARIABLE_REFERENCE'           BACnetConstructedDataControlledVariableReference]
        //[*, 'CONTROLLED_VARIABLE_UNITS'               BACnetConstructedDataControlledVariableUnits]
        //[*, 'CONTROLLED_VARIABLE_VALUE'               BACnetConstructedDataControlledVariableValue]
        //[*, 'COUNT'                                   BACnetConstructedDataCount]
        //[*, 'COUNT_BEFORE_CHANGE'                     BACnetConstructedDataCountBeforeChange]
        //[*, 'COUNT_CHANGE_TIME'                       BACnetConstructedDataCountChangeTime]
        //[*, 'COV_INCREMENT'                           BACnetConstructedDataCOVIncrement]
        //[*, 'COV_PERIOD'                              BACnetConstructedDataCOVPeriod]
        //[*, 'COV_RESUBSCRIPTION_INTERVAL'             BACnetConstructedDataCOVResubscriptionInterval]
        //[*, 'COVU_PERIOD'                             BACnetConstructedDataCOVUPeriod]
        //[*, 'COVU_RECIPIENTS'                         BACnetConstructedDataCOVURecipients]
        //[*, 'CREDENTIAL_DISABLE'                      BACnetConstructedDataCredentialDisable]
        //[*, 'CREDENTIAL_STATUS'                       BACnetConstructedDataCredentialStatus]
        //[*, 'CREDENTIALS'                             BACnetConstructedDataCredentials]
        //[*, 'CREDENTIALS_IN_ZONE'                     BACnetConstructedDataCredentialsInZone]
        //[*, 'CURRENT_COMMAND_PRIORITY'                BACnetConstructedDataCurrentCommandPriority]
        //[*, 'DATABASE_REVISION'                       BACnetConstructedDataDatabaseRevision]
        //[*, 'DATE_LIST'                               BACnetConstructedDataDateList]
        //[*, 'DAYLIGHT_SAVINGS_STATUS'                 BACnetConstructedDataDaylightSavingsStatus]
        //[*, 'DAYS_REMAINING'                          BACnetConstructedDataDaysRemaining]
        //[*, 'DEADBAND'                                BACnetConstructedDataDeadband]
        //[*, 'DEFAULT_FADE_TIME'                       BACnetConstructedDataDefaultFadeTime]
        //[*, 'DEFAULT_RAMP_RATE'                       BACnetConstructedDataDefaultRampRate]
        //[*, 'DEFAULT_STEP_INCREMENT'                  BACnetConstructedDataDefaultStepIncrement]
        //[*, 'DEFAULT_SUBORDINATE_RELATIONSHIP'        BACnetConstructedDataDefaultSubordinateRelationship]
        //[*, 'DEFAULT_TIMEOUT'                         BACnetConstructedDataDefaultTimeout]
        //[*, 'DEPLOYED_PROFILE_LOCATION'               BACnetConstructedDataDeployedProfileLocation]
        //[*, 'DERIVATIVE_CONSTANT'                     BACnetConstructedDataDerivativeConstant]
        //[*, 'DERIVATIVE_CONSTANT_UNITS'               BACnetConstructedDataDerivativeConstantUnits]
        //[*, 'DESCRIPTION'                             BACnetConstructedDataDescription]
        //[*, 'DESCRIPTION_OF_HALT'                     BACnetConstructedDataDescriptionOfHalt]
        //[*, 'DEVICE_ADDRESS_BINDING'                  BACnetConstructedDataDeviceAddressBinding]
        //[*, 'DEVICE_TYPE'                             BACnetConstructedDataDeviceType]
        //[*, 'DIRECT_READING'                          BACnetConstructedDataDirectReading]
        //[*, 'DISTRIBUTION_KEY_REVISION'               BACnetConstructedDataDistributionKeyRevision]
        //[*, 'DO_NOT_HIDE'                             BACnetConstructedDataDoNotHide]
        //[*, 'DOOR_ALARM_STATE'                        BACnetConstructedDataDoorAlarmState]
        //[*, 'DOOR_EXTENDED_PULSE_TIME'                BACnetConstructedDataDoorExtendedPulseTime]
        //[*, 'DOOR_MEMBERS'                            BACnetConstructedDataDoorMembers]
        //[*, 'DOOR_OPEN_TOO_LONG_TIME'                 BACnetConstructedDataDoorOpenTooLongTime]
        //[*, 'DOOR_PULSE_TIME'                         BACnetConstructedDataDoorPuleTime]
        //[*, 'DOOR_STATUS'                             BACnetConstructedDataDoorStatus]
        //[*, 'DOOR_UNLOCK_DELAY_TIME'                  BACnetConstructedDataDoorUnlockDelayTime]
        //[*, 'DUTY_WINDOW'                             BACnetConstructedDataDutyWindow]
        //[*, 'EFFECTIVE_PERIOD'                        BACnetConstructedDataEffectivePeriod]
        //[*, 'EGRESS_ACTIVE'                           BACnetConstructedDataEgressActive]
        //[*, 'EGRESS_TIME'                             BACnetConstructedDataEgressTime]
        //[*, 'ELAPSED_ACTIVE_TIME'                     BACnetConstructedDataElapsedActiveTime]
        //[*, 'ELEVATOR_GROUP'                          BACnetConstructedDataElevatorGroup]
        //[*, 'ENABLE'                                  BACnetConstructedDataEnable]
        //[*, 'ENERGY_METER'                            BACnetConstructedDataEnergyMeter]
        //[*, 'ENERGY_METER_REF'                        BACnetConstructedDataEnergyMeterRef]
        //[*, 'ENTRY_POINTS'                            BACnetConstructedDataEntryPoints]
        //[*, 'ERROR_LIMIT'                             BACnetConstructedDataErrorLimit]
        //[*, 'ESCALATOR_MODE'                          BACnetConstructedDataEscalatorMode]
        //[*, 'EVENT_ALGORITHM_INHIBIT'                 BACnetConstructedDataEventAlgorithmInhibit]
        //[*, 'EVENT_ALGORITHM_INHIBIT_REF'             BACnetConstructedDataEventAlgorithmInhibitRef]
        //[*, 'EVENT_DETECTION_ENABLE'                  BACnetConstructedDataEventDetectionEnable]
        //[*, 'EVENT_ENABLE'                            BACnetConstructedDataEventEnable]
        //[*, 'EVENT_MESSAGE_TEXTS'                     BACnetConstructedDataEventMessageTexts]
        //[*, 'EVENT_MESSAGE_TEXTS_CONFIG'              BACnetConstructedDataEventMessageTextsConfig]
        //[*, 'EVENT_PARAMETERS'                        BACnetConstructedDataEventParameters]
        //[*, 'EVENT_STATE'                             BACnetConstructedDataEventState]
        [*, 'EVENT_TIME_STAMPS'                         BACnetConstructedDataEventTimestamps
            [simple BACnetContextTagTime('0', 'BACnetDataType.TIME')
                    toOffnormal                                                                                 ]
            [simple BACnetContextTagUnsignedInteger('1', 'BACnetDataType.UNSIGNED_INTEGER')
                    toFault                                                                                     ]
            [simple BACnetDateTimeEnclosed('2')
                    toNormal                                                                                    ]
        ]
        //[*, 'EVENT_TYPE'                              BACnetConstructedDataEventType]
        //[*, 'EXCEPTION_SCHEDULE'                      BACnetConstructedDataExceptionSchedule]
        //[*, 'EXECUTION_DELAY'                         BACnetConstructedDataExecutionDelay]
        //[*, 'EXIT_POINTS'                             BACnetConstructedDataExitPoints]
        //[*, 'EXPECTED_SHED_LEVEL'                     BACnetConstructedDataExpectedShedLevel]
        //[*, 'EXPIRATION_TIME'                         BACnetConstructedDataExpirationTime]
        //[*, 'EXTENDED_TIME_ENABLE'                    BACnetConstructedDataExtendedTimeEnable]
        //[*, 'FAILED_ATTEMPT_EVENTS'                   BACnetConstructedDataFailedAttemptEvents]
        //[*, 'FAILED_ATTEMPTS'                         BACnetConstructedDataFailedAttempts]
        //[*, 'FAILED_ATTEMPTS_TIME'                    BACnetConstructedDataFailedAttemptsTime]
        //[*, 'FAULT_HIGH_LIMIT'                        BACnetConstructedDataFaultHighLimit]
        //[*, 'FAULT_LOW_LIMIT'                         BACnetConstructedDataFaultLowLimit]
        //[*, 'FAULT_PARAMETERS'                        BACnetConstructedDataFaultParameters]
        //[*, 'FAULT_SIGNALS'                           BACnetConstructedDataFaultSignals]
        //[*, 'FAULT_TYPE'                              BACnetConstructedDataFaultType]
        //[*, 'FAULT_VALUES'                            BACnetConstructedDataFaultValues]
        //[*, 'FD_BBMD_ADDRESS'                         BACnetConstructedDataFdBbmdAddress]
        //[*, 'FD_SUBSCRIPTION_LIFETIME'                BACnetConstructedDataFdSubscriptionLifetime]
        //[*, 'FEEDBACK_VALUE'                          BACnetConstructedDataFeedbackValue]
        //[*, 'FILE_ACCESS_METHOD'                      BACnetConstructedDataFileAccessMethod]
        //[*, 'FILE_SIZE'                               BACnetConstructedDataFileSize]
        //[*, 'FILE_TYPE'                               BACnetConstructedDataFileType]
        //[*, 'FIRMWARE_REVISION'                       BACnetConstructedDataFirmwareRevision]
        //[*, 'FLOOR_TEXT'                              BACnetConstructedDataFloorText]
        //[*, 'FULL_DUTY_BASELINE'                      BACnetConstructedDataFullDutyBaseline]
        //[*, 'GLOBAL_IDENTIFIER'                       BACnetConstructedDataGlobalIdentifier]
        //[*, 'GROUP_ID'                                BACnetConstructedDataGroupId]
        //[*, 'GROUP_MEMBER_NAMES'                      BACnetConstructedDataGroupMemberNames]
        //[*, 'GROUP_MEMBERS'                           BACnetConstructedDataGroupMembers]
        //[*, 'GROUP_MODE'                              BACnetConstructedDataGroupMode]
        //[*, 'HIGH_LIMIT'                              BACnetConstructedDataHighLimit]
        //[*, 'HIGHER_DECK'                             BACnetConstructedDataHigherDeck]
        //[*, 'IN_PROCESS'                              BACnetConstructedDataInProcess]
        //[*, 'IN_PROGRESS'                             BACnetConstructedDataInProgress]
        //[*, 'INACTIVE_TEXT'                           BACnetConstructedDataInactiveText]
        //[*, 'INITIAL_TIMEOUT'                         BACnetConstructedDataInitialTimeout]
        //[*, 'INPUT_REFERENCE'                         BACnetConstructedDataInputReference]
        //[*, 'INSTALLATION_ID'                         BACnetConstructedDataInstallationId]
        //[*, 'INSTANCE_OF'                             BACnetConstructedDataInstanceOf]
        //[*, 'INSTANTANEOUS_POWER'                     BACnetConstructedDataInstantaneousPower]
        //[*, 'INTEGRAL_CONSTANT'                       BACnetConstructedDataIntegralConstant]
        //[*, 'INTEGRAL_CONSTANT_UNITS'                 BACnetConstructedDataIntegralConstantUnits]
        //[*, 'INTERFACE_VALUE'                         BACnetConstructedDataInterfaceValue]
        //[*, 'INTERVAL_OFFSET'                         BACnetConstructedDataIntervalOffset]
        //[*, 'IP_ADDRESS'                              BACnetConstructedDataIpAddress]
        //[*, 'IP_DEFAULT_GATEWAY'                      BACnetConstructedDataIpDefaultGateway]
        //[*, 'IP_DHCP_ENABLE'                          BACnetConstructedDataIpDhcpEnable]
        //[*, 'IP_DHCP_LEASE_TIME'                      BACnetConstructedDataIpDhcpLeaseTime]
        //[*, 'IP_DHCP_LEASE_TIME_REMAINING'            BACnetConstructedDataIpDhcpLeaseTimeRemaining]
        //[*, 'IP_DHCP_SERVER'                          BACnetConstructedDataIpDhcpServer]
        //[*, 'IP_DNS_SERVER'                           BACnetConstructedDataIpDnsServer]
        //[*, 'IP_SUBNET_MASK'                          BACnetConstructedDataIpSubnetMask]
        //[*, 'IPV6_ADDRESS'                            BACnetConstructedDataIpv6Address]
        //[*, 'IPV6_AUTO_ADDRESSING_ENABLE'             BACnetConstructedDataIpv6AutoAddressingEnable]
        //[*, 'IPV6_DEFAULT_GATEWAY'                    BACnetConstructedDataIpv6DefaultGateway]
        //[*, 'IPV6_DHCP_LEASE_TIME'                    BACnetConstructedDataIpv6DhcpLeaseTime]
        //[*, 'IPV6_DHCP_LEASE_TIME_REMAINING'          BACnetConstructedDataIpv6DhcpLeaseTimeRemaining]
        //[*, 'IPV6_DHCP_SERVER'                        BACnetConstructedDataIpv6DhcpServer]
        //[*, 'IPV6_DNS_SERVER'                         BACnetConstructedDataIpv6DnsServer]
        //[*, 'IPV6_PREFIX_LENGTH'                      BACnetConstructedDataIpv6PrefixLength]
        //[*, 'IPV6_ZONE_INDEX'                         BACnetConstructedDataIpv6ZoneIndex]
        //[*, 'IS_UTC'                                  BACnetConstructedDataIsUtc]
        //[*, 'KEY_SETS'                                BACnetConstructedDataKeySets]
        //[*, 'LANDING_CALL_CONTROL'                    BACnetConstructedDataLandingCallControl]
        //[*, 'LANDING_CALLS'                           BACnetConstructedDataLandingCalls]
        //[*, 'LANDING_DOOR_STATUS'                     BACnetConstructedDataLandingDoorStatus]
        //[*, 'LAST_ACCESS_EVENT'                       BACnetConstructedDataLastAccessEvent]
        //[*, 'LAST_ACCESS_POINT'                       BACnetConstructedDataLastAccessPoint]
        //[*, 'LAST_COMMAND_TIME'                       BACnetConstructedDataLastCommandTime]
        //[*, 'LAST_CREDENTIAL_ADDED'                   BACnetConstructedDataLastCredentialAdded]
        //[*, 'LAST_CREDENTIAL_ADDED_TIME'              BACnetConstructedDataLastCredentialAddedTime]
        //[*, 'LAST_CREDENTIAL_REMOVED'                 BACnetConstructedDataLastCredentialRemoved]
        //[*, 'LAST_CREDENTIAL_REMOVED_TIME'            BACnetConstructedDataLastCredentialRemovedTime]
        //[*, 'LAST_KEY_SERVER'                         BACnetConstructedDataLastKeyServer]
        //[*, 'LAST_NOTIFY_RECORD'                      BACnetConstructedDataLastNotifyRecord]
        //[*, 'LAST_PRIORITY'                           BACnetConstructedDataLastPriority]
        //[*, 'LAST_RESTART_REASON'                     BACnetConstructedDataLastRestartReason]
        //[*, 'LAST_RESTORE_TIME'                       BACnetConstructedDataLastRestoreTime]
        //[*, 'LAST_STATE_CHANGE'                       BACnetConstructedDataLastStateChange]
        //[*, 'LAST_USE_TIME'                           BACnetConstructedDataLastUseTime]
        //[*, 'LIFE_SAFETY_ALARM_VALUES'                BACnetConstructedDataLifeSafetyAlarmValues]
        //[*, 'LIGHTING_COMMAND'                        BACnetConstructedDataLightingCommand]
        //[*, 'LIGHTING_COMMAND_DEFAULT_PRIORITY'       BACnetConstructedDataLightingCommandDefaultPriority]
        //[*, 'LIMIT_ENABLE'                            BACnetConstructedDataLimitEnable]
        //[*, 'LIMIT_MONITORING_INTERVAL'               BACnetConstructedDataLimitMonitoringInterval]
        //[*, 'LINK_SPEED'                              BACnetConstructedDataLinkSpeed]
        //[*, 'LINK_SPEED_AUTONEGOTIATE'                BACnetConstructedDataLinkSpeedAutonegotiate]
        //[*, 'LINK_SPEEDS'                             BACnetConstructedDataLinkSpeeds]
        //[*, 'LIST_OF_GROUP_MEMBERS'                   BACnetConstructedDataListOfGroupMembers]
        [*, 'LIST_OF_OBJECT_PROPERTY_REFERENCES'        BACnetConstructedDataListOfObjectPropertyReferences
            [array    BACnetDeviceObjectPropertyReference
                            references              terminated
                                'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, tagNumber)']
        ]
        //[*, 'LOCAL_DATE'                              BACnetConstructedDataLocalDate]
        //[*, 'LOCAL_FORWARDING_ONLY'                   BACnetConstructedDataLocalForwardingOnly]
        //[*, 'LOCAL_TIME'                              BACnetConstructedDataLocalTime]
        //[*, 'LOCATION'                                BACnetConstructedDataLocation]
        //[*, 'LOCK_STATUS'                             BACnetConstructedDataLockStatus]
        //[*, 'LOCKOUT'                                 BACnetConstructedDataLockout]
        //[*, 'LOCKOUT_RELINQUISH_TIME'                 BACnetConstructedDataLockoutRelinquishTime]
        //[*, 'LOG_BUFFER'                              BACnetConstructedDataLogBuffer]
        //[*, 'LOG_DEVICE_OBJECT_PROPERTY'              BACnetConstructedDataLogDeviceObjectProperty]
        //[*, 'LOG_INTERVAL'                            BACnetConstructedDataLogInterval]
        //[*, 'LOGGING_OBJECT'                          BACnetConstructedDataLoggingObject]
        //[*, 'LOGGING_RECORD'                          BACnetConstructedDataLoggingRecord]
        //[*, 'LOGGING_TYPE'                            BACnetConstructedDataLoggingType]
        //[*, 'LOW_DIFF_LIMIT'                          BACnetConstructedDataLowDiffLimit]
        //[*, 'LOW_LIMIT'                               BACnetConstructedDataLowLimit]
        //[*, 'LOWER_DECK'                              BACnetConstructedDataLowerDeck]
        //[*, 'MAC_ADDRESS'                             BACnetConstructedDataMacAddress]
        //[*, 'MACHINE_ROOM_ID'                         BACnetConstructedDataMachineRoomId]
        //[*, 'MAINTENANCE_REQUIRED'                    BACnetConstructedDataMaintenanceRequired]
        //[*, 'MAKING_CAR_CALL'                         BACnetConstructedDataMakingCarCall]
        //[*, 'MANIPULATED_VARIABLE_REFERENCE'          BACnetConstructedDataManipulatedVariableReference]
        //[*, 'MANUAL_SLAVE_ADDRESS_BINDING'            BACnetConstructedDataManualSlaveAddressBinding]
        //[*, 'MASKED_ALARM_VALUES'                     BACnetConstructedDataMaskedAlarmValues]
        //[*, 'MAX_ACTUAL_VALUE'                        BACnetConstructedDataMaxActualValue]
        //[*, 'MAX_APDU_LENGTH_ACCEPTED'                BACnetConstructedDataMaxApduLengthAccepted]
        //[*, 'MAX_FAILED_ATTEMPTS'                     BACnetConstructedDataMaxFailedAttempts]
        //[*, 'MAX_INFO_FRAMES'                         BACnetConstructedDataMaxInfoFrames]
        //[*, 'MAX_MASTER'                              BACnetConstructedDataMaxMaster]
        //[*, 'MAX_PRES_VALUE'                          BACnetConstructedDataMaxPresValue]
        //[*, 'MAX_SEGMENTS_ACCEPTED'                   BACnetConstructedDataMaxSegmentsAccepted]
        //[*, 'MAXIMUM_OUTPUT'                          BACnetConstructedDataMaximumOutput]
        //[*, 'MAXIMUM_VALUE'                           BACnetConstructedDataMaximumValue]
        //[*, 'MAXIMUM_VALUE_TIMESTAMP'                 BACnetConstructedDataMaximumValueTimestamp]
        [*, 'MEMBER_OF' BACnetConstructedDataMemberOf
            [array  BACnetDeviceObjectReference
                    zones
                            terminated
                            'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, tagNumber)'            ]
        ]
        //[*, 'MEMBER_STATUS_FLAGS'                     BACnetConstructedDataMemberStatusFlags]
        //[*, 'MEMBERS'                                 BACnetConstructedDataMembers]
        //[*, 'MIN_ACTUAL_VALUE'                        BACnetConstructedDataMinActualValue]
        //[*, 'MIN_PRES_VALUE'                          BACnetConstructedDataMinPresValue]
        //[*, 'MINIMUM_OFF_TIME'                        BACnetConstructedDataMinimumOffTime]
        //[*, 'MINIMUM_ON_TIME'                         BACnetConstructedDataMinimumOnTime]
        //[*, 'MINIMUM_OUTPUT'                          BACnetConstructedDataMinimumOutput]
        //[*, 'MINIMUM_VALUE'                           BACnetConstructedDataMinimumValue]
        //[*, 'MINIMUM_VALUE_TIMESTAMP'                 BACnetConstructedDataMinimumValueTimestamp]
        //[*, 'MODE'                                    BACnetConstructedDataMode]
        //[*, 'MODEL_NAME'                              BACnetConstructedDataModelName]
        //[*, 'MODIFICATION_DATE'                       BACnetConstructedDataModificationDate]
        //[*, 'MUSTER_POINT'                            BACnetConstructedDataMusterPoint]
        //[*, 'NEGATIVE_ACCESS_RULES'                   BACnetConstructedDataNegativeAccessRules]
        //[*, 'NETWORK_ACCESS_SECURITY_POLICIES'        BACnetConstructedDataNetworkAccessSecurityPolicies]
        //[*, 'NETWORK_INTERFACE_NAME'                  BACnetConstructedDataNetworkInterfaceName]
        //[*, 'NETWORK_NUMBER'                          BACnetConstructedDataNetworkNumber]
        //[*, 'NETWORK_NUMBER_QUALITY'                  BACnetConstructedDataNetworkNumberQuality]
        //[*, 'NETWORK_TYPE'                            BACnetConstructedDataNetworkType]
        //[*, 'NEXT_STOPPING_FLOOR'                     BACnetConstructedDataNextStoppingFloor]
        //[*, 'NODE_SUBTYPE'                            BACnetConstructedDataNodeSubtype]
        //[*, 'NODE_TYPE'                               BACnetConstructedDataNodeType]
        //[*, 'NOTIFICATION_CLASS'                      BACnetConstructedDataNotificationClass]
        //[*, 'NOTIFICATION_THRESHOLD'                  BACnetConstructedDataNotificationThreshold]
        //[*, 'NOTIFY_TYPE'                             BACnetConstructedDataNotifyType]
        //[*, 'NUMBER_OF_APDU_RETRIES'                  BACnetConstructedDataNumberOfApduRetries]
        //[*, 'NUMBER_OF_AUTHENTICATION_POLICIES'       BACnetConstructedDataNumberOfAuthenticationPolicies]
        //[*, 'NUMBER_OF_STATES'                        BACnetConstructedDataNumberOfStates]
        //[*, 'OBJECT_IDENTIFIER'                       BACnetConstructedDataObjectIdentifier]
        //[*, 'OBJECT_LIST'                             BACnetConstructedDataObjectList]
        //[*, 'OBJECT_NAME'                             BACnetConstructedDataObjectName]
        //[*, 'OBJECT_PROPERTY_REFERENCE'               BACnetConstructedDataObjectPropertyReference]
        //[*, 'OBJECT_TYPE'                             BACnetConstructedDataObjectType]
        //[*, 'OCCUPANCY_COUNT'                         BACnetConstructedDataOccupancyCount]
        //[*, 'OCCUPANCY_COUNT_ADJUST'                  BACnetConstructedDataOccupancyCountAdjust]
        //[*, 'OCCUPANCY_COUNT_ENABLE'                  BACnetConstructedDataOccupancyCountEnable]
        //[*, 'OCCUPANCY_LOWER_LIMIT'                   BACnetConstructedDataOccupancyLowerLimit]
        //[*, 'OCCUPANCY_LOWER_LIMIT_ENFORCED'          BACnetConstructedDataOccupancyLowerLimitEnforced]
        //[*, 'OCCUPANCY_STATE'                         BACnetConstructedDataOccupancyState]
        //[*, 'OCCUPANCY_UPPER_LIMIT'                   BACnetConstructedDataOccupancyUpperLimit]
        //[*, 'OCCUPANCY_UPPER_LIMIT_ENFORCED'          BACnetConstructedDataOccupancyUpperLimitEnforced]
        //[*, 'OPERATION_DIRECTION'                     BACnetConstructedDataOperationDirection]
        //[*, 'OPERATION_EXPECTED'                      BACnetConstructedDataOperationExpected]
        //[*, 'OPTIONAL'                                BACnetConstructedDataOptional]
        //[*, 'OUT_OF_SERVICE'                          BACnetConstructedDataOutOfService]
        //[*, 'OUTPUT_UNITS'                            BACnetConstructedDataOutputUnits]
        //[*, 'PACKET_REORDER_TIME'                     BACnetConstructedDataPacketReorderTime]
        //[*, 'PASSBACK_MODE'                           BACnetConstructedDataPassbackMode]
        //[*, 'PASSBACK_TIMEOUT'                        BACnetConstructedDataPassbackTimeout]
        //[*, 'PASSENGER_ALARM'                         BACnetConstructedDataPassengerAlarm]
        //[*, 'POLARITY'                                BACnetConstructedDataPolarity]
        //[*, 'PORT_FILTER'                             BACnetConstructedDataPortFilter]
        //[*, 'POSITIVE_ACCESS_RULES'                   BACnetConstructedDataPositiveAccessRules]
        //[*, 'POWER'                                   BACnetConstructedDataPower]
        //[*, 'POWER_MODE'                              BACnetConstructedDataPowerMode]
        //[*, 'PRESCALE'                                BACnetConstructedDataPrescale]
        //[*, 'PRESENT_VALUE'                           BACnetConstructedDataPresentValue]
        //[*, 'PRIORITY'                                BACnetConstructedDataPriority]
        //[*, 'PRIORITY_ARRAY'                          BACnetConstructedDataPriorityArray]
        //[*, 'PRIORITY_FOR_WRITING'                    BACnetConstructedDataPriorityForWriting]
        //[*, 'PROCESS_IDENTIFIER'                      BACnetConstructedDataProcessIdentifier]
        //[*, 'PROCESS_IDENTIFIER_FILTER'               BACnetConstructedDataProcessIdentifierFilter]
        //[*, 'PROFILE_LOCATION'                        BACnetConstructedDataProfileLocation]
        //[*, 'PROFILE_NAME'                            BACnetConstructedDataProfileName]
        //[*, 'PROGRAM_CHANGE'                          BACnetConstructedDataProgramChange]
        //[*, 'PROGRAM_LOCATION'                        BACnetConstructedDataProgramLocation]
        //[*, 'PROGRAM_STATE'                           BACnetConstructedDataProgramState]
        //[*, 'PROPERTY_LIST'                           BACnetConstructedDataPropertyList]
        //[*, 'PROPORTIONAL_CONSTANT'                   BACnetConstructedDataProportionalConstant]
        //[*, 'PROPORTIONAL_CONSTANT_UNITS'             BACnetConstructedDataProportionalConstantUnits]
        //[*, 'PROTOCOL_LEVEL'                          BACnetConstructedDataProtocolLevel]
        //[*, 'PROTOCOL_CONFORMANCE_CLASS'              BACnetConstructedDataProtocolConformanceClass]
        //[*, 'PROTOCOL_OBJECT_TYPES_SUPPORTED'         BACnetConstructedDataProtocolObjectTypesSupported]
        //[*, 'PROTOCOL_REVISION'                       BACnetConstructedDataProtocolRevision]
        //[*, 'PROTOCOL_SERVICES_SUPPORTED'             BACnetConstructedDataProtocolServicesSupported]
        //[*, 'PROTOCOL_VERSION'                        BACnetConstructedDataProtocolVersion]
        //[*, 'PULSE_RATE'                              BACnetConstructedDataPulseRate]
        //[*, 'READ_ONLY'                               BACnetConstructedDataReadOnly]
        //[*, 'REASON_FOR_DISABLE'                      BACnetConstructedDataReasonForDisable]
        //[*, 'REASON_FOR_HALT'                         BACnetConstructedDataReasonForHalt]
        //[*, 'RECIPIENT_LIST'                          BACnetConstructedDataRecipientList]
        //[*, 'RECORD_COUNT'                            BACnetConstructedDataRecordCount]
        //[*, 'RECORDS_SINCE_NOTIFICATION'              BACnetConstructedDataRecordsSinceNotification]
        //[*, 'REFERENCE_PORT'                          BACnetConstructedDataReferencePort]
        //[*, 'REGISTERED_CAR_CALL'                     BACnetConstructedDataRegisteredCarCall]
        //[*, 'RELIABILITY'                             BACnetConstructedDataReliability]
        //[*, 'RELIABILITY_EVALUATION_INHIBIT'          BACnetConstructedDataReliabilityEvaluationInhibit]
        //[*, 'RELINQUISH_DEFAULT'                      BACnetConstructedDataRelinquishDefault]
        //[*, 'REPRESENTS'                              BACnetConstructedDataRepresents]
        //[*, 'REQUESTED_SHED_LEVEL'                    BACnetConstructedDataRequestedShedLevel]
        //[*, 'REQUESTED_UPDATE_INTERVAL'               BACnetConstructedDataRequestedUpdateInterval]
        //[*, 'REQUIRED'                                BACnetConstructedDataRequired]
        //[*, 'RESOLUTION'                              BACnetConstructedDataResolution]
        //[*, 'RESTART_NOTIFICATION_RECIPIENTS'         BACnetConstructedDataRestartNotificationRecipients]
        //[*, 'RESTORE_COMPLETION_TIME'                 BACnetConstructedDataRestoreCompletionTime]
        //[*, 'RESTORE_PREPARATION_TIME'                BACnetConstructedDataRestorePreparationTime]
        //[*, 'ROUTING_TABLE'                           BACnetConstructedDataRoutingTable]
        //[*, 'SCALE'                                   BACnetConstructedDataScale]
        //[*, 'SCALE_FACTOR'                            BACnetConstructedDataScaleFactor]
        //[*, 'SCHEDULE_DEFAULT'                        BACnetConstructedDataScheduleDefault]
        //[*, 'SECURED_STATUS'                          BACnetConstructedDataSecuredStatus]
        //[*, 'SECURITY_PDU_TIMEOUT'                    BACnetConstructedDataSecurityPduTimeout]
        //[*, 'SECURITY_TIME_WINDOW'                    BACnetConstructedDataSecurityTimeWindow]
        //[*, 'SEGMENTATION_SUPPORTED'                  BACnetConstructedDataSegmentationSupported]
        //[*, 'SERIAL_NUMBER'                           BACnetConstructedDataSerialNumber]
        //[*, 'SETPOINT'                                BACnetConstructedDataSetpoint]
        //[*, 'SETPOINT_REFERENCE'                      BACnetConstructedDataSetpointReference]
        //[*, 'SETTING'                                 BACnetConstructedDataSetting]
        //[*, 'SHED_DURATION'                           BACnetConstructedDataShedDuration]
        //[*, 'SHED_LEVEL_DESCRIPTIONS'                 BACnetConstructedDataShedLevelDescriptions]
        //[*, 'SHED_LEVELS'                             BACnetConstructedDataShedLevels]
        //[*, 'SILENCED'                                BACnetConstructedDataSilenced]
        //[*, 'SLAVE_ADDRESS_BINDING'                   BACnetConstructedDataSlaveAddressBinding]
        //[*, 'SLAVE_PROXY_ENABLE'                      BACnetConstructedDataSlaveProxyEnable]
        //[*, 'START_TIME'                              BACnetConstructedDataStartTime]
        //[*, 'STATE_CHANGE_VALUES'                     BACnetConstructedDataStateChangeValues]
        //[*, 'STATE_DESCRIPTION'                       BACnetConstructedDataStateDescription]
        //[*, 'STATE_TEXT'                              BACnetConstructedDataStateText]
        //[*, 'STATUS_FLAGS'                            BACnetConstructedDataStatusFlags]
        //[*, 'STOP_TIME'                               BACnetConstructedDataStopTime]
        //[*, 'STOP_WHEN_FULL'                          BACnetConstructedDataStopWhenFull]
        //[*, 'STRIKE_COUNT'                            BACnetConstructedDataStrikeCount]
        //[*, 'STRUCTURED_OBJECT_LIST'                  BACnetConstructedDataStructuredObjectList]
        //[*, 'SUBORDINATE_ANNOTATIONS'                 BACnetConstructedDataSubordinateAnnotations]
        //[*, 'SUBORDINATE_LIST'                        BACnetConstructedDataSubordinateList]
        //[*, 'SUBORDINATE_NODE_TYPES'                  BACnetConstructedDataSubordinateNodeTypes]
        //[*, 'SUBORDINATE_RELATIONSHIPS'               BACnetConstructedDataSubordinateRelationships]
        //[*, 'SUBORDINATE_TAGS'                        BACnetConstructedDataSubordinateTags]
        //[*, 'SUBSCRIBED_RECIPIENTS'                   BACnetConstructedDataSubscribedRecipients]
        //[*, 'SUPPORTED_FORMAT_CLASSES'                BACnetConstructedDataSupportedFormatClasses]
        //[*, 'SUPPORTED_FORMATS'                       BACnetConstructedDataSupportedFormats]
        //[*, 'SUPPORTED_SECURITY_ALGORITHMS'           BACnetConstructedDataSupportedSecurityAlgorithms]
        //[*, 'SYSTEM_STATUS'                           BACnetConstructedDataSystemStatus]
        //[*, 'TAGS'                                    BACnetConstructedDataTags]
        //[*, 'THREAT_AUTHORITY'                        BACnetConstructedDataThreatAuthority]
        //[*, 'THREAT_LEVEL'                            BACnetConstructedDataThreatLevel]
        //[*, 'TIME_DELAY'                              BACnetConstructedDataTimeDelay]
        //[*, 'TIME_DELAY_NORMAL'                       BACnetConstructedDataTimeDelayNormal]
        //[*, 'TIME_OF_ACTIVE_TIME_RESET'               BACnetConstructedDataTimeOfActiveTimeReset]
        //[*, 'TIME_OF_DEVICE_RESTART'                  BACnetConstructedDataTimeOfDeviceRestart]
        //[*, 'TIME_OF_STATE_COUNT_RESET'               BACnetConstructedDataTimeOfStateCountReset]
        //[*, 'TIME_OF_STRIKE_COUNT_RESET'              BACnetConstructedDataTimeOfStrikeCountReset]
        //[*, 'TIME_SYNCHRONIZATION_INTERVAL'           BACnetConstructedDataTimeSynchronizationInterval]
        //[*, 'TIME_SYNCHRONIZATION_RECIPIENTS'         BACnetConstructedDataTimeSynchronizationRecipients]
        //[*, 'TIMER_RUNNING'                           BACnetConstructedDataTimerRunning]
        //[*, 'TIMER_STATE'                             BACnetConstructedDataTimerState]
        //[*, 'TOTAL_RECORD_COUNT'                      BACnetConstructedDataTotalRecordCount]
        //[*, 'TRACE_FLAG'                              BACnetConstructedDataTraceFlag]
        //[*, 'TRACKING_VALUE'                          BACnetConstructedDataTrackingValue]
        //[*, 'TRANSACTION_NOTIFICATION_CLASS'          BACnetConstructedDataTransactionNotificationClass]
        //[*, 'TRANSITION'                              BACnetConstructedDataTransition]
        //[*, 'TRIGGER'                                 BACnetConstructedDataTrigger]
        //[*, 'UNITS'                                   BACnetConstructedDataUnits]
        //[*, 'UPDATE_INTERVAL'                         BACnetConstructedDataUpdateInterval]
        //[*, 'UPDATE_KEY_SET_TIMEOUT'                  BACnetConstructedDataUpdateKeySetTimeout]
        //[*, 'UPDATE_TIME'                             BACnetConstructedDataUpdateTime]
        //[*, 'USER_EXTERNAL_IDENTIFIER'                BACnetConstructedDataUserExternalIdentifier]
        //[*, 'USER_INFORMATION_REFERENCE'              BACnetConstructedDataUserInformationReference]
        //[*, 'USER_NAME'                               BACnetConstructedDataUserName]
        //[*, 'USER_TYPE'                               BACnetConstructedDataUserType]
        //[*, 'USES_REMAINING'                          BACnetConstructedDataUsesRemaining]
        //[*, 'UTC_OFFSET'                              BACnetConstructedDataUtcOffset]
        //[*, 'UTC_TIME_SYNCHRONIZATION_RECIPIENTS'     BACnetConstructedDataUtcTimeSynchronizationRecipients]
        //[*, 'VALID_SAMPLES'                           BACnetConstructedDataValidSamples]
        //[*, 'VALUE_BEFORE_CHANGE'                     BACnetConstructedDataValueBeforeChange]
        //[*, 'VALUE_CHANGE_TIME'                       BACnetConstructedDataValueChangeTime]
        //[*, 'VALUE_SET'                               BACnetConstructedDataValueSet]
        //[*, 'VALUE_SOURCE'                            BACnetConstructedDataValueSource]
        //[*, 'VALUE_SOURCE_ARRAY'                      BACnetConstructedDataValueSourceArray]
        //[*, 'VARIANCE_VALUE'                          BACnetConstructedDataVarianceValue]
        //[*, 'VENDOR_IDENTIFIER'                       BACnetConstructedDataVendorIdentifier]
        //[*, 'VENDOR_NAME'                             BACnetConstructedDataVendorName]
        //[*, 'VERIFICATION_TIME'                       BACnetConstructedDataVerificationTime]
        //[*, 'VIRTUAL_MAC_ADDRESS_TABLE'               BACnetConstructedDataVirtualMacAddressTable]
        //[*, 'VT_CLASSES_SUPPORTED'                    BACnetConstructedDataVtClassesSupported]
        //[*, 'WEEKLY_SCHEDULE'                         BACnetConstructedDataWeeklySchedule]
        //[*, 'WINDOW_INTERVAL'                         BACnetConstructedDataWindowInterval]
        //[*, 'WINDOW_SAMPLES'                          BACnetConstructedDataWindowSamples]
        //[*, 'WRITE_STATUS'                            BACnetConstructedDataWriteStatus]
        //[*, 'ZONE_FROM'                               BACnetConstructedDataZoneFrom]
        [*, 'ZONE_MEMBERS'                              BACnetConstructedDataZoneMembers
            [array  BACnetDeviceObjectReference
                    members
                            terminated
                            'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, tagNumber)'            ]
        ]
        //[*, 'ZONE_TO'                                 BACnetConstructedDataZoneTo]
        // BACnetConstructedDataUnspecified is used for unmapped properties
        [BACnetConstructedDataUnspecified
            [array    BACnetConstructedDataElement('objectType', 'propertyIdentifierArgument')
                            data                    terminated
                                'STATIC_CALL("isBACnetConstructedDataClosingTag", readBuffer, false, tagNumber)']
        ]
        //
        /////
    ]
    [simple       BACnetClosingTag('tagNumber', 'BACnetDataType.CLOSING_TAG')
                        closingTag                                                                              ]
]

[type BACnetDeviceObjectReference
    [optional BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                        deviceIdentifier                                                                        ]
    [simple   BACnetContextTagObjectIdentifier('1', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                        objectIdentifier                                                                        ]
]

[type BACnetDeviceObjectPropertyReference
    [simple   BACnetContextTagObjectIdentifier('0', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                        objectIdentifier                                                                        ]
    [simple   BACnetContextTagPropertyIdentifier('1', 'BACnetDataType.BACNET_PROPERTY_IDENTIFIER')
                        propertyIdentifier                                                                      ]
    [optional BACnetContextTagUnsignedInteger('2', 'BACnetDataType.UNSIGNED_INTEGER')
                        arrayIndex                                                                              ]
    [optional BACnetContextTagObjectIdentifier('3', 'BACnetDataType.BACNET_OBJECT_IDENTIFIER')
                        deviceIdentifier                                                                        ]
]

[type BACnetConstructedDataElement(BACnetObjectType objectType, BACnetContextTagPropertyIdentifier propertyIdentifier)
    [peek       BACnetTagHeader
                            peekedTagHeader                                                                     ]
    [virtual    uint 8      peekedTagNumber     'peekedTagHeader.actualTagNumber']
    [virtual    bit         isApplicationTag    'peekedTagHeader.tagClass == TagClass.APPLICATION_TAGS'         ]
    [virtual    bit         isConstructedData   '!isApplicationTag && peekedTagHeader.lengthValueType == 0x6'   ]
    [virtual    bit         isContextTag        '!isConstructedData && !isApplicationTag'                       ]
    [validation '!isContextTag || (isContextTag && peekedTagHeader.lengthValueType != 0x7)'
                "unexpected closing tag"                                                                        ]
    [optional   BACnetApplicationTag
                            applicationTag      'isApplicationTag'                                              ]
    [optional   BACnetContextTag('peekedTagNumber',
                                 'STATIC_CALL("guessDataType", objectType, propertyIdentifier)')
                            contextTag          'isContextTag'                                                  ]
    [optional   BACnetConstructedData('peekedTagNumber', 'objectType', 'propertyIdentifier')
                            constructedData     'isConstructedData'                                             ]
    [validation 'isApplicationTag || isContextTag || isConstructedData'
                "BACnetConstructedDataElement could not parse anything"                                         ]
]

[enum uint 16 BVLCResultCode
    ['0x0000' SUCCESSFUL_COMPLETION                     ]
    ['0x0010' WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK    ]
    ['0x0020' READ_BROADCAST_DISTRIBUTION_TABLE_NAK     ]
    ['0x0030' REGISTER_FOREIGN_DEVICE_NAK               ]
    ['0x0040' READ_FOREIGN_DEVICE_TABLE_NAK             ]
    ['0x0050' DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK     ]
    ['0x0060' DISTRIBUTE_BROADCAST_TO_NETWORK_NAK       ]
]

[enum uint 1 TagClass
    ['0x0' APPLICATION_TAGS                     ]
    ['0x1' CONTEXT_SPECIFIC_TAGS                ]
]

[enum uint 8 BACnetDataType
    ['0' NULL                                   ]
    ['1' BOOLEAN                                ]
    ['2' UNSIGNED_INTEGER                       ]
    ['3' SIGNED_INTEGER                         ]
    ['4' REAL                                   ]
    ['5' DOUBLE                                 ]
    ['6' OCTET_STRING                           ]
    ['7' CHARACTER_STRING                       ]
    ['8' BIT_STRING                             ]
    ['9' ENUMERATED                             ]
    ['10' DATE                                  ]
    ['11' TIME                                  ]
    ['12' BACNET_OBJECT_IDENTIFIER              ]
    //////////
    //////////
    //
    // Custom plc4x helper values below here
    ['16' BACNET_PROPERTY_IDENTIFIER            ]
    ['17' BACNET_DEVICE_STATE                   ]
    ['20' OPENING_TAG                           ]
    ['21' CLOSING_TAG                           ]
    // Event Related tags
    ['30' EVENT_TYPE                            ]
    ['31' EVENT_STATE                           ]
    ['32' NOTIFY_TYPE                           ]
    //
    //////////
    //////////
]

[enum byte BACnetCharacterEncoding
    ['0x0' ISO_10646                            ] // UTF-8
    ['0x1' IBM_Microsoft_DBCS                   ]
    ['0x2' JIS_X_0208                           ]
    ['0x3' ISO_10646_4                          ] // (UCS-4)
    ['0x4' ISO_10646_2                          ] // (UCS-2)
    ['0x5' ISO_8859_1                           ]
]

[enum uint 4 BACnetNetworkType
    ['0x0' ETHERNET                             ]
    ['0x1' ARCNET                               ]
    ['0x2' MSTP                                 ]
    ['0x3' PTP                                  ]
    ['0x4' LONTALK                              ]
    ['0x5' IPV4                                 ]
    ['0x6' ZIGBEE                               ]
    ['0x7' VIRTUAL                              ]
    ['0x8' REMOVED_NON_BACNET                   ]
    ['0x9' IPV6                                 ]
    ['0xA' SERIAL                               ]
]

[enum uint 8 BACnetDeviceState
    ['0x0' COLDSTART                            ]
    ['0x1' WARMSTART                            ]
    ['0x2' ACTIVATE_CHANGES                     ]
    ['0x3' STARTBACKUP                          ]
    ['0x4' ENDBACKUP                            ]
    ['0x5' STARTRESTORE                         ]
    ['0x6' ENDRESTORE                           ]
    ['0x7' ABORTRESTORE                         ]
]

[enum uint 8 BACnetNodeType
    ['0x00' UNKNOWN                             ]
    ['0x01' SYSTEM                              ]
    ['0x02' NETWORK                             ]
    ['0x03' DEVICE                              ]
    ['0x04' ORGANIZATIONAL                      ]
    ['0x05' AREA                                ]
    ['0x06' EQUIPMENT                           ]
    ['0x07' POINT                               ]
    ['0x08' COLLECTION                          ]
    ['0x09' PROPERTY                            ]
    ['0x0A' FUNCTIONAL                          ]
    ['0x0B' OTHER                               ]
    ['0x0C' SUBSYSTEM                           ]
    ['0x0D' BUILDING                            ]
    ['0x0E' FLOOR                               ]
    ['0x0F' SECTION                             ]
    ['0x10' MODULE                              ]
    ['0x11' TREE                                ]
    ['0x12' MEMBER                              ]
    ['0x13' PROTOCOL                            ]
    ['0x14' ROOM                                ]
    ['0x15' ZONE                                ]
]

[enum uint 8 BACnetNotifyType
    ['0x0' ALARM                                ]
    ['0x1' EVENT                                ]
    ['0x2' ACK_NOTIFICATION                     ]
]

[enum uint 9 BACnetEventType
    ['0'  CHANGE_OF_BITSTRING                   ]
    ['1'  CHANGE_OF_STATE                       ]
    ['2'  CHANGE_OF_VALUE                       ]
    ['3'  COMMAND_FAILURE                       ]
    ['4'  FLOATING_LIMIT                        ]
    ['5'  OUT_OF_RANGE                          ]
    ['8'  CHANGE_OF_LIFE_SAFETY                 ]
    ['9'  EXTENDED                              ]
    ['10' BUFFER_READY                          ]
    ['11' UNSIGNED_RANGE                        ]
    ['13' ACCESS_EVENT                          ]
    ['14' DOUBLE_OUT_OF_RANGE                   ]
    ['15' SIGNED_OUT_OF_RANGE                   ]
    ['16' UNSIGNED_OUT_OF_RANGE                 ]
    ['17' CHANGE_OF_CHARACTERSTRING             ]
    ['18' CHANGE_OF_STATUS_FLAGS                ]
    ['19' CHANGE_OF_RELIABILITY                 ]
    ['20' NONE                                  ]
    ['21' CHANGE_OF_DISCRETE_VALUE              ]
    ['22' CHANGE_OF_TIMER                       ]
    
    // plc4x definition
    ['0xFFFF' VENDOR_PROPRIETARY_VALUE          ]
]

[enum uint 9 BACnetEventState
    ['0'  NORMAL                                ]
    ['1'  FAULT                                 ]
    ['2'  OFFNORMAL                             ]
    ['3'  HIGH_LIMIT                            ]
    ['4'  LOW_LIMIT                             ]
    ['5'  LIFE_SAVETY_ALARM                     ]
    
    // plc4x definition
    ['0xFFFF' VENDOR_PROPRIETARY_VALUE          ]
]

[enum uint 10 BACnetObjectType
    ['32' ACCESS_CREDENTIAL                     ]
    ['30' ACCESS_DOOR                           ]
    ['33' ACCESS_POINT                          ]
    ['34' ACCESS_RIGHTS                         ]
    ['35' ACCESS_USER                           ]
    ['36' ACCESS_ZONE                           ]
    ['23' ACCUMULATOR                           ]
    ['52' ALERT_ENROLLMENT                      ]
    ['0'  ANALOG_INPUT                          ]
    ['1'  ANALOG_OUTPUT                         ]
    ['2'  ANALOG_VALUE                          ]
    ['18' AVERAGING                             ]
    ['3'  BINARY_INPUT                          ]
    ['55' BINARY_LIGHTING_OUTPUT                ]
    ['4'  BINARY_OUTPUT                         ]
    ['5'  BINARY_VALUE                          ]
    ['39' BITSTRING_VALUE                       ]
    ['6'  CALENDAR                              ]
    ['53' CHANNEL                               ]
    ['40' CHARACTERSTRING_VALUE                 ]
    ['7'  COMMAND                               ]
    ['37' CREDENTIAL_DATA_INPUT                 ]
    ['41' DATEPATTERN_VALUE                     ]
    ['42' DATE_VALUE                            ]
    ['43' DATETIMEPATTERN_VALUE                 ]
    ['44' DATETIME_VALUE                        ]
    ['8'  DEVICE                                ]
    ['57' ELEVATOR_GROUP                        ]
    ['58' ESCALATOR                             ]
    ['9'  EVENT_ENROLLMENT                      ]
    ['25' EVENT_LOG                             ]
    ['10' FILE                                  ]
    ['26' GLOBAL_GROUP                          ]
    ['11' GROUP                                 ]
    ['45' INTEGER_VALUE                         ]
    ['46' LARGE_ANALOG_VALUE                    ]
    ['21' LIFE_SAFETY_POINT                     ]
    ['22' LIFE_SAFETY_ZONE                      ]
    ['59' LIFT                                  ]
    ['54' LIGHTING_OUTPUT                       ]
    ['28' LOAD_CONTROL                          ]
    ['12' LOOP                                  ]
    ['13' MULTI_STATE_INPUT                     ]
    ['14' MULTI_STATE_OUTPUT                    ]
    ['19' MULTI_STATE_VALUE                     ]
    ['56' NETWORK_PORT                          ]
    ['38' NETWORK_SECURITY                      ]
    ['15' NOTIFICATION_CLASS                    ]
    ['51' NOTIFICATION_FORWARDER                ]
    ['47' OCTETSTRING_VALUE                     ]
    ['48' POSITIVE_INTEGER_VALUE                ]
    ['16' PROGRAM                               ]
    ['24' PULSE_CONVERTER                       ]
    ['17' SCHEDULE                              ]
    ['29' STRUCTURED_VIEW                       ]
    ['49' TIMEPATTERN_VALUE                     ]
    ['50' TIME_VALUE                            ]
    ['31' TIMER                                 ]
    ['20' TREND_LOG                             ]
    ['27' TREND_LOG_MULTIPLE                    ]

    // plc4x Value
    ['0x3FF' VENDOR_PROPRIETARY_VALUE           ]
]

[enum uint 32 BACnetPropertyIdentifier
    ['244' ABSENTEE_LIMIT                       ]
    ['175' ACCEPTED_MODES                       ]
    ['245' ACCESS_ALARM_EVENTS                  ]
    ['246' ACCESS_DOORS                         ]
    ['247' ACCESS_EVENT                         ]
    ['248' ACCESS_EVENT_AUTHENTICATION_FACTOR   ]
    ['249' ACCESS_EVENT_CREDENTIAL              ]
    ['322' ACCESS_EVENT_TAG                     ]
    ['250' ACCESS_EVENT_TIME                    ]
    ['251' ACCESS_TRANSACTION_EVENTS            ]
    ['252' ACCOMPANIMENT                        ]
    ['253' ACCOMPANIMENT_TIME                   ]
    ['1'   ACK_REQUIRED                         ]
    ['0'   ACKED_TRANSITIONS                    ]
    ['2'   ACTION                               ]
    ['3'   ACTION_TEXT                          ]
    ['254' ACTIVATION_TIME                      ]
    ['255' ACTIVE_AUTHENTICATION_POLICY         ]
    ['481' ACTIVE_COV_MULTIPLE_SUBSCRIPTIONS    ]
    ['152' ACTIVE_COV_SUBSCRIPTIONS             ]
    ['4'   ACTIVE_TEXT                          ]
    ['5'   ACTIVE_VT_SESSIONS                   ]
    ['212' ACTUAL_SHED_LEVEL                    ]
    ['176' ADJUST_VALUE                         ]
    ['6'   ALARM_VALUE                          ]
    ['7'   ALARM_VALUES                         ]
    ['193' ALIGN_INTERVALS                      ]
    ['8'   ALL                                  ]
    ['9'   ALL_WRITES_SUCCESSFUL                ]
    ['365' ALLOW_GROUP_DELAY_INHIBIT            ]
    ['399' APDU_LENGTH                          ]
    ['10'  APDU_SEGMENT_TIMEOUT                 ]
    ['11'  APDU_TIMEOUT                         ]
    ['12'  APPLICATION_SOFTWARE_VERSION         ]
    ['13'  ARCHIVE                              ]
    ['256' ASSIGNED_ACCESS_RIGHTS               ]
    ['447' ASSIGNED_LANDING_CALLS               ]
    ['124' ATTEMPTED_SAMPLES                    ]
    ['257' AUTHENTICATION_FACTORS               ]
    ['258' AUTHENTICATION_POLICY_LIST           ]
    ['259' AUTHENTICATION_POLICY_NAMES          ]
    ['260' AUTHENTICATION_STATUS                ]
    ['364' AUTHORIZATION_EXEMPTIONS             ]
    ['261' AUTHORIZATION_MODE                   ]
    ['169' AUTO_SLAVE_DISCOVERY                 ]
    ['125' AVERAGE_VALUE                        ]
    ['338' BACKUP_AND_RESTORE_STATE             ]
    ['153' BACKUP_FAILURE_TIMEOUT               ]
    ['339' BACKUP_PREPARATION_TIME              ]
    ['407' BACNET_IP_GLOBAL_ADDRESS             ]
    ['408' BACNET_IP_MODE                       ]
    ['409' BACNET_IP_MULTICAST_ADDRESS          ]
    ['410' BACNET_IP_NAT_TRAVERSAL              ]
    ['412' BACNET_IP_UDP_PORT                   ]
    ['435' BACNET_IPV6_MODE                     ]
    ['438' BACNET_IPV6_UDP_PORT                 ]
    ['440' BACNET_IPV6_MULTICAST_ADDRESS        ]
    ['327' BASE_DEVICE_SECURITY_POLICY          ]
    ['413' BBMD_ACCEPT_FD_REGISTRATIONS         ]
    ['414' BBMD_BROADCAST_DISTRIBUTION_TABLE    ]
    ['415' BBMD_FOREIGN_DEVICE_TABLE            ]
    ['262' BELONGS_TO                           ]
    ['14'  BIAS                                 ]
    ['342' BIT_MASK                             ]
    ['343' BIT_TEXT                             ]
    ['373' BLINK_WARN_ENABLE                    ]
    ['126' BUFFER_SIZE                          ]
    ['448' CAR_ASSIGNED_DIRECTION               ]
    ['449' CAR_DOOR_COMMAND                     ]
    ['450' CAR_DOOR_STATUS                      ]
    ['451' CAR_DOOR_TEXT                        ]
    ['452' CAR_DOOR_ZONE                        ]
    ['453' CAR_DRIVE_STATUS                     ]
    ['454' CAR_LOAD                             ]
    ['455' CAR_LOAD_UNITS                       ]
    ['456' CAR_MODE                             ]
    ['457' CAR_MOVING_DIRECTION                 ]
    ['458' CAR_POSITION                         ]
    ['15'  CHANGE_OF_STATE_COUNT                ]
    ['16'  CHANGE_OF_STATE_TIME                 ]
    ['416' CHANGES_PENDING                      ]
    ['366' CHANNEL_NUMBER                       ]
    ['127' CLIENT_COV_INCREMENT                 ]
    ['417' COMMAND                              ]
    ['430' COMMAND_TIME_ARRAY                   ]
    ['154' CONFIGURATION_FILES                  ]
    ['367' CONTROL_GROUPS                       ]
    ['19'  CONTROLLED_VARIABLE_REFERENCE        ]
    ['20'  CONTROLLED_VARIABLE_UNITS            ]
    ['21'  CONTROLLED_VARIABLE_VALUE            ]
    ['177' COUNT                                ]
    ['178' COUNT_BEFORE_CHANGE                  ]
    ['179' COUNT_CHANGE_TIME                    ]
    ['22'  COV_INCREMENT                        ]
    ['180' COV_PERIOD                           ]
    ['128' COV_RESUBSCRIPTION_INTERVAL          ]
    ['349' COVU_PERIOD                          ]
    ['350' COVU_RECIPIENTS                      ]
    ['263' CREDENTIAL_DISABLE                   ]
    ['264' CREDENTIAL_STATUS                    ]
    ['265' CREDENTIALS                          ]
    ['266' CREDENTIALS_IN_ZONE                  ]
    ['431' CURRENT_COMMAND_PRIORITY             ]
    ['155' DATABASE_REVISION                    ]
    ['23'  DATE_LIST                            ]
    ['24'  DAYLIGHT_SAVINGS_STATUS              ]
    ['267' DAYS_REMAINING                       ]
    ['25'  DEADBAND                             ]
    ['374' DEFAULT_FADE_TIME                    ]
    ['375' DEFAULT_RAMP_RATE                    ]
    ['376' DEFAULT_STEP_INCREMENT               ]
    ['490' DEFAULT_SUBORDINATE_RELATIONSHIP     ]
    ['393' DEFAULT_TIMEOUT                      ]
    ['484' DEPLOYED_PROFILE_LOCATION            ]
    ['26'  DERIVATIVE_CONSTANT                  ]
    ['27'  DERIVATIVE_CONSTANT_UNITS            ]
    ['28'  DESCRIPTION                          ]
    ['29'  DESCRIPTION_OF_HALT                  ]
    ['30'  DEVICE_ADDRESS_BINDING               ]
    ['31'  DEVICE_TYPE                          ]
    ['156' DIRECT_READING                       ]
    ['328' DISTRIBUTION_KEY_REVISION            ]
    ['329' DO_NOT_HIDE                          ]
    ['226' DOOR_ALARM_STATE                     ]
    ['227' DOOR_EXTENDED_PULSE_TIME             ]
    ['228' DOOR_MEMBERS                         ]
    ['229' DOOR_OPEN_TOO_LONG_TIME              ]
    ['230' DOOR_PULSE_TIME                      ]
    ['231' DOOR_STATUS                          ]
    ['232' DOOR_UNLOCK_DELAY_TIME               ]
    ['213' DUTY_WINDOW                          ]
    ['32'  EFFECTIVE_PERIOD                     ]
    ['386' EGRESS_ACTIVE                        ]
    ['377' EGRESS_TIME                          ]
    ['33'  ELAPSED_ACTIVE_TIME                  ]
    ['459' ELEVATOR_GROUP                       ]
    ['133' ENABLE                               ]
    ['460' ENERGY_METER                         ]
    ['461' ENERGY_METER_REF                     ]
    ['268' ENTRY_POINTS                         ]
    ['34'  ERROR_LIMIT                          ]
    ['462' ESCALATOR_MODE                       ]
    ['354' EVENT_ALGORITHM_INHIBIT              ]
    ['355' EVENT_ALGORITHM_INHIBIT_REF          ]
    ['353' EVENT_DETECTION_ENABLE               ]
    ['35'  EVENT_ENABLE                         ]
    ['351' EVENT_MESSAGE_TEXTS                  ]
    ['352' EVENT_MESSAGE_TEXTS_CONFIG           ]
    ['83'  EVENT_PARAMETERS                     ]
    ['36'  EVENT_STATE                          ]
    ['130' EVENT_TIME_STAMPS                    ]
    ['37'  EVENT_TYPE                           ]
    ['38'  EXCEPTION_SCHEDULE                   ]
    ['368' EXECUTION_DELAY                      ]
    ['269' EXIT_POINTS                          ]
    ['214' EXPECTED_SHED_LEVEL                  ]
    ['270' EXPIRATION_TIME                      ]
    ['271' EXTENDED_TIME_ENABLE                 ]
    ['272' FAILED_ATTEMPT_EVENTS                ]
    ['273' FAILED_ATTEMPTS                      ]
    ['274' FAILED_ATTEMPTS_TIME                 ]
    ['388' FAULT_HIGH_LIMIT                     ]
    ['389' FAULT_LOW_LIMIT                      ]
    ['358' FAULT_PARAMETERS                     ]
    ['463' FAULT_SIGNALS                        ]
    ['359' FAULT_TYPE                           ]
    ['39'  FAULT_VALUES                         ]
    ['418' FD_BBMD_ADDRESS                      ]
    ['419' FD_SUBSCRIPTION_LIFETIME             ]
    ['40'  FEEDBACK_VALUE                       ]
    ['41'  FILE_ACCESS_METHOD                   ]
    ['42'  FILE_SIZE                            ]
    ['43'  FILE_TYPE                            ]
    ['44'  FIRMWARE_REVISION                    ]
    ['464' FLOOR_TEXT                           ]
    ['215' FULL_DUTY_BASELINE                   ]
    ['323' GLOBAL_IDENTIFIER                    ]
    ['465' GROUP_ID                             ]
    ['346' GROUP_MEMBER_NAMES                   ]
    ['345' GROUP_MEMBERS                        ]
    ['467' GROUP_MODE                           ]
    ['45'  HIGH_LIMIT                           ]
    ['468' HIGHER_DECK                          ]
    ['47'  IN_PROCESS                           ]
    ['378' IN_PROGRESS                          ]
    ['46'  INACTIVE_TEXT                        ]
    ['394' INITIAL_TIMEOUT                      ]
    ['181' INPUT_REFERENCE                      ]
    ['469' INSTALLATION_ID                      ]
    ['48'  INSTANCE_OF                          ]
    ['379' INSTANTANEOUS_POWER                  ]
    ['49'  INTEGRAL_CONSTANT                    ]
    ['50'  INTEGRAL_CONSTANT_UNITS              ]
    ['387' INTERFACE_VALUE                      ]
    ['195' INTERVAL_OFFSET                      ]
    ['400' IP_ADDRESS                           ]
    ['401' IP_DEFAULT_GATEWAY                   ]
    ['402' IP_DHCP_ENABLE                       ]
    ['403' IP_DHCP_LEASE_TIME                   ]
    ['404' IP_DHCP_LEASE_TIME_REMAINING         ]
    ['405' IP_DHCP_SERVER                       ]
    ['406' IP_DNS_SERVER                        ]
    ['411' IP_SUBNET_MASK                       ]
    ['436' IPV6_ADDRESS                         ]
    ['442' IPV6_AUTO_ADDRESSING_ENABLE          ]
    ['439' IPV6_DEFAULT_GATEWAY                 ]
    ['443' IPV6_DHCP_LEASE_TIME                 ]
    ['444' IPV6_DHCP_LEASE_TIME_REMAINING       ]
    ['445' IPV6_DHCP_SERVER                     ]
    ['441' IPV6_DNS_SERVER                      ]
    ['437' IPV6_PREFIX_LENGTH                   ]
    ['446' IPV6_ZONE_INDEX                      ]
    ['344' IS_UTC                               ]
    ['330' KEY_SETS                             ]
    ['471' LANDING_CALL_CONTROL                 ]
    ['470' LANDING_CALLS                        ]
    ['472' LANDING_DOOR_STATUS                  ]
    ['275' LAST_ACCESS_EVENT                    ]
    ['276' LAST_ACCESS_POINT                    ]
    ['432' LAST_COMMAND_TIME                    ]
    ['277' LAST_CREDENTIAL_ADDED                ]
    ['278' LAST_CREDENTIAL_ADDED_TIME           ]
    ['279' LAST_CREDENTIAL_REMOVED              ]
    ['280' LAST_CREDENTIAL_REMOVED_TIME         ]
    ['331' LAST_KEY_SERVER                      ]
    ['173' LAST_NOTIFY_RECORD                   ]
    ['369' LAST_PRIORITY                        ]
    ['196' LAST_RESTART_REASON                  ]
    ['157' LAST_RESTORE_TIME                    ]
    ['395' LAST_STATE_CHANGE                    ]
    ['281' LAST_USE_TIME                        ]
    ['166' LIFE_SAFETY_ALARM_VALUES             ]
    ['380' LIGHTING_COMMAND                     ]
    ['381' LIGHTING_COMMAND_DEFAULT_PRIORITY    ]
    ['52'  LIMIT_ENABLE                         ]
    ['182' LIMIT_MONITORING_INTERVAL            ]
    ['420' LINK_SPEED                           ]
    ['422' LINK_SPEED_AUTONEGOTIATE             ]
    ['421' LINK_SPEEDS                          ]
    ['53'  LIST_OF_GROUP_MEMBERS                ]
    ['54'  LIST_OF_OBJECT_PROPERTY_REFERENCES   ]
    ['56'  LOCAL_DATE                           ]
    ['360' LOCAL_FORWARDING_ONLY                ]
    ['57'  LOCAL_TIME                           ]
    ['58'  LOCATION                             ]
    ['233' LOCK_STATUS                          ]
    ['282' LOCKOUT                              ]
    ['283' LOCKOUT_RELINQUISH_TIME              ]
    ['131' LOG_BUFFER                           ]
    ['132' LOG_DEVICE_OBJECT_PROPERTY           ]
    ['134' LOG_INTERVAL                         ]
    ['183' LOGGING_OBJECT                       ]
    ['184' LOGGING_RECORD                       ]
    ['197' LOGGING_TYPE                         ]
    ['390' LOW_DIFF_LIMIT                       ]
    ['59'  LOW_LIMIT                            ]
    ['473' LOWER_DECK                           ]
    ['423' MAC_ADDRESS                          ]
    ['474' MACHINE_ROOM_ID                      ]
    ['158' MAINTENANCE_REQUIRED                 ]
    ['475' MAKING_CAR_CALL                      ]
    ['60'  MANIPULATED_VARIABLE_REFERENCE       ]
    ['170' MANUAL_SLAVE_ADDRESS_BINDING         ]
    ['234' MASKED_ALARM_VALUES                  ]
    ['382' MAX_ACTUAL_VALUE                     ]
    ['62'  MAX_APDU_LENGTH_ACCEPTED             ]
    ['285' MAX_FAILED_ATTEMPTS                  ]
    ['63'  MAX_INFO_FRAMES                      ]
    ['64'  MAX_MASTER                           ]
    ['65'  MAX_PRES_VALUE                       ]
    ['167' MAX_SEGMENTS_ACCEPTED                ]
    ['61'  MAXIMUM_OUTPUT                       ]
    ['135' MAXIMUM_VALUE                        ]
    ['149' MAXIMUM_VALUE_TIMESTAMP              ]
    ['159' MEMBER_OF                            ]
    ['347' MEMBER_STATUS_FLAGS                  ]
    ['286' MEMBERS                              ]
    ['383' MIN_ACTUAL_VALUE                     ]
    ['69'  MIN_PRES_VALUE                       ]
    ['66'  MINIMUM_OFF_TIME                     ]
    ['67'  MINIMUM_ON_TIME                      ]
    ['68'  MINIMUM_OUTPUT                       ]
    ['136' MINIMUM_VALUE                        ]
    ['150' MINIMUM_VALUE_TIMESTAMP              ]
    ['160' MODE                                 ]
    ['70'  MODEL_NAME                           ]
    ['71'  MODIFICATION_DATE                    ]
    ['287' MUSTER_POINT                         ]
    ['288' NEGATIVE_ACCESS_RULES                ]
    ['332' NETWORK_ACCESS_SECURITY_POLICIES     ]
    ['424' NETWORK_INTERFACE_NAME               ]
    ['425' NETWORK_NUMBER                       ]
    ['426' NETWORK_NUMBER_QUALITY               ]
    ['427' NETWORK_TYPE                         ]
    ['476' NEXT_STOPPING_FLOOR                  ]
    ['207' NODE_SUBTYPE                         ]
    ['208' NODE_TYPE                            ]
    ['17'  NOTIFICATION_CLASS                   ]
    ['137' NOTIFICATION_THRESHOLD               ]
    ['72'  NOTIFY_TYPE                          ]
    ['73'  NUMBER_OF_APDU_RETRIES               ]
    ['289' NUMBER_OF_AUTHENTICATION_POLICIES    ]
    ['74'  NUMBER_OF_STATES                     ]
    ['75'  OBJECT_IDENTIFIER                    ]
    ['76'  OBJECT_LIST                          ]
    ['77'  OBJECT_NAME                          ]
    ['78'  OBJECT_PROPERTY_REFERENCE            ]
    ['79'  OBJECT_TYPE                          ]
    ['290' OCCUPANCY_COUNT                      ]
    ['291' OCCUPANCY_COUNT_ADJUST               ]
    ['292' OCCUPANCY_COUNT_ENABLE               ]
    ['294' OCCUPANCY_LOWER_LIMIT                ]
    ['295' OCCUPANCY_LOWER_LIMIT_ENFORCED       ]
    ['296' OCCUPANCY_STATE                      ]
    ['297' OCCUPANCY_UPPER_LIMIT                ]
    ['298' OCCUPANCY_UPPER_LIMIT_ENFORCED       ]
    ['477' OPERATION_DIRECTION                  ]
    ['161' OPERATION_EXPECTED                   ]
    ['80'  OPTIONAL                             ]
    ['81'  OUT_OF_SERVICE                       ]
    ['82'  OUTPUT_UNITS                         ]
    ['333' PACKET_REORDER_TIME                  ]
    ['300' PASSBACK_MODE                        ]
    ['301' PASSBACK_TIMEOUT                     ]
    ['478' PASSENGER_ALARM                      ]
    ['84'  POLARITY                             ]
    ['363' PORT_FILTER                          ]
    ['302' POSITIVE_ACCESS_RULES                ]
    ['384' POWER                                ]
    ['479' POWER_MODE                           ]
    ['185' PRESCALE                             ]
    ['85'  PRESENT_VALUE                        ]
    ['86'  PRIORITY                             ]
    ['87'  PRIORITY_ARRAY                       ]
    ['88'  PRIORITY_FOR_WRITING                 ]
    ['89'  PROCESS_IDENTIFIER                   ]
    ['361' PROCESS_IDENTIFIER_FILTER            ]
    ['485' PROFILE_LOCATION                     ]
    ['168' PROFILE_NAME                         ]
    ['90'  PROGRAM_CHANGE                       ]
    ['91'  PROGRAM_LOCATION                     ]
    ['92'  PROGRAM_STATE                        ]
    ['371' PROPERTY_LIST                        ]
    ['93'  PROPORTIONAL_CONSTANT                ]
    ['94'  PROPORTIONAL_CONSTANT_UNITS          ]
    ['482' PROTOCOL_LEVEL                       ]
    ['95'  PROTOCOL_CONFORMANCE_CLASS           ]
    ['96'  PROTOCOL_OBJECT_TYPES_SUPPORTED      ]
    ['139' PROTOCOL_REVISION                    ]
    ['97'  PROTOCOL_SERVICES_SUPPORTED          ]
    ['98'  PROTOCOL_VERSION                     ]
    ['186' PULSE_RATE                           ]
    ['99'  READ_ONLY                            ]
    ['303' REASON_FOR_DISABLE                   ]
    ['100' REASON_FOR_HALT                      ]
    ['102' RECIPIENT_LIST                       ]
    ['141' RECORD_COUNT                         ]
    ['140' RECORDS_SINCE_NOTIFICATION           ]
    ['483' REFERENCE_PORT                       ]
    ['480' REGISTERED_CAR_CALL                  ]
    ['103' RELIABILITY                          ]
    ['357' RELIABILITY_EVALUATION_INHIBIT       ]
    ['104' RELINQUISH_DEFAULT                   ]
    ['491' REPRESENTS                           ]
    ['218' REQUESTED_SHED_LEVEL                 ]
    ['348' REQUESTED_UPDATE_INTERVAL            ]
    ['105' REQUIRED                             ]
    ['106' RESOLUTION                           ]
    ['202' RESTART_NOTIFICATION_RECIPIENTS      ]
    ['340' RESTORE_COMPLETION_TIME              ]
    ['341' RESTORE_PREPARATION_TIME             ]
    ['428' ROUTING_TABLE                        ]
    ['187' SCALE                                ]
    ['188' SCALE_FACTOR                         ]
    ['174' SCHEDULE_DEFAULT                     ]
    ['235' SECURED_STATUS                       ]
    ['334' SECURITY_PDU_TIMEOUT                 ]
    ['335' SECURITY_TIME_WINDOW                 ]
    ['107' SEGMENTATION_SUPPORTED               ]
    ['372' SERIAL_NUMBER                        ]
    ['108' SETPOINT                             ]
    ['109' SETPOINT_REFERENCE                   ]
    ['162' SETTING                              ]
    ['219' SHED_DURATION                        ]
    ['220' SHED_LEVEL_DESCRIPTIONS              ]
    ['221' SHED_LEVELS                          ]
    ['163' SILENCED                             ]
    ['171' SLAVE_ADDRESS_BINDING                ]
    ['172' SLAVE_PROXY_ENABLE                   ]
    ['142' START_TIME                           ]
    ['396' STATE_CHANGE_VALUES                  ]
    ['222' STATE_DESCRIPTION                    ]
    ['110' STATE_TEXT                           ]
    ['111' STATUS_FLAGS                         ]
    ['143' STOP_TIME                            ]
    ['144' STOP_WHEN_FULL                       ]
    ['391' STRIKE_COUNT                         ]
    ['209' STRUCTURED_OBJECT_LIST               ]
    ['210' SUBORDINATE_ANNOTATIONS              ]
    ['211' SUBORDINATE_LIST                     ]
    ['487' SUBORDINATE_NODE_TYPES               ]
    ['489' SUBORDINATE_RELATIONSHIPS            ]
    ['488' SUBORDINATE_TAGS                     ]
    ['362' SUBSCRIBED_RECIPIENTS                ]
    ['305' SUPPORTED_FORMAT_CLASSES             ]
    ['304' SUPPORTED_FORMATS                    ]
    ['336' SUPPORTED_SECURITY_ALGORITHMS        ]
    ['112' SYSTEM_STATUS                        ]
    ['486' TAGS                                 ]
    ['306' THREAT_AUTHORITY                     ]
    ['307' THREAT_LEVEL                         ]
    ['113' TIME_DELAY                           ]
    ['356' TIME_DELAY_NORMAL                    ]
    ['114' TIME_OF_ACTIVE_TIME_RESET            ]
    ['203' TIME_OF_DEVICE_RESTART               ]
    ['115' TIME_OF_STATE_COUNT_RESET            ]
    ['392' TIME_OF_STRIKE_COUNT_RESET           ]
    ['204' TIME_SYNCHRONIZATION_INTERVAL        ]
    ['116' TIME_SYNCHRONIZATION_RECIPIENTS      ]
    ['397' TIMER_RUNNING                        ]
    ['398' TIMER_STATE                          ]
    ['145' TOTAL_RECORD_COUNT                   ]
    ['308' TRACE_FLAG                           ]
    ['164' TRACKING_VALUE                       ]
    ['309' TRANSACTION_NOTIFICATION_CLASS       ]
    ['385' TRANSITION                           ]
    ['205' TRIGGER                              ]
    ['117' UNITS                                ]
    ['118' UPDATE_INTERVAL                      ]
    ['337' UPDATE_KEY_SET_TIMEOUT               ]
    ['189' UPDATE_TIME                          ]
    ['310' USER_EXTERNAL_IDENTIFIER             ]
    ['311' USER_INFORMATION_REFERENCE           ]
    ['317' USER_NAME                            ]
    ['318' USER_TYPE                            ]
    ['319' USES_REMAINING                       ]
    ['119' UTC_OFFSET                           ]
    ['206' UTC_TIME_SYNCHRONIZATION_RECIPIENTS  ]
    ['146' VALID_SAMPLES                        ]
    ['190' VALUE_BEFORE_CHANGE                  ]
    ['192' VALUE_CHANGE_TIME                    ]
    ['191' VALUE_SET                            ]
    ['433' VALUE_SOURCE                         ]
    ['434' VALUE_SOURCE_ARRAY                   ]
    ['151' VARIANCE_VALUE                       ]
    ['120' VENDOR_IDENTIFIER                    ]
    ['121' VENDOR_NAME                          ]
    ['326' VERIFICATION_TIME                    ]
    ['429' VIRTUAL_MAC_ADDRESS_TABLE            ]
    ['122' VT_CLASSES_SUPPORTED                 ]
    ['123' WEEKLY_SCHEDULE                      ]
    ['147' WINDOW_INTERVAL                      ]
    ['148' WINDOW_SAMPLES                       ]
    ['370' WRITE_STATUS                         ]
    ['320' ZONE_FROM                            ]
    ['165' ZONE_MEMBERS                         ]
    ['321' ZONE_TO                              ]

    // plc4x definition
    ['9999' VENDOR_PROPRIETARY_VALUE            ]
]
