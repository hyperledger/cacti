/**
 * Hyperledger Core API
 *
 * Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..
 *
 * The version of the OpenAPI document: 0.2.0
 * 
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models


import com.squareup.moshi.Json

/**
 * Enumerates a list of consensus algorithm families in existence. Does not intend to be an exhaustive list, just a practical one, meaning that we only include items here that are relevant to Hyperledger Cactus in fulfilling its own duties. This can be extended later as more sophisticated features of Cactus get implemented. This enum is meant to be first and foremost a useful abstraction for achieving practical tasks, not an encyclopedia and therefore we ask of everyone that this to be extended only in ways that serve a practical purpose for the runtime behavior of Cactus or Cactus plugins in general. The bottom line is that we can accept this enum being not 100% accurate as long as it 100% satisfies what it was designed to do.
 *
 * Values: aUTHORITY,sTAKE,wORK
 */

enum class ConsensusAlgorithmFamily(val value: kotlin.String) {

    @Json(name = "org.hyperledger.cactus.consensusalgorithm.PROOF_OF_AUTHORITY")
    aUTHORITY("org.hyperledger.cactus.consensusalgorithm.PROOF_OF_AUTHORITY"),

    @Json(name = "org.hyperledger.cactus.consensusalgorithm.PROOF_OF_STAKE")
    sTAKE("org.hyperledger.cactus.consensusalgorithm.PROOF_OF_STAKE"),

    @Json(name = "org.hyperledger.cactus.consensusalgorithm.PROOF_OF_WORK")
    wORK("org.hyperledger.cactus.consensusalgorithm.PROOF_OF_WORK");

    /**
     * Override toString() to avoid using the enum variable name as the value, and instead use
     * the actual value defined in the API spec file.
     *
     * This solves a problem when the variable name and its value are different, and ensures that
     * the client sends the correct enum values to the server always.
     */
    override fun toString(): String = value

    companion object {
        /**
         * Converts the provided [data] to a [String] on success, null otherwise.
         */
        fun encode(data: Any?): kotlin.String? = if (data is ConsensusAlgorithmFamily) "$data" else null

        /**
         * Returns a valid [ConsensusAlgorithmFamily] for [data], null otherwise.
         */
        fun decode(data: Any?): ConsensusAlgorithmFamily? = data?.let {
          val normalizedData = "$it".lowercase()
          values().firstOrNull { value ->
            it == value || normalizedData == "$value".lowercase()
          }
        }
    }
}

