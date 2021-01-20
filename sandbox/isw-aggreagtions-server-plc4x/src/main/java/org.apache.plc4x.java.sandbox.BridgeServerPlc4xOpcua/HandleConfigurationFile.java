/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.sandbox.BridgeServerPlc4xOpcua;

import org.eclipse.milo.opcua.sdk.core.ValueRanks;
import org.eclipse.milo.opcua.sdk.server.api.methods.AbstractMethodInvocationHandler;
import org.eclipse.milo.opcua.sdk.server.nodes.UaMethodNode;
import org.eclipse.milo.opcua.stack.core.Identifiers;
import org.eclipse.milo.opcua.stack.core.types.builtin.LocalizedText;
import org.eclipse.milo.opcua.stack.core.types.builtin.Variant;
import org.eclipse.milo.opcua.stack.core.types.structured.Argument;
import org.json.JSONObject;

public class HandleConfigurationFile extends AbstractMethodInvocationHandler {

    private static final Argument ConfigurationFile = new Argument(
        "ConfigurationFile",
        Identifiers.String,
        ValueRanks.Scalar,
        null,
        new LocalizedText("Load the JSON Configurationfile here")
    );

    protected HandleConfigurationFile(UaMethodNode node) {
        super(node);
    }

    @Override
    public Argument[] getInputArguments() {
        return new Argument[]{ConfigurationFile};
    }

    @Override
    public Argument[] getOutputArguments() {
        return new Argument[]{};
    }

    @Override
    protected Variant[] invoke(InvocationContext invocationContext, Variant[] inputValues) {

        String jsonAsString = (String) inputValues[0].getValue();
        JSONObject jsonFile = new JSONObject(jsonAsString);
        Boolean success = AggregationsNamespace.get().handleConfigurationFile(jsonFile);

        if (!success) {
            String[] intendedNullPointer = null;
            String x = intendedNullPointer[42];
        }

        return new Variant[]{new Variant(true)};
    }

}