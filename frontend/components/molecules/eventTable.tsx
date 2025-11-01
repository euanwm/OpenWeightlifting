import {
  Table,
  TableHeader,
  TableColumn,
  TableCell,
  TableRow,
  TableBody, Link,
} from '@nextui-org/react'
import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'
import { useState } from 'react'

export const EventTable = ({ 
                               history,
                              }: {
                                history: LifterResult[]
                              }) => {
  const [sortKey, setSortKey] = useState<'bodyweight' | 'total' | 'sinclair' | null>(null)
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('asc')

  const sortedHistory = [...history].sort((a, b) => {
    if (!sortKey) return 0
    const aValue = a[sortKey] ?? 0
    const bValue = b[sortKey] ?? 0
    return sortOrder === 'asc' ? aValue - bValue : bValue - aValue
  })

  const handleSort = (key: 'bodyweight' | 'total' | 'sinclair') => {
    if (sortKey === key) {
      setSortOrder(sortOrder === 'asc' ? 'desc' : 'asc')
    } else {
      setSortKey(key)
      setSortOrder('asc')
    }
  }

  const renderSortIcon = (key: 'bodyweight' | 'total' | 'sinclair') => {
    if (sortKey !== key) return null
    return sortOrder === 'asc' ? '▲' : '▼'
  }
  

  const renderSortableHeader = (label: string, key: 'bodyweight' | 'total' | 'sinclair') => (
    <div
      onClick={() => handleSort(key)}
      style={{
        display: 'flex',
        alignItems: 'center',
        cursor: 'pointer',
        userSelect: 'none',
        fontWeight: sortKey === key ? 'bold' : 'normal',
        backgroundColor: sortKey === key ? 'rgba(0, 123, 255, 0.1)' : 'transparent',
        padding: '4px 8px',
        borderRadius: 4,
      }}
    >
      {label}
      {renderSortIcon(key)}
    </div>
  )

  return (
    <Table>
      <TableHeader>
        <TableColumn>Date</TableColumn>
        <TableColumn>Name</TableColumn>
        <TableColumn>{renderSortableHeader('Bodyweight', 'bodyweight')}</TableColumn>
        <TableColumn>1st Snatch</TableColumn>
        <TableColumn>2nd Snatch</TableColumn>
        <TableColumn>3rd Snatch</TableColumn>
        <TableColumn>1st C&J</TableColumn>
        <TableColumn>2nd C&J</TableColumn>
        <TableColumn>3rd C&J</TableColumn>
        <TableColumn>{renderSortableHeader('Total', 'total')}</TableColumn>
        <TableColumn>{renderSortableHeader('Sinclair', 'sinclair')}</TableColumn>
      </TableHeader>
      <TableBody>
        {sortedHistory.map((lift, index) => {
          const {
            date,
            lifter_name,
            bodyweight,
            snatch_1,
            snatch_2,
            snatch_3,
            cj_1,
            cj_2,
            cj_3,
            total,
            sinclair,
          } = lift

          const lifter_page = '../lifter?name=' + lifter_name

          return (
            <TableRow key={`history-${index}`}>
              <TableCell>{date}</TableCell>
              <TableCell>
                <Link href={lifter_page}>
                {lifter_name}
                </Link>
              </TableCell>
              <TableCell>{bodyweight}</TableCell>
              <TableCell>{snatch_1}</TableCell>
              <TableCell>{snatch_2}</TableCell>
              <TableCell>{snatch_3}</TableCell>
              <TableCell>{cj_1}</TableCell>
              <TableCell>{cj_2}</TableCell>
              <TableCell>{cj_3}</TableCell>
              <TableCell>{total}</TableCell>
              <TableCell>{sinclair}</TableCell>
            </TableRow>
          )
        })}
      </TableBody>
    </Table>
  )
}
