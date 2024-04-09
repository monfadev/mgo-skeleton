# Monfadev Golang Skeleton

Terjemahan:

* Indonesia

## Ringkasan

Berikut ini merupakan tata letak dasar untuk proyek aplikasi Go.


# Architecture

DTO (Data transfer object)

Im adoption repository pattern, the pattern is layer per layer. as example:

Models: data object
Repository: only communication with database
Services: communication between bussiness logic (features)
Handler: communication between routes and service