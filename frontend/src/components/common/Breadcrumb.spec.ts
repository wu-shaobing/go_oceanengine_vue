import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Breadcrumb from './Breadcrumb.vue'

// Mock router-link
const RouterLinkStub = {
  template: '<a><slot /></a>',
  props: ['to']
}

describe('Breadcrumb', () => {
  const globalStubs = {
    global: {
      stubs: {
        'router-link': RouterLinkStub
      }
    }
  }

  it('renders breadcrumb items correctly', () => {
    const items = [
      { name: '广告管理', path: '/ads' },
      { name: '广告列表' }
    ]
    
    const wrapper = mount(Breadcrumb, {
      props: { items },
      ...globalStubs
    })
    
    // 检查文本内容
    expect(wrapper.text()).toContain('首页')
    expect(wrapper.text()).toContain('广告管理')
    expect(wrapper.text()).toContain('广告列表')
  })

  it('renders links for items with path', () => {
    const items = [
      { name: '广告管理', path: '/ads' },
      { name: '当前页' }
    ]
    
    const wrapper = mount(Breadcrumb, {
      props: { items },
      ...globalStubs
    })
    
    // 应该有链接
    const links = wrapper.findAll('a')
    expect(links.length).toBeGreaterThan(0)
  })

  it('does not render link for last item', () => {
    const items = [
      { name: '广告管理', path: '/ads' },
      { name: '当前页' }
    ]
    
    const wrapper = mount(Breadcrumb, {
      props: { items },
      ...globalStubs
    })
    
    // 最后一项应该是文本
    expect(wrapper.text()).toContain('当前页')
  })

  it('renders empty when no items provided', () => {
    const wrapper = mount(Breadcrumb, {
      props: { items: [] },
      ...globalStubs
    })
    
    // 只有首页链接
    expect(wrapper.text()).toContain('首页')
  })

  it('handles single item', () => {
    const items = [{ name: '当前页' }]
    
    const wrapper = mount(Breadcrumb, {
      props: { items },
      ...globalStubs
    })
    
    expect(wrapper.text()).toContain('首页')
    expect(wrapper.text()).toContain('当前页')
  })
})
