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
#include "plc4x_field.h"

// Code generated by code-generation. DO NOT EDIT.


// Parse function.
plc4c_return_code plc4c_plc4x_read_write_plc4x_field_parse(plc4c_spi_read_buffer* readBuffer, plc4c_plc4x_read_write_plc4x_field** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_plc4x_read_write_plc4x_field));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Implicit Field (nameLen) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t nameLen = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &nameLen);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (name)
  char* name = "";
  _res = plc4c_spi_read_string(readBuffer, (nameLen) * (8), "UTF-8", (char**) &name);
  if(_res != OK) {
    return _res;
  }
  (*_message)->name = name;

  // Implicit Field (fieldQueryLen) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t fieldQueryLen = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &fieldQueryLen);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (fieldQuery)
  char* fieldQuery = "";
  _res = plc4c_spi_read_string(readBuffer, (fieldQueryLen) * (8), "UTF-8", (char**) &fieldQuery);
  if(_res != OK) {
    return _res;
  }
  (*_message)->field_query = fieldQuery;

  return OK;
}

plc4c_return_code plc4c_plc4x_read_write_plc4x_field_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_plc4x_read_write_plc4x_field* _message) {
  plc4c_return_code _res = OK;

  // Implicit Field (nameLen) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t nameLen = plc4c_spi_evaluation_helper_str_len(_message->name);
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, nameLen);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (name)
  char* name = _message->name;
  _res = plc4c_spi_write_string(writeBuffer, (nameLen) * (8), "UTF-8", name);
  if(_res != OK) {
    return _res;
  }

  // Implicit Field (fieldQueryLen) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t fieldQueryLen = plc4c_spi_evaluation_helper_str_len(_message->field_query);
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, fieldQueryLen);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (fieldQuery)
  char* fieldQuery = _message->field_query;
  _res = plc4c_spi_write_string(writeBuffer, (fieldQueryLen) * (8), "UTF-8", fieldQuery);
  if(_res != OK) {
    return _res;
  }

  return OK;
}

uint16_t plc4c_plc4x_read_write_plc4x_field_length_in_bytes(plc4c_plc4x_read_write_plc4x_field* _message) {
  return plc4c_plc4x_read_write_plc4x_field_length_in_bits(_message) / 8;
}

uint16_t plc4c_plc4x_read_write_plc4x_field_length_in_bits(plc4c_plc4x_read_write_plc4x_field* _message) {
  uint16_t lengthInBits = 0;

  // Implicit Field (nameLen)
  lengthInBits += 8;

  // Simple field (name)
  lengthInBits += plc4c_spi_evaluation_helper_str_len(_message->name) * 8;

  // Implicit Field (fieldQueryLen)
  lengthInBits += 8;

  // Simple field (fieldQuery)
  lengthInBits += plc4c_spi_evaluation_helper_str_len(_message->field_query) * 8;

  return lengthInBits;
}

