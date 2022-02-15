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
package org.apache.plc4x.plugins.codegenerator.language.mspec.model.definitions;

import org.apache.plc4x.plugins.codegenerator.types.definitions.Argument;
import org.apache.plc4x.plugins.codegenerator.types.definitions.ComplexTypeDefinition;
import org.apache.plc4x.plugins.codegenerator.types.definitions.DataIoTypeDefinition;
import org.apache.plc4x.plugins.codegenerator.types.fields.*;
import org.apache.plc4x.plugins.codegenerator.types.references.DefaultComplexTypeReference;
import org.apache.plc4x.plugins.codegenerator.types.references.TypeReference;
import org.apache.plc4x.plugins.codegenerator.types.terms.Term;

import java.util.*;

public class DefaultDataIoTypeDefinition extends DefaultTypeDefinition implements DataIoTypeDefinition {

    private final SwitchField switchField;
    private final TypeReference type;

    public DefaultDataIoTypeDefinition(String name, Map<String, Term> attributes, List<Argument> parserArguments, SwitchField switchField) {
        super(name, attributes, parserArguments);
        this.switchField = Objects.requireNonNull(switchField);
        if (parserArguments.size() < 1) {
            throw new IllegalStateException();
        }
        this.type = Objects.requireNonNull(parserArguments.get(0).getType());
    }

    public SwitchField getSwitchField() {
        return switchField;
    }

    public TypeReference getType() {
        return this.type;
    }

    @Override
    public String toString() {
        return "DefaultDataIoTypeDefinition{" +
            "switchField=" + switchField +
            ", type=" + type +
            "} " + super.toString();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        if (!super.equals(o)) return false;
        DefaultDataIoTypeDefinition that = (DefaultDataIoTypeDefinition) o;
        return Objects.equals(switchField, that.switchField) && Objects.equals(type, that.type);
    }

    @Override
    public int hashCode() {
        return Objects.hash(super.hashCode(), switchField, type);
    }
}
