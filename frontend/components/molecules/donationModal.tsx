import { Modal, ModalBody, ModalContent, ModalHeader } from '@nextui-org/react'
import posthog from 'posthog-js'
import { useState } from 'react'

export default function DonationModal() {
  let [modalOpen, setModalOpen] = useState(true)

  const handleClick = () => {
    posthog.capture('buy_me_a_coffee_visited', { link_name: 'BMAC Link' })
  }

  return (
    <Modal
      isOpen={modalOpen}
      onOpenChange={setModalOpen}
      className="border border-[#00B0F0]"
    >
      <ModalContent>
        <ModalHeader>
          <h2 className="text-2xl font-bold">Support OpenWeightlifting</h2>
        </ModalHeader>
        <ModalBody>
          <div className="flex flex-col items-center space-y-4">
            <p className="text-lg text-center">
              If you find our OpenWeightlifting useful, consider supporting us
              on Buy Me a Coffee.
            </p>
            <a
              href="https://www.buymeacoffee.com/openweightlifting"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-block"
              onClick={handleClick}
            >
              <img
                src="https://img.buymeacoffee.com/button-api/?text=Buy us a coffee&emoji=&slug=openweightlifting&button_colour=00B0F0&font_colour=000000&font_family=Cookie&outline_colour=000000&coffee_colour=FFDD00"
                alt="Buy us a coffee"
                className="h-12 w-auto"
                width="217"
                height="60"
              />
            </a>
          </div>
        </ModalBody>
      </ModalContent>
    </Modal>
  )
}
