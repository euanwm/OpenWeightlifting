import styles from './filters.module.css'

const Filters = ({ currentGender, handleGenderChange }) => (
  <div className={styles.filters}>
    <div className={styles.genderWrapper}>
      <label htmlFor="genderSelect">Gender</label>
      <div className={styles.selectContainer}>
        <select className={styles.select} id="genderSelect" value={currentGender} onChange={(e) => handleGenderChange(e)}>
          <option value="male">Male</option>
          <option value="female">Female</option>
        </select>
      </div>
    </div>
  </div>
)

export default Filters