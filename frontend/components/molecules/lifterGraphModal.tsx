import useSWR from 'swr'
import { Modal, ModalContent, ModalHeader, Spinner } from '@nextui-org/react'
import fetchLifterGraphData from '@/api/fetchLifterGraphData/fetchLifterGraphData'
import { LifterGraph } from './lifterGraph'

function LifterGraphModal({
  lifterName,
  onClose,
}: {
  lifterName: string
  onClose: () => void
}) {
  const { data, isLoading } = useSWR(lifterName, fetchLifterGraphData)

  return (
    <Modal
      closeButton
      isOpen={true}
      size={'4xl'}
      onClose={onClose}
      placement={'center'}
    >
      <ModalContent>
        <ModalHeader>{lifterName}</ModalHeader>
        {isLoading ? (
          <div className="w-full h-full z-10 flex justify-center items-center">
            <Spinner size="lg" label="Loading..." />
          </div>
        ) : (
          <LifterGraph lifterHistory={data} setRatio={1} />
        )}
      </ModalContent>
    </Modal>
  )
}

export default LifterGraphModal
