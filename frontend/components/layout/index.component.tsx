import Image from 'next/image'
import { Switch, Link } from '@nextui-org/react'
import { FaSun, FaMoon, FaGithub, FaInstagram, FaCalculator } from 'react-icons/fa'

import styles from './layout.module.css'

import Logo from '../../public/OWL-logo.png'

// I have no idea if that's the right way to do it, but it gets rid of the error
const Layout = ({ children }: { children: React.ReactNode }) => {

  return (
      <nav className={styles.navbar}>
        <div className={styles.logo}>
          <Image src={Logo} fill={true} alt='Open Weightlifting' sizes='auto' placeholder='blur' blurDataURL={"/OWL-logo.png"} />
        </div>
        <div className={styles.linkContainer}>
          <Link href="/sinclair"><FaCalculator size='30px' /></Link>
          <a href="https://www.instagram.com/openweightlifting/" target="_blank" rel="noreferrer"><FaInstagram size='30px' /></a>
          <a href="https://github.com/euanwm/OpenWeightlifting" target="_blank" rel="noreferrer"><FaGithub size='30px' /></a>
          <div className={styles.themeSelector}>
            <Switch
              defaultSelected
              startContent={<FaSun />}
              endContent={<FaMoon />}
            />
          </div>
          {children}
        </div>
      </nav>
  )
}

export default Layout