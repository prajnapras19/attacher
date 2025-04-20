# Attacher
---

## Description
This is a service to serve file for users. Current features:
- Login as user or admin
- Authorized user can list active files.
- Authorized user can download active files.
- Admin user can bulk upsert to tables.

## Deployment
See `docker-compose.yml` file, adjust environment variables to your need, and run:
```
docker compose up -d
```

Migrate the database in `migrations` folder, and you're all set.

It is preferable to also define the attached volume, so that we can transfer the file from the volume.

## Bulk Upsert
We need to upsert conveniently to `users` and `attachments` table. Therefore, an endpoint for that is[Click here](https://docs.google.com/spreadsheets/d/18FEA3Hl09_AMeXCcxe4fRyDKvM2jP4fnaUkTaTUaDqo/edit?usp=sharing) to see bulk upsert example.

## Notes
### Compose File
The compose file use prebuilt image in order to avoid long build time in low-spec machine.
```
docker build . -t prajnapras19/attacher:1.0.0
docker push prajnapras19/attacher:1.0.0
```

### Local MySQL
In order to develop locally, we may see errors on creating trigger due to permission error. Try the following in mysql:
```
GRANT TRIGGER ON db_name.* TO 'user'@'localhost';
SHOW GRANTS FOR 'user'@'localhost';

SET GLOBAL log_bin_trust_function_creators = 1;
SHOW VARIABLES LIKE 'log_bin_trust_function_creators';
```

Then try to migrate:
```
make mysql-migrate-up
```

Remember to revoke access again when unused.