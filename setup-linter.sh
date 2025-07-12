#!/bin/bash

# Setup script for Go linting and formatting
# This script installs golangci-lint and configures your Go project

set -e

echo "🚀 Setting up Go linting and formatting..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

# Install golangci-lint
echo "📦 Installing golangci-lint..."
if ! command -v golangci-lint &> /dev/null; then
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
    echo "✅ golangci-lint installed successfully"
else
    echo "✅ golangci-lint is already installed"
fi

# Install goimports
echo "📦 Installing goimports..."
if ! command -v goimports &> /dev/null; then
    go install golang.org/x/tools/cmd/goimports@latest
    echo "✅ goimports installed successfully"
else
    echo "✅ goimports is already installed"
fi

# Check if GOPATH/bin is in PATH
if [[ ":$PATH:" != *":$(go env GOPATH)/bin:"* ]]; then
    echo "⚠️  Warning: $(go env GOPATH)/bin is not in your PATH"
    echo "Add the following line to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    echo "export PATH=\$PATH:$(go env GOPATH)/bin"
fi

# Run initial formatting
echo "🔧 Running initial code formatting..."
if command -v goimports &> /dev/null; then
    goimports -w .
    echo "✅ Code formatted with goimports"
fi

if command -v gofmt &> /dev/null; then
    gofmt -w .
    echo "✅ Code formatted with gofmt"
fi

# Run linter check
echo "🔍 Running linter check..."
if command -v golangci-lint &> /dev/null; then
    golangci-lint run || echo "⚠️  Found linting issues. Run 'make lint-fix' to auto-fix some of them."
else
    echo "❌ golangci-lint not found in PATH"
fi

echo ""
echo "🎉 Setup complete! Here's what you can do now:"
echo ""
echo "📋 Available commands:"
echo "  make fmt              - Format your code"
echo "  make lint             - Run linter"
echo "  make lint-fix         - Auto-fix linting issues"
echo "  make check-line-length - Check line length specifically"
echo ""
echo "🔧 VS Code/Cursor Integration:"
echo "  - Install the Go extension if not already installed"
echo "  - The .vscode/settings.json file has been configured"
echo "  - Lines longer than 100 characters will be highlighted"
echo "  - Code will be automatically formatted on save"
echo ""
echo "📏 Line Length Enforcement:"
echo "  - Maximum line length is set to 100 characters"
echo "  - The 'lll' linter will flag lines that are too long"
echo "  - Editor ruler at 100 characters will be visible"
echo ""
echo "Happy coding! 🎯" 