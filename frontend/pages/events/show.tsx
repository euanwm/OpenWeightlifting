import ShowEvent from '@/components/organisms/eventPage'

function Event(props: EventParams) {
  return <ShowEvent {...props} />
}

export interface EventParams {
  [key: string]: string
}

export async function getServerSideProps(context: { query: EventParams }) {
  const query = context.query
    return { props: { query } }
}

export default Event