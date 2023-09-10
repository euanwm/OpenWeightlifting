import Image from 'next/image'
import { Navbar, NavbarBrand, Link, NavbarMenuToggle, NavbarMenu, NavbarMenuItem } from "@nextui-org/react";
import { FaGithub, FaInstagram, FaCalculator } from 'react-icons/fa'

import Logo from '../public/OWL-logo.png'
import { useState } from 'react'

const HeaderBar = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  return (
    <Navbar isBordered>
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
    </Navbar>
  )
}

export default HeaderBar