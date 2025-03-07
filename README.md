# Monfadev Golang Skeleton

Terjemahan:

- Indonesia

## Ringkasan

Berikut ini merupakan tata letak dasar untuk proyek aplikasi Go.

# Architecture

DTO (Data transfer object)

Im adoption repository pattern, the pattern is layer per layer. as example:

Models: data object
Repository: only communication with database
Services: communication between bussiness logic (features)
Handler: communication between routes and service

# Migration

migrate create -ext sql -dir db/migrations -seq create_users_table

## Migration Up

if there are changes, then run

## Migration Down

rollback from migration up

## Steps

- Ensure database created
- migrate -database "postgres://postgres:postgres@localhost:5432/mgoskeleton?sslmode=disable&TimeZone=Asia/Jakarta" -path migrations up

<img width="925" alt="Screenshot 2025-03-02 at 13 47 01" src="https://github.com/user-attachments/assets/331197d1-d7af-4ffa-a350-d4e9f4c57510" />

