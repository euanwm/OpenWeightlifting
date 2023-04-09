import Image from 'next/image'
import { Container, Switch, useTheme } from '@nextui-org/react'
import { useTheme as useNextTheme } from 'next-themes'
import { FaSun, FaMoon, FaGithub, FaDiscord } from 'react-icons/fa'

import styles from './layout.module.css'

import Logo from '../../public/OWL-logo.png'

// I have no idea if that's the right way to do it, but it gets rid of the error
const Layout = ({ children }: { children: React.ReactNode }) => {
  console.log('Rendering Layout: ', children)
  const { setTheme } = useNextTheme();
  const { isDark } = useTheme();
  const themeIconClass = isDark ? styles.themeIconDark : styles.themeIconLight

  return (
    <Container xl>
      <nav className={styles.navbar}>
        <div className={styles.logo}>
          <Image src={Logo} fill={true} alt='Open Weightlifting' sizes='auto' placeholder='blur' blurDataURL={"/OWL-logo.png"} />
        </div>
        <div className={styles.linkContainer}>
          <a href="https://discord.gg/kqnBqdktgr" target="_blank" rel="noreferrer"><FaDiscord size='30px' className={themeIconClass} /></a>
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