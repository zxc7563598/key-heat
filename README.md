# Key Heat

<p align="center">
    <img src="https://raw.githubusercontent.com/zxc7563598/key-heat/main/build/appicon.png" width="100">
</p>

<div align="center">
  <a href="./README.md">English</a>｜<a href="./README.zh-CN.md">简体中文</a>
  <hr width="50%"/>
</div>

<div align="center">
  Transform your keyboard into an observable heatmap
</div>

<p align="center">
    <img src="https://raw.githubusercontent.com/zxc7563598/key-heat/main/assets/Screenshot01.png" width="40%">
    <img src="https://raw.githubusercontent.com/zxc7563598/key-heat/main/assets/Screenshot02.png" width="40%">
</p>

---

Key Heat is a cross platform desktop application that can transform your daily keyboard usage into an interactive visual heatmap.

It runs silently in the system tray, listening to every keystroke input in all applications, and aggregates this data according to the "key dimension" for statistical analysis.  
Ultimately, these data will be rendered as an intuitive, color coded keyboard heatmap, allowing you to clearly see your input habits.
Many times, we type a lot every day, but we hardly realize:

- Which keys are frequently used
- What operations have formed muscle memory
- Is there any bias in one's input behavior

Key Heat simply turns these 'emotionless behaviors' into observable data.

---

## Features

- Global keyboard monitoring (cross application)
- Frequency statistics at the key level
- Real time visualization of heat map
- Local data storage
- Extremely low performance occupancy
- Cross platform support (macOS/Windows)

---

## Download

The installation package is provided through [GitHub Releases](https://github.com/zxc7563598/key-heat/releases).

### Windows

- Installation version (recommended)
- Free installation version (ready to use after decompression)

### macOS

1. Download `. app`
2. Drag into the 'Applications' directory
3. Execute the following command on the terminal to lift system restrictions:

```bash
xattr -cr /Applications/key-heat.app
```

## Privacy & Security

The design principle of Key Heat is very clear: \* \* only for statistics, not for records. \*\*

- Not able to connect to the internet
- Will not upload any data
- Do not record specific input content
- Do not record the order of key presses
- Only count 'how many times a certain button has been pressed'

All data is stored on local devices, without external transmission or privacy collection.

---

## Performance

Key Heat can run in the background for a long time and maintain very low resource usage.

Data flashing adopts a dual triggering mechanism:

- Based on time: Automatically refresh every 10 seconds
- Based on counting: write triggered every 500 keystrokes

This strategy ensures data reliability while avoiding the performance overhead caused by frequent IO.

---

## Development

If you wish to build or participate in development on your own:

### 1. Install dependencies

```bash
go install github.com/wailsapp/wails/v3/cmd/wails3@latest
```

### 2. Start the development environment

```bash
wails3 dev
```

### 3. package the application

```bash
wails3 package
```

---

## Notes

- MacOS needs to enable Accessibility to listen for keyboard events
- Windows may trigger system security prompts (which is a normal behavior)

---

## About

The starting point of this project is actually quite simple.

We spend a lot of time on the keyboard every day, but we hardly really observe how we use it.  
Which keys are frequently tapped and which combinations have become subconscious reactions, these behaviors are usually invisible.

Key Heat simply records these behaviors and presents them in a sufficiently intuitive way.

Sometimes the results can be a bit unexpected.
