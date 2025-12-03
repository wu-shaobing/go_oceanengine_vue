#!/usr/bin/env node

/**
 * 巨量引擎页面批量生成脚本
 * 根据 pages-config.json 自动生成所有缺失页面
 * 
 * 使用方法：
 * node generate-pages.js
 * 或: chmod +x generate-pages.js && ./generate-pages.js
 */

const fs = require('fs');
const path = require('path');

// ============ 配置 ============
const CONFIG_FILE = './pages-config.json';
const OUTPUT_DIR = './';

// ============ HTML 模板 ============

// 列表页模板
const LIST_TEMPLATE = (title, description, moduleName) => `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="${description} - 巨量引擎管理平台">
    <meta name="theme-color" content="#3b82f6">
    <title>${title} - 巨量引擎管理平台</title>
    <link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='.9em' font-size='90'>⚡</text></svg>">
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        ::-webkit-scrollbar { width: 8px; height: 8px; }
        ::-webkit-scrollbar-track { background: #f1f5f9; }
        ::-webkit-scrollbar-thumb { background: #cbd5e1; border-radius: 4px; }
        ::-webkit-scrollbar-thumb:hover { background: #94a3b8; }
    </style>
</head>
<body class="bg-gray-50">
    <!-- Header -->
    <header class="sticky top-0 z-40 bg-white/95 backdrop-blur border-b border-gray-200">
        <div class="flex h-16 items-center justify-between px-6">
            <div class="flex items-center gap-3">
                <div class="flex h-10 w-10 items-center justify-center rounded-lg bg-blue-600">
                    <svg class="h-6 w-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                    </svg>
                </div>
                <div class="hidden md:block">
                    <h1 class="text-lg font-bold text-gray-900">巨量引擎管理平台</h1>
                    <p class="text-xs text-gray-500">Ad Engine Management</p>
                </div>
            </div>
            <div class="flex items-center gap-3">
                <button class="relative flex items-center justify-center w-10 h-10 rounded-lg hover:bg-gray-100">
                    <svg class="h-5 w-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/>
                    </svg>
                </button>
                <div class="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100">
                    <div class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-600 text-white">
                        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                        </svg>
                    </div>
                    <span class="hidden text-sm font-medium md:inline-block">管理员</span>
                </div>
            </div>
        </div>
    </header>

    <div class="flex">
        <aside class="w-64 bg-white border-r border-gray-200 min-h-[calc(100vh-64px)] sticky top-16">
            <nav class="p-4 space-y-6">
                <div>
                    <h3 class="mb-3 px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">${moduleName || '功能模块'}</h3>
                    <ul class="space-y-1">
                        <li>
                            <a href="dashboard.html" class="group flex items-center gap-3 px-3 py-2.5 rounded-lg text-gray-700 hover:bg-gray-100">
                                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
                                </svg>
                                <span>工作台</span>
                            </a>
                        </li>
                    </ul>
                </div>
            </nav>
        </aside>

        <main id="main-content" class="flex-1 p-8">
            <div class="max-w-7xl mx-auto space-y-6">
                <div class="flex items-center text-sm text-gray-500 mb-2">
                    <a href="index.html" class="hover:text-blue-600">首页</a>
                    <svg class="w-4 h-4 mx-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                    </svg>
                    <span class="text-gray-900">${title}</span>
                </div>

                <div class="flex items-center justify-between">
                    <div>
                        <h1 class="text-3xl font-bold text-gray-900">${title}</h1>
                        <p class="mt-2 text-gray-600">${description}</p>
                    </div>
                    <button class="px-4 py-2 text-sm text-white bg-blue-600 rounded-lg hover:bg-blue-700 shadow-sm">
                        <svg class="w-4 h-4 inline mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
                        </svg>
                        新建
                    </button>
                </div>

                <div class="bg-white rounded-lg border border-gray-200 shadow-sm">
                    <div class="p-4 border-b border-gray-200">
                        <div class="flex gap-4">
                            <input type="search" placeholder="搜索..." class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
                            <select class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
                                <option>全部状态</option>
                                <option>启用</option>
                                <option>禁用</option>
                            </select>
                            <button class="px-4 py-2 text-sm text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50">
                                筛选
                            </button>
                        </div>
                    </div>

                    <div class="overflow-x-auto">
                        <table class="w-full">
                            <thead class="bg-gray-50 border-b border-gray-200">
                                <tr>
                                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                        <input type="checkbox" class="w-4 h-4 rounded border-gray-300" />
                                    </th>
                                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">名称</th>
                                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">状态</th>
                                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">创建时间</th>
                                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">操作</th>
                                </tr>
                            </thead>
                            <tbody class="bg-white divide-y divide-gray-200">
                                <tr class="hover:bg-gray-50">
                                    <td class="px-6 py-4 whitespace-nowrap">
                                        <input type="checkbox" class="w-4 h-4 rounded border-gray-300" />
                                    </td>
                                    <td class="px-6 py-4 whitespace-nowrap">
                                        <div class="text-sm font-medium text-gray-900">示例数据 #1</div>
                                        <div class="text-sm text-gray-500">ID: 1234567890</div>
                                    </td>
                                    <td class="px-6 py-4 whitespace-nowrap">
                                        <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">启用</span>
                                    </td>
                                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">2025-11-10 15:30:00</td>
                                    <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                                        <button class="text-blue-600 hover:text-blue-900 mr-3">编辑</button>
                                        <button class="text-red-600 hover:text-red-900">删除</button>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>

                    <div class="px-6 py-4 border-t border-gray-200 flex items-center justify-between">
                        <div class="text-sm text-gray-700">
                            显示 <span class="font-medium">1</span> 到 <span class="font-medium">10</span> 条，共 <span class="font-medium">100</span> 条
                        </div>
                        <div class="flex gap-2">
                            <button class="px-3 py-1 text-sm border border-gray-300 rounded hover:bg-gray-50">上一页</button>
                            <button class="px-3 py-1 text-sm bg-blue-600 text-white rounded">1</button>
                            <button class="px-3 py-1 text-sm border border-gray-300 rounded hover:bg-gray-50">2</button>
                            <button class="px-3 py-1 text-sm border border-gray-300 rounded hover:bg-gray-50">下一页</button>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </div>
</body>
</html>`;

// 表单页模板
const FORM_TEMPLATE = (title, description, moduleName) => `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="${description} - 巨量引擎管理平台">
    <meta name="theme-color" content="#3b82f6">
    <title>${title} - 巨量引擎管理平台</title>
    <link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='.9em' font-size='90'>⚡</text></svg>">
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        ::-webkit-scrollbar { width: 8px; height: 8px; }
        ::-webkit-scrollbar-track { background: #f1f5f9; }
        ::-webkit-scrollbar-thumb { background: #cbd5e1; border-radius: 4px; }
        ::-webkit-scrollbar-thumb:hover { background: #94a3b8; }
    </style>
</head>
<body class="bg-gray-50">
    <header class="sticky top-0 z-40 bg-white/95 backdrop-blur border-b border-gray-200">
        <div class="flex h-16 items-center justify-between px-6">
            <div class="flex items-center gap-3">
                <div class="flex h-10 w-10 items-center justify-center rounded-lg bg-blue-600">
                    <svg class="h-6 w-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                    </svg>
                </div>
                <div class="hidden md:block">
                    <h1 class="text-lg font-bold text-gray-900">巨量引擎管理平台</h1>
                    <p class="text-xs text-gray-500">Ad Engine Management</p>
                </div>
            </div>
            <div class="flex items-center gap-3">
                <div class="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100">
                    <div class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-600 text-white">
                        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                        </svg>
                    </div>
                    <span class="hidden text-sm font-medium md:inline-block">管理员</span>
                </div>
            </div>
        </div>
    </header>

    <div class="flex">
        <aside class="w-64 bg-white border-r border-gray-200 min-h-[calc(100vh-64px)] sticky top-16">
            <nav class="p-4 space-y-6">
                <div>
                    <h3 class="mb-3 px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">${moduleName || '功能模块'}</h3>
                    <ul class="space-y-1">
                        <li>
                            <a href="dashboard.html" class="group flex items-center gap-3 px-3 py-2.5 rounded-lg text-gray-700 hover:bg-gray-100">
                                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
                                </svg>
                                <span>工作台</span>
                            </a>
                        </li>
                    </ul>
                </div>
            </nav>
        </aside>

        <main id="main-content" class="flex-1 p-8">
            <div class="max-w-4xl mx-auto space-y-6">
                <div class="flex items-center text-sm text-gray-500 mb-2">
                    <a href="index.html" class="hover:text-blue-600">首页</a>
                    <svg class="w-4 h-4 mx-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                    </svg>
                    <span class="text-gray-900">${title}</span>
                </div>

                <div>
                    <h1 class="text-3xl font-bold text-gray-900">${title}</h1>
                    <p class="mt-2 text-gray-600">${description}</p>
                </div>

                <div class="bg-white rounded-lg border border-gray-200 shadow-sm">
                    <form class="divide-y divide-gray-200">
                        <div class="p-6">
                            <h2 class="text-lg font-semibold text-gray-900 mb-4">基本信息</h2>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                                <div>
                                    <label class="block text-sm font-medium text-gray-700 mb-2">
                                        名称 <span class="text-red-500">*</span>
                                    </label>
                                    <input type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="请输入名称" />
                                </div>
                                <div>
                                    <label class="block text-sm font-medium text-gray-700 mb-2">
                                        类型 <span class="text-red-500">*</span>
                                    </label>
                                    <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                                        <option>请选择类型</option>
                                        <option>类型A</option>
                                        <option>类型B</option>
                                    </select>
                                </div>
                                <div class="md:col-span-2">
                                    <label class="block text-sm font-medium text-gray-700 mb-2">
                                        描述
                                    </label>
                                    <textarea rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="请输入描述信息"></textarea>
                                </div>
                            </div>
                        </div>

                        <div class="p-6 bg-gray-50 flex justify-end gap-3">
                            <button type="button" class="px-6 py-2 text-sm text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50">
                                取消
                            </button>
                            <button type="submit" class="px-6 py-2 text-sm text-white bg-blue-600 rounded-lg hover:bg-blue-700 shadow-sm">
                                保存
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </main>
    </div>
</body>
</html>`;

// 详情页模板
const DETAIL_TEMPLATE = (title, description, moduleName) => `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="${description} - 巨量引擎管理平台">
    <meta name="theme-color" content="#3b82f6">
    <title>${title} - 巨量引擎管理平台</title>
    <link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='.9em' font-size='90'>⚡</text></svg>">
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        ::-webkit-scrollbar { width: 8px; height: 8px; }
        ::-webkit-scrollbar-track { background: #f1f5f9; }
        ::-webkit-scrollbar-thumb { background: #cbd5e1; border-radius: 4px; }
        ::-webkit-scrollbar-thumb:hover { background: #94a3b8; }
    </style>
</head>
<body class="bg-gray-50">
    <header class="sticky top-0 z-40 bg-white/95 backdrop-blur border-b border-gray-200">
        <div class="flex h-16 items-center justify-between px-6">
            <div class="flex items-center gap-3">
                <div class="flex h-10 w-10 items-center justify-center rounded-lg bg-blue-600">
                    <svg class="h-6 w-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                    </svg>
                </div>
                <div class="hidden md:block">
                    <h1 class="text-lg font-bold text-gray-900">巨量引擎管理平台</h1>
                    <p class="text-xs text-gray-500">Ad Engine Management</p>
                </div>
            </div>
            <div class="flex items-center gap-3">
                <div class="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100">
                    <div class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-600 text-white">
                        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                        </svg>
                    </div>
                    <span class="hidden text-sm font-medium md:inline-block">管理员</span>
                </div>
            </div>
        </div>
    </header>

    <div class="flex">
        <aside class="w-64 bg-white border-r border-gray-200 min-h-[calc(100vh-64px)] sticky top-16">
            <nav class="p-4 space-y-6">
                <div>
                    <h3 class="mb-3 px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">${moduleName || '功能模块'}</h3>
                    <ul class="space-y-1">
                        <li>
                            <a href="dashboard.html" class="group flex items-center gap-3 px-3 py-2.5 rounded-lg text-gray-700 hover:bg-gray-100">
                                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
                                </svg>
                                <span>工作台</span>
                            </a>
                        </li>
                    </ul>
                </div>
            </nav>
        </aside>

        <main id="main-content" class="flex-1 p-8">
            <div class="max-w-5xl mx-auto space-y-6">
                <div class="flex items-center text-sm text-gray-500 mb-2">
                    <a href="index.html" class="hover:text-blue-600">首页</a>
                    <svg class="w-4 h-4 mx-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                    </svg>
                    <span class="text-gray-900">${title}</span>
                </div>

                <div class="flex items-center justify-between">
                    <div>
                        <h1 class="text-3xl font-bold text-gray-900">${title}</h1>
                        <p class="mt-2 text-gray-600">${description}</p>
                    </div>
                    <div class="flex gap-2">
                        <button class="px-4 py-2 text-sm text-white bg-blue-600 rounded-lg hover:bg-blue-700">编辑</button>
                        <button class="px-4 py-2 text-sm text-gray-700 border border-gray-300 rounded-lg hover:bg-gray-50">删除</button>
                    </div>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <div class="text-sm text-gray-500 mb-1">状态</div>
                        <div class="flex items-center gap-2">
                            <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">启用</span>
                        </div>
                    </div>
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <div class="text-sm text-gray-500 mb-1">创建时间</div>
                        <div class="text-lg font-semibold text-gray-900">2025-11-10</div>
                    </div>
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <div class="text-sm text-gray-500 mb-1">更新时间</div>
                        <div class="text-lg font-semibold text-gray-900">2025-11-11</div>
                    </div>
                </div>

                <div class="bg-white rounded-lg border border-gray-200 shadow-sm">
                    <div class="border-b border-gray-200 p-6">
                        <h2 class="text-lg font-semibold text-gray-900">详细信息</h2>
                    </div>
                    <div class="p-6">
                        <dl class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-6">
                            <div>
                                <dt class="text-sm font-medium text-gray-500">名称</dt>
                                <dd class="mt-1 text-sm text-gray-900">示例数据</dd>
                            </div>
                            <div>
                                <dt class="text-sm font-medium text-gray-500">类型</dt>
                                <dd class="mt-1 text-sm text-gray-900">类型A</dd>
                            </div>
                            <div class="md:col-span-2">
                                <dt class="text-sm font-medium text-gray-500">描述</dt>
                                <dd class="mt-1 text-sm text-gray-900">这是一个详细描述示例</dd>
                            </div>
                        </dl>
                    </div>
                </div>
            </div>
        </main>
    </div>
</body>
</html>`;

// 工具页模板 (简化版)
const TOOL_TEMPLATE = (title, description) => LIST_TEMPLATE(title, description, '工具中心');

// 仪表盘模板 (带图表)
const DASHBOARD_TEMPLATE = (title, description, moduleName) => `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="${description} - 巨量引擎管理平台">
    <meta name="theme-color" content="#3b82f6">
    <title>${title} - 巨量引擎管理平台</title>
    <link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='.9em' font-size='90'>⚡</text></svg>">
    <script src="https://cdn.tailwindcss.com"></script>
    <script src="https://cdn.jsdelivr.net/npm/chart.js@4.4.0/dist/chart.umd.min.js"></script>
    <style>
        ::-webkit-scrollbar { width: 8px; height: 8px; }
        ::-webkit-scrollbar-track { background: #f1f5f9; }
        ::-webkit-scrollbar-thumb { background: #cbd5e1; border-radius: 4px; }
        ::-webkit-scrollbar-thumb:hover { background: #94a3b8; }
    </style>
</head>
<body class="bg-gray-50">
    <header class="sticky top-0 z-40 bg-white/95 backdrop-blur border-b border-gray-200">
        <div class="flex h-16 items-center justify-between px-6">
            <div class="flex items-center gap-3">
                <div class="flex h-10 w-10 items-center justify-center rounded-lg bg-blue-600">
                    <svg class="h-6 w-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/></svg>
                </div>
                <div class="hidden md:block">
                    <h1 class="text-lg font-bold text-gray-900">巨量引擎管理平台</h1>
                    <p class="text-xs text-gray-500">Ad Engine Management</p>
                </div>
            </div>
            <div class="flex items-center gap-3">
                <div class="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100">
                    <div class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-600 text-white">
                        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                    </div>
                    <span class="hidden text-sm font-medium md:inline-block">管理员</span>
                </div>
            </div>
        </div>
    </header>

    <div class="flex">
        <aside class="w-64 bg-white border-r border-gray-200 min-h-[calc(100vh-64px)] sticky top-16">
            <nav class="p-4 space-y-6">
                <div>
                    <h3 class="mb-3 px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">${moduleName || '数据报表'}</h3>
                    <ul class="space-y-1">
                        <li><a href="dashboard.html" class="group flex items-center gap-3 px-3 py-2.5 rounded-lg text-gray-700 hover:bg-gray-100"><svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/></svg><span>工作台</span></a></li>
                        <li><a href="reports.html" class="group flex items-center gap-3 px-3 py-2.5 rounded-lg text-gray-700 hover:bg-gray-100"><svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/></svg><span>数据报表</span></a></li>
                    </ul>
                </div>
            </nav>
        </aside>

        <main id="main-content" class="flex-1 p-8">
            <div class="max-w-7xl mx-auto space-y-6">
                <div class="flex items-center text-sm text-gray-500 mb-2">
                    <a href="index.html" class="hover:text-blue-600">首页</a>
                    <svg class="w-4 h-4 mx-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
                    <a href="reports.html" class="hover:text-blue-600">数据报表</a>
                    <svg class="w-4 h-4 mx-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
                    <span class="text-gray-900">${title}</span>
                </div>

                <div class="flex items-center justify-between">
                    <div>
                        <h1 class="text-3xl font-bold text-gray-900">${title}</h1>
                        <p class="mt-2 text-gray-600">${description}</p>
                    </div>
                    <div class="flex gap-2">
                        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
                            <option>最近7天</option>
                            <option>最近30天</option>
                            <option>本月</option>
                        </select>
                        <button class="px-4 py-2 text-sm text-white bg-blue-600 rounded-lg hover:bg-blue-700">导出报表</button>
                    </div>
                </div>

                <!-- 统计卡片 -->
                <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <div class="flex items-center justify-between">
                            <div>
                                <p class="text-sm font-medium text-gray-600">总计数量</p>
                                <p class="text-2xl font-bold text-gray-900">12,845</p>
                                <p class="text-sm text-green-600 mt-1">+12.5%</p>
                            </div>
                            <div class="p-3 bg-blue-100 rounded-lg"><svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/></svg></div>
                        </div>
                    </div>
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <div class="flex items-center justify-between">
                            <div>
                                <p class="text-sm font-medium text-gray-600">活跃数量</p>
                                <p class="text-2xl font-bold text-gray-900">8,234</p>
                                <p class="text-sm text-green-600 mt-1">+8.3%</p>
                            </div>
                            <div class="p-3 bg-green-100 rounded-lg"><svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"/></svg></div>
                        </div>
                    </div>
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <div class="flex items-center justify-between">
                            <div>
                                <p class="text-sm font-medium text-gray-600">转化率</p>
                                <p class="text-2xl font-bold text-gray-900">64.1%</p>
                                <p class="text-sm text-red-600 mt-1">-2.1%</p>
                            </div>
                            <div class="p-3 bg-purple-100 rounded-lg"><svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"/></svg></div>
                        </div>
                    </div>
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <div class="flex items-center justify-between">
                            <div>
                                <p class="text-sm font-medium text-gray-600">平均值</p>
                                <p class="text-2xl font-bold text-gray-900">¥156</p>
                                <p class="text-sm text-gray-500 mt-1">较上期持平</p>
                            </div>
                            <div class="p-3 bg-orange-100 rounded-lg"><svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/></svg></div>
                        </div>
                    </div>
                </div>

                <!-- 图表区域 -->
                <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <h3 class="text-lg font-semibold text-gray-900 mb-4">趋势分析</h3>
                        <div style="height: 300px;"><canvas id="trendChart"></canvas></div>
                    </div>
                    <div class="bg-white rounded-lg border border-gray-200 p-6">
                        <h3 class="text-lg font-semibold text-gray-900 mb-4">分布占比</h3>
                        <div style="height: 300px;"><canvas id="pieChart"></canvas></div>
                    </div>
                </div>

                <!-- 数据明细 -->
                <div class="bg-white rounded-lg border border-gray-200">
                    <div class="p-4 border-b border-gray-200"><h3 class="text-lg font-semibold text-gray-900">数据明细</h3></div>
                    <div class="overflow-x-auto">
                        <table class="w-full">
                            <thead class="bg-gray-50"><tr>
                                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">日期</th>
                                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">指标1</th>
                                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">指标2</th>
                                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">变化</th>
                            </tr></thead>
                            <tbody class="divide-y divide-gray-200">
                                <tr class="hover:bg-gray-50"><td class="px-6 py-4 text-sm text-gray-900">2025-11-11</td><td class="px-6 py-4 text-sm text-gray-900">12,845</td><td class="px-6 py-4 text-sm text-gray-900">8,234</td><td class="px-6 py-4 text-sm text-green-600">+12.5%</td></tr>
                                <tr class="hover:bg-gray-50"><td class="px-6 py-4 text-sm text-gray-900">2025-11-10</td><td class="px-6 py-4 text-sm text-gray-900">11,420</td><td class="px-6 py-4 text-sm text-gray-900">7,820</td><td class="px-6 py-4 text-sm text-green-600">+8.3%</td></tr>
                                <tr class="hover:bg-gray-50"><td class="px-6 py-4 text-sm text-gray-900">2025-11-09</td><td class="px-6 py-4 text-sm text-gray-900">10,890</td><td class="px-6 py-4 text-sm text-gray-900">7,100</td><td class="px-6 py-4 text-sm text-red-600">-2.1%</td></tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </main>
    </div>

    <script>
        const trendCtx = document.getElementById('trendChart').getContext('2d');
        new Chart(trendCtx, {
            type: 'line',
            data: {
                labels: ['11/05', '11/06', '11/07', '11/08', '11/09', '11/10', '11/11'],
                datasets: [{
                    label: '指标1',
                    data: [9800, 11200, 10500, 13200, 10890, 11420, 12845],
                    borderColor: '#3b82f6',
                    backgroundColor: 'rgba(59, 130, 246, 0.1)',
                    fill: true,
                    tension: 0.4
                }, {
                    label: '指标2',
                    data: [6500, 7100, 6800, 7500, 7100, 7820, 8234],
                    borderColor: '#10b981',
                    backgroundColor: 'rgba(16, 185, 129, 0.1)',
                    fill: true,
                    tension: 0.4
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: { legend: { position: 'top' } },
                scales: { y: { beginAtZero: true } }
            }
        });

        const pieCtx = document.getElementById('pieChart').getContext('2d');
        new Chart(pieCtx, {
            type: 'doughnut',
            data: {
                labels: ['分类A', '分类B', '分类C', '其他'],
                datasets: [{ data: [45, 25, 18, 12], backgroundColor: ['#3b82f6', '#10b981', '#f59e0b', '#6b7280'] }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: { legend: { position: 'right' } }
            }
        });
    </script>
</body>
</html>`;

// ============ 模板映射 ============
const TEMPLATES = {
    list: LIST_TEMPLATE,
    form: FORM_TEMPLATE,
    detail: DETAIL_TEMPLATE,
    tool: TOOL_TEMPLATE,
    dashboard: DASHBOARD_TEMPLATE
};

// ============ 主逻辑 ============
function generatePages() {
    console.log('🚀 开始批量生成页面...\n');

    // 读取配置文件
    const config = JSON.parse(fs.readFileSync(CONFIG_FILE, 'utf8'));
    let totalGenerated = 0;
    let skipped = 0;

    // 遍历所有模块和页面
    config.pages.forEach(module => {
        module.pages.forEach(page => {
            const fileName = `${page.name}.html`;
            const filePath = path.join(OUTPUT_DIR, fileName);

            // 检查文件是否已存在
            if (fs.existsSync(filePath)) {
                console.log(`⏭️  跳过（已存在）: ${fileName}`);
                skipped++;
                return;
            }

            // 获取模板函数
            const templateFn = TEMPLATES[page.template];
            if (!templateFn) {
                console.error(`❌ 未知模板类型: ${page.template} (${fileName})`);
                return;
            }

            // 生成HTML内容
            const html = templateFn(page.title, page.description, module.module);

            // 写入文件
            fs.writeFileSync(filePath, html, 'utf8');
            console.log(`✅ 已生成: ${fileName} [${page.template}] - ${page.title}`);
            totalGenerated++;
        });
    });

    console.log(`\n🎉 生成完成！`);
    console.log(`   新增页面: ${totalGenerated} 个`);
    console.log(`   跳过页面: ${skipped} 个`);
    console.log(`   总计应有: ${config.metadata.total_pages} 个页面\n`);
}

// 执行生成
try {
    generatePages();
} catch (error) {
    console.error('❌ 生成失败:', error.message);
    process.exit(1);
}
