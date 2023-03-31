/**
 * Hyperledger Cactus API
 *
 * Interact with a Cactus deployment through HTTP.
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
 * 
 *
 * Values: Subscribe,Next,Unsubscribe,Error,Complete
 */

enum class WatchHealthcheckV1(val value: kotlin.String) {

    @Json(name = "org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Subscribe")
    Subscribe("org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Subscribe"),

    @Json(name = "org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Next")
    Next("org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Next"),

    @Json(name = "org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Unsubscribe")
    Unsubscribe("org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Unsubscribe"),

    @Json(name = "org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Error")
    Error("org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Error"),

    @Json(name = "org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Complete")
    Complete("org.hyperledger.cactus.api.async.besu.WatchHealthcheckV1.Complete");

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
        fun encode(data: Any?): kotlin.String? = if (data is WatchHealthcheckV1) "$data" else null

        /**
         * Returns a valid [WatchHealthcheckV1] for [data], null otherwise.
         */
        fun decode(data: Any?): WatchHealthcheckV1? = data?.let {
          val normalizedData = "$it".lowercase()
          values().firstOrNull { value ->
            it == value || normalizedData == "$value".lowercase()
          }
        }
    }
}

