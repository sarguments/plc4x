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
package org.apache.plc4x.java.modbus;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcDriver;
import org.apache.plc4x.java.api.messages.PlcDiscoveryResponse;
import org.junit.jupiter.api.Disabled;

@Disabled("Manual Test")
public class ManualModbusDiscoveryTest {

    public static void main(String[] args) throws Exception {
        final PlcDriver modbusDriver = new PlcDriverManager().getDriver("modbus-tcp");
        final PlcDiscoveryResponse plcDiscoveryResponse = modbusDriver.discoveryRequestBuilder().build().execute().get();
        System.out.println(plcDiscoveryResponse);
    }

}
