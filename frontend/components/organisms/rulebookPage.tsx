import HeaderBar from '@/components/molecules/head'

import {Accordion, AccordionItem} from "@nextui-org/react";

function RulebookPage() {
  return (
    <div>
      <HeaderBar />
      <h1 className={'text-4xl font-bold text-center mt-8 mb-4'}>
        The Competition Rulebook
      </h1>
      <div className="text-center mx-4">
        <i className="text-center text-lg mt-4">
          This is the start of the Federation of OpenWeightlifting, casually referred to as FOWL.
        </i>
      </div>
      <br></br>
      <div className="text-center mx-4">
        <p className="text-lg">
          We are currently working on a full version of a Competition Rulebook required for Olympic Weightlifting Gyms or Clubs to host
          their own competitions and for the results of those competitions to be at a suitable standard for submission
          to the OpenWeightlifting database. These rules are based on the IWF and British Weightlifting Technical and Competition Rules & Regulations, however they have been reduced significantly to encourage participation and increase spectator satisfaction. Lastly, we are not able to allow any athlete currently serving a sport-related ban or sanction to compete in any FOWL competition.
        </p>
      </div>
      <br></br>
      <h1 className="text-2xl font-bold text-center mt-8 mb-4">
        What's the main differences between the IWF / BWL and FOWL rules?
      </h1>
      <div className="mx-4">
        <Accordion>
          <AccordionItem title="The Weigh-In" subtitle="Revised">
            <p>
              Reverting back to a simpler time, a lifter is no longer required to weigh-in in a singlet. It is up to their personal discretion, and the discretion of the organisers, as to what level of dressed is appropriate for weigh-in. Alternatively, organisers may also be allowed to deduct 1kg from the bodyweight of a fully dressed lifter, without shoes.
            </p>
          </AccordionItem>
          <AccordionItem title="The Press-Out Rule" subtitle="Revised">
            <p>
              The press-out rule has been revised and elbow movement is allowed up to an angle of 30 degrees. In layman terms, this is the angle between numbers on a clock face from 12 to 1.
            </p>
          </AccordionItem>
          <AccordionItem title="The Jury" subtitle="Revised">
            <p>
              After much debate, the power of the jury has been reduced. The jury will only be able to participate by the request of the centre referee should a decision be too close to call. The jury will not be able to overturn a decision made by the centre referee. A split jury defaults to being in-favour of a good lift result. Ultimately, the jury are there to support the centre referee and not to undermine them.
            </p>
          </AccordionItem>
          <AccordionItem title="Commentator Sheets & Venue Music" subtitle="Encouraged">
            <p>
              At the discretion of the organiser, commentator sheets may be used to provide the audience with more information about the lifters and the competition. This is not a requirement, but it is encouraged. Lifters may also request a particular walk-out song to be played before and/or during their attempts.
            </p>
          </AccordionItem>
          <AccordionItem title="Anti-Doping" subtitle="Removed">
            <p>
              The anti-doping section has been removed from the rulebook. While we do not condone the use of performance enhancing drugs, we believe that an anti-doping policy within Olympic Weightlifting leads to a culture of bribery and corruption.
            </p>
          </AccordionItem>
        </Accordion>
      </div>
      <br></br>
      <div className="text-center mx-4">
        <i className="text-lg">
          We hope to have the full rulebook available for download in the coming months. In the meantime, if you have any questions or concerns, please do not hesitate to contact us at support@openweightlifting.org
        </i>
      </div>
    </div>
  )}

export default RulebookPage