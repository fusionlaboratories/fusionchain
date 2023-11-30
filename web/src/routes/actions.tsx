import { useKeplrAddress } from "../keplr";
import Actions from "@/components/actions";

function ActionsPage() {
  const addr = useKeplrAddress();
  if (!addr) {
    return (
      <div className="px-6 mt-10">
        <h1 className="text-lg font-bold">Your actions</h1>
        <p>Connect your wallet to see your actions</p>
      </div>
    );
  }

  // return (
  //   <div className="flex-col flex-1 hidden h-full p-8 space-y-8 md:flex">
  //     <div className="flex items-center justify-between space-y-2">
  //       <div>
  //         <h2 className="text-2xl font-bold tracking-tight">Your actions</h2>
  //         <p className="text-muted-foreground">
  //           Actions that interest you
  //         </p>
  //       </div>
  //     </div>

  //     <Actions />
  //   </div>
  // )

  return (
    <div className="flex flex-col flex-1 h-full px-8 py-4 space-y-8">
      <div className="flex items-center justify-between pb-4 space-y-2 border-b">
        <div>
          <h2 className="text-4xl">Actions</h2>
          <p className="text-muted-foreground">
            Actions that interest you.
          </p>
        </div>
      </div>
      <Actions />
    </div>
  )
}

export default ActionsPage;
