# 🚀 Go URL Checker

![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge)

A lightning-fast URL validator built with Go that checks website availability and provides detailed feedback.

## ✨ Features

- 🎯 **Smart URL Validation** - Detects malformed URLs instantly
- ⚡ **30-Second Timeout** - Never hangs on slow websites
- 🎨 **Beautiful CLI Output** - Uses emojis and colors for better UX
- 🔒 **Error Handling** - Graceful error messages for all edge cases
- 📊 **Status Code Check** - Validates HTTP response codes (2xx, 3xx)

## 🛠️ Installation

```bash
# Clone the repository
git clone https://github.com/TUMHARA_USERNAME/url-checker.git

# Navigate to project
cd url-checker

# Build the project
go build -o urlchecker

# Run it!
./urlchecker
```

## 🎮 Usage

```bash
# Run the program
./urlchecker

# Or directly with go run
go run main.go
```

### Example:
```
Enter your url: https://google.com
✅ Valid & Working URL! (Status: 200)

Enter your url: invalid-url
❌ Format galat hai! Example: https://google.com
```

## 📸 Screenshots

```
╔════════════════════════════════════╗
║      URL CHECKER v1.0              ║
╠════════════════════════════════════╣
║ Enter your url: https://github.com ║
║ ✅ Valid & Working URL! (200)      ║
╚════════════════════════════════════╝
```

## 🧪 Tech Stack

- **Language:** Go 1.20+
- **Libraries:** Standard library only (no external dependencies!)
- **Key Packages:** `net/http`, `net/url`, `bufio`, `time`

## 📝 Code Structure

```
url-checker/
│
├── main.go          # Main application logic
├── README.md        # Project documentation
└── LICENSE          # MIT License
```

## 🚀 Quick Start (One-Liner)

```bash
echo "https://google.com" | go run main.go
```

## 🔥 Why This Project?

- ✅ Zero external dependencies
- ✅ Production-ready error handling
- ✅ Clean and readable code
- ✅ Perfect for learning HTTP clients in Go
- ✅ Beginner-friendly yet professional

## 🤝 Contributing

Contributions are welcome! Feel free to:
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## 📄 License

MIT License - feel free to use, modify, and distribute!

### Made with ❤️ by [Nikhil Khalane]

**⭐ Don't forget to star this repo if you found it useful!**

### Made with ❤️ by [Your Name]

**⭐ Don't forget to star this repo if you found it useful!**
