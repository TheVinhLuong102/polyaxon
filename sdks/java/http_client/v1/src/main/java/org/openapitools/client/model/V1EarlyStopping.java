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
import org.openapitools.client.model.V1DiffStoppingPolicy;
import org.openapitools.client.model.V1FailureEarlyStopping;
import org.openapitools.client.model.V1MedianStoppingPolicy;
import org.openapitools.client.model.V1MetricEarlyStopping;
import org.openapitools.client.model.V1TruncationStoppingPolicy;

/**
 * V1EarlyStopping
 */

public class V1EarlyStopping {
  public static final String SERIALIZED_NAME_MEDIAN = "median";
  @SerializedName(SERIALIZED_NAME_MEDIAN)
  private V1MedianStoppingPolicy median;

  public static final String SERIALIZED_NAME_DIFF = "diff";
  @SerializedName(SERIALIZED_NAME_DIFF)
  private V1DiffStoppingPolicy diff;

  public static final String SERIALIZED_NAME_TRUNCATION = "truncation";
  @SerializedName(SERIALIZED_NAME_TRUNCATION)
  private V1TruncationStoppingPolicy truncation;

  public static final String SERIALIZED_NAME_METRIC = "metric";
  @SerializedName(SERIALIZED_NAME_METRIC)
  private V1MetricEarlyStopping metric;

  public static final String SERIALIZED_NAME_FAILURE = "failure";
  @SerializedName(SERIALIZED_NAME_FAILURE)
  private V1FailureEarlyStopping failure;


  public V1EarlyStopping median(V1MedianStoppingPolicy median) {
    
    this.median = median;
    return this;
  }

   /**
   * Get median
   * @return median
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1MedianStoppingPolicy getMedian() {
    return median;
  }


  public void setMedian(V1MedianStoppingPolicy median) {
    this.median = median;
  }


  public V1EarlyStopping diff(V1DiffStoppingPolicy diff) {
    
    this.diff = diff;
    return this;
  }

   /**
   * Get diff
   * @return diff
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1DiffStoppingPolicy getDiff() {
    return diff;
  }


  public void setDiff(V1DiffStoppingPolicy diff) {
    this.diff = diff;
  }


  public V1EarlyStopping truncation(V1TruncationStoppingPolicy truncation) {
    
    this.truncation = truncation;
    return this;
  }

   /**
   * Get truncation
   * @return truncation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1TruncationStoppingPolicy getTruncation() {
    return truncation;
  }


  public void setTruncation(V1TruncationStoppingPolicy truncation) {
    this.truncation = truncation;
  }


  public V1EarlyStopping metric(V1MetricEarlyStopping metric) {
    
    this.metric = metric;
    return this;
  }

   /**
   * Get metric
   * @return metric
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1MetricEarlyStopping getMetric() {
    return metric;
  }


  public void setMetric(V1MetricEarlyStopping metric) {
    this.metric = metric;
  }


  public V1EarlyStopping failure(V1FailureEarlyStopping failure) {
    
    this.failure = failure;
    return this;
  }

   /**
   * Get failure
   * @return failure
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1FailureEarlyStopping getFailure() {
    return failure;
  }


  public void setFailure(V1FailureEarlyStopping failure) {
    this.failure = failure;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1EarlyStopping v1EarlyStopping = (V1EarlyStopping) o;
    return Objects.equals(this.median, v1EarlyStopping.median) &&
        Objects.equals(this.diff, v1EarlyStopping.diff) &&
        Objects.equals(this.truncation, v1EarlyStopping.truncation) &&
        Objects.equals(this.metric, v1EarlyStopping.metric) &&
        Objects.equals(this.failure, v1EarlyStopping.failure);
  }

  @Override
  public int hashCode() {
    return Objects.hash(median, diff, truncation, metric, failure);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1EarlyStopping {\n");
    sb.append("    median: ").append(toIndentedString(median)).append("\n");
    sb.append("    diff: ").append(toIndentedString(diff)).append("\n");
    sb.append("    truncation: ").append(toIndentedString(truncation)).append("\n");
    sb.append("    metric: ").append(toIndentedString(metric)).append("\n");
    sb.append("    failure: ").append(toIndentedString(failure)).append("\n");
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

