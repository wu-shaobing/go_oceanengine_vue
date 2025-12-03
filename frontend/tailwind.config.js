/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  // 启用 safelist 确保动态类名正确生成
  safelist: [
    // 产品线颜色
    'bg-orange-500', 'bg-green-500', 'bg-purple-500', 'bg-yellow-500', 'bg-blue-500', 'bg-gray-500',
    'text-orange-600', 'text-green-600', 'text-purple-600', 'text-yellow-600', 'text-blue-600', 'text-gray-600',
    'border-orange-200', 'border-green-200', 'border-purple-200', 'border-yellow-200', 'border-blue-200', 'border-gray-200',
    'bg-orange-50', 'bg-green-50', 'bg-purple-50', 'bg-yellow-50', 'bg-blue-50', 'bg-gray-50',
    'bg-orange-100', 'bg-green-100', 'bg-purple-100', 'bg-yellow-100', 'bg-blue-100', 'bg-gray-100',
    'hover:border-orange-400', 'hover:border-green-400', 'hover:border-purple-400', 'hover:border-yellow-400', 'hover:border-blue-400', 'hover:border-gray-400',
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#3b82f6',
          50: '#eff6ff',
          100: '#dbeafe',
          200: '#bfdbfe',
          300: '#93c5fd',
          400: '#60a5fa',
          500: '#3b82f6',
          600: '#2563eb',
          700: '#1d4ed8',
          800: '#1e40af',
          900: '#1e3a8a',
        }
      },
      borderRadius: {
        '2xl': '1rem',
        '3xl': '1.5rem',
      },
      boxShadow: {
        'card': '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)',
        'card-hover': '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
      },
      spacing: {
        '18': '4.5rem',
        '22': '5.5rem',
      }
    },
  },
  plugins: [],
}
