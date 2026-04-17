# Key Heat

<p align="center">
    <img src="https://raw.githubusercontent.com/zxc7563598/key-heat/main/build/appicon.png" width="100">
</p>

<div align="center">
  <a href="./README.md">English</a>｜<a href="./README.zh-CN.md">简体中文</a>
  <hr width="50%"/>
</div>

<div align="center">
  将你的键盘使用，变成一张可以被观察的热力图
</div>

<p align="center">
    <img src="https://raw.githubusercontent.com/zxc7563598/key-heat/main/assets/Screenshot01.png" width="40%">
    <img src="https://raw.githubusercontent.com/zxc7563598/key-heat/main/assets/Screenshot02.png" width="40%">
</p>

---

Key Heat 是一款跨平台桌面应用，能够将你日常的键盘使用情况转化为交互式的可视化热力图。

它在系统托盘中静默运行，监听所有应用中的每一次按键输入，并将这些数据按“按键维度”进行聚合统计。  
最终，这些数据会被渲染为一张直观的、颜色编码的键盘热力图，让你可以清晰地看到自己的输入习惯。

很多时候，我们每天都在大量打字，但几乎不会意识到：

- 哪些键被高频使用
- 哪些操作已经形成肌肉记忆
- 自己的输入行为是否存在某种偏向

Key Heat 只是把这些“无感行为”，变成了可以被观察的数据。

---

## Features

- 全局键盘监听（跨应用）
- 按键级别的频率统计
- 实时热力图可视化
- 本地数据存储
- 极低性能占用
- 跨平台支持（macOS / Windows）

---

## Download

安装包通过 [GitHub Releases](https://github.com/zxc7563598/key-heat/releases) 提供。

### Windows

- 安装版（推荐）
- 免安装版（解压即用）

### macOS

1. 下载 `.app`
2. 拖入「应用程序」目录
3. 在终端执行以下命令解除系统限制：

```bash
xattr -cr /Applications/key-heat.app
```

## Privacy & Security

Key Heat 的设计原则很明确：**只做统计，不做记录。**

- 不会联网
- 不会上传任何数据
- 不记录具体输入内容
- 不记录按键顺序
- 仅统计「某个按键被按了多少次」

所有数据都保存在本地设备中，不存在外部传输，也不存在隐私收集。

---

## Performance

Key Heat 可以长期在后台运行，并保持非常低的资源占用。

数据刷写采用双重触发机制：

- 基于时间：每 10 秒自动刷新
- 基于计数：每 500 次按键触发一次写入

这种策略在保证数据可靠性的同时，避免了频繁 IO 带来的性能开销。

---

## Development

如果你希望自行构建或参与开发：

### 1. 安装依赖

```bash
go install github.com/wailsapp/wails/v3/cmd/wails3@latest
```

### 2. 启动开发环境

```bash
wails3 dev
```

### 3. 打包应用

```bash
wails3 package
```

---

## Notes

- macOS 需要开启「辅助功能权限（Accessibility）」以监听键盘事件
- Windows 可能会触发系统安全提示（属于正常行为）

---

## About

这个项目的出发点其实很简单。

我们每天都会花很多时间在键盘上，但几乎不会真正去观察“自己是怎么使用键盘的”。  
哪些键被频繁敲击，哪些组合已经成为下意识反应，这些行为通常都是隐形的。

Key Heat 做的事情，只是把这些行为记录下来，并用一种足够直观的方式展示出来。

有时候结果会有点意外。
