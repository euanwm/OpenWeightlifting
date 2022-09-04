import styles from './filters.module.css'

const Filters = ({ currentGender, handleGenderChange }) => {
  // const { setTheme } = useNextTheme();
  // const { isDark } = useTheme();
  // const themeIconClass = isDark ? styles.themeIconDark : styles.themeIconLight

  return (
    <div className={styles.filters}>
      <select value={currentGender} onChange={(e) => handleGenderChange(e)}>
        <option value="male">Male</option>
        <option value="female">Female</option>
      </select>
    </div>
  )
}

export default Filters