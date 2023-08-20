import Image from 'next/image'
import { Navbar, NavbarBrand, NavbarContent, Switch, Link } from '@nextui-org/react'
import { FaSun, FaMoon, FaGithub, FaInstagram, FaCalculator } from 'react-icons/fa'

import Logo from '../public/OWL-logo.png'

const HeaderBar = () => {
  return (
    <Navbar>
      <title>OpenWeightlifting - The Biggest Weightlifting Database</title>
      <NavbarBrand>
        <Image src={Logo} alt='OpenWeightlifting' width={200}/>
      </NavbarBrand>
      <NavbarContent>
        <Link href="/sinclair"><FaCalculator size='30px' /></Link>
      </NavbarContent>
      <a href="https://www.instagram.com/openweightlifting/"><FaInstagram size='30px' /></a>
      <a href="https://github.com/euanwm/OpenWeightlifting"><FaGithub size='30px' /></a>
        <Switch
          defaultSelected
          startContent={<FaSun />}
          endContent={<FaMoon />}
        />
    </Navbar>
  )
}

export default HeaderBar