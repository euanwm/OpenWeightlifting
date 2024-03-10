import HeaderBar from '@/components/molecules/head'

function AboutPage() {
  return (
    <div>
      <HeaderBar />
      <h1 className={'text-4xl font-bold text-center mt-8 mb-4'}>
        About OpenWeightlifting
      </h1>
      <div className="text-center mx-4">
        <i className="text-center text-lg mt-4">
          What is our mission?
        </i>
      </div>
      <br></br>
      <div className="text-center mx-4">
        <ul className="text-lg">
          <li>
            <b>Provide an accessible and open results platform</b> for the sport of weightlifting.
          </li>
          <li>
            <b>Make historical weightlifting data accessible</b> to athletes, coaches, and fans.
          </li>
          <li>
            <b>Enable the development of new tools and services</b> for the weightlifting community.
          </li>
          <li>
            <b>Establish a federation</b> that truly represents the interests of athletes and coaches.
          </li>
        </ul>
        <br/>
        <p>
          We are a community-driven project and welcome contributions from anyone who shares our goals. No matter your skill level, there are many ways to get involved. We are always looking for help with data entry, software development, and community outreach.
        </p>
      </div>
    </div>
  )
}

export default AboutPage