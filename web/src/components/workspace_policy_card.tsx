import { Workspace } from "@/proto/fusionchain/identity/workspace_pb";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "./ui/card";
import CardRow from "./card_row";
import { useState } from "react";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { useBroadcaster } from "@/hooks/keplr";
import { useQuery } from "@tanstack/react-query";
import { policyById } from "@/client/policy";
import { useKeplrAddress } from "@/keplr";
import Policy from "./policy";
import { MsgUpdateWorkspace } from "@/proto/fusionchain/identity/tx_pb";

export default function WorkspacePolicyCard({ workspace }: { workspace: Workspace }) {
  const addr = useKeplrAddress();
  const { broadcast } = useBroadcaster();
  const [editMode, setEditMode] = useState(false);

  return (
    <Card>
      <CardHeader>
        <CardTitle>Policies</CardTitle>
        <CardDescription>Policies define who can operate on this workspace or use its keys to generate and sign transactions.</CardDescription>
      </CardHeader>
      {editMode ? (
        <EditCardContent workspace={workspace} onSave={async (adminPolicyId, signPolicyId) => {
          setEditMode(false);
          await broadcast([
            new MsgUpdateWorkspace({
              creator: addr,
              workspaceAddr: workspace.address,
              adminPolicyId: BigInt(adminPolicyId),
              signPolicyId: BigInt(signPolicyId),
            })
          ]);
        }} />
      ) : (
        <ViewCardContent workspace={workspace} onEdit={() => setEditMode(true)} />
      )}
    </Card>
  )
}

function ViewCardContent({ workspace, onEdit }: { workspace: Workspace, onEdit: () => void }) {
  return (
    <>
      <CardContent className="flex flex-col gap-4">
        <CardRow label="Admin policy">
          <PreviewPolicyCard id={workspace.adminPolicyId.toString()} />
        </CardRow>
        <CardRow label="Sign policy">
          <PreviewPolicyCard id={workspace.signPolicyId.toString()} />
        </CardRow>
      </CardContent>
      <CardFooter>
        <Button onClick={onEdit}>Edit</Button>
      </CardFooter>
    </>
  );
}

function EditCardContent({ workspace, onSave }: { workspace: Workspace, onSave: (adminPolicyId: string, signPolicyId: string) => void | Promise<void> }) {
  const [adminPolicyId, setAdminPolicyId] = useState(workspace.adminPolicyId.toString());
  const [signPolicyId, setSignPolicyId] = useState(workspace.signPolicyId.toString());

  return (
    <>
      <CardContent className="flex flex-col gap-6">
        <div className="flex flex-col gap-2">
          <div className="flex flex-row gap-3 items-center">
            <Label className="w-32">Admin policy ID:</Label>
            <Input value={adminPolicyId} onChange={e => setAdminPolicyId(e.target.value)} />
          </div>

          <PreviewPolicyCard id={adminPolicyId} />
        </div>

        <div className="flex flex-col gap-2">
          <div className="flex flex-row gap-3 items-center">
            <Label className="w-32">Sign policy ID:</Label>
            <Input value={signPolicyId} onChange={e => setSignPolicyId(e.target.value)} />
          </div>

          <PreviewPolicyCard id={signPolicyId} />
        </div>
      </CardContent>
      <CardFooter>
        <Button onClick={() => onSave(adminPolicyId, signPolicyId)}>Save</Button>
      </CardFooter>
    </>
  );
}

function PreviewPolicyCard({ id }: { id: string }) {
  const q = useQuery(["policy", id], () => policyById(id), {
    refetchInterval: Infinity,
    retry: false,
  });

  if (id === "0") {
    return (
      <Card>
        <CardHeader>
          <CardDescription>Default policy applied</CardDescription>
        </CardHeader>
      </Card>
    );
  }

  if (q.status === "loading") {
    return (
      <Card>
        <CardHeader>
          <CardDescription>Loading policy #{id}...</CardDescription>
        </CardHeader>
      </Card>
    );
  }

  if (q.status === "error") {
    return (
      <Card>
        <CardHeader>
          <CardDescription>Error loading policy</CardDescription>
        </CardHeader>
      </Card>
    );
  }

  if (!q.data?.policy) {
    return null;
  }

  return (
    <Policy response={q.data.policy} />
  );
}
