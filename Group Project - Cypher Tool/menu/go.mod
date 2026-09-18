module hive.fi/cypher/menu

go 1.26.0

replace hive.fi/cypher/encryption => ../encryption

replace hive.fi/cypher/decryption => ../decryption

require (
	hive.fi/cypher/decryption v0.0.0-00010101000000-000000000000
	hive.fi/cypher/encryption v0.0.0-00010101000000-000000000000
)
