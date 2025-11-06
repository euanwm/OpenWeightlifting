import { Modal, ModalBody, ModalContent, ModalHeader } from '@nextui-org/react'
import posthog from 'posthog-js'
import { useState } from 'react'

export default function DonationModal() {
  const randomNumber = Math.floor(Math.random() * 10) + 1
  console.log(randomNumber)
  let [modalOpen, setModalOpen] = useState(randomNumber === 8)

  const handleClick = () => {
    posthog.capture('early_access_visited', { link_name: 'Alpha Site Link' })
  }

  return (
    <Modal
      isOpen={modalOpen}
      onOpenChange={setModalOpen}
      className="border border-[#00B0F0]"
    >
      <ModalContent>
        <ModalHeader className="justify-center">
          <h2 className="text-2xl font-bold">You're an Alpha (tester)</h2>
        </ModalHeader>
        <ModalBody>
          <div className="flex flex-col items-center space-y-4">
            <p className="text-lg text-center">
              Be among the first to break stuff on the new version of
              OpenWeightlifting. Try our alpha version and help shape the future
              of Olympic Weightlifting data.
            </p>
            <a
              href="https://alpha.openweightlifting.org"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-block bg-[#00B0F0] hover:bg-[#0099D6] text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200"
              onClick={handleClick}
            >
              Take me to the new version!
            </a>
          </div>
        </ModalBody>
      </ModalContent>
    </Modal>
  )
}
