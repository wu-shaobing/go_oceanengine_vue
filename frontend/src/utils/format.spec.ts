import { describe, it, expect } from 'vitest'

// 模拟格式化工具函数
const formatNumber = (num: number): string => {
  if (num >= 100000000) {
    return (num / 100000000).toFixed(2) + '亿'
  }
  if (num >= 10000) {
    return (num / 10000).toFixed(2) + '万'
  }
  return num.toLocaleString()
}

const formatMoney = (amount: number): string => {
  return '¥' + amount.toFixed(2).replace(/\B(?=(\d{3})+(?!\d))/g, ',')
}

const formatPercent = (value: number): string => {
  return (value * 100).toFixed(2) + '%'
}

const formatDate = (date: Date | string): string => {
  const d = typeof date === 'string' ? new Date(date) : date
  const year = d.getFullYear()
  const month = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
}

const formatDateTime = (date: Date | string): string => {
  const d = typeof date === 'string' ? new Date(date) : date
  const dateStr = formatDate(d)
  const hours = String(d.getHours()).padStart(2, '0')
  const minutes = String(d.getMinutes()).padStart(2, '0')
  const seconds = String(d.getSeconds()).padStart(2, '0')
  return `${dateStr} ${hours}:${minutes}:${seconds}`
}

describe('formatNumber', () => {
  it('formats numbers less than 10000', () => {
    expect(formatNumber(999)).toBe('999')
    expect(formatNumber(1234)).toBe('1,234')
    expect(formatNumber(9999)).toBe('9,999')
  })

  it('formats numbers in 万 (10k+)', () => {
    expect(formatNumber(10000)).toBe('1.00万')
    expect(formatNumber(12345)).toBe('1.23万')
    expect(formatNumber(99999900)).toBe('9999.99万')
  })

  it('formats numbers in 亿 (100M+)', () => {
    expect(formatNumber(100000000)).toBe('1.00亿')
    expect(formatNumber(123456789)).toBe('1.23亿')
  })

  it('handles zero', () => {
    expect(formatNumber(0)).toBe('0')
  })
})

describe('formatMoney', () => {
  it('formats money with currency symbol', () => {
    expect(formatMoney(0)).toBe('¥0.00')
    expect(formatMoney(100)).toBe('¥100.00')
    expect(formatMoney(1234.56)).toBe('¥1,234.56')
    expect(formatMoney(1000000)).toBe('¥1,000,000.00')
  })

  it('rounds to 2 decimal places', () => {
    expect(formatMoney(99.999)).toBe('¥100.00')
    expect(formatMoney(99.994)).toBe('¥99.99')
  })
})

describe('formatPercent', () => {
  it('formats decimal to percentage', () => {
    expect(formatPercent(0)).toBe('0.00%')
    expect(formatPercent(0.5)).toBe('50.00%')
    expect(formatPercent(1)).toBe('100.00%')
    expect(formatPercent(0.1234)).toBe('12.34%')
  })

  it('handles values greater than 1', () => {
    expect(formatPercent(1.5)).toBe('150.00%')
  })
})

describe('formatDate', () => {
  it('formats Date object', () => {
    const date = new Date(2024, 5, 15) // June 15, 2024
    expect(formatDate(date)).toBe('2024-06-15')
  })

  it('formats date string', () => {
    expect(formatDate('2024-01-01')).toBe('2024-01-01')
  })

  it('pads single digit months and days', () => {
    const date = new Date(2024, 0, 5) // January 5, 2024
    expect(formatDate(date)).toBe('2024-01-05')
  })
})

describe('formatDateTime', () => {
  it('formats date with time', () => {
    const date = new Date(2024, 5, 15, 14, 30, 45)
    expect(formatDateTime(date)).toBe('2024-06-15 14:30:45')
  })

  it('pads single digit time components', () => {
    const date = new Date(2024, 0, 1, 9, 5, 3)
    expect(formatDateTime(date)).toBe('2024-01-01 09:05:03')
  })
})

describe('Edge cases', () => {
  it('handles negative numbers in formatNumber', () => {
    expect(formatNumber(-1000)).toBe('-1,000')
  })

  it('handles negative money', () => {
    expect(formatMoney(-100)).toBe('¥-100.00')
  })

  it('handles negative percentages', () => {
    expect(formatPercent(-0.1)).toBe('-10.00%')
  })
})
