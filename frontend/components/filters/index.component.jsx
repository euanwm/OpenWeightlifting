import NativeSelect from '@mui/material/NativeSelect'
import InputLabel from '@mui/material/InputLabel'
import FormControl from '@mui/material/FormControl'

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

const weightClassList = [
  { value: 'allcats', label: 'All Categories' },
  { value: 'M55', label: 'Men\'s 55kg' },
  { value: 'M61', label: 'Men\'s 61kg' },
  { value: 'M67', label: 'Men\'s 67kg' },
  { value: 'M73', label: 'Men\'s 73kg' },
  { value: 'M81', label: 'Men\'s 81kg' },
  { value: 'M89', label: 'Men\'s 89kg' },
  { value: 'M96', label: 'Men\'s 96kg' },
  { value: 'M102', label: 'Men\'s 102kg' },
  { value: 'M109', label: 'Men\'s 109kg' },
  { value: 'M109+', label: 'Men\'s +109kg' },
  { value: 'F45', label: 'Women\'s 45kg' },
  { value: 'F49', label: 'Women\'s 49kg' },
  { value: 'F55', label: 'Women\'s 55kg' },
  { value: 'F59', label: 'Women\'s 59kg' },
  { value: 'F64', label: 'Women\'s 64kg' },
  { value: 'F71', label: 'Women\'s 71kg' },
  { value: 'F76', label: 'Women\'s 76kg' },
  { value: 'F81', label: 'Women\'s 81kg' },
  { value: 'F87', label: 'Women\'s 87kg' },
  { value: 'F87+', label: 'Women\'s +87kg' }
]

const Filters = ({ currentGender, sortBy, federation, handleGenderChange, weightClass }) => (
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

    <FormControl className={styles.selectContainer}>
      <InputLabel variant="standard" htmlFor="weightSelect">
        Weight Class
      </InputLabel>
      <NativeSelect
          id="weightSelect"
          label="Weight Class"
          value={weightClass}
          onChange={e =>
              handleGenderChange({ type: 'weightclass', value: e.target.value })
          }
      >
        {weightClassList.map(item => (
            <option key={item.value} value={item.value}>
              {item.label}
            </option>
        ))}
      </NativeSelect>
    </FormControl>
  </div>
)

export default Filters
