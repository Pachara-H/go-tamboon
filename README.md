# Go-Tamboon Donation System

Command-line application ที่ประมวลผลการบริจาคจากไฟล์ CSV ที่เข้ารหัสด้วย ROT-128 และสร้าง charges ผ่าน Omise API

## Features

- 🔐 ROT-128 decryption สำหรับไฟล์ CSV
- 💳 ประมวลผลการชำระเงินผ่าน Omise API
- 🚀 Concurrent processing ด้วย goroutines
- 🛡️ Secure memory management สำหรับข้อมูล sensitive
- ⚡ Rate limiting และ exponential backoff
- 📊 Summary report

## Project Structure

```
go-tamboon/
├── main.go                          # Entry point
├── go.mod                           # Go module
├── .env                             # Environment variables
├── test_custom_1000.csv.rot128      # Test data
├── pkg/                             # Public packages
│   ├── constants/                   # Application constants
│   ├── errors/                      # Error handling
│   └── utilities/                   # Utility functions
└── internal/                        # Private packages
    ├── adapter/                     # External adapters
    │   ├── csv/                     # CSV parsing
    │   └── omise/                   # Omise API client
    ├── cipher/                      # ROT-128 encryption/decryption
    ├── configs/                     # Configuration management
    ├── domains/                     # Domain entities
    │   └── entities/                # Business entities
    ├── errorcode/                   # Error codes
    ├── reporter/                    # Report generation
    ├── services/                    # Application services
    └── validator/                   # Data validation
```

## Installation

```bash
# Install dependencies
go mod tidy

# Build and install
go install .
```

## Configuration

สร้างไฟล์ `.env`:

```bash
# Omise API keys (Base64 encoded)
OMISE_PUBLIC_KEY=<base64-encoded-public-key>
OMISE_SECRET_KEY=<base64-encoded-secret-key>

# Optional
ENV=development
TZ=Asia/Bangkok
```

## Usage

```bash
# Run the application
$GOPATH/bin/go-tamboon <csv-file-path>

# Example
go-tamboon test_custom_1000.csv.rot128
```

### Input Format

CSV file (หลังถอดรหัส ROT-128):
```csv
Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear
John Doe,10000,4242424242424242,123,12,2025
Jane Smith,5000,5555555555554444,456,01,2026
```

### Output Example

```
performing donations...
Configuration loaded successfully
Target .CSV file: test_custom_1000.csv.rot128

[SUCCESS]: charge transaction Seq.0 was success with amount: 100.00
[SUCCESS]: charge transaction Seq.1 was success with amount: 50.00

done.
Total received: 15000
successfully donated: 12000
faulty donation: 3000
average per person: 400.00

top donors:
John Doe
Jane Smith
Bob Wilson

Time usage = 2.345000 sec.
```

## Architecture

Clean Architecture pattern:
- **CLI Layer**: main.go
- **Service Layer**: internal/services
- **Domain Layer**: internal/domains
- **Adapter Layer**: internal/adapter
- **Infrastructure**: internal/cipher, validator, reporter, configs

## Development

```bash
# Build
go build -o bin/go-tamboon .

# Test
go test ./...

# Lint
golangci-lint run
```

## Security

- Secure memory management ด้วย `SecureString` และ `SecureByte`
- Auto-clear sensitive data หลังใช้งาน
- Base64 encoded API keys
- HTTPS-only communication

## Performance

- Concurrent processing ด้วย goroutines
- Rate limiting เพื่อป้องกัน API limits
- Exponential backoff สำหรับ retry logic
- Graceful shutdown และ timeout handling