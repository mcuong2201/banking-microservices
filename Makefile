# Makefile gốc – điều phối toàn hệ thống

all: docker-up run-auth run-profile run-ledger run-risk

docker-up:
	@echo "🐳 Starting Docker infra..."
	cd infra && docker compose up -d

docker-down:
	cd infra && docker compose down

run-auth:
	cd auth-service && make run

run-profile:
	cd user-profile-service && make run

run-ledger:
	cd transaction-ledger && make run

run-risk:
	cd risk-detection && make run

format-all:
	@echo "🎨 Format toàn bộ code..."
	cd auth-service && make format
	cd user-profile-service && make format
	cd transaction-ledger && make format
	cd risk-detection && make format

clean-all:
	@echo "🧹 Cleaning all..."
	cd auth-service && make clean
	cd user-profile-service && make clean
	cd transaction-ledger && make clean
	cd risk-detection && make clean

build-all:
	@echo "⚙️ Build tất cả service..."
	cd auth-service && make build
	cd user-profile-service && make build
	cd transaction-ledger && make build
	cd risk-detection && make build

test-all:
	@echo "🧪 Test toàn bộ service..."
	cd auth-service && make test
	cd user-profile-service && make test
	cd transaction-ledger && make test
	cd risk-detection && make test

format-all:
	@echo "🎨 Format toàn bộ code..."
	cd auth-service && make fmt
	cd user-profile-service && make fmt
	cd transaction-ledger && make fmt
	cd risk-detection && make fmt

clean-all:
	@echo "🧹 Cleaning all..."
	cd auth-service && make clean
	cd user-profile-service && make clean
	cd transaction-ledger && make clean
	cd risk-detection && make clean