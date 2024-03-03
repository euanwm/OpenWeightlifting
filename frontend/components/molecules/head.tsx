import Image from 'next/image'
import {
  Navbar,
  NavbarBrand,
  Link,
  NavbarMenuToggle,
  NavbarMenu,
  NavbarMenuItem,
} from '@nextui-org/react'
import { FaGithub, FaInstagram } from 'react-icons/fa'
import { IoCalendarNumberOutline } from "react-icons/io5";
import { SlCalculator } from 'react-icons/sl'
import { MdOutlinePersonSearch } from 'react-icons/md'
import { FiHome } from 'react-icons/fi'

import Logo from '../../public/OWL-logo.png'
import { useState } from 'react'
import { TbBrandPatreon } from "react-icons/tb";
import { LuBookOpenCheck } from "react-icons/lu";


const HeaderBar = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false)

  return (
    <Navbar isBordered className="py-2">
      <NavbarBrand>
        <Link href="/">
          <Image src={Logo} alt="OpenWeightlifting" width={130} />
        </Link>
      </NavbarBrand>

      <NavbarMenuToggle aria-label={isMenuOpen ? 'Close menu' : 'Open menu'} />

      <NavbarMenu>
        <NavbarMenuItem>
          <Link href="/">
            <FiHome size="30px" className="mt-2"/>
            <span className="ml-2">Home</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="/events">
            <IoCalendarNumberOutline size="30px" className="mt-2"/>
            <span className="ml-2">Recent Events</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="/search">
            <MdOutlinePersonSearch size="30px" />
            <span className="ml-2">Lifter search</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="/rulebook">
            <LuBookOpenCheck size="30px" />
            <span className="ml-2">The Rulebook</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="/sinclair">
            <SlCalculator size="30px" />
            <span className="ml-2">Sinclair Calculator</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="https://www.patreon.com/OpenWeightlifting">
            <TbBrandPatreon size="30px" />
            <span className="ml-2">Support Us</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="https://www.instagram.com/openweightlifting/">
            <FaInstagram size="30px" />
            <span className="ml-2">Instagram</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="https://github.com/euanwm/OpenWeightlifting">
            <FaGithub size="30px" />
            <span className="ml-2">GitHub</span>
          </Link>
        </NavbarMenuItem>
      </NavbarMenu>
    </Navbar>
  )
}

export default HeaderBar
