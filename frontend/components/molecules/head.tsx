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
import { IoCalendarNumberOutline } from 'react-icons/io5'
import { SlCalculator } from 'react-icons/sl'
import { MdOutlinePersonSearch } from 'react-icons/md'
import { FiHome } from 'react-icons/fi'

import Logo from '../../public/OWL-logo.png'
import OPLLogo from '../../public/opl-logo.png'
import { useState } from 'react'

const HeaderBar = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false)

  return (
    <Navbar isBordered className="py-2">
      <NavbarBrand>
        <Link href="/">
          <Image
            src={Logo}
            alt="OpenWeightlifting"
            height={70}
            priority={true}
          />
        </Link>
      </NavbarBrand>

      <NavbarMenuToggle aria-label={isMenuOpen ? 'Close menu' : 'Open menu'} />

      <NavbarMenu>
        <NavbarMenuItem>
          <Link href="/">
            <FiHome size="30px" className="mt-2" />
            <span className="ml-2">Home</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="/events">
            <IoCalendarNumberOutline size="30px" className="mt-2" />
            <span className="ml-2">Recent Events</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="/search">
            <MdOutlinePersonSearch size="30px" />
            <span className="ml-2">Lifter Search</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link href="/calculator">
            <SlCalculator size="30px" />
            <span className="ml-2">Calculator</span>
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
        <NavbarMenuItem>
          <Link href="https://www.openpowerlifting.org/">
            <Image src={OPLLogo} alt="OpenPowerlifting" height={30} />
            <span className="ml-2">OpenPowerlifting</span>
          </Link>
        </NavbarMenuItem>
        <NavbarMenuItem>
          <Link
            href="https://www.buymeacoffee.com/openweightlifting"
            target="_blank"
            rel="noopener noreferrer"
          >
            <img
              src="https://img.buymeacoffee.com/button-api/?text=Support OpenWeightlifting&emoji=&slug=openweightlifting&button_colour=000000&font_colour=00B0F0&font_family=Inter&outline_colour=00B0F0&coffee_colour=FFDD00"
              alt="Buy us a coffee"
              className="h-12"
            />
          </Link>
        </NavbarMenuItem>
      </NavbarMenu>
    </Navbar>
  )
}

export default HeaderBar
