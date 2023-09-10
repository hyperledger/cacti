/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package org.hyperledger.cacti.weaver.imodule.corda

import org.hyperledger.cacti.weaver.imodule.corda.flows.*
import org.hyperledger.cacti.weaver.imodule.corda.states.*
import net.corda.core.utilities.getOrThrow
import net.corda.testing.node.MockNetwork
import net.corda.testing.node.MockNetworkParameters
import net.corda.testing.node.StartedMockNode
import net.corda.testing.node.TestCordapp
import org.junit.AfterClass
import org.junit.BeforeClass
import org.junit.Test
import kotlin.test.assertEquals
import kotlin.test.assertTrue
import com.google.gson.Gson
import com.google.gson.annotations.SerializedName
import org.hyperledger.cacti.weaver.imodule.corda.states.InvocationSpec

class WriteExternalStateTest {
    companion object {
        lateinit var network: MockNetwork
        lateinit var partyA: StartedMockNode
        lateinit var partyB: StartedMockNode

        @BeforeClass
        @JvmStatic
        fun setup() {
            network = MockNetwork(MockNetworkParameters(cordappsForAllNodes = listOf(
                    TestCordapp.findCordapp("org.hyperledger.cacti.weaver.imodule.corda.contracts"),
                    TestCordapp.findCordapp("org.hyperledger.cacti.weaver.imodule.corda.flows"),
                    TestCordapp.findCordapp("org.hyperledger.cacti.weaver.imodule.corda.test")
            )))
            partyA = network.createPartyNode()
            partyB = network.createPartyNode()
            network.runNetwork()
        }

        @AfterClass
        @JvmStatic
        fun tearDown() {
            network.stopNodes()
            System.setProperty("net.corda.node.dbtransactionsresolver.InMemoryResolutionLimit", "0")
        }
    }
    
    val partya = partyA.info.legalIdentities.first()
    val partyb = partyB.info.legalIdentities.first()
    
    val fabricNetwork = "network1"
    val fabricRelayEndpoint = "relay-network1:9080"
    val fabricViewAddress = "mychannel:simplestate:Read:a"
    
    data class TestData(
        @SerializedName("view64") val B64View: String,
        @SerializedName("confidential_view64") val B64ViewConfidential: String,
        @SerializedName("confidential_view_content64") val B64ViewContents: List<String>
    )
    val fabricTestDataJSON = javaClass.getResource("/test_data/fabric_viewdata_1_org.json").readText(Charsets.UTF_8)
    val fabricTestData = Gson().fromJson(fabricTestDataJSON, TestData::class.java)
    
    val fabricCert = javaClass.getResource("/test_data/fabric_cacert_org1.pem").readText(Charsets.UTF_8)

    val fabricVerificationPolicy = VerificationPolicyState(
            securityDomain = fabricNetwork,
            identifiers = listOf(Identifier(
                    fabricViewAddress,
                    Policy("signature", listOf("Org1MSP"))
            )),
            participants = listOf(partya, partyb)
    )

    val fabricMembership = MembershipState(
            securityDomain = fabricNetwork,
            members = mapOf("Org1MSP" to Member(
                    value = fabricCert,
                    type = "ca",
                    chain = listOf("")
            )),
            participants = listOf(partya, partyb)
    )
    
    val cordaTestDataJSON = javaClass.getResource("/test_data/corda_viewdata.json").readText(Charsets.UTF_8)
    val cordaTestData = Gson().fromJson(cordaTestDataJSON, TestData::class.java)

    val cordaVerificationPolicy = VerificationPolicyState(
            securityDomain = "Corda_Network",
            identifiers = listOf(Identifier(
                    "localhost:10006#com.cordaSimpleApplication.flow.GetStateByKey:*",
                    Policy("signature", listOf("PartyA"))
            )),
            participants = listOf(partya, partyb)
    )
    
    val rootCACert = javaClass.getResource("/test_data/corda_cacert_root.pem").readText(Charsets.UTF_8)
    val doormanCACert = javaClass.getResource("/test_data/corda_cacert_doorman.pem").readText(Charsets.UTF_8)
    val nodeCACert = javaClass.getResource("/test_data/corda_cacert_node.pem").readText(Charsets.UTF_8)
    val certChain = listOf(
            rootCACert,
            doormanCACert,
            nodeCACert)
            
    /* val cordaCert = javaClass.getResource("/test_data/corda_cacert.pem").readText(Charsets.UTF_8) */
    
    val cordaMembership = MembershipState(
            securityDomain = "Corda_Network",
            members = mapOf("PartyA" to Member(
                    value = "",
                    type = "certificate",
                    chain = certChain
            )),
            participants = listOf(partya, partyb)
    )
    val disableInvokeObject = InvocationSpec()
    val invokeObject = InvocationSpec(
        disableInvocation = false,
        invokeFlowName = "org.hyperledger.cacti.weaver.imodule.corda.test.UserFlow",
        invokeFlowArgs = listOf(arrayOf<String>()),
        interopArgsIndex = 0
    )
    
    fun initCordaPolicies() {
        // Create corda membership and verificationPolicy in vault
        val future = partyA.startFlow(CreateVerificationPolicyState(cordaVerificationPolicy))
        network.runNetwork()
        val linearId = future.getOrThrow()
        assert(linearId.isRight()) { "CreateVerificationPolicyState should return a Right(UniqueIdentifier)" }
        
        val futureb = partyB.startFlow(CreateVerificationPolicyState(cordaVerificationPolicy))
        network.runNetwork()
        val linearIdb = futureb.getOrThrow()
        assert(linearIdb.isRight()) { "CreateVerificationPolicyState should return a Right(UniqueIdentifier)" }

        val future2 = partyA.startFlow(CreateMembershipState(cordaMembership))
        network.runNetwork()
        val linearId2 = future2.getOrThrow()
        assert(linearId2.isRight()) { "CreateMembershipState should return a Right(UniqueIdentifier)" }

        val future2b = partyB.startFlow(CreateMembershipState(cordaMembership))
        network.runNetwork()
        val linearId2b = future2b.getOrThrow()
        assert(linearId2b.isRight()) { "CreateMembershipState should return a Right(UniqueIdentifier)" }
    }
    
    
    fun initFabricPolicies() {
        // Create fabric membership and verificationPolicy in vault
        val future3 = partyA.startFlow(CreateVerificationPolicyState(fabricVerificationPolicy))
        network.runNetwork()
        val linearId3 = future3.getOrThrow()
        assert(linearId3.isRight()) { "CreateVerificationPolicyState should return a Right(UniqueIdentifier)" }

        val future4 = partyA.startFlow(CreateMembershipState(fabricMembership))
        network.runNetwork()
        val linearId4 = future4.getOrThrow()
        assert(linearId4.isRight()) { "CreateMembershipState should return a Right(UniqueIdentifier)" }
        
        // Create fabric membership and verificationPolicy in vault
        val future3b = partyB.startFlow(CreateVerificationPolicyState(fabricVerificationPolicy))
        network.runNetwork()
        val linearId3b = future3b.getOrThrow()
        assert(linearId3b.isRight()) { "CreateVerificationPolicyState should return a Right(UniqueIdentifier)" }

        val future4b = partyB.startFlow(CreateMembershipState(fabricMembership))
        network.runNetwork()
        val linearId4b = future4b.getOrThrow()
        assert(linearId4b.isRight()) { "CreateMembershipState should return a Right(UniqueIdentifier)" }
    }

    @Test
    fun `WriteExternalState tests`() {
        // Corda happy case
        initCordaPolicies()

        val happyFuture = partyA.startFlow(WriteExternalStateInitiator(arrayOf(cordaTestData.B64View), arrayOf("localhost:9081/Corda_Network/localhost:10006#com.cordaSimpleApplication.flow.GetStateByKey:H"), disableInvokeObject, listOf(partyb)))
        network.runNetwork()
        val happyLinearId = happyFuture.getOrThrow()
        assertTrue(happyLinearId.isRight())
        
        // Corda happy case2
        val happyFuture_1 = partyA.startFlow(WriteExternalStateInitiator(arrayOf(cordaTestData.B64View), arrayOf("localhost:9081/Corda_Network/localhost:10006#com.cordaSimpleApplication.flow.GetStateByKey:H"), invokeObject, listOf(partyb)))
        network.runNetwork()
        val happyResponse_1 = happyFuture_1.getOrThrow()
        assertTrue(happyResponse_1.isRight())
        happyResponse_1.fold({ println("Error") },{
            assertEquals(it, true)
        })

        // Fabric happy case
        initFabricPolicies()

        val happyFuture2 = partyA.startFlow(WriteExternalStateInitiator(arrayOf(fabricTestData.B64View), arrayOf("${fabricRelayEndpoint}/${fabricNetwork}/${fabricViewAddress}"), disableInvokeObject, listOf(partyb)))
        network.runNetwork()
        val happyLinearId2 = happyFuture2.getOrThrow()
        assertTrue(happyLinearId2.isRight())

        // Test case: Invalid cert in Membership
        val invalidMembership = cordaMembership.copy(members = mapOf(("PartyA" to Member(
                value = "invalid_cert",
                type = "ca",
                chain = listOf("")
        ))))
        val future5 = partyA.startFlow(UpdateMembershipState(invalidMembership))
        network.runNetwork()
        val linearId5 = future5.getOrThrow()
        assert(linearId5.isRight()) { "UpdateMembershipState should return a Right(UniqueIdentifier)" }
        
        val future5b = partyB.startFlow(UpdateMembershipState(invalidMembership))
        network.runNetwork()
        val linearId5b = future5b.getOrThrow()
        assert(linearId5b.isRight()) { "UpdateMembershipState should return a Right(UniqueIdentifier)" }

        val unhappyFuture = partyA.startFlow(WriteExternalStateInitiator(arrayOf(cordaTestData.B64View), arrayOf("localhost:9081/Corda_Network/localhost:10006#com.cordaSimpleApplication.flow.GetStateByKey:H"), disableInvokeObject, listOf(partyb)))
        network.runNetwork()
        val unhappyLinearId = unhappyFuture.getOrThrow()
        assertTrue(unhappyLinearId.isLeft())
        assertEquals("View verification failed with error: Parse Error: failed to parse requester certificate: Illegal base64 character 5f", unhappyLinearId.fold({ it.message }, { "" }))


        // Test case: Invalid policy in verification policy


    }
}
