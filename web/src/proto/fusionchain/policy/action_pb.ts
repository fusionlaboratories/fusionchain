// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file fusionchain/policy/action.proto (package fusionchain.policy, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * Current status of an action.
 *
 * @generated from enum fusionchain.policy.ActionStatus
 */
export enum ActionStatus {
  /**
   * Unspecified status.
   *
   * @generated from enum value: ACTION_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * Action is pending approval. This is the initial status.
   *
   * @generated from enum value: ACTION_STATUS_PENDING = 1;
   */
  PENDING = 1,

  /**
   * Policy has been satified, action has been executed.
   *
   * @generated from enum value: ACTION_STATUS_COMPLETED = 2;
   */
  COMPLETED = 2,

  /**
   * Action has been revoked by its creator.
   *
   * @generated from enum value: ACTION_STATUS_REVOKED = 3;
   */
  REVOKED = 3,

  /**
   * Action has been rejected since Btl is expired
   *
   * @generated from enum value: ACTION_STATUS_TIMEOUT = 4;
   */
  TIMEOUT = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(ActionStatus)
proto3.util.setEnumType(ActionStatus, "fusionchain.policy.ActionStatus", [
  { no: 0, name: "ACTION_STATUS_UNSPECIFIED" },
  { no: 1, name: "ACTION_STATUS_PENDING" },
  { no: 2, name: "ACTION_STATUS_COMPLETED" },
  { no: 3, name: "ACTION_STATUS_REVOKED" },
  { no: 4, name: "ACTION_STATUS_TIMEOUT" },
]);

/**
 * Action wraps a message that needs to be approved by a set of approvers.
 *
 * @generated from message fusionchain.policy.Action
 */
export class Action extends Message<Action> {
  /**
   * @generated from field: uint64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: repeated string approvers = 2;
   */
  approvers: string[] = [];

  /**
   * @generated from field: fusionchain.policy.ActionStatus status = 3;
   */
  status = ActionStatus.UNSPECIFIED;

  /**
   * Optional policy id that must be satisfied by the approvers.
   * If not specified, it's up to the creator of the action to decide what to
   * apply.
   *
   * @generated from field: uint64 policy_id = 4;
   */
  policyId = protoInt64.zero;

  /**
   * Original message that started the action, it will be executed when the
   * policy is satisfied.
   *
   * @generated from field: google.protobuf.Any msg = 5;
   */
  msg?: Any;

  /**
   * @generated from field: string creator = 6;
   */
  creator = "";

  /**
   * BTL (blocks to live) is the block height up until this action can be
   * approved or rejected.
   *
   * @generated from field: uint64 btl = 7;
   */
  btl = protoInt64.zero;

  /**
   * @generated from field: repeated fusionchain.policy.KeyValue policy_data = 8;
   */
  policyData: KeyValue[] = [];

  constructor(data?: PartialMessage<Action>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.policy.Action";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "approvers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(ActionStatus) },
    { no: 4, name: "policy_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 5, name: "msg", kind: "message", T: Any },
    { no: 6, name: "creator", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "btl", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 8, name: "policy_data", kind: "message", T: KeyValue, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Action {
    return new Action().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Action {
    return new Action().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Action {
    return new Action().fromJsonString(jsonString, options);
  }

  static equals(a: Action | PlainMessage<Action> | undefined, b: Action | PlainMessage<Action> | undefined): boolean {
    return proto3.util.equals(Action, a, b);
  }
}

/**
 * KeyValue is a simple key/value pair.
 *
 * @generated from message fusionchain.policy.KeyValue
 */
export class KeyValue extends Message<KeyValue> {
  /**
   * @generated from field: string key = 1;
   */
  key = "";

  /**
   * @generated from field: bytes value = 2;
   */
  value = new Uint8Array(0);

  constructor(data?: PartialMessage<KeyValue>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.policy.KeyValue";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KeyValue {
    return new KeyValue().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KeyValue {
    return new KeyValue().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KeyValue {
    return new KeyValue().fromJsonString(jsonString, options);
  }

  static equals(a: KeyValue | PlainMessage<KeyValue> | undefined, b: KeyValue | PlainMessage<KeyValue> | undefined): boolean {
    return proto3.util.equals(KeyValue, a, b);
  }
}

