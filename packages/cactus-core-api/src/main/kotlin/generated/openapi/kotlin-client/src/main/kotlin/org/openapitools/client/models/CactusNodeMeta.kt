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
 * A Cactus node meta information
 *
 * @param nodeApiHost 
 * @param publicKeyPem The PEM encoded public key that was used to generate the JWS included in the response (the jws property)
 */

data class CactusNodeMeta (

    @Json(name = "nodeApiHost")
    val nodeApiHost: kotlin.String,

    /* The PEM encoded public key that was used to generate the JWS included in the response (the jws property) */
    @Json(name = "publicKeyPem")
    val publicKeyPem: kotlin.String

)

