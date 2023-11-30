import { useKeplrAddress } from "../keplr";
import { useQuery } from "@tanstack/react-query";
import { balances } from "../client/bank";
import FaucetButton from "./faucet_button";
import { PlusIcon, XMarkIcon } from '@heroicons/react/24/outline'
import { Fragment, useState } from 'react'
import { Dialog, Transition } from '@headlessui/react'
import { Button } from "./ui/button";

function AccountInfo() {

  const [open, setOpen] = useState(false)

  const addr = useKeplrAddress();
  const bq = useQuery({ queryKey: ["balances", addr], queryFn: () => balances(addr) });
  const nqrdo = bq.data?.balances.find((b) => b.denom === "nQRDO")?.amount || "0";
  const qrdo = parseInt(nqrdo) / 10 ** 9;

  // return (
  //   <div className="flex flex-row items-center gap-6">
  //     <div className="flex flex-col items-end">
  //       <span className="text-sm font-monospace">{addr}</span>
  //       <span className="text-xs">{qrdo.toFixed(2)} QRDO</span>
  //     </div>
  //     <FaucetButton />
  //   </div>
  // );
  return (
    <>
      <button
        type="button"
        onClick={() => setOpen(true)}
        className="flex items-center justify-between w-full p-2 space-x-2 text-left border border-gray-300 rounded-full group hover:bg-gray-50"
      >
        <span className="flex items-center flex-1 min-w-0 space-x-2">
          <span className="flex-shrink-0 block">
            <span className="block w-10 h-10 rounded-full bg-primary"></span>
          </span>
          <span className="flex-1 block min-w-0">
            <span className="block text-sm text-gray-900 truncate font-display">{addr}</span>
            <span className="block text-sm text-gray-500 truncate font-display">{qrdo.toFixed(2)} QRDO</span>
          </span>
        </span>
        <span className="inline-flex items-center justify-center flex-shrink-0 w-10 h-10">
          <PlusIcon className="w-5 h-5 text-gray-400 group-hover:text-gray-500" aria-hidden="true" />
        </span>
      </button>

      <Transition.Root show={open} as={Fragment}>
        <Dialog as="div" className="relative z-50" onClose={setOpen}>
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0 transition-opacity bg-white bg-opacity-75 backdrop-blur-sm" />
          </Transition.Child>

          <div className="fixed inset-0 w-screen overflow-y-auto">
            <div className="flex items-end justify-center min-h-full p-4 text-center sm:items-center sm:p-0">
              <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                enterTo="opacity-100 translate-y-0 sm:scale-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              >
                <Dialog.Panel className="relative px-4 pt-5 pb-4 overflow-hidden text-left transition-all transform bg-white border shadow-xl rounded-2xl sm:my-8 sm:w-full sm:max-w-2xl sm:p-14">
                  <div className="absolute top-0 right-0 hidden pt-4 pr-4 sm:block">
                    <button
                      type="button"
                      className="text-gray-400 bg-white rounded-md hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                      onClick={() => setOpen(false)}
                    >
                      <span className="sr-only">Close</span>
                      <XMarkIcon className="w-6 h-6" aria-hidden="true" />
                    </button>
                  </div>
                  <div>
                    {/* <div className="flex items-center justify-center w-12 h-12 mx-auto rounded-full bg-primary">
                      <CheckIcon className="w-6 h-6 text-green-600" aria-hidden="true" />
                    </div> */}
                    <div className="mt-3 text-center sm:mt-5">
                      <Dialog.Title as="h3" className="text-xl leading-6 font-display">
                        {addr}
                      </Dialog.Title>
                      <div className="mt-2">
                        <p className="text-lg tracking-normal text-gray-500 font-display">
                          {qrdo.toFixed(2)} QRDO
                        </p>
                      </div>
                    </div>
                  </div>
                  <div className="flex mt-5 sm:mt-6 place-content-center">
                    <FaucetButton />
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition.Root>
    </>
  )
}

export default AccountInfo;
