// @generated by protoc-gen-es v1.4.2 with parameter "target=ts"
// @generated from file fusionchain/identity/genesis.proto (package fusionchain.identity, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Params } from "./params_pb.js";
import { Keyring } from "./keyring_pb.js";
import { Workspace } from "./workspace_pb.js";

/**
 * GenesisState defines the identity module's genesis state.
 *
 * @generated from message fusionchain.identity.GenesisState
 */
export class GenesisState extends Message<GenesisState> {
  /**
   * @generated from field: fusionchain.identity.Params params = 1;
   */
  params?: Params;

  /**
   * @generated from field: repeated fusionchain.identity.Keyring keyrings = 2;
   */
  keyrings: Keyring[] = [];

  /**
   * @generated from field: repeated fusionchain.identity.Workspace workspaces = 3;
   */
  workspaces: Workspace[] = [];

  constructor(data?: PartialMessage<GenesisState>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.identity.GenesisState";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "params", kind: "message", T: Params },
    { no: 2, name: "keyrings", kind: "message", T: Keyring, repeated: true },
    { no: 3, name: "workspaces", kind: "message", T: Workspace, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenesisState {
    return new GenesisState().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenesisState {
    return new GenesisState().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenesisState {
    return new GenesisState().fromJsonString(jsonString, options);
  }

  static equals(a: GenesisState | PlainMessage<GenesisState> | undefined, b: GenesisState | PlainMessage<GenesisState> | undefined): boolean {
    return proto3.util.equals(GenesisState, a, b);
  }
}

