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
package org.apache.plc4x.camel;

import org.apache.camel.Component;
import org.apache.camel.Processor;
import org.apache.camel.impl.DefaultCamelContext;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.hamcrest.core.Is.is;
import static org.hamcrest.core.IsNull.notNullValue;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.mockito.Mockito.*;

public class Plc4XEndpointTest {

    Plc4XEndpoint SUT;

    @BeforeEach
    public void setUp() throws Exception {
        Component mockComponent = mock(Component.class, RETURNS_DEEP_STUBS);
        when(mockComponent.getCamelContext()).thenReturn(new DefaultCamelContext());
        SUT = new Plc4XEndpoint("plc4x:mock:10.10.10.1/1/1", mockComponent);
    }

    // TODO: figure out what this is
    @Test
    public void createProducer() throws Exception {
        assertThat(SUT.createProducer(), notNullValue());
    }

    @Test
    public void createConsumer() throws Exception {
        assertThat(SUT.createConsumer(mock(Processor.class)), notNullValue());
    }

    @Test
    public void isSingleton() {
        assertThat(SUT.isSingleton(), is(true));
    }

}