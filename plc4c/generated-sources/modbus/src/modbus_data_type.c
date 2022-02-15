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

#include "modbus_data_type.h"
#include <string.h>

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_null_const;

plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_null() {
  return plc4c_modbus_read_write_modbus_data_type_null_const;
}

uint8_t plc4c_modbus_read_write_modbus_data_type_get_value(plc4c_modbus_read_write_modbus_data_type value) {
    switch(value) {
        case plc4c_modbus_read_write_modbus_data_type_BOOL:
            return (uint8_t) 1;
        case plc4c_modbus_read_write_modbus_data_type_BYTE:
            return (uint8_t) 2;
        case plc4c_modbus_read_write_modbus_data_type_WORD:
            return (uint8_t) 3;
        case plc4c_modbus_read_write_modbus_data_type_DWORD:
            return (uint8_t) 4;
        case plc4c_modbus_read_write_modbus_data_type_LWORD:
            return (uint8_t) 5;
        case plc4c_modbus_read_write_modbus_data_type_SINT:
            return (uint8_t) 6;
        case plc4c_modbus_read_write_modbus_data_type_INT:
            return (uint8_t) 7;
        case plc4c_modbus_read_write_modbus_data_type_DINT:
            return (uint8_t) 8;
        case plc4c_modbus_read_write_modbus_data_type_LINT:
            return (uint8_t) 9;
        case plc4c_modbus_read_write_modbus_data_type_USINT:
            return (uint8_t) 10;
        case plc4c_modbus_read_write_modbus_data_type_UINT:
            return (uint8_t) 11;
        case plc4c_modbus_read_write_modbus_data_type_UDINT:
            return (uint8_t) 12;
        case plc4c_modbus_read_write_modbus_data_type_ULINT:
            return (uint8_t) 13;
        case plc4c_modbus_read_write_modbus_data_type_REAL:
            return (uint8_t) 14;
        case plc4c_modbus_read_write_modbus_data_type_LREAL:
            return (uint8_t) 15;
        case plc4c_modbus_read_write_modbus_data_type_TIME:
            return (uint8_t) 16;
        case plc4c_modbus_read_write_modbus_data_type_LTIME:
            return (uint8_t) 17;
        case plc4c_modbus_read_write_modbus_data_type_DATE:
            return (uint8_t) 18;
        case plc4c_modbus_read_write_modbus_data_type_LDATE:
            return (uint8_t) 19;
        case plc4c_modbus_read_write_modbus_data_type_TIME_OF_DAY:
            return (uint8_t) 20;
        case plc4c_modbus_read_write_modbus_data_type_LTIME_OF_DAY:
            return (uint8_t) 21;
        case plc4c_modbus_read_write_modbus_data_type_DATE_AND_TIME:
            return (uint8_t) 22;
        case plc4c_modbus_read_write_modbus_data_type_LDATE_AND_TIME:
            return (uint8_t) 23;
        case plc4c_modbus_read_write_modbus_data_type_CHAR:
            return (uint8_t) 24;
        case plc4c_modbus_read_write_modbus_data_type_WCHAR:
            return (uint8_t) 25;
        case plc4c_modbus_read_write_modbus_data_type_STRING:
            return (uint8_t) 26;
        case plc4c_modbus_read_write_modbus_data_type_WSTRING:
            return (uint8_t) 27;
    }
    return 0;
}

plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_for_value(uint8_t value) {
    switch(value) {
        case (uint8_t) 1:
            return plc4c_modbus_read_write_modbus_data_type_BOOL;
        case (uint8_t) 2:
            return plc4c_modbus_read_write_modbus_data_type_BYTE;
        case (uint8_t) 3:
            return plc4c_modbus_read_write_modbus_data_type_WORD;
        case (uint8_t) 4:
            return plc4c_modbus_read_write_modbus_data_type_DWORD;
        case (uint8_t) 5:
            return plc4c_modbus_read_write_modbus_data_type_LWORD;
        case (uint8_t) 6:
            return plc4c_modbus_read_write_modbus_data_type_SINT;
        case (uint8_t) 7:
            return plc4c_modbus_read_write_modbus_data_type_INT;
        case (uint8_t) 8:
            return plc4c_modbus_read_write_modbus_data_type_DINT;
        case (uint8_t) 9:
            return plc4c_modbus_read_write_modbus_data_type_LINT;
        case (uint8_t) 10:
            return plc4c_modbus_read_write_modbus_data_type_USINT;
        case (uint8_t) 11:
            return plc4c_modbus_read_write_modbus_data_type_UINT;
        case (uint8_t) 12:
            return plc4c_modbus_read_write_modbus_data_type_UDINT;
        case (uint8_t) 13:
            return plc4c_modbus_read_write_modbus_data_type_ULINT;
        case (uint8_t) 14:
            return plc4c_modbus_read_write_modbus_data_type_REAL;
        case (uint8_t) 15:
            return plc4c_modbus_read_write_modbus_data_type_LREAL;
        case (uint8_t) 16:
            return plc4c_modbus_read_write_modbus_data_type_TIME;
        case (uint8_t) 17:
            return plc4c_modbus_read_write_modbus_data_type_LTIME;
        case (uint8_t) 18:
            return plc4c_modbus_read_write_modbus_data_type_DATE;
        case (uint8_t) 19:
            return plc4c_modbus_read_write_modbus_data_type_LDATE;
        case (uint8_t) 20:
            return plc4c_modbus_read_write_modbus_data_type_TIME_OF_DAY;
        case (uint8_t) 21:
            return plc4c_modbus_read_write_modbus_data_type_LTIME_OF_DAY;
        case (uint8_t) 22:
            return plc4c_modbus_read_write_modbus_data_type_DATE_AND_TIME;
        case (uint8_t) 23:
            return plc4c_modbus_read_write_modbus_data_type_LDATE_AND_TIME;
        case (uint8_t) 24:
            return plc4c_modbus_read_write_modbus_data_type_CHAR;
        case (uint8_t) 25:
            return plc4c_modbus_read_write_modbus_data_type_WCHAR;
        case (uint8_t) 26:
            return plc4c_modbus_read_write_modbus_data_type_STRING;
        case (uint8_t) 27:
            return plc4c_modbus_read_write_modbus_data_type_WSTRING;
    }
    return 0;
}

    // Parse function.
plc4c_return_code plc4c_modbus_read_write_modbus_data_type_parse(plc4c_spi_read_buffer* readBuffer, plc4c_modbus_read_write_modbus_data_type* _message) {
    plc4c_return_code _res = OK;

    uint8_t value;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
    *_message = plc4c_modbus_read_write_modbus_data_type_for_value(value);

    return _res;
}

plc4c_return_code plc4c_modbus_read_write_modbus_data_type_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_modbus_read_write_modbus_data_type* _message) {
    plc4c_return_code _res = OK;

    uint8_t value = plc4c_modbus_read_write_modbus_data_type_get_value(*_message);
    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, value);

    return _res;
}

plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_value_of(char* value_string) {
    if(strcmp(value_string, "BOOL") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_BOOL;
    }
    if(strcmp(value_string, "BYTE") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_BYTE;
    }
    if(strcmp(value_string, "WORD") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_WORD;
    }
    if(strcmp(value_string, "DWORD") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_DWORD;
    }
    if(strcmp(value_string, "LWORD") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_LWORD;
    }
    if(strcmp(value_string, "SINT") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_SINT;
    }
    if(strcmp(value_string, "INT") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_INT;
    }
    if(strcmp(value_string, "DINT") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_DINT;
    }
    if(strcmp(value_string, "LINT") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_LINT;
    }
    if(strcmp(value_string, "USINT") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_USINT;
    }
    if(strcmp(value_string, "UINT") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_UINT;
    }
    if(strcmp(value_string, "UDINT") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_UDINT;
    }
    if(strcmp(value_string, "ULINT") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_ULINT;
    }
    if(strcmp(value_string, "REAL") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_REAL;
    }
    if(strcmp(value_string, "LREAL") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_LREAL;
    }
    if(strcmp(value_string, "TIME") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_TIME;
    }
    if(strcmp(value_string, "LTIME") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_LTIME;
    }
    if(strcmp(value_string, "DATE") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_DATE;
    }
    if(strcmp(value_string, "LDATE") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_LDATE;
    }
    if(strcmp(value_string, "TIME_OF_DAY") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_TIME_OF_DAY;
    }
    if(strcmp(value_string, "LTIME_OF_DAY") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_LTIME_OF_DAY;
    }
    if(strcmp(value_string, "DATE_AND_TIME") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_DATE_AND_TIME;
    }
    if(strcmp(value_string, "LDATE_AND_TIME") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_LDATE_AND_TIME;
    }
    if(strcmp(value_string, "CHAR") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_CHAR;
    }
    if(strcmp(value_string, "WCHAR") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_WCHAR;
    }
    if(strcmp(value_string, "STRING") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_STRING;
    }
    if(strcmp(value_string, "WSTRING") == 0) {
        return plc4c_modbus_read_write_modbus_data_type_WSTRING;
    }
    return -1;
}

int plc4c_modbus_read_write_modbus_data_type_num_values() {
  return 27;
}

plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_modbus_read_write_modbus_data_type_BOOL;
      }
      case 1: {
        return plc4c_modbus_read_write_modbus_data_type_BYTE;
      }
      case 2: {
        return plc4c_modbus_read_write_modbus_data_type_WORD;
      }
      case 3: {
        return plc4c_modbus_read_write_modbus_data_type_DWORD;
      }
      case 4: {
        return plc4c_modbus_read_write_modbus_data_type_LWORD;
      }
      case 5: {
        return plc4c_modbus_read_write_modbus_data_type_SINT;
      }
      case 6: {
        return plc4c_modbus_read_write_modbus_data_type_INT;
      }
      case 7: {
        return plc4c_modbus_read_write_modbus_data_type_DINT;
      }
      case 8: {
        return plc4c_modbus_read_write_modbus_data_type_LINT;
      }
      case 9: {
        return plc4c_modbus_read_write_modbus_data_type_USINT;
      }
      case 10: {
        return plc4c_modbus_read_write_modbus_data_type_UINT;
      }
      case 11: {
        return plc4c_modbus_read_write_modbus_data_type_UDINT;
      }
      case 12: {
        return plc4c_modbus_read_write_modbus_data_type_ULINT;
      }
      case 13: {
        return plc4c_modbus_read_write_modbus_data_type_REAL;
      }
      case 14: {
        return plc4c_modbus_read_write_modbus_data_type_LREAL;
      }
      case 15: {
        return plc4c_modbus_read_write_modbus_data_type_TIME;
      }
      case 16: {
        return plc4c_modbus_read_write_modbus_data_type_LTIME;
      }
      case 17: {
        return plc4c_modbus_read_write_modbus_data_type_DATE;
      }
      case 18: {
        return plc4c_modbus_read_write_modbus_data_type_LDATE;
      }
      case 19: {
        return plc4c_modbus_read_write_modbus_data_type_TIME_OF_DAY;
      }
      case 20: {
        return plc4c_modbus_read_write_modbus_data_type_LTIME_OF_DAY;
      }
      case 21: {
        return plc4c_modbus_read_write_modbus_data_type_DATE_AND_TIME;
      }
      case 22: {
        return plc4c_modbus_read_write_modbus_data_type_LDATE_AND_TIME;
      }
      case 23: {
        return plc4c_modbus_read_write_modbus_data_type_CHAR;
      }
      case 24: {
        return plc4c_modbus_read_write_modbus_data_type_WCHAR;
      }
      case 25: {
        return plc4c_modbus_read_write_modbus_data_type_STRING;
      }
      case 26: {
        return plc4c_modbus_read_write_modbus_data_type_WSTRING;
      }
      default: {
        return -1;
      }
    }
}

uint8_t plc4c_modbus_read_write_modbus_data_type_get_data_type_size(plc4c_modbus_read_write_modbus_data_type value) {
  switch(value) {
    case plc4c_modbus_read_write_modbus_data_type_BOOL: { /* '1' */
      return 1;
    }
    case plc4c_modbus_read_write_modbus_data_type_USINT: { /* '10' */
      return 1;
    }
    case plc4c_modbus_read_write_modbus_data_type_UINT: { /* '11' */
      return 2;
    }
    case plc4c_modbus_read_write_modbus_data_type_UDINT: { /* '12' */
      return 4;
    }
    case plc4c_modbus_read_write_modbus_data_type_ULINT: { /* '13' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_REAL: { /* '14' */
      return 4;
    }
    case plc4c_modbus_read_write_modbus_data_type_LREAL: { /* '15' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_TIME: { /* '16' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_LTIME: { /* '17' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_DATE: { /* '18' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_LDATE: { /* '19' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_BYTE: { /* '2' */
      return 1;
    }
    case plc4c_modbus_read_write_modbus_data_type_TIME_OF_DAY: { /* '20' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_LTIME_OF_DAY: { /* '21' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_DATE_AND_TIME: { /* '22' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_LDATE_AND_TIME: { /* '23' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_CHAR: { /* '24' */
      return 1;
    }
    case plc4c_modbus_read_write_modbus_data_type_WCHAR: { /* '25' */
      return 2;
    }
    case plc4c_modbus_read_write_modbus_data_type_STRING: { /* '26' */
      return 1;
    }
    case plc4c_modbus_read_write_modbus_data_type_WSTRING: { /* '27' */
      return 2;
    }
    case plc4c_modbus_read_write_modbus_data_type_WORD: { /* '3' */
      return 2;
    }
    case plc4c_modbus_read_write_modbus_data_type_DWORD: { /* '4' */
      return 4;
    }
    case plc4c_modbus_read_write_modbus_data_type_LWORD: { /* '5' */
      return 8;
    }
    case plc4c_modbus_read_write_modbus_data_type_SINT: { /* '6' */
      return 1;
    }
    case plc4c_modbus_read_write_modbus_data_type_INT: { /* '7' */
      return 2;
    }
    case plc4c_modbus_read_write_modbus_data_type_DINT: { /* '8' */
      return 4;
    }
    case plc4c_modbus_read_write_modbus_data_type_LINT: { /* '9' */
      return 8;
    }
    default: {
      return 0;
    }
  }
}

plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_get_first_enum_for_field_data_type_size(uint8_t value) {
    switch(value) {
        case 1: {
            return plc4c_modbus_read_write_modbus_data_type_BOOL;
        }
        case 2: {
            return plc4c_modbus_read_write_modbus_data_type_WORD;
        }
        case 4: {
            return plc4c_modbus_read_write_modbus_data_type_DWORD;
        }
        case 8: {
            return plc4c_modbus_read_write_modbus_data_type_LWORD;
        }
        default: {
            return -1;
        }
    }
}

uint16_t plc4c_modbus_read_write_modbus_data_type_length_in_bytes(plc4c_modbus_read_write_modbus_data_type* _message) {
    return plc4c_modbus_read_write_modbus_data_type_length_in_bits(_message) / 8;
}

uint16_t plc4c_modbus_read_write_modbus_data_type_length_in_bits(plc4c_modbus_read_write_modbus_data_type* _message) {
    return 8;
}
