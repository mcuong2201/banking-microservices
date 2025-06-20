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
