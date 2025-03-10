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
package org.apache.plc4x.protocol.plc4x.v0;

import org.apache.plc4x.plugins.codegenerator.language.mspec.parser.MessageFormatParser;
import org.apache.plc4x.plugins.codegenerator.protocol.Protocol;
import org.apache.plc4x.plugins.codegenerator.protocol.TypeContext;
import org.apache.plc4x.plugins.codegenerator.types.definitions.TypeDefinition;
import org.apache.plc4x.plugins.codegenerator.types.exceptions.GenerationException;

import java.io.InputStream;
import java.util.Map;
import java.util.Optional;

public class Plc4xProtocol implements Protocol {

    @Override
    public String getName() {
        return "plc4x";
    }

    @Override
    public Optional<String> getVersion() {
        return Optional.of("0");
    }

    @Override
    public TypeContext getTypeContext() throws GenerationException {
        InputStream schemaInputStream = Plc4xProtocol.class.getResourceAsStream("/protocols/plc4x/v0/plc4x.mspec");
        if(schemaInputStream == null) {
            throw new GenerationException("Error loading message-format schema for protocol '" + getName() + "'");
        }
        TypeContext typeContext = new MessageFormatParser().parse(schemaInputStream);
        if (typeContext.getUnresolvedTypeReferences().size() > 0) {
            throw new GenerationException("Unresolved types left: " + typeContext.getUnresolvedTypeReferences());
        }
        return typeContext;
    }

}
