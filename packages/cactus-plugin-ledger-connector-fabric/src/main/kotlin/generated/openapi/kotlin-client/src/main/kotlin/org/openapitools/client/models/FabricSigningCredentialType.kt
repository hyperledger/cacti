/**
 * Hyperledger Cactus Plugin - Connector Fabric
 *
 * Can perform basic tasks on a fabric ledger
 *
 * The version of the OpenAPI document: 0.0.1
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
 * different type of identity provider for singing fabric messages supported by this package
 *
 * Values: xPeriod509,vaultMinusXPeriod509,wSMinusXPeriod509
 */

enum class FabricSigningCredentialType(val value: kotlin.String) {

    @Json(name = "X.509")
    xPeriod509("X.509"),

    @Json(name = "Vault-X.509")
    vaultMinusXPeriod509("Vault-X.509"),

    @Json(name = "WS-X.509")
    wSMinusXPeriod509("WS-X.509");

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
        fun encode(data: Any?): kotlin.String? = if (data is FabricSigningCredentialType) "$data" else null

        /**
         * Returns a valid [FabricSigningCredentialType] for [data], null otherwise.
         */
        fun decode(data: Any?): FabricSigningCredentialType? = data?.let {
          val normalizedData = "$it".lowercase()
          values().firstOrNull { value ->
            it == value || normalizedData == "$value".lowercase()
          }
        }
    }
}

