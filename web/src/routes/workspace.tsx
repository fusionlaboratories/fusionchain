import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { useLoaderData } from "react-router";
import { Link, Params } from "react-router-dom";
import { useKeplrAddress } from "../keplr";
import { MsgNewKeyRequest, MsgNewKeyRequestResponse } from "../proto/fusionchain/treasury/tx_pb";
import { KeyRequestStatus, KeyType } from "../proto/fusionchain/treasury/key_pb";
import Keys from "../components/keys";
import KeyRequests from "../components/key_requests";
import { workspaceByAddress } from "../client/identity";
import { useQuery } from "@tanstack/react-query";
import Address from "../components/address";
import { MsgRemoveWorkspaceOwner } from "../proto/fusionchain/identity/tx_pb";
import AddWorkspaceOwnerForm from "@/components/add_workspace_owner_form";
import { Button } from "@/components/ui/button";
import { Breadcrumb, BreadcrumbItem, BreadcrumbLink } from "@/components/ui/breadcrumb";
import { useBroadcaster } from "@/hooks/keplr";
import WorkspacePolicyCard from "@/components/workspace_policy_card";
import useKeyringAddress from "@/hooks/useKeyringAddress";
import { AlertDialog, AlertDialogContent, AlertDialogDescription, AlertDialogHeader, AlertDialogTitle } from "@/components/ui/alert-dialog";
import { Progress } from "@/components/ui/progress";
import { useState } from "react";
import { TxMsgData } from "cosmjs-types/cosmos/base/abci/v1beta1/abci";
import { keyRequestById } from "@/client/treasury";

function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

enum KeyRequesterState {
  IDLE = "idle",
  BROADCAST_KEY_REQUEST = "broadcast_key_request",
  WAITING_KEYRING = "waiting_keyring",
  KEY_FULFILLED = "key_fulfilled",
  ERROR = "error",
}

function useKeyRequester() {
  const { broadcast } = useBroadcaster();
  const [state, setState] = useState<KeyRequesterState>(KeyRequesterState.IDLE);
  const [error, setError] = useState<string | undefined>(undefined);

  return {
    state,
    error,
    requestKey: async (keyringAddress: string, addr: string, workspaceAddr: string) => {
      try {
        setState(KeyRequesterState.BROADCAST_KEY_REQUEST);

        const res = await broadcast([
          new MsgNewKeyRequest({ keyringAddr: keyringAddress, creator: addr, workspaceAddr, keyType: KeyType.ECDSA_SECP256K1 }),
        ]);

        setState(KeyRequesterState.WAITING_KEYRING);

        if (!res || !res.result) {
          throw new Error('failed to broadcast tx');
        }

        if (res.result?.tx_result.code) {
          throw new Error(`tx failed with code ${res.result?.tx_result.code}`);
        }

        // parse tx msg response
        const bytes = Uint8Array.from(atob(res.result.tx_result.data), c => c.charCodeAt(0));
        const msgData = TxMsgData.decode(bytes);
        const newKeyRequestResponse = MsgNewKeyRequestResponse.fromBinary(msgData.msgResponses[0].value);
        const keyRequestId = newKeyRequestResponse.id;

        // wait for sign request to be processed
        while (true) {
          const res = await keyRequestById(keyRequestId);
          if (res?.keyRequest?.status === KeyRequestStatus.PENDING) {
            await sleep(1000);
            continue;
          }

          if (res.keyRequest?.status === KeyRequestStatus.FULFILLED) {
            setState(KeyRequesterState.KEY_FULFILLED);
            return
          }

          throw new Error(`key request rejected with reason: ${res.keyRequest?.rejectReason}`);
        }
      } catch (e) {
        setError(`${e}`);
        setState(KeyRequesterState.ERROR);
      }
    },
    reset: () => {
      if (state === KeyRequesterState.KEY_FULFILLED) {
        setState(KeyRequesterState.IDLE);
      }
    },
  }
}

function textForState(state: KeyRequesterState) {
  switch (state) {
    case KeyRequesterState.IDLE:
      return "";
    case KeyRequesterState.BROADCAST_KEY_REQUEST:
      return "Waiting for Keplr to broadcast the request to Fusion Chain...";
    case KeyRequesterState.WAITING_KEYRING:
      return "Waiting for keyring to accept the request...";
    case KeyRequesterState.KEY_FULFILLED:
      return "Key request fulfilled!";
    case KeyRequesterState.ERROR:
      return "Error!";
  }
}

function progressForState(state: KeyRequesterState) {
  switch (state) {
    case KeyRequesterState.IDLE:
      return 0;
    case KeyRequesterState.BROADCAST_KEY_REQUEST:
      return 10;
    case KeyRequesterState.WAITING_KEYRING:
      return 50;
    case KeyRequesterState.KEY_FULFILLED:
      return 100;
  }
}

function Workspace() {
  const addr = useKeplrAddress();
  const [keyringAddress, _] = useKeyringAddress();
  const { broadcast } = useBroadcaster();
  const { state, error, requestKey, reset } = useKeyRequester();
  const { workspaceAddr } = useLoaderData() as Awaited<ReturnType<typeof loader>>;
  const wsQuery = useQuery({
    queryKey: ["workspace", workspaceAddr],
    queryFn: () => workspaceByAddress(workspaceAddr)
  });
  const ws = wsQuery.data?.workspace;

  if (!ws) {
    return (
      <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
        <div className="flex items-center justify-between space-y-2">
          <div>
            <h2 className="text-2xl font-bold tracking-tight">Workspace {workspaceAddr} not found</h2>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <Breadcrumb>
        <BreadcrumbItem>
          <BreadcrumbLink href="/">Home</BreadcrumbLink>
        </BreadcrumbItem>
        <BreadcrumbItem isCurrentPage>
          <BreadcrumbLink href={`/workspaces/${workspaceAddr}`}>Workspace {workspaceAddr}</BreadcrumbLink>
        </BreadcrumbItem>
      </Breadcrumb>

      <div className="flex items-center justify-between space-y-2">
        <div>
          <h2 className="text-2xl font-bold tracking-tight">Workspace {workspaceAddr}</h2>
          <p className="text-muted-foreground">
            Created by <Address address={ws.creator} />.
          </p>
        </div>
      </div>

      <WorkspacePolicyCard workspace={ws} />

      <Card>
        <CardHeader>
          <CardTitle>Owners</CardTitle>
          <CardDescription>With default policies, owners will be able to perform actions such as adding other owners or signing transactions.</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="grid w-full items-center gap-4">
            <ul className="flex flex-col space-y-1">
              {ws.owners.map((owner) => (
                <li key={owner} className="list-disc list-inside group">
                  <Address address={owner} />
                  <Button variant="destructive" className="opacity-20 px-2 py-0.5 ml-2 h-auto w-auto group-hover:opacity-100" onClick={() => {
                    broadcast([
                      new MsgRemoveWorkspaceOwner({ creator: addr, workspaceAddr, owner }),
                    ]);
                  }}>
                    X
                  </Button>
                </li>
              ))}
            </ul>
          </div>
        </CardContent>
        <CardFooter>
          <AddWorkspaceOwnerForm addr={addr} workspaceAddr={workspaceAddr} />
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Keys</CardTitle>
          <CardDescription>Keys are used to derive blockchain addresses and sign transactions.</CardDescription>
        </CardHeader>
        <CardContent>
          {
            keyringAddress ? (
              <>
                <Button
                  className="flex flex-col"
                  onClick={() => requestKey(keyringAddress, addr, workspaceAddr)}>
                  <span>
                    Request a new key
                  </span>
                  <span className="text-xs">
                    ({keyringAddress})
                  </span>
                </Button>
                <AlertDialog open={state !== "idle"}>
                  <AlertDialogContent>
                    <AlertDialogHeader>
                      <AlertDialogTitle>New key request</AlertDialogTitle>
                      <AlertDialogDescription>
                        {textForState(state)}
                        {
                          progressForState(state) ? (
                            <Progress value={progressForState(state)} />
                          ) : null
                        }
                        {
                          state === KeyRequesterState.KEY_FULFILLED ? (
                            <div className="mt-4">
                              <Button onClick={() => reset()}>
                                Okay
                              </Button>
                            </div>
                          ) : null
                        }
                        {
                          state === KeyRequesterState.ERROR ? (
                            <div>
                              <p>{error}</p>
                              <div className="mt-4">
                                <Button onClick={() => reset()}>
                                  Okay
                                </Button>
                              </div>
                            </div>
                          ) : null
                        }
                      </AlertDialogDescription>
                    </AlertDialogHeader>
                  </AlertDialogContent>
                </AlertDialog>
              </>
            ) : (
              <Link to={`/keyrings`}>
                <Button>
                  Select a keyring
                </Button>
              </Link>
            )}

          <KeyRequests workspaceAddr={workspaceAddr} />
          <Keys workspaceAddr={workspaceAddr} />
        </CardContent>
      </Card>
    </div>
  );
}

export async function loader({ params }: { params: Params<string> }) {
  if (!params.workspaceAddr) {
    throw new Error("No workspace address provided");
  }
  return {
    workspaceAddr: params.workspaceAddr,
  };
}

export default Workspace;
