// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.0.94
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * V1GcsType
 */

public class V1GcsType {
  public static final String SERIALIZED_NAME_BUCKET = "bucket";
  @SerializedName(SERIALIZED_NAME_BUCKET)
  private String bucket;

  public static final String SERIALIZED_NAME_BLOB = "blob";
  @SerializedName(SERIALIZED_NAME_BLOB)
  private String blob;


  public V1GcsType bucket(String bucket) {
    
    this.bucket = bucket;
    return this;
  }

   /**
   * Get bucket
   * @return bucket
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBucket() {
    return bucket;
  }


  public void setBucket(String bucket) {
    this.bucket = bucket;
  }


  public V1GcsType blob(String blob) {
    
    this.blob = blob;
    return this;
  }

   /**
   * Get blob
   * @return blob
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBlob() {
    return blob;
  }


  public void setBlob(String blob) {
    this.blob = blob;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1GcsType v1GcsType = (V1GcsType) o;
    return Objects.equals(this.bucket, v1GcsType.bucket) &&
        Objects.equals(this.blob, v1GcsType.blob);
  }

  @Override
  public int hashCode() {
    return Objects.hash(bucket, blob);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1GcsType {\n");
    sb.append("    bucket: ").append(toIndentedString(bucket)).append("\n");
    sb.append("    blob: ").append(toIndentedString(blob)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

