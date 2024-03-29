import { MsgRemoveKeyringParty } from "@/proto/fusionchain/identity/tx_pb";
import { useState } from "react";
import { Button } from "./ui/button";
import { useBroadcaster } from "@/hooks/keplr";

function RemoveKeyringPartyForm({ addr, keyringAddr }: { addr: string, keyringAddr: string }) {
  const { broadcast } = useBroadcaster();
  const [newParty, setNewParty] = useState("");

  return (
    <div className="flex flex-row items-center mt-4 gap-2">
      <input className="px-3 py-2 border border-slate-200 rounded-lg" type="text" placeholder="Remove party" value={newParty} onChange={(v) => setNewParty(v.target.value)} />
      <Button onClick={async () => {
        await broadcast([
          new MsgRemoveKeyringParty({ creator: addr, keyringAddr, party: newParty }),
        ]);
        setNewParty("");
      }}>
        Remove
      </Button>
    </div>
  );
}

export default RemoveKeyringPartyForm;
