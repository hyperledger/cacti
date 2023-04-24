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
 * 
 *
 * @param id 
 * @param name 
 * @param mainApiHost 
 * @param memberIds The collection (array) of primary keys of consortium member entities that belong to this Consortium.
 */

data class Consortium (

    @Json(name = "id")
    val id: kotlin.String,

    @Json(name = "name")
    val name: kotlin.String,

    @Json(name = "mainApiHost")
    val mainApiHost: kotlin.String,

    /* The collection (array) of primary keys of consortium member entities that belong to this Consortium. */
    @Json(name = "memberIds")
    val memberIds: kotlin.collections.List<kotlin.String> = arrayListOf()

)

