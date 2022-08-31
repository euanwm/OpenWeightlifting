import Image from 'next/image'
import { Container, Switch, useTheme } from '@nextui-org/react'
import { useTheme as useNextTheme } from 'next-themes'
import { FaSun, FaMoon } from 'react-icons/fa'

import styles from './layout.module.css'

import Logo from '../../public/OWL-logo.png'

const Layout = ({ children }) => {
  const { setTheme, theme } = useNextTheme();
  const { isDark } = useTheme();
  const themeIconClass = theme === 'dark' ? styles.themeIconDark : styles.themeIconLight

  return (
    <Container xl>
      <nav className={styles.navbar}>
        <div className={styles.logo}>
          <Image src={Logo} layout='fill' objectFit='contain' />
        </div>
        <div className={styles.themeSelector}>
          <FaSun size='24px' className={themeIconClass} />
          <Switch
            checked={isDark}
            onChange={(e) => setTheme(e.target.checked ? 'dark' : 'light')}
          />
          <FaMoon size='20px' className={themeIconClass} />
        </div>
      </nav>
      {children}
    </Container>
  )
}

export default Layout