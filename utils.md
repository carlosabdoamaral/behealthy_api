## DB

If database is dirty caused by the migrations, run the following commands

```
task migration-down
```

```
UPDATE public.schema_migrations SET dirty = false WHERE version = 1;
```
