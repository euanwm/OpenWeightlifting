import Image from 'next/image'
import { Container, Switch, useTheme } from '@nextui-org/react'
import { useTheme as useNextTheme } from 'next-themes'
import { FaSun, FaMoon, FaGithub } from 'react-icons/fa'

import styles from './layout.module.css'

import Logo from '../../public/OWL-logo.png'

const Layout = ({ children }) => {
  const { setTheme } = useNextTheme();
  const { isDark } = useTheme();
  const themeIconClass = isDark ? styles.themeIconDark : styles.themeIconLight

  return (
    <Container xl>
      <nav className={styles.navbar}>
        <div className={styles.logo}>
          <Image src={Logo} layout='fill' objectFit='contain' alt='Open Weight Lifting' />
        </div>
        <div className={styles.linkContainer}>
          <a href="https://github.com/euanwm/OpenWeightlifting" target="_blank" rel="noreferrer"><FaGithub size='30px' className={themeIconClass} /></a>
          <div className={styles.themeSelector}>
            <FaSun size='24px' className={themeIconClass} />
            <Switch
              checked={isDark}
              onChange={(e) => setTheme(e.target.checked ? 'dark' : 'light')}
            />
            <FaMoon size='20px' className={themeIconClass} />
          </div>
        </div>
      </nav>
      {children}
    </Container>
  )
}

export default Layout