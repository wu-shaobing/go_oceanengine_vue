import type { Directive, DirectiveBinding } from 'vue'

interface LoadingElement extends HTMLElement {
  _loadingMask?: HTMLElement
  _loadingText?: string
}

const createLoadingMask = (text: string = '加载中...'): HTMLElement => {
  const mask = document.createElement('div')
  mask.className = 'loading-mask'
  mask.innerHTML = `
    <div class="loading-mask-content">
      <div class="loading-spinner">
        <svg class="animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
      </div>
      <span class="loading-text">${text}</span>
    </div>
  `

  // 内联样式
  mask.style.cssText = `
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(255, 255, 255, 0.9);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 100;
    transition: opacity 0.2s ease;
  `

  const content = mask.querySelector('.loading-mask-content') as HTMLElement
  if (content) {
    content.style.cssText = `
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 12px;
    `
  }

  const spinner = mask.querySelector('.loading-spinner') as HTMLElement
  if (spinner) {
    spinner.style.cssText = `
      width: 32px;
      height: 32px;
      color: #3b82f6;
    `
  }

  const svg = mask.querySelector('svg') as SVGElement
  if (svg) {
    svg.style.cssText = `
      width: 100%;
      height: 100%;
      animation: spin 1s linear infinite;
    `
  }

  const textEl = mask.querySelector('.loading-text') as HTMLElement
  if (textEl) {
    textEl.style.cssText = `
      font-size: 14px;
      color: #6b7280;
    `
  }

  // 添加动画样式
  if (!document.getElementById('loading-directive-styles')) {
    const style = document.createElement('style')
    style.id = 'loading-directive-styles'
    style.textContent = `
      @keyframes spin {
        from { transform: rotate(0deg); }
        to { transform: rotate(360deg); }
      }
    `
    document.head.appendChild(style)
  }

  return mask
}

const showLoading = (el: LoadingElement, text?: string) => {
  // 确保父元素有相对定位
  const position = getComputedStyle(el).position
  if (position === 'static') {
    el.style.position = 'relative'
  }

  // 如果已存在 loading，先移除
  if (el._loadingMask) {
    el._loadingMask.remove()
  }

  // 创建新的 loading mask
  const mask = createLoadingMask(text || el._loadingText)
  el._loadingMask = mask
  el.appendChild(mask)
}

const hideLoading = (el: LoadingElement) => {
  if (el._loadingMask) {
    el._loadingMask.style.opacity = '0'
    setTimeout(() => {
      el._loadingMask?.remove()
      el._loadingMask = undefined
    }, 200)
  }
}

export const loading: Directive = {
  mounted(el: LoadingElement, binding: DirectiveBinding) {
    // 保存文本配置
    el._loadingText = typeof binding.value === 'string' ? binding.value : binding.arg || '加载中...'
    
    // 如果初始值为 true，显示 loading
    if (binding.value === true || (typeof binding.value === 'object' && binding.value?.loading)) {
      showLoading(el, el._loadingText)
    }
  },
  
  updated(el: LoadingElement, binding: DirectiveBinding) {
    const isLoading = binding.value === true || (typeof binding.value === 'object' && binding.value?.loading)
    const wasLoading = binding.oldValue === true || (typeof binding.oldValue === 'object' && binding.oldValue?.loading)
    
    // 更新文本
    if (typeof binding.value === 'string') {
      el._loadingText = binding.value
    } else if (typeof binding.value === 'object' && binding.value?.text) {
      el._loadingText = binding.value.text
    }
    
    if (isLoading && !wasLoading) {
      showLoading(el, el._loadingText)
    } else if (!isLoading && wasLoading) {
      hideLoading(el)
    }
  },
  
  unmounted(el: LoadingElement) {
    if (el._loadingMask) {
      el._loadingMask.remove()
      el._loadingMask = undefined
    }
  }
}

export default loading
