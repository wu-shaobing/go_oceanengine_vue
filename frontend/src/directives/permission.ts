import type { Directive } from 'vue'

const getPermissions = (): string[] => {
  const permStr = localStorage.getItem('permissions')
  return permStr ? JSON.parse(permStr) : []
}

export const permissionDirective: Directive<HTMLElement, string | string[]> = {
  mounted(el, binding) {
    const { value } = binding
    const permissions = getPermissions()

    if (value) {
      const requiredPermissions = Array.isArray(value) ? value : [value]
      const hasPermission = requiredPermissions.some(p => permissions.includes(p))

      if (!hasPermission) {
        el.parentNode?.removeChild(el)
      }
    }
  }
}

export const hasPermission = (permission: string | string[]): boolean => {
  const permissions = getPermissions()
  const requiredPermissions = Array.isArray(permission) ? permission : [permission]
  return requiredPermissions.some(p => permissions.includes(p))
}
