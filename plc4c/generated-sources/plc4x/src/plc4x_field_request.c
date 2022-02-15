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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "plc4x_field_request.h"

// Code generated by code-generation. DO NOT EDIT.


// Parse function.
plc4c_return_code plc4c_plc4x_read_write_plc4x_field_request_parse(plc4c_spi_read_buffer* readBuffer, plc4c_plc4x_read_write_plc4x_field_request** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_plc4x_read_write_plc4x_field_request));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (field)
  plc4c_plc4x_read_write_plc4x_field* field;
  _res = plc4c_plc4x_read_write_plc4x_field_parse(readBuffer, (void*) &field);
  if(_res != OK) {
    return _res;
  }
  (*_message)->field = field;

  return OK;
}

plc4c_return_code plc4c_plc4x_read_write_plc4x_field_request_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_plc4x_read_write_plc4x_field_request* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (field)
  plc4c_plc4x_read_write_plc4x_field* field = _message->field;
  _res = plc4c_plc4x_read_write_plc4x_field_serialize(writeBuffer, field);
  if(_res != OK) {
    return _res;
  }

  return OK;
}

uint16_t plc4c_plc4x_read_write_plc4x_field_request_length_in_bytes(plc4c_plc4x_read_write_plc4x_field_request* _message) {
  return plc4c_plc4x_read_write_plc4x_field_request_length_in_bits(_message) / 8;
}

uint16_t plc4c_plc4x_read_write_plc4x_field_request_length_in_bits(plc4c_plc4x_read_write_plc4x_field_request* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (field)
  lengthInBits += plc4c_plc4x_read_write_plc4x_field_length_in_bits(_message->field);

  return lengthInBits;
}

