module hive.fi/cypher

go 1.26.0

replace hive.fi/cypher/menu => ./menu

replace hive.fi/cypher/encryption => ./encryption

replace hive.fi/cypher/decryption => ./decryption

require hive.fi/cypher/menu v0.0.0-00010101000000-000000000000

require (
	hive.fi/cypher/decryption v0.0.0-00010101000000-000000000000 // indirect
	hive.fi/cypher/encryption v0.0.0-00010101000000-000000000000 // indirect
)
