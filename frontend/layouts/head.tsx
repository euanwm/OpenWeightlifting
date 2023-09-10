import Image from 'next/image'
import { Navbar, NavbarBrand, NavbarContent, Switch, Link, NavbarMenuToggle, NavbarMenu, NavbarMenuItem } from "@nextui-org/react";
import { FaSun, FaMoon, FaGithub, FaInstagram, FaCalculator } from 'react-icons/fa'

import Logo from '../public/OWL-logo.png'
import { useState } from 'react'

const HeaderBar = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  return (
    <Navbar isBordered>
      <title>OpenWeightlifting</title>
      <NavbarBrand>
        <Link href="/">
          <Image src={Logo} alt='OpenWeightlifting' width={130} />
        </Link>
      </NavbarBrand>

      <NavbarMenuToggle
        aria-label={isMenuOpen ? 'Close menu' : 'Open menu'}
      />

      <NavbarMenu>
        <NavbarMenuItem>
          <Link href="/sinclair"><FaCalculator size='30px' />Sinclair Calculator</Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="https://www.instagram.com/openweightlifting/"><FaInstagram size='30px' />Instagram</Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="https://github.com/euanwm/OpenWeightlifting"><FaGithub size='30px' />GitHub</Link>
        </NavbarMenuItem>
      </NavbarMenu>


        <Switch
          startContent={<FaSun />}
          endContent={<FaMoon />}
        />
    </Navbar>
  )
}

export default HeaderBar