import styles from './filters.module.css'

const Filters = ({ currentGender, sortBy, selectFederation, handleGenderChange }) => (
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
      <div className={styles.sinclairWrapper}>
        <label htmlFor="sortBy">Total/Sinclair</label>
        <div className={styles.selectContainer}>
            <select className={styles.select} id="sortBy" value={sortBy} onChange={(e) => handleGenderChange(e)}>
                <option value="total">Total</option>
                <option value="sinclair">Sinclair</option>
        </select>
        </div>
      </div>
    </div>
)

export default Filters