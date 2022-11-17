import NativeSelect from '@mui/material/NativeSelect'
import InputLabel from '@mui/material/InputLabel'
import FormControl from '@mui/material/FormControl'

import SearchFilter from '../search/index.component'

import styles from './filters.module.css'

const genderList = [
  { value: 'male', label: 'Male' },
  { value: 'female', label: 'Female' },
]

const sortByList = [
  { value: 'total', label: 'Total' },
  { value: 'sinclair', label: 'Sinclair' },
]

const federationList = [
  { value: 'allfeds', label: 'ALL' },
  { value: 'UK', label: 'UK' },
  { value: 'US', label: 'US' },
  { value: 'AUS', label: 'AUS' },
  { value: 'IWF', label: 'IWF' },
]

const Filters = ({ currentGender, sortBy, federation, handleGenderChange }) => (
  <div className={styles.filterContainer}>
    <div className={styles.filters}>
      <FormControl className={styles.selectContainer}>
        <InputLabel variant="standard" htmlFor="genderSelect">
          Gender
        </InputLabel>
        <NativeSelect
          id="genderSelect"
          label="Gender"
          value={currentGender}
          onChange={e =>
            handleGenderChange({ type: 'gender', value: e.target.value })
          }
        >
          {genderList.map(item => (
            <option key={item.value} value={item.value}>
              {item.label}
            </option>
          ))}
        </NativeSelect>
      </FormControl>

      <FormControl className={styles.selectContainer}>
        <InputLabel variant="standard" htmlFor="sortBySelect">
          Total/Sinclair
        </InputLabel>
        <NativeSelect
          id="sortBySelect"
          label="Total/Sinclair"
          value={sortBy}
          onChange={e =>
            handleGenderChange({ type: 'sortBy', value: e.target.value })
          }
        >
          {sortByList.map(item => (
            <option key={item.value} value={item.value}>
              {item.label}
            </option>
          ))}
        </NativeSelect>
      </FormControl>

      <FormControl className={styles.selectContainer}>
        <InputLabel variant="standard" htmlFor="federationSelect">
          Federation
        </InputLabel>
        <NativeSelect
          id="federationSelect"
          label="Federation"
          value={federation}
          onChange={e =>
            handleGenderChange({ type: 'federation', value: e.target.value })
          }
        >
          {federationList.map(item => (
            <option key={item.value} value={item.value}>
              {item.label}
            </option>
          ))}
        </NativeSelect>
      </FormControl>
    </div>
    <SearchFilter />
  </div>
)

export default Filters
