/**
 * Hyperledger Cactus Plugin - Keychain Memory 
 *
 * Contains/describes the Hyperledger Cactus Keychain Memory plugin.
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
 * 
 *
 * @param key The key for the entry to set on the keychain.
 * @param `value` The value that will be associated with the key on the keychain.
 */

data class SetKeychainEntryRequest (

    /* The key for the entry to set on the keychain. */
    @Json(name = "key")
    val key: kotlin.String,

    /* The value that will be associated with the key on the keychain. */
    @Json(name = "value")
    val `value`: kotlin.String

)

