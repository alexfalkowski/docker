#!/usr/bin/env bash

set -euo pipefail

oteldel_start() {
  readonly OTELDEL_PIPELINE_NAME="$1"
  local start_time
  start_time="$(date +%s.%N)"
  readonly OTELDEL_START_TIME="$start_time"
  OTELDEL_RESULT="success"
  OTELDEL_ERROR_TYPE=""
}

oteldel_mark_skip() {
  OTELDEL_RESULT="skip"
}

oteldel_set_error_type() {
  OTELDEL_ERROR_TYPE="$1"
}

oteldel_now() {
  date +%s.%N
}

oteldel_elapsed_since() {
  local start="$1"
  local finish
  finish="$(oteldel_now)"

  awk "BEGIN { printf \"%.6f\", $finish - $start }"
}

oteldel_emit() {
  local metric_type="$1"
  local metric_name="$2"
  local value="$3"
  shift 3

  if ! command -v oteldel >/dev/null 2>&1; then
    return 0
  fi

  local args=()
  local attr
  for attr in "$@"; do
    args+=("--attr" "$attr")
  done

  oteldel emit "$metric_type" "$metric_name" "$value" "${args[@]}" || true
}

oteldel_emit_artifact_publish() {
  local result="$1"
  local duration="$2"

  oteldel_emit counter cicd.artifact.publish.count 1 \
    "artifact.kind=github_release" \
    "result=$result"
  oteldel_emit histogram cicd.artifact.publish.duration "$duration" \
    "artifact.kind=github_release" \
    "result=$result"
}

oteldel_emit_deployment_request() {
  local result="$1"
  local duration="$2"

  oteldel_emit counter cicd.deployment.request.count 1 \
    "target.system=infraops" \
    "environment=production" \
    "result=$result"
  oteldel_emit histogram cicd.deployment.request.duration "$duration" \
    "target.system=infraops" \
    "environment=production" \
    "result=$result"
}

oteldel_finalize() {
  local exit_code="$1"
  local duration
  duration="$(oteldel_elapsed_since "$OTELDEL_START_TIME")"
  local attrs=(
    "cicd.pipeline.name=$OTELDEL_PIPELINE_NAME"
    "cicd.pipeline.run.state=executing"
  )

  if [[ "$exit_code" -ne 0 ]]; then
    OTELDEL_RESULT="failure"
    if [[ -z "$OTELDEL_ERROR_TYPE" ]]; then
      OTELDEL_ERROR_TYPE="unknown"
    fi
  fi

  attrs+=("cicd.pipeline.result=$OTELDEL_RESULT")

  if [[ "$OTELDEL_RESULT" == "failure" ]]; then
    attrs+=("error.type=$OTELDEL_ERROR_TYPE")
  fi

  oteldel_emit histogram cicd.pipeline.run.duration "$duration" "${attrs[@]}"

  if [[ "$OTELDEL_RESULT" == "failure" ]]; then
    oteldel_emit counter cicd.pipeline.run.errors 1 \
      "cicd.pipeline.name=$OTELDEL_PIPELINE_NAME" \
      "error.type=$OTELDEL_ERROR_TYPE"
  fi
}
