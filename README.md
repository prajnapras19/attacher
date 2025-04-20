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

It is preferable to also define the attached volume, so that we can transfer the file from the volume.

## Bulk Upsert
We need to upsert conveniently to `users` and `attachments` table. Therefore, an endpoint for that is[Click here](https://docs.google.com/spreadsheets/d/18FEA3Hl09_AMeXCcxe4fRyDKvM2jP4fnaUkTaTUaDqo/edit?usp=sharing) to see bulk upsert example.

## Notes
The compose file use prebuilt image in order to avoid long build time in low-spec machine.
```
docker build . -t prajnapras19/attacher:1.0.0
docker push prajnapras19/attacher:1.0.0
```