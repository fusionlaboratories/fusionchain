import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { useLoaderData } from "react-router";
import { Params } from "react-router-dom";
import { useKeplrAddress } from "../keplr";
import { useQuery } from "@tanstack/react-query";
import Address from "../components/address";
import { Breadcrumb, BreadcrumbItem, BreadcrumbLink } from "@/components/ui/breadcrumb";
import { keyringByAddress } from "@/client/identity";
import AddKeyringPartyForm from "@/components/add_keyring_party_form";
import CardRow from "@/components/card_row";

function Keyring() {
  const addr = useKeplrAddress();
  const { keyringAddr } = useLoaderData() as Awaited<ReturnType<typeof loader>>;
  const krQuery = useQuery({
    queryKey: ["keyring", keyringAddr],
    queryFn: () => keyringByAddress(keyringAddr)
  });
  const kr = krQuery.data?.keyring;

  if (!kr) {
    return (
      <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
        <div className="flex items-center justify-between space-y-2">
          <div>
            <h2 className="text-2xl font-bold tracking-tight">Keyring {keyringAddr} not found</h2>
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
        <BreadcrumbItem>
          <BreadcrumbLink href="/keyrings">Keyrings</BreadcrumbLink>
        </BreadcrumbItem>
        <BreadcrumbItem isCurrentPage>
          <BreadcrumbLink href={`/keyrings/${keyringAddr}`}>{keyringAddr}</BreadcrumbLink>
        </BreadcrumbItem>
      </Breadcrumb>

      <div className="flex items-center justify-between space-y-2">
        <div>
          <h2 className="text-2xl font-bold tracking-tight">Keyring {keyringAddr}</h2>
          <p className="text-muted-foreground">
            Created by <Address address={kr.creator} />.
          </p>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Status</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid w-full items-center gap-4">
            <CardRow label="Description">
              {kr.description}
            </CardRow>

            <CardRow label="Creator">
              <Address address={kr.creator} />
            </CardRow>

            <CardRow label="Active">
              {kr.isActive ? <span className="font-bold text-green-600">Active</span> : "Inactive"}
            </CardRow>

            <CardRow label="Admins">
              <ul>
                {kr.admins.map((admin) => (
                  <li key={admin}>
                    <Address address={admin} />
                  </li>
                ))}
              </ul>
            </CardRow>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Parties</CardTitle>
          <CardDescription>These accounts will be allowed to respond to keys and signatures requests for this keyring.</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="grid w-full items-center gap-4">
            <ul className="flex flex-col space-y-1">
              {kr.parties.map((p) => (
                <li key={p} className="list-disc list-inside group">
                  <Address address={p} />
                  {/** this tx doesn't exist yet
                  <Button variant="destructive" className="opacity-20 px-2 py-0.5 ml-2 h-auto w-auto group-hover:opacity-100" onClick={() => {
                    broadcast([
                    ]);
                  }}>
                    X
                  </Button>
                  **/ }
                </li>
              ))}
            </ul>
          </div>
        </CardContent>
        <CardFooter>
          <AddKeyringPartyForm addr={addr} keyringAddr={keyringAddr} />
        </CardFooter>
      </Card>
    </div>
  );
}

export async function loader({ params }: { params: Params<string> }) {
  if (!params.keyringAddr) {
    throw new Error("No keyring address provided");
  }
  return {
    keyringAddr: params.keyringAddr,
  };
}

export default Keyring;
