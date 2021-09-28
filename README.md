**Status:**

Auth +
Get balance +
Make transaction -/+ (need to add timestamp)
Dockerization -/+ (Docker compose done, need to add automigration to use all of this out of box)
Migrations -/+ (Need to execute sql script manually)

How-To:

1. Start Postgres server.

2. Execute migrations/initial.sql on database

3. Fill environment with vars: 
   "POSTGRES_USER",
   "POSTGRES_PASSWORD"
   "POSTGRES_HOST"
   "POSTGRES_PORT"
   "POSTGRES_DB"
4. Install dependencies and run server.
5. You can use /api/balance and /api/transaction routes with header Authorization and its values for

**Alice**: "abd52071ae09da447b528cfe27fe47ae5e4788285ef04cea0359cea649fdfc0f84e1ee62368c937511164854fa53715e1f4abbc333ee7c0b042c975cc944aed6"

or

**Bob**: "d6dc3cbea6c9ed0e8e126759f4b156516cc9af7da69ee9f8bbde53f8feafb270e88a5c7b6b553cf942187c0f1a97b83ee030cf67d96ea0d92c342abab2664e5b"
